INSERT INTO password_reset (
    id,
    editor_id,
    token,
    expire_time
) VALUES (
    @id,
    @editor_id,
    @token,
    @expire_time
);