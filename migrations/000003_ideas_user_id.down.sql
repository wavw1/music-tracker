ALTER TABLE tracker.ideas
DROP CONSTRAINT fk_tracker_ideas_user_id;

ALTER TABLE tracker.ideas
DROP COLUMN user_id;