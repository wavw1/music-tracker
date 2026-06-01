CREATE TABLE IF NOT EXISTS tracker.ideas (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    bpm INTEGER,
    tonic_key VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL,
    tags TEXT[]
);