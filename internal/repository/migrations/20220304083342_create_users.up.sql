CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    email VARCHAR (10) UNIQUE NOT NULL,
    password VARCHAR (100) NOT NULL
);
