DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS genres;

CREATE TABLE genres (
  value         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`value`)
);

CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  genre     INT NOT NULL,
  amount     INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`genre`) REFERENCES genres(`value`)
);

INSERT INTO genres
  (name, value)
VALUES
  ('Adventure', 1),
  ('Classics', 2),
  ('Fantasy', 3);
