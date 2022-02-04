-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS products;
CREATE TABLE products (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  category   VARCHAR(128) NOT NULL,
  quantity   INT NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE products CONVERT TO CHARACTER SET utf8;

INSERT INTO products
  (title, category, quantity, price)
VALUES
  ('Супа Топчета', 'Супи', 1000, 1.49),
  ('Сос Бешамел', 'Сосове', 1000, 0.99),
  ('Салата Цезар', 'Салати', 1000, 5.99),
  ('Френска селска торта', 'Десерти', 1000, 3.49),
  ('Панираин ролца от раци', 'Предястия', 1000, 3.99),
  ('Руло Стефани', 'Основни', 1000, 6.99),
  ('Шоколадов мъфин', 'Десерти', 1000, 1.49),
  ('Спагети Болонезе', 'Основни', 1000, 7.49),
  ('Пилешко филе с гъбен сос', 'Основни', 1000, 6.99),
  ('Гръцка Салата', 'Салати', 1000, 5.49),
  ('Сос Холандез', 'Сосове', 1000, 0.99),
  ('Пилешки хапки с корнфлейкс', 'Предястия', 1000, 5.99);