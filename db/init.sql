-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Create the 'tasks' table
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    user_id UUID REFERENCES users(id),
    completed_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

