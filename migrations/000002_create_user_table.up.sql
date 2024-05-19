CREATE TABLE IF NOT EXISTS "user" (
  "id" UUID                   PRIMARY KEY DEFAULT uuid_generate_v4(),
  "mail"                      TEXT NOT NULL UNIQUE,
  "created_at"                TIMESTAMP default current_timestamp,
  "updated_at"                TIMESTAMP default current_timestamp
);