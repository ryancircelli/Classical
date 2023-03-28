DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS class;
CREATE TABLE class (
  id         INT AUTO_INCREMENT NOT NULL,
  className  VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE post (
  postID    INT AUTO_INCREMENT NOT NULL,
  classID   INT,
  FOREIGN KEY (classID) REFERENCES class(id),
  postName  VARCHAR(128) NOT NULL,
  postContent VARCHAR(128) NOT NULL,
  postVotes INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`postID`)
);

INSERT INTO class
  (className)
VALUES
  ('COP5000');

INSERT INTO class
  (className)
VALUES
  ('CIS4930');

INSERT INTO class
  (className)
VALUES
  ('CGS3065');

INSERT INTO post
  (classID,postName,postContent,postVotes)
VALUES
  (1,"Facebook Link", "www.facebook.com", 3);

INSERT INTO post
  (classID,postName,postContent,postVotes)
VALUES
  (1,"Discord Link", "www.discord.com", 2);

INSERT INTO post
  (classID,postName,postContent,postVotes)
VALUES
  (2,"GroupMe Link", "www.groupme.com", 9)  