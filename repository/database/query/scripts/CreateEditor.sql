INSERT INTO editor (
    id,
    password,
    email,
    created_at,
    updated_at
) VALUES
    (@id, @password, @email, @created_at, @updated_at);