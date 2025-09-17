create table subscriptions (
    id serial primary key,
    sport_id int not null,
    team_id int null,
    user_id int references account(id) not null,
    is_active boolean not null default true,
    billing_status varchar(255),
    subscription_type varchar(255),
    date_created timestamp not null default current_timestamp,
    date_updated timestamp not null default current_timestamp,
    date_deleted timestamp null
);


create index idx_subscriptions_user_id on subscriptions(user_id);
create index idx_subscriptions_user_id_team_id on subscriptions(user_id, team_id);
create index idx_subscriptions_user_id_sport_id on subscriptions(user_id, sport_id);