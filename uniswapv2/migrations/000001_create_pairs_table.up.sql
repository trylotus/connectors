CREATE TABLE IF NOT EXISTS v2_pairs (
    number INT8 PRIMARY KEY,
    address TEXT NOT NULL,
    token0 TEXT NOT NULL,
    token1 TEXT NOT NULL
);

CREATE UNIQUE INDEX ON v2_pairs (address);
