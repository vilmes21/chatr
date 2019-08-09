CREATE TABLE sentence
(
  chat_speaker_id INTEGER REFERENCES chat_speaker(id),
  content text NOT NULL,
  time timestamp with time zone NOT NULL
)