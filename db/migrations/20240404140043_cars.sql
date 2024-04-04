-- +goose Up
-- +goose StatementBegin
CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name CHAR(100) NOT NULL DEFAULT '',
    surname CHAR(100) NOT NULL DEFAULT '',
    patronymic CHAR(100) NOT NULL DEFAULT ''
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    reqnum CHAR(9) NOT NULL DEFAULT '',
    mark CHAR(20) NOT NULL DEFAULT '',
    model CHAR(100) NOT NULL DEFAULT '',
    year INTEGER NOT NULL DEFAULT 0,
    owner_id INTEGER NOT NULL REFERENCES owners(id)
);

INSERT INTO owners (name, surname, patronymic) VALUES ('Иван', 'Иванов', 'Иванович');
INSERT INTO owners (name, surname, patronymic) VALUES ('Петр', 'Петров', 'Петрович');
INSERT INTO owners (name, surname, patronymic) VALUES ('Мария', 'Сидорова', 'Ивановна');

INSERT INTO cars (reqnum, mark, model, year, owner_id) VALUES ('X123XX150', 'Lada', 'Vesta', 2002, 1);
INSERT INTO cars (reqnum, mark, model, year, owner_id) VALUES ('A456BC789', 'Toyota', 'Camry', 2010, 2);
INSERT INTO cars (reqnum, mark, model, year, owner_id) VALUES ('E777OP111', 'BMW', 'X5', 2015, 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
DROP TABLE IF EXISTS owners;
-- +goose StatementEnd
