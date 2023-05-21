CREATE TABLE role
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO role (id, name)
VALUES (1, 'Admin');
INSERT INTO role (id, name)
VALUES (2, 'Teacher');
INSERT INTO role (id, name)
VALUES (3, 'Student');

CREATE TABLE user
(
    id        TEXT PRIMARY KEY,
    email     TEXT    NOT NULL,
    firstname TEXT    NOT NULL,
    lastname  TEXT    NOT NULL,
    active    INTEGER NOT NULL DEFAULT 1,
    role_id   INTEGER NOT NULL,

    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);
