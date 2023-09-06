ALTER TABLE products
ADD COLUMN "StoreID" integer;

ALTER TABLE products
ADD CONSTRAINT fk_store
FOREIGN KEY ("StoreID")
REFERENCES "store" ("StoreID");