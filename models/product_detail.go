package models

type ProductDetails struct {
	Product
	ArriveToday  bool    `json:"arrive_today"`
	TotalReviews int     `json:"total_reviews"`
	ReviewScore  float64 `json:"review_score"`
}
