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
  ('bc264186-9c2e-4533-6ba5-705c160303c1', 'Men Blue T-shirt', 'Men shirts', 1000, 19.99),
  ('376d4186-4533-9c2e-6ba5-484c1604543c', 'Men Trousers', 'Men Trousers', 1000, 35.99),
  ('0b45c13b-376d-4a1c-6941-4a7bd9ba03fc', 'Men White Jeans', 'Men Trousers', 1000, 39.99),
  ('1ac5fc58-8360-48a1-526a-6dfbf8c3fb59', 'Women Red Bag', 'Women bags', 1000, 39.99);