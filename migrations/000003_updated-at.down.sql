DROP TRIGGER IF EXISTS users_update_updated_at ON users;
ALTER TABLE users DROP COLUMN IF EXISTS updated_at;
