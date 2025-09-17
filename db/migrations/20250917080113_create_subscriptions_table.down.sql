Drop index if exists idx_subscriptions_user_id on subscriptions;
Drop index if exists idx_subscriptions_user_id_team_id on subscriptions;
Drop index if exists idx_subscriptions_user_id_sport_id on subscriptions;

Drop table if exists subscriptions;