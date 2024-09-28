CREATE TABLE IF NOT EXISTS assets (
    id SERIAL PRIMARY KEY,
    wallet_id BIGINT REFERENCES wallets(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL,  -- e.g., 'gold', 'USD'
    balance NUMERIC(20, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (wallet_id, type)  -- Each wallet can have one entry per asset type
);
