CREATE TABLE IF NOT EXISTS users(
   id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
   uuid VARCHAR (20) UNIQUE,
   user_id VARCHAR (80) NOT NULL,
   password VARCHAR (200) NOT NULL
);
ALTER TABLE users ADD INDEX uuid_index(uuid);
ALTER TABLE users ADD INDEX user_index(user_id, password);