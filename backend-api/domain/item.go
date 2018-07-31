package domain

// Item is 商品
type Item struct {
	ID          int64
	Name        string
	Description string
	Amount      int64
}

// Items are 商品のリスト
type Items []Item
