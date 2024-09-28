CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    wallet_id BIGINT REFERENCES wallets(id) ON DELETE CASCADE, -- Reference to the wallet
    asset_id BIGINT REFERENCES assets(id) ON DELETE CASCADE, -- Reference to the asset
    type VARCHAR(10) NOT NULL,  -- 'credit' or 'debit'
    amount NUMERIC(20, 2) NOT NULL,
    description TEXT,
    status VARCHAR(10) NOT NULL, -- 'pending', 'completed', 'failed'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
