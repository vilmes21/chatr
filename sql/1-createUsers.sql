CREATE TABLE users
(
  id SERIAL PRIMARY KEY,
  age integer,
  first_name character varying(255),
  last_name character varying(255),
  email text
)