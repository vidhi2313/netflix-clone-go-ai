
CREATE TABLE users (
id SERIAL PRIMARY KEY,
name TEXT,
email TEXT,
password TEXT
);

CREATE TABLE movies (
id SERIAL PRIMARY KEY,
title TEXT,
genre TEXT,
description TEXT,
video_url TEXT
);
