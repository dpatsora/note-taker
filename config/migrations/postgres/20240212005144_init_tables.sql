-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notes (
id                   UUID UNIQUE NOT NULL,
title       VARCHAR(255),
description     TEXT,
timestamp  TIMESTAMP(0) DEFAULT NOW(),
PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
-- +goose StatementEnd
