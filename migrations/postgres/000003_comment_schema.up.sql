CREATE TABLE IF NOT EXISTS comments(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    blog_id INT NOT NULL,
    content TEXT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
                ON DELETE CASCADE,
    CONSTRAINT fk_blog
        FOREIGN KEY(blog_id)
            REFERENCES blogs(id)
                ON DELETE CASCADE
)
