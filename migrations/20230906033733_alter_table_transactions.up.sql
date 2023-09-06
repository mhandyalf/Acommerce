ALTER TABLE transactions
ADD store_id INTEGER REFERENCES Store(store_id);