-- +goose Up
-- SQL in this section is executed when the migration is applied.

alter table event add column column1 char(255) not null;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

alter table event drop column column1;
