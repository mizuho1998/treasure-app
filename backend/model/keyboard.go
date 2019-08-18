package model

type RequestSerchKeyboard struct {
	IDs        []int  `db:"id" json:"ids"`
	IsSplit    int64  `db:"is_split" json:"is_split"`
	Color      string `db:"color" json:"color"`
	KeyNum     int64  `db:"key_num" json:"key_num"`
	Matrix     string `db:"matrix" json:"matrix"`
	KeyProfile string `db:"key_profile" json:"key_profile"`
}

type RequestKeyboard struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	CreaterName string `db:"creater_name" json:"creater_name"`
	URL         string `db:"url" json:"url"`
}

type Keyboard struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	CreaterName string `db:"creater_name" json:"creater_name"`
	URL         string `db:"url" json:"url"`
}
