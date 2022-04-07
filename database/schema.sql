CREATE TYPE authmode AS ENUM ('LOCAL', 'OAUTH2');

CREATE TABLE IF NOT EXISTS users (
  username      VARCHAR(24) PRIMARY KEY NOT NULL UNIQUE,
  password      VARCHAR(100) DEFAULT NULL,
  auth_mode     authmode NOT NULL DEFAULT 'LOCAL',
  f_root_folder VARCHAR(256) DEFAULT NULL,
  private_key   VARCHAR(256) NOT NULL,
  public_key    VARCHAR(256) NOT NULL
);


CREATE TABLE IF NOT EXISTS folders (
  id        VARCHAR(256) PRIMARY KEY NOT NULL UNIQUE,
  name      VARCHAR(128) NOT NULL,
  f_owner   VARCHAR(24) DEFAULT NULL,
  f_parent  VARCHAR(256) DEFAULT NULL,

  FOREIGN KEY (f_parent)
    REFERENCES folders(id),

  FOREIGN KEY (f_owner)
    REFERENCES users(username)
);

ALTER TABLE users ADD
  FOREIGN KEY (f_root_folder)
    REFERENCES folders(id);

CREATE TABLE IF NOT EXISTS files (
  id          VARCHAR(256) PRIMARY KEY NOT NULL UNIQUE,
  name        VARCHAR(128) NOT NULL,
  size        BIGINT NOT NULL,
  type        VARCHAR(128) NOT NULL,
  is_favorite BOOLEAN NOT NULL DEFAULT FALSE,
  key         VARCHAR(256) NOT NULL,
  f_parent    VARCHAR(256) NOT NULL,
  f_owner     VARCHAR(24) NOT NULL,

  CONSTRAINT c_owner
    FOREIGN KEY (f_owner)
      REFERENCES users(username),

  CONSTRAINT c_parent
    FOREIGN KEY (f_parent)
      REFERENCES folders(id)
        ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS shares (
  id          VARCHAR(256) PRIMARY KEY NOT NULL UNIQUE,
  key         VARCHAR(256) DEFAULT NULL,
  is_favorite BOOLEAN NOT NULL DEFAULT FALSE,
  is_file     BOOLEAN NOT NULL,  
  f_file      VARCHAR(256) DEFAULT NULL,
  f_folder    VARCHAR(256) DEFAULT NULL,
  f_owner     VARCHAR(24) NOT NULL,
  f_parent    VARCHAR(256) NOT NULL,

  CONSTRAINT c_file
    FOREIGN KEY (f_file)
      REFERENCES files(id),

  CONSTRAINT c_folder
    FOREIGN KEY (f_folder)
      REFERENCES folders(id),

  CONSTRAINT c_parent
    FOREIGN KEY (f_parent)
      REFERENCES folders(id),

  CONSTRAINT c_owner
    FOREIGN KEY (f_owner)
      REFERENCES users(username)
);