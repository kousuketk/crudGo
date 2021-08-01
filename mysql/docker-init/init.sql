DROP database if exists godb;
CREATE DATABASE godb;
USE godb;

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  self_introduction text,
  email varchar(255) NOT NULL,
  password_digest varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  phone_number varchar(255) NOT NULL,
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,
  deleted_at datetime,
  PRIMARY KEY (id)
);

INSERT INTO users(id, name, self_introduction, email, password_digest, address, phone_number) VALUES(1, 'test_user1', 'test1_self_introduction', 'testuser1@test.com', 'testuser1_passowrd_digest', 'testuser1_address', '00000000001');
INSERT INTO users(id, name, self_introduction, email, password_digest, address, phone_number) VALUES(2, 'test_user2', 'test2_self_introduction', 'testuser2@test.com', 'testuser2_passowrd_digest', 'testuser2_address', '00000000002');
INSERT INTO users(id, name, self_introduction, email, password_digest, address, phone_number) VALUES(3, 'test_user3', 'test3_self_introduction', 'testuser3@test.com', 'testuser3_passowrd_digest', 'testuser3_address', '00000000003');
INSERT INTO users(id, name, self_introduction, email, password_digest, address, phone_number) VALUES(4, 'test_user4', 'test4_self_introduction', 'testuser4@test.com', 'testuser4_passowrd_digest', 'testuser4_address', '00000000004');
