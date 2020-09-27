ALTER TABLE users DROP INDEX user_index;
ALTER TABLE users ADD COLUMN attendance_group_users VARCHAR (20) NOT NULL AFTER login_id;
ALTER TABLE users ADD FOREIGN KEY fk_attendance_group_id(attendance_group_users) REFERENCES attendance_groups(id);
ALTER TABLE users ADD UNIQUE belong_index(attendance_group_users, login_id);
