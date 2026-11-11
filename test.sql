CREATE TABLE users (
  id  TEXT UNIQUE NOT NULL,
  admin BOOL,
  created DATETIME,

  PRIMARY KEY(id)
);

CREATE TABLE invoices (
  id        INT AUTO_INCREMENT NOT NULL,
  user_id   TEXT NOT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
)

CREATE TABLE mailboxes (
  id        INT AUTO_INCREMENT NOT NULL,
  user_id   TEXT NOT NULL,
  prefix    TEXT NOT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);