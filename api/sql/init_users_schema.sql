create table users (
    id int auto_increment not null,
    username varchar(24) not null,
    password_hash varchar(24) not null,
    email varchar(256) not null,
    primary key(id)
);