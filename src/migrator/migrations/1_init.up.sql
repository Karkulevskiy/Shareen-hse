--creating lobbies
CREATE TABLE IF NOT EXISTS lobbies
(
    id SERIAL PRIMARY KEY,
    lobby_url varchar(10) UNIQUE NOT NULL,
    video_url varchar(500) DEFAULT NULL,
    pause BOOLEAN DEFAULT FALSE,
    timing VARCHAR(8)
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

CREATE TABLE IF NOT EXISTS chats
(
    id SERIAL PRIMARY KEY,
    user_login VARCHAR(20) REFERENCES users (login) ON DELETE CASCADE,
    lobby_id SERIAL REFERENCES lobbies ON DELETE CASCADE,
    time TIME,
    message VARCHAR(1000)
);

CREATE TABLE IF NOT EXISTS users_secrets
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) REFERENCES users (login) ON DELETE CASCADE, 
    pass_hash BYTEA NOT NULL
);