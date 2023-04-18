DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS class;
CREATE TABLE class (
  className  VARCHAR(128) NOT NULL,
  lastUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  totalVotes INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`className`)
);

CREATE TABLE post (
  postID    INT AUTO_INCREMENT NOT NULL,
  postClassName   VARCHAR(128) NOT NULL,
  FOREIGN KEY (postClassName) REFERENCES class(className),
  postName  VARCHAR(128) NOT NULL,
  postContent VARCHAR(128) NOT NULL,
  postVotes INT NOT NULL DEFAULT 0,
  timePosted TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  (postClassName,postName,postContent,postVotes)
VALUES
  ("COP5000","Facebook Link", "www.facebook.com", 3);

INSERT INTO post
  (postClassName,postName,postContent,postVotes)
VALUES
  ("COP5000","Discord Link", "www.discord.com", 2);

INSERT INTO post
  (postClassName,postName,postContent,postVotes)
VALUES
  ("CIS4930","GroupMe Link", "www.groupme.com", 9)  