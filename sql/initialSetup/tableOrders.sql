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
  ('a980ced5-4acb-4be8-51e3-4565e06082ea', 'Georgi Georgiev', 'Vratsa Pop Kiro 15', '0888888856', 95.97, 'On the way'),
  ('b8aeed64-1c6d-40bb-6d5e-6e1efbab5dcb', 'Kristina Mincheva', 'Sofia Nadejda 1', '0888848853', 119.97, 'Accepted');