SELECT
    id,
    editor_id,
    token,
    expire_time
FROM password_reset
WHERE token = @token;