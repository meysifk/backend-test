CREATE DATABASE simple_db;
USE simple_db;

CREATE TABLE users (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  userName VARCHAR(30) NOT NULL,
  parentId INT DEFAULT NULL UNIQUE,
  FOREIGN KEY (parentId) REFERENCES users(id)
);

INSERT INTO users (userName) VALUES
('Ali'),
('Budi'),
('Cecep');

UPDATE users SET parentId=2 WHERE id=1;
UPDATE users SET parentId=1 WHERE id=3;

SELECT * FROM users;
SELECT u.id, u.userName, p.userName AS "ParentUserName" FROM users u LEFT JOIN users p ON p.id = u.parentId;
