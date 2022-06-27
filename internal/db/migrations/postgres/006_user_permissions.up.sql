alter table user_countries
    drop column permission;

alter table user_countries
    add column read boolean not null default false;

alter table user_countries
    add column write boolean not null default false;

alter table user_countries
    add column admin boolean not null default false;