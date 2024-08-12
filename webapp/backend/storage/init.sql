CREATE TABLE mytodos (
    id SERIAL PRIMARY KEY,
    task varchar(100) NOT NULL,
    status status_opts NOT NULL
);