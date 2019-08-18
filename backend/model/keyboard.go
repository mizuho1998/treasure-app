package model

type KeyboardID struct {
	ID int `db:"id" json:"id"`
}

type SerchKeyboard struct {
	IDs    []KeyboardID `db:"id" json:"ids"`
	Answer string       `json:"answer"`
}

type RequestSerchKeyboard struct {
	IDs    []KeyboardID `db:"id" json:"ids"`
	QID    int          `json:"qid"`
	Answer string       `json:"answer"`
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
	Pricestring string `db:"price" json:"price"`
	URL         string `db:"url" json:"url"`
	ImageURL    string `db:"image_url" json:"image_url"`
}
