CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    duration INT NOT NULL
);

CREATE TABLE products (
    code varchar PRIMARY KEY,
    name varchar NOT NULL,
    weight numeric NOT NULL,
    description varchar NOT NULL
);
