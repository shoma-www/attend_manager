ALTER TABLE users DROP FOREIGN KEY users_ibfk_1;
ALTER TABLE users DROP INDEX belong_index;
ALTER TABLE users DROP COLUMN attendance_group_users;
ALTER TABLE users ADD UNIQUE user_index (login_id);
