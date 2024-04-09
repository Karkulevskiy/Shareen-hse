--creating extension for unique identifiers, check github.com/google/uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;

--creating table lobby
CREATE TABLE lobbies(
    id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    lobby_url varchar(255) NOT NULL,
    video_url varchar(255),
    created_at varchar(255) NOT NULL ,
    changed_at varchar(255) ,
    PRIMARY KEY (id)
);

--creating table user
CREATE TABLE users(
    id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    lobby_id uuid,
    name VARCHAR(20) NOT NULL,
    FOREIGN KEY (lobby_id) REFERENCES lobbies (id),
    PRIMARY KEY(id)
);

--creating lobby_user
CREATE TABLE lobbies_users(
    id BIGSERIAL PRIMARY KEY,
    users_id uuid REFERENCES users ON DELETE SET NULL, --Проверить поведение БД, при удалении пользователя и лобби
    lobbies_id uuid REFERENCES lobbies ON DELETE CASCADE
)

--TODO Сделать таблицу для чата
