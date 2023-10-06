CREATE TABLE "users" (
  "username" varchar  PRIMARY KEY NOT NULL DEFAULT 0,
  "password" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username", "phone_number")
