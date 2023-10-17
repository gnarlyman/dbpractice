UPDATE users
SET updated_at = CURRENT_TIMESTAMP
WHERE updated_at IS NULL;
