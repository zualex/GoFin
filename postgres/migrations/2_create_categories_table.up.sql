CREATE TABLE categories (
   id serial PRIMARY KEY,
   name VARCHAR (255) UNIQUE NOT NULL,
   currency VARCHAR (3) NOT NULL
);