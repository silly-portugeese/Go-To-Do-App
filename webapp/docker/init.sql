CREATE TYPE status_opts AS ENUM('Pending', 'In Progress', 'Completed');

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    task varchar(100) NOT NULL,
    status status_opts NOT NULL
);