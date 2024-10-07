package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"products_server/models"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id string) (*models.ProductDetails, error) {
	query := `
		SELECT p.id, p.name, p.brand, p.price, p.discount, p.discounted_price, p.installment_value, p.max_installments, p.has_interest, p.highlight_status, p.store_name, p.image_link, p.category,
		       s.arrive_today, r.total_reviews, r.review_score
		FROM product p
		LEFT JOIN shipping_info s ON p.id = s.product_id
		LEFT JOIN product_review r ON p.id = r.product_id
		WHERE p.id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var product models.ProductDetails
	if err := row.Scan(&product.ID, &product.Name, &product.Brand, &product.Price, &product.Discount, &product.DiscountedPrice, &product.InstallmentValue, &product.MaxInstallments, &product.HasInterest, &product.HighlightStatus, &product.StoreName, &product.ImageLink, &product.Category, &product.ArriveToday, &product.TotalReviews, &product.ReviewScore); err != nil {
		return nil, fmt.Errorf("failed to scan product: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context, page, pageSize int, textFilter string) ([]models.ProductDetails, int, error) {
	offset := (page - 1) * pageSize

	var totalRecords int
	countQuery := `SELECT COUNT(*) FROM product`
	if textFilter != "" {
		countQuery += " WHERE name ILIKE '%' || $1 || '%'"
	}

	if textFilter != "" {
		if err := r.db.QueryRow(ctx, countQuery, textFilter).Scan(&totalRecords); err != nil {
			return nil, 0, fmt.Errorf("failed to get total records: %w", err)
		}
	} else {
		if err := r.db.QueryRow(ctx, countQuery).Scan(&totalRecords); err != nil {
			return nil, 0, fmt.Errorf("failed to get total records: %w", err)
		}
	}

	query := `
		SELECT p.id, p.name, p.brand, p.price, p.discount, p.discounted_price, p.installment_value, p.max_installments, p.has_interest, p.highlight_status, p.store_name, p.image_link, p.category,
		       s.arrive_today, r.total_reviews, r.review_score
		FROM product p
		LEFT JOIN shipping_info s ON p.id = s.product_id
		LEFT JOIN product_review r ON p.id = r.product_id`

	if textFilter != "" {
		query += " WHERE p.name ILIKE '%' || $1 || '%'"
		query += " LIMIT $2 OFFSET $3"
		rows, err := r.db.Query(ctx, query, textFilter, pageSize, offset)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		var products []models.ProductDetails
		for rows.Next() {
			var p models.ProductDetails
			if err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Price, &p.Discount, &p.DiscountedPrice, &p.InstallmentValue, &p.MaxInstallments, &p.HasInterest, &p.HighlightStatus, &p.StoreName, &p.ImageLink, &p.Category, &p.ArriveToday, &p.TotalReviews, &p.ReviewScore); err != nil {
				return nil, 0, fmt.Errorf("failed to scan row: %w", err)
			}
			products = append(products, p)
		}

		if rows.Err() != nil {
			return nil, 0, fmt.Errorf("failed to iterate rows: %w", rows.Err())
		}

		return products, totalRecords, nil
	}

	query += " LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var products []models.ProductDetails
	for rows.Next() {
		var p models.ProductDetails
		if err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Price, &p.Discount, &p.DiscountedPrice, &p.InstallmentValue, &p.MaxInstallments, &p.HasInterest, &p.HighlightStatus, &p.StoreName, &p.ImageLink, &p.Category, &p.ArriveToday, &p.TotalReviews, &p.ReviewScore); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		products = append(products, p)
	}

	if rows.Err() != nil {
		return nil, 0, fmt.Errorf("failed to iterate rows: %w", rows.Err())
	}

	return products, totalRecords, nil
}
