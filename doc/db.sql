CREATE DATABASE fcms;


-- 用户表user
CREATE TABLE users(
  user_id serial,
  user_name varchar(255),
  pass_word text,
  reg_date timestamp,
  PRIMARY KEY(user_id)
)

-- 文件存储表
CREATE TABLE files(
  file_id bigserial,
  file_name varchar(255),
  file_data  text,
  file_upload_date timestamp,
  PRIMARY KEY(file_id)
)