CREATE TABLE IF NOT EXISTS v2_pairs (
    number INT8 NOT NULL,
    address BYTEA PRIMARY KEY,
    token0 BYTEA NOT NULL,
    token1 BYTEA NOT NULL,
    block_number INT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS v2_pairs_scanned_block (
    id INT8 PRIMARY KEY,
    number INT8 NOT NULL
);
