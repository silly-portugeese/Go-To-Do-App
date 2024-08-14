CREATE TYPE status_opts AS ENUM('Pending', 'In Progress', 'Completed');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar(50) NOT NULL UNIQUE,
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    task varchar(100) NOT NULL,
    status status_opts NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);