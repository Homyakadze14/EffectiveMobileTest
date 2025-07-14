CREATE TABLE IF NOT EXISTS subscriptions(
    id SERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price INT NOT NULL,
    user_id TEXT NOT NULL,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP
);