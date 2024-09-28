CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    asset_type VARCHAR(50) NOT NULL,  -- e.g., 'gold'
    total_quantity NUMERIC(20, 6) DEFAULT 0 NOT NULL,  -- Total quantity available
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (asset_type)
);
