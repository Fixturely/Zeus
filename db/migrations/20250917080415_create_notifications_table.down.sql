drop index if exists idx_notifications_user_id;
drop index if exists idx_notifications_subscription_id;
drop index if exists idx_notifications_sent_at;
drop index if exists idx_notifications_user_id_sent_at;

drop table if exists notifications;