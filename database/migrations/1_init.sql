-- +goose Up
CREATE TABLE keyboards (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  creater_name VARCHAR(255),
  price int(10) NOT NULL,
  url VARCHAR(255) NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  split int(10) UNSIGNED NOT NULL,
  led int(10) UNSIGNED NOT NULL,
  key_num int(10) UNSIGNED NOT NULL,
  matrix VARCHAR(255) NOT NULL,
  key_profile VARCHAR(255) NOT NULL,
  color VARCHAR(255) NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE questions (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  question VARCHAR(255) NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +goose Down
DROP TABLE keyboards;
DROP TABLE questions;
