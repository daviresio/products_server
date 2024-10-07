CREATE TABLE product (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         name VARCHAR(255) NOT NULL,
                         brand VARCHAR(255) NOT NULL,
                         price DECIMAL(10, 2) NOT NULL,
                         discount DECIMAL(5, 2) DEFAULT 0,
                         discounted_price DECIMAL(10, 2) NOT NULL,
                         installment_value DECIMAL(10, 2),
                         max_installments INTEGER,
                         has_interest BOOLEAN,
                         highlight_status VARCHAR(50),
                         store_name VARCHAR(255) NOT NULL,
                         image_link TEXT,
                         category VARCHAR(50) NOT NULL,
                        is_deleted BOOLEAN DEFAULT FALSE,
                        created_at TIMESTAMP DEFAULT now(),
                        updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE shipping_info (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               product_id UUID REFERENCES product(id) ON DELETE CASCADE,
                               arrive_today BOOLEAN DEFAULT FALSE,
                               is_deleted BOOLEAN DEFAULT FALSE,
                               created_at TIMESTAMP DEFAULT now(),
                               updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE product_review (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                product_id UUID REFERENCES product(id) ON DELETE CASCADE,
                                total_reviews INTEGER DEFAULT 0,
                                review_score DECIMAL(3, 2) DEFAULT 0.00,
                                is_deleted BOOLEAN DEFAULT FALSE,
                                created_at TIMESTAMP DEFAULT now(),
                                updated_at TIMESTAMP DEFAULT now()
);
