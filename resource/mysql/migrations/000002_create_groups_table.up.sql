CREATE TABLE IF NOT EXISTS attendance_groups(
   id VARCHAR (20) NOT NULL,
   name VARCHAR (20),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE attendance_groups ADD PRIMARY KEY (id);
