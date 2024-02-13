CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "title" varchar NOT NULL,
  "category" varchar NOT NULL,
  "description" varchar,
  "completed" boolean NOT NULL DEFAULT false
);
