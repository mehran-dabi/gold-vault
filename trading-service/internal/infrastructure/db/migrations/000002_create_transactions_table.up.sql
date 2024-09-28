CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,  -- Foreign key referencing the user who initiated the trade
    asset_type VARCHAR(50) NOT NULL,  -- e.g., 'gold'
    quantity NUMERIC(20, 6) NOT NULL,  -- Quantity of the product traded
    price NUMERIC(20, 6) NOT NULL,  -- Price at which the trade was executed
    transaction_type VARCHAR(10) NOT NULL,  -- 'buy' or 'sell'
    status VARCHAR(20) DEFAULT 'completed',  -- 'pending', 'completed', 'failed'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
