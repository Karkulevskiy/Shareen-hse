CREATE TABLE IF NOT EXISTS lobbies
(
    id SERIAL PRIMARY KEY,
    lobby_url varchar(255) UNIQUE NOT NULL,
    video_url varchar(255)
);

--creating table user
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL
);

--creating lobby_user
CREATE TABLE IF NOT EXISTS lobbies_users
(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES users ON DELETE SET NULL, --Проверить поведение БД, при удалении пользователя и лобби
    lobby_id SERIAL REFERENCES lobbies ON DELETE CASCADE,
    UNIQUE(user_id, lobby_id)
);

--creating chats for lobbies
CREATE TABLE IF NOT EXISTS chats 
(
    id SERIAL PRIMARY KEY,
    lobby_id SERIAL REFERENCES lobbies ON DELETE CASCADE
);