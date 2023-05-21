UPDATE newsletter
SET name = @name,
    description = @description,
    updated_at = @updated_at
WHERE id = @id;
