CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username VARCHAR(255),
  firstName VARCHAR(255),
  lastName VARCHAR(255),
  pass varchar(255)
);