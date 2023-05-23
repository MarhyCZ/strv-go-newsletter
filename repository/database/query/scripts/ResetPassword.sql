UPDATE editor
SET password = @password,
    updated_at = @updated_at
WHERE id = @id;