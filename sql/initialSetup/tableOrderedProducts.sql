-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS orderedProduct;

CREATE TABLE orderedProduct (
  id          VARCHAR(256) NOT NULL,
  product_id  VARCHAR(256) NOT NULL,
  quantity    INT NOT NULL,
  order_id    VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE orderedProduct CONVERT TO CHARACTER SET utf8;

INSERT INTO orderedProduct
  (id, product_id, quantity,  order_id)
VALUES
  ('UUID', 'PRODUCT-123', 10, 'ORDER-123');