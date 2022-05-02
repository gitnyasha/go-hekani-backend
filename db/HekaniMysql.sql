CREATE TABLE `answers` (
  `id` int PRIMARY KEY,
  `user_id` int UNIQUE NOT NULL,
  `question_id` int,
  `title` varchar(255),
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `article_categories` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `question_categories` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `articles` (
  `id` int PRIMARY KEY,
  `title` varchar(255),
  `link` varchar(255),
  `image` varchar(255),
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now()),
  `article_category_id` int
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `email` varchar(255) UNIQUE,
  `hashed_password` varchar(255) NOT NULL,
  `bio` varchar(255),
  `birth` date,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now()),
  `image` varchar(255)
);

CREATE TABLE `comments` (
  `id` int PRIMARY KEY,
  `title` varchar(255),
  `user_id` int UNIQUE NOT NULL,
  `answer_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `likes` (
  `id` int PRIMARY KEY,
  `user_id` int UNIQUE NOT NULL,
  `article_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `votes` (
  `id` int PRIMARY KEY,
  `user_id` int UNIQUE NOT NULL,
  `answer_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `questions` (
  `id` int PRIMARY KEY,
  `title` varchar(255),
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now()),
  `user_id` int UNIQUE NOT NULL,
  `question_category_id` int
);

CREATE TABLE `relationships` (
  `id` int PRIMARY KEY,
  `follower_id` int,
  `followed_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `replies` (
  `id` int PRIMARY KEY,
  `title` varchar(255),
  `article_id` int,
  `user_id` int UNIQUE NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `user_article_relationships` (
  `id` int PRIMARY KEY,
  `follower_id` int,
  `followed_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `user_question_relations` (
  `id` int PRIMARY KEY,
  `follower_id` int,
  `followed_id` int,
  `created_at` timestamptz NOT NULL DEFAULT (now()),
  `updated_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX `index_likes_on_article_id` ON `likes` (`article_id`);

CREATE INDEX `index_likes_on_user_id` ON `likes` (`user_id`);

CREATE INDEX `index_votes_on_answer_id` ON `votes` (`answer_id`);

CREATE INDEX `index_votes_on_user_id` ON `votes` (`user_id`);

CREATE INDEX `index_relationships_on_followed_id` ON `relationships` (`followed_id`);

CREATE UNIQUE INDEX `index_relationships_on_follower_id_and_followed_id` ON `relationships` (`follower_id`, `followed_id`);

CREATE INDEX `index_relationships_on_follower_id` ON `relationships` (`follower_id`);

CREATE INDEX `index_user_article_relationships_on_followed_id` ON `user_article_relationships` (`followed_id`);

CREATE UNIQUE INDEX `index_user_article_relationships_on_follower_id_and_followed_id` ON `user_article_relationships` (`follower_id`, `followed_id`);

CREATE INDEX `index_user_article_relationships_on_follower_id` ON `user_article_relationships` (`follower_id`);

CREATE INDEX `index_user_question_relations_on_followed_id` ON `user_question_relations` (`followed_id`);

CREATE UNIQUE INDEX `index_user_question_relations_on_follower_id_and_followed_id` ON `user_question_relations` (`follower_id`, `followed_id`);

CREATE INDEX `index_user_question_relations_on_follower_id` ON `user_question_relations` (`follower_id`);

ALTER TABLE `answers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `votes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `replies` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
