CREATE TABLE friend
(
  user_id INTEGER REFERENCES users(id),
  user2_id INTEGER REFERENCES users(id),
  PRIMARY KEY (user_id, user2_id)
)