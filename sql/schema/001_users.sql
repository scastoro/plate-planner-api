-- +goose Up
CREATE SCHEMA "Admin";

CREATE TABLE "Admin"."Users" (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  first_name TEXT,
  last_name TEXT,
  body_weight DECIMAL,
  username TEXT,
  email TEXT,
  password TEXT,
  lastLoggedIn TIMESTAMP
);

-- +goose Down
DROP TABLE "Admin"."Users";

DROP SCHEMA "Admin";