CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username VARCHAR(150),
  firstName VARCHAR(100),
  lastName VARCHAR(100),
  pass varchar(255) NOT NULL
);