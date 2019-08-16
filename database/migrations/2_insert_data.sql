-- +goose Up
INSERT INTO keyboards(name, creater_name, url, is_split, color, key_num, matrix, key_profile) VALUES("keyboard_name", "creater_name", "http://keyboard.com", 1, "black", 64, "koushi", "low"), ("keyboard_name2", "creater_name2", "http://keyboard.com", 1, "red", 48, "1/3u", "high");

INSERT INTO questions(question) VALUES("question!!");

-- +goose Down
DELETE FROM keyboards;
DELETE FROM questions;
