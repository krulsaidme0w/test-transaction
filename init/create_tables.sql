CREATE TABLE test_user
(
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    balance BIGINT NOT NULL CHECK (balance > 0)
);

CREATE TABLE transactions
(
    user_id VARCHAR(64) REFERENCES test_user(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT id PRIMARY KEY (user_id, created_at) 
);
