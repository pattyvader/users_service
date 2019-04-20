--Database: users_service
CREATE DATABASE users_service;

-- Table: users
CREATE TABLE users 
(  id serial PRIMARY KEY, 
   name VARCHAR(50) UNIQUE NOT NULL, 
   email VARCHAR(255) UNIQUE NOT NULL, 
   password VARCHAR(50) NOT NULL, 
   admin BOOLEAN NOT NULL
);