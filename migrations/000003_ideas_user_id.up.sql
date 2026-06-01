ALTER TABLE tracker.ideas
ADD COLUMN user_id INT NOT NULL;

ALTER TABLE tracker.ideas
ADD CONSTRAINT fk_tracker_ideas_user_id 
FOREIGN KEY (user_id)
REFERENCES tracker.users (id)
ON DELETE CASCADE;