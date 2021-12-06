DROP TABLE IF EXISTS books;
CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  genre     INT NOT NULL,
  amount     INT NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO books
  (name, price, genre, amount)
VALUES
  ('Blue Train', 56.99, 1, 23),
  ('Giant Steps', 63.99, 2, 12),
  ('Jeru', 17.99, 3, 42),
  ('Sarah Vaughan', 34.98, 2, 3);
