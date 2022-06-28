create table if not exists user_permissions
(
    user_id         varchar(32) not null primary key,
    is_global_admin boolean     not null default false,
    foreign key (user_id) references users (id)
);

update user_permissions set is_global_admin = true where user_id in (select id from users order by id limit 1);