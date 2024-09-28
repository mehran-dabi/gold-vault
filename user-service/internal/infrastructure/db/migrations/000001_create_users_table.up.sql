-- Create an enum type for user roles
CREATE TYPE user_role AS ENUM ('admin', 'customer');

CREATE TABLE IF NOT EXISTS users (
   id BIGSERIAL PRIMARY KEY,
   phone VARCHAR(20) UNIQUE NOT NULL,
   first_name VARCHAR(100),
   last_name VARCHAR(100),
   national_code VARCHAR(20) UNIQUE,
   birthday DATE,
   role user_role DEFAULT 'customer',
   is_verified BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
