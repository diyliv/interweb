CREATE TABLE IF NOT EXISTS user_info(
    user_id SERIAL PRIMARY KEY,
    user_telegram_id BIGINT NOT NULL,
    user_first_request TIMESTAMP NOT NULL, 
    user_requests_count INTEGER NOT NULL
);