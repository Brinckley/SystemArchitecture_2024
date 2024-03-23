CREATE SCHEMA IF NOT EXISTS social_network;

CREATE TABLE IF NOT EXISTS social_network.account (
  id            SERIAL PRIMARY KEY, 
  username      VARCHAR (50) UNIQUE NOT NULL, 
  password      VARCHAR (255) NOT NULL, 
  first_name    VARCHAR (255) NOT NULL, 
  last_name     VARCHAR (255) NOT NULL, 
  email         VARCHAR (255) UNIQUE NOT NULL
);

CREATE TABLE social_network.post (
  id            SERIAL PRIMARY KEY,
  account_id       INT NOT NULL,
  content       TEXT NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(account_id) REFERENCES social_network.account(id)
);

CREATE TABLE social_network.message (
  id            SERIAL PRIMARY KEY,
  sender_id     INT NOT NULL,
  receiver_id   INT NOT NULL,
  content       TEXT NOT NULL,
  CONSTRAINT fk_sender   FOREIGN KEY(sender_id) REFERENCES social_network.account(id),
  CONSTRAINT fk_receiver FOREIGN KEY(receiver_id) REFERENCES social_network.account(id)
);
