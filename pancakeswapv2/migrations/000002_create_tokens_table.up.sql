CREATE TABLE IF NOT EXISTS bep20_tokens (
    address BYTEA PRIMARY KEY,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    decimals INT8 NOT NULL
);
