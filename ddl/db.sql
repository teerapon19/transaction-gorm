create table if not exists fighter.user(
    id varchar(36) primary key,
    name varchar(255) not null
);

create table if not exists fighter.vehicle(
    id varchar(36) primary key,
    name varchar(255) not null
);

create table if not exists fighter.weapon(
    id varchar(36) primary key,
    name varchar(255) not null
);