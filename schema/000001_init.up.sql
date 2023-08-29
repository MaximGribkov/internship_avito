CREATE TABLE users
(
    user_id serial primary key
);

CREATE TABLE segments
(
    segments_name varchar(255) primary key,
    time_create timestamp
);

CREATE TABLE segments_user
(
    user_id int REFERENCES users(user_id) on delete cascade,
    segments_name varchar(255) REFERENCES segments(segments_name) on delete cascade,
    time_create timestamp,

    CONSTRAINT user_segments_pkey PRIMARY KEY (user_id, segments_name)
);