
CREATE DATABASE testdb CHARACTER SET utf8mb4;

CREATE TABLE testdb.user (
  id int(11) not null AUTO_INCREMENT,
  name VARCHAR(20) not null,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;