DROP TABLE IF EXISTS class;
CREATE TABLE class (
  id         INT AUTO_INCREMENT NOT NULL,
  className  VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  postID    INT AUTO_INCREMENT NOT NULL,
  classID   INT REFERENCES class(id),
  postName  VARCHAR(128) NOT NULL,
  postContent VARCHAR(128) NOT NULL,
  PRIMARY KEY (`postID`)
);

INSERT INTO class
  (className)
VALUES
  ('CEN3031');

INSERT INTO posts
  (classID, postName, postContent)
VALUES
  (1, 'This class is awesome', 'Take this class!'),
  (1, 'This class good', 'yes' );