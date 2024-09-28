CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,  -- Foreign key referencing the user who placed the order
    asset_type VARCHAR(50) NOT NULL,  -- e.g., 'gold'
    quantity NUMERIC(20, 6) NOT NULL,  -- Quantity requested in the order
    price NUMERIC(20, 6) NOT NULL,  -- Requested price for the order
    order_type VARCHAR(10) NOT NULL,  -- 'buy' or 'sell'
    status VARCHAR(20) DEFAULT 'pending',  -- 'pending', 'completed', 'cancelled'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
