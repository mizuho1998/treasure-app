-- +goose Up
INSERT INTO keyboards(name, creater_name, price, url, image_url, split, led, key_num, matrix, key_profile, color)
VALUES
("Gingham", "", 8000, "https://yushakobo.jp/shop/gingham/", "https://yushakobo.jp/wp-content/uploads/2019/08/A0100GH-3.jpg", 0, 0, 60, "RS", "high", "black"),
("UT47.2", "", 7500, "https://yushakobo.jp/shop/ut472/", "https://yushakobo.jp/wp-content/uploads/2019/06/A0100UT.jpg", 0, 1, 47, "RS", "high", "purple"),
("UT47.2", "", 7500, "https://yushakobo.jp/shop/ut472/", "https://yushakobo.jp/wp-content/uploads/2019/06/A0100UT.jpg", 0, 1, 47, "RS", "high", "black"),
("NumAtreus", "", 11000, "https://yushakobo.jp/shop/numatreus-kit/", "https://yushakobo.jp/wp-content/uploads/2019/06/A0100NA-e1561107206291.jpg", 0, 0, 42, "CS", "high", "black"),
("Nomu30", "", 8000, "https://yushakobo.jp/shop/nomu30kit/", "https://yushakobo.jp/wp-content/uploads/2019/05/photo-2-e1558743129481.jpg", 0, 0, 31, "RS", "high", "black"),
("Lunar B-stock", "", 49000, "https://yushakobo.jp/shop/lunar/", "https://yushakobo.jp/wp-content/uploads/2019/05/lunar_main-e1557987621397.jpg", 0, 0, 70, "RS", "high", "silver"),
("Lunar B-stock", "", 49000, "https://yushakobo.jp/shop/lunar/", "https://yushakobo.jp/wp-content/uploads/2019/05/lunar_main-e1557987621397.jpg", 1, 1, 70, "RS", "high", "grey"),
("Lunar B-stock", "", 49000, "https://yushakobo.jp/shop/lunar/", "https://yushakobo.jp/wp-content/uploads/2019/05/lunar_main-e1557987621397.jpg", 1, 1, 70, "RS", "high", "black"),
("Plaid", "", 7000, "https://yushakobo.jp/shop/plaid/", "https://yushakobo.jp/wp-content/uploads/2019/04/plaid.jpg", 0, 0, 48, "RS", "high", "white"),
("Fortitude60", "", 13000, "https://yushakobo.jp/shop/fortitude60/", "https://yushakobo.jp/wp-content/uploads/2019/03/otakublack.jpg", 1, 1, 60, "CS", "high", "black"),
("Fortitude60", "", 13000, "https://yushakobo.jp/shop/fortitude60/", "https://yushakobo.jp/wp-content/uploads/2019/03/otakublack.jpg", 1, 1, 60, "CS", "high", "white"),
("Corne Cherry", "", 12800, "https://yushakobo.jp/shop/corne-cherry/", "https://yushakobo.jp/wp-content/uploads/2019/03/CorneCherry.jpg", 1, 1, 42, "CS", "high", "black"),
("Lily58 Pro", "", 14800, "https://yushakobo.jp/shop/lily58-pro/", "https://yushakobo.jp/wp-content/uploads/2019/03/White_MX1.jpg", 1, 1, 58, "CS", "high", "white"),
("Lily58 Pro", "", 14800, "https://yushakobo.jp/shop/lily58-pro/", "https://yushakobo.jp/wp-content/uploads/2019/03/White_MX1.jpg", 1, 1, 58, "CS", "low", "white"),
("Ergo42 Towel", "", 10500, "https://yushakobo.jp/shop/ergo42-towel/", "https://yushakobo.jp/wp-content/uploads/2019/03/abe21d7d-7acb-4a30-89eb-ae4cfb6f6998_base_resized.jpg", 1, 0, 56, "K", "high", "black"),
("Ergo42 Towel", "", 10500, "https://yushakobo.jp/shop/ergo42-towel/", "https://yushakobo.jp/wp-content/uploads/2019/03/abe21d7d-7acb-4a30-89eb-ae4cfb6f6998_base_resized.jpg", 1, 0, 56, "K", "high", "clear"),
("ErgoDash mini", "", 12000, "https://yushakobo.jp/shop/ergodash-mini/", "https://yushakobo.jp/wp-content/uploads/2019/02/155e315b-1b2a-4618-a9c3-d459cbb83053_base_resized.jpg", 1, 1, 56, "CS", "high", "black"),
("ErgoDash", "", 13000, "https://yushakobo.jp/shop/ergodash/", "https://yushakobo.jp/wp-content/uploads/2019/02/4f25af9c-da88-40b7-aebc-5663cb36cff6_base_resized.jpg", 1, 1, 56, "CS", "high", "black"),
("[GB] UT47.2", "", 6400, "https://yushakobo.jp/shop/gb-ut47-2/", "https://yushakobo.jp/wp-content/uploads/2019/02/ty3AYac.jpg", 0, 0, 47, "RS", "high", "purple"),
("[GB] UT47.2", "", 6400, "https://yushakobo.jp/shop/gb-ut47-2/", "https://yushakobo.jp/wp-content/uploads/2019/02/ty3AYac.jpg", 0, 0, 47, "RS", "high", "black"),
("Helix", "", 11600, "https://yushakobo.jp/shop/helix-keyboard-kit/", "https://yushakobo.jp/wp-content/uploads/2018/05/YKB0002-1.jpg", 1, 1, 64, "K", "low", "black"),
("Helix", "", 11600, "https://yushakobo.jp/shop/helix-keyboard-kit/", "https://yushakobo.jp/wp-content/uploads/2018/05/YKB0002-1.jpg", 1, 1, 64, "K", "high", "black"),
("Helix", "", 11600, "https://yushakobo.jp/shop/helix-keyboard-kit/", "https://yushakobo.jp/wp-content/uploads/2018/05/YKB0002-1.jpg", 1, 1, 50, "K", "low", "black"),
("Helix", "", 11600, "https://yushakobo.jp/shop/helix-keyboard-kit/", "https://yushakobo.jp/wp-content/uploads/2018/05/YKB0002-1.jpg", 1, 1, 50, "K", "high", "black"),
("HelixPico", "", 8200, "https://yushakobo.jp/shop/helixpico/", "https://yushakobo.jp/wp-content/uploads/2018/07/HelixPico_keikou.jpg", 1, 1, 50, "K", "low", "red"),
("HelixPico", "", 8200, "https://yushakobo.jp/shop/helixpico/", "https://yushakobo.jp/wp-content/uploads/2018/07/HelixPico_keikou.jpg", 1, 1, 50, "K", "low", "yellow"),
("HelixPico", "", 8200, "https://yushakobo.jp/shop/helixpico/", "https://yushakobo.jp/wp-content/uploads/2018/07/HelixPico_keikou.jpg", 1, 1, 50, "K", "low", "green");


INSERT INTO questions(id, question, answer1, answer2, answer3, answer1_val, answer2_val, answer3_val) 
VALUES
(1, "左右分離が良い?", "yes", "no", NULL,"1","0", NULL),
(2, "LEDで光らせたい?", "yes", "no", NULL,"1","0", NULL),
(3, "キーの配列どれがいい?", "Row-Staggered", "Column-Staggered", "格子","RS","CS", "K"),
(4, "キーの数は多いほうがいい?", "yes", "no", NULL,"1","0", NULL),
(5, "キーストロークは浅いほうが好き?", "yes", "no", NULL,"1","0", NULL)
;

-- +goose Down
DELETE FROM keyboards;
DELETE FROM questions;
