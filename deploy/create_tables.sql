
CREATE TABLE IF NOT EXISTS url_tab (
  id INT(11) NOT NULL AUTO_INCREMENT,
  short_url VARCHAR(128) NOT NULL,
  full_url VARCHAR(1024) NOT NULL,
  active SMALLINT NOT NULL,
  create_time INT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (short_url)
) ENGINE=InnoDB