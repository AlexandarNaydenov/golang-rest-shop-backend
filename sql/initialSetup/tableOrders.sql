-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
  id         INT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  address   VARCHAR(256) NOT NULL,
  phone     VARCHAR(256) NOT NULL,
  products  VARCHAR(512) NOT NULL,
  price     DECIMAL(5,2) NOT NULL,
  status     VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE orders CONVERT TO CHARACTER SET utf8;

INSERT INTO orders
  (id, name, address, phone, products, price, status)
VALUES
  (23145, 'Кристина Йотова', 'Мездра', '0888888856', 'Супа Топчета - 2;Салата Цезар - 3;', 11.49, 'On the way');