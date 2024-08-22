CREATE TABLE repository_events (
    id SERIAL PRIMARY KEY,
    action VARCHAR(50),
    repo_name VARCHAR(100),
    repo_owner VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
