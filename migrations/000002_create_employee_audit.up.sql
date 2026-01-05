CREATE TABLE IF NOT EXISTS admins (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    position VARCHAR(100) NOT NULL,
    salary DOUBLE PRECISION NOT NULL,
    profile_picture VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Seed Admin
INSERT INTO admins (name, email, password) 
VALUES ('admin', 'admin@example.com', 'SecredPassword123!') 
ON CONFLICT (email) DO NOTHING;
