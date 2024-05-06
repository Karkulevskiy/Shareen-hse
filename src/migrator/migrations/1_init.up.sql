--creating lobbies
CREATE TABLE IF NOT EXISTS lobbies
(
    id SERIAL PRIMARY KEY,
    lobby_url varchar(255) UNIQUE NOT NULL,
    video_url varchar(500)
);
--creating index for lobby_url
CREATE INDEX IF NOT EXISTS lobby_url_idx ON lobbies (lobby_url);

--creating table user
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(20) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS user_login_idx ON users (login);

--creating chats for lobbies
CREATE TABLE IF NOT EXISTS chats 
(
    id SERIAL PRIMARY KEY,
    lobby_url VARCHAR(255) REFERENCES lobbies (lobby_url) ON DELETE CASCADE
);
--creating lobby_user
CREATE TABLE IF NOT EXISTS lobbies_users
(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES users ON DELETE CASCADE, --Проверить поведение БД, при удалении пользователя и лобби
    lobby_url VARCHAR(255) REFERENCES lobbies (lobby_url) ON DELETE CASCADE,
    UNIQUE(user_id, lobby_url)
);

CREATE TABLE IF NOT EXISTS users_secrets
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) REFERENCES users (login) ON DELETE CASCADE, 
    pass_hash BYTEA NOT NULL
);