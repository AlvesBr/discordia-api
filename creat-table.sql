    CREATE TABLE product (
         id SERIAL PRIMARY KEY,
         name VARCHAR(50) NOT NULL,
         price NUMERIC(10, 2) NOT NULL
     );

    CREATE TABLE posts (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        handle VARCHAR(50) NOT NULL,
        initials VARCHAR(10) NOT NULL,
        avatar_gradient VARCHAR(100) NOT NULL,
        body TEXT NOT NULL,
        likes INT DEFAULT 0,
        reposts INT DEFAULT 0,
        liked BOOLEAN DEFAULT FALSE,
        reposted BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE comments (
        id SERIAL PRIMARY KEY,
        post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
        name VARCHAR(100) NOT NULL DEFAULT 'Anônimo',
        handle VARCHAR(50) NOT NULL,
        initials VARCHAR(5) NOT NULL,
        avatar_gradient VARCHAR(100) NOT NULL,
        body TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

    CREATE INDEX idx_comments_post_id ON comments(post_id);