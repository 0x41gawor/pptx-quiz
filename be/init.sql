-- init.sql
DROP TABLE IF EXISTS completions;

CREATE TABLE completions (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    id_passport_number VARCHAR(50),
    company VARCHAR(150),
    date DATE NOT NULL DEFAULT CURRENT_DATE
);
