-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS products;
CREATE TABLE products (
  id         VARCHAR(256) NOT NULL,
  name      VARCHAR(128) NOT NULL,
  category   VARCHAR(128) NOT NULL,
  quantity   INT NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE products CONVERT TO CHARACTER SET utf8;

INSERT INTO products
  (id, name, category, quantity, price)
VALUES
  ('PRODUCT-123', 'Супа Топчета', 'Супи', 1000, 1.49);