-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
  id        VARCHAR(256) NOT NULL,
  name      VARCHAR(128) NOT NULL,
  address   VARCHAR(256) NOT NULL,
  phone     VARCHAR(256) NOT NULL,
  price     DECIMAL(5,2) NOT NULL,
  status     VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE orders CONVERT TO CHARACTER SET utf8;

INSERT INTO orders
  (id, name, address, phone, price, status)
VALUES
  ('ORDER-123', 'Кристина Йотова', 'Мездра', '0888888856', 11.49, 'On the way');