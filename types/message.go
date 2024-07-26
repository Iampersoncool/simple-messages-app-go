package types

type Message struct {
	Id      int    `json:"-"`
	Content string `json:"content"`
	Color   string `color:"color"`
}
