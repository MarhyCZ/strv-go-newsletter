CREATE TABLE "editor"(
    id        UUID PRIMARY KEY,
    password  VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);
CREATE TABLE "newsletter"(
    id            UUID PRIMARY KEY,
    editor_id     UUID NOT NULL REFERENCES editor(id),
    name          VARCHAR(255) NOT NULL,
    description   VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);

CREATE TABLE "password_reset"(
    id          UUID PRIMARY KEY,
    editor_id   UUID NOT NULL REFERENCES editor(id),
    expire_time TIMESTAMP NOT NULL
);