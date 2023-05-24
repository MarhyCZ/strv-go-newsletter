SELECT (
        id,
        editor_id,
        name,
        description,
        created_at,
        updated_at
           )
FROM newsletter
WHERE id = @id;