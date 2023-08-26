CREATE TABLE users
(
    user_id serial primary key
);

CREATE TABLE segments
(
    segments_id serial primary key,
    segments_name varchar(255) not null
);

CREATE TABLE segments_user
(
    user_id int REFERENCES users(user_id),
    segments_id int REFERENCES segments(segments_id),
    time timestamp not null,

    CONSTRAINT user_segments_pkey PRIMARY KEY (user_id, segments_id)
);