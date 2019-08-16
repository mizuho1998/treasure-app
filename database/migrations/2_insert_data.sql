-- +goose Up
INSERT INTO keyboards(name, creater_name, url) VALUES("keyboard_name", "creater_name", "http://keyboard.com");
INSERT INTO questions(question) VALUES("question!!");

-- +goose Down
DELETE FROM keyboards;
DELETE FROM questions;
