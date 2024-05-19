CREATE TABLE IF NOT EXISTS "rate" (
  "id" UUID                   PRIMARY KEY DEFAULT uuid_generate_v4(),
  "date"                      TEXT DEFAULT to_char( now(), 'DD.MM.YYYY' ),
  "rate"                      FLOAT DEFAULT 0,
  "val_code"                  TEXT NOT NULL UNIQUE,
  "code"                      INTEGER,
  "created_at"                TIMESTAMP default current_timestamp,
  "updated_at"                TIMESTAMP default current_timestamp
);