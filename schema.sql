CREATE TABLE product (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "sku" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "price" float NOT NULL,
  PRIMARY KEY(id)
);