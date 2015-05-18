-- Reset
DROP TABLE IF EXISTS users;

-- Tables
CREATE TABLE IF NOT EXISTS Users (
  ID SERIAL PRIMARY KEY,
  first_name VARCHAR(80) NOT NULL,
  middle_name VARCHAR(80) NOT NULL,
  last_name VARCHAR(80) NOT NULL,
  encrypted_password VARCHAR(255),
  facebook_user_id VARCHAR(255),
  google_user_id VARCHAR(255),
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);
