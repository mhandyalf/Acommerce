CREATE TABLE store (
    store_id serial PRIMARY KEY,
    nama_store VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    longitude DECIMAL(10, 6) NOT NULL,
    latitude DECIMAL(10, 6) NOT NULL,
    rating DECIMAL(3, 2)
);

