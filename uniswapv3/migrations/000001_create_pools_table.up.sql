CREATE TABLE IF NOT EXISTS v3_pools (
    address TEXT PRIMARY KEY,
    token0 TEXT NOT NULL,
    token1 TEXT NOT NULL,
    fee INT8 NOT NULL,
    tick_spacing INT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS v3_pools_scanned_block (
    id INT8 PRIMARY KEY,
    number INT8 NOT NULL
);
