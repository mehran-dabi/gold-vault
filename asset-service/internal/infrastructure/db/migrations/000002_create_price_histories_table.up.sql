CREATE TABLE IF NOT EXISTS price_histories(
    id SERIAL PRIMARY KEY,
    asset_type VARCHAR(50) NOT NULL,  -- e.g., 'gold'
    buy_price NUMERIC(20, 2) NOT NULL,  -- Price at which the asset is bought
    sell_price NUMERIC(20, 2) NOT NULL,  -- Price at which the asset is sold
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Timestamp of when the price was recorded
);
