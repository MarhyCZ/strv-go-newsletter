SELECT (
    id,
    editor_id,
    name,
    description,
    created_at,
    updated_at
) FROM newsletter
WHERE editor_id = @editor_id;