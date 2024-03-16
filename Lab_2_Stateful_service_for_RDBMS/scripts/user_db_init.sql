CREATE SCHEMA IF NOT EXISTS social_network;

CREATE TABLE IF NOT EXISTS social_network.account (
  id             SERIAL PRIMARY KEY, 
  username       VARCHAR (50) UNIQUE NOT NULL, 
  password       VARCHAR (50) NOT NULL, 
  first_name     VARCHAR (255) NOT NULL, 
  last_name      VARCHAR (255) NOT NULL, 
  email          VARCHAR (255) UNIQUE NOT NULL
);