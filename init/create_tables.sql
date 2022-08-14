CREATE TABLE test_user
(
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    balance BIGINT NOT NULL CHECK (balance > 0)
);

CREATE TABLE transaction
(
    user_id VARCHAR(64) REFERENCES test_user(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    status VARCHAR(32) NOT NULL,
    key VARCHAR(64) NOT NULL,
    CONSTRAINT id PRIMARY KEY (user_id, created_at) 
);

CREATE INDEX shorturl_key_hash_index ON transaction USING hash(key);
