CREATE TABLE chat_speaker
(
  id SERIAL PRIMARY KEY,
  chat_id INTEGER REFERENCES chat(id),
  user_id INTEGER REFERENCES users(id),
  nickname varchar(255),
  CONSTRAINT unique_chat_speaker UNIQUE(chat_id, user_id)
)