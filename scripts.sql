CREATE TABLE IF NOT EXISTS "user" (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS drug (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    approved BOOLEAN NOT NULL,
    minDose INT,
    maxDose INT,
    availableAt DATE
);

CREATE TABLE IF NOT EXISTS vaccination (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    drugId BIGINT REFERENCES drug(id),
    dose INT,
    date DATE
);