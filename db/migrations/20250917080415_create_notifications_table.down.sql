drop index if exists idx_notifications_user_id on notifications;
drop index if exists idx_notifications_subscription_id on notifications;
drop index if exists idx_notifications_sent_at on notifications;
drop index if exists idx_notifications_user_id_sent_at on notifications;

drop table if exists notifications;