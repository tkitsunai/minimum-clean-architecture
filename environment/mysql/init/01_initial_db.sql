ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'test';
ALTER USER 'test'@'%' IDENTIFIED WITH mysql_native_password BY 'test';

CREATE DATABASE testdb CHARACTER SET utf8mb4;

CREATE TABLE testdb.user (
  id int(11) not null AUTO_INCREMENT,
  name VARCHAR(20) not null,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO testdb.user values (null, 'ichiro yamada');
INSERT INTO testdb.user values (null, 'jiro yamada');
INSERT INTO testdb.user values (null, 'saburo yamada');