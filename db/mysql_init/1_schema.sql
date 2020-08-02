USE go_rest_api;

CREATE TABLE IF NOT EXISTS articles (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title       VARCHAR(256) NOT NULL,
  url         VARCHAR(256) NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS categories (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name        VARCHAR(256) NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;