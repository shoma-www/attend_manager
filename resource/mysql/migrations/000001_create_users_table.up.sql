CREATE TABLE IF NOT EXISTS users(
   id VARCHAR (20) PRIMARY KEY,
   user_id VARCHAR (100) NOT NULL,
   password VARCHAR (250) NOT NULL
);
ALTER TABLE users ADD INDEX login_index(user_id, password)