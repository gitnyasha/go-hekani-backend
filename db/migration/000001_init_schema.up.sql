CREATE TABLE "answers" (
  "id" bigserial PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "question_id" int NOT NULL,
  "title" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "article_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "question_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "articles" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "link" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "article_category_id" int NOT NULL,
  "user_id" int UNIQUE NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "bio" varchar NOT NULL,
  "birth" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "image" varchar NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "user_id" int UNIQUE NOT NULL,
  "answer_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "likes" (
  "id" bigserial PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "article_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "votes" (
  "id" bigserial PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "answer_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "questions" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" int UNIQUE NOT NULL,
  "question_category_id" int NOT NULL
);

CREATE TABLE "relationships" (
  "id" bigserial PRIMARY KEY,
  "follower_id" int NOT NULL,
  "followed_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "replies" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "article_id" int NOT NULL,
  "user_id" int UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_article_relationships" (
  "id" bigserial PRIMARY KEY,
  "follower_id" int NOT NULL,
  "followed_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_question_relations" (
  "id" bigserial PRIMARY KEY,
  "follower_id" int NOT NULL,
  "followed_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "index_likes_on_article_id" ON "likes" ("article_id");

CREATE INDEX "index_likes_on_user_id" ON "likes" ("user_id");

CREATE INDEX "index_votes_on_answer_id" ON "votes" ("answer_id");

CREATE INDEX "index_votes_on_user_id" ON "votes" ("user_id");

CREATE INDEX "index_relationships_on_followed_id" ON "relationships" ("followed_id");

CREATE UNIQUE INDEX "index_relationships_on_follower_id_and_followed_id" ON "relationships" ("follower_id", "followed_id");

CREATE INDEX "index_relationships_on_follower_id" ON "relationships" ("follower_id");

CREATE INDEX "index_user_article_relationships_on_followed_id" ON "user_article_relationships" ("followed_id");

CREATE UNIQUE INDEX "index_user_article_relationships_on_follower_id_and_followed_id" ON "user_article_relationships" ("follower_id", "followed_id");

CREATE INDEX "index_user_article_relationships_on_follower_id" ON "user_article_relationships" ("follower_id");

CREATE INDEX "index_user_question_relations_on_followed_id" ON "user_question_relations" ("followed_id");

CREATE UNIQUE INDEX "index_user_question_relations_on_follower_id_and_followed_id" ON "user_question_relations" ("follower_id", "followed_id");

CREATE INDEX "index_user_question_relations_on_follower_id" ON "user_question_relations" ("follower_id");

ALTER TABLE "answers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "articles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "replies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
