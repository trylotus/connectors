CREATE TABLE IF NOT EXISTS erc20_tokens (
    address TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    decimals INT4 NOT NULL
);
