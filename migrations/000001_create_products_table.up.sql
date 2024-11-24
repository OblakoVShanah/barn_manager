CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    weight_per_pkg INT UNSIGNED NOT NULL,
    amount INT UNSIGNED NOT NULL,
    price_per_pkg DECIMAL(10,2) NOT NULL,
    expiration_date DATETIME NOT NULL,
    present_in_fridge BOOLEAN NOT NULL,
    proteins INT NOT NULL,
    fats INT NOT NULL,
    carbohydrates INT NOT NULL,
    calories INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_name (name),
    INDEX idx_expiration_date (expiration_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 

