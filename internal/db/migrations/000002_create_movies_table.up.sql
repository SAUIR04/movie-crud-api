CREATE TABLE IF NOT EXISTS movies (
                                      id SERIAL PRIMARY KEY,
                                      title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    genre VARCHAR(100),
    release_date VARCHAR(100)
    );
