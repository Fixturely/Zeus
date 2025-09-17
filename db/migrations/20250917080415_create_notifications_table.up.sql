create table notifications (
    id serial primary key,
    subscription_id int references subscriptions(id) not null,
    user_id int references account(id) not null,
    sent_at timestamp null,
    status varchar(255) check (status in ('pending', 'sent', 'failed')) not null default 'pending',
    date_created timestamp not null default current_timestamp,
    date_updated timestamp not null default current_timestamp,
    date_deleted timestamp null
);

create index idx_notifications_user_id on notifications(user_id);
create index idx_notifications_subscription_id on notifications(subscription_id);
create index idx_notifications_sent_at on notifications(sent_at);
create index idx_notifications_user_id_sent_at on notifications(user_id, sent_at);
