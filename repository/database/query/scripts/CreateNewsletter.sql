INSERT INTO newsletter (
    id,
    editor_id,
    name,
    description,
    created_at,
    updated_at
)
VALUES
    (@id, @editor_id, @name, @description, @created_at, @updated_at);
