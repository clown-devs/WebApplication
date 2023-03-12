package model

type Client struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Inn  string `json:"inn"` // max length - 12
}

type Contact struct {
	ID       uint64  `json:"id"`
	Fullname string  `json:"fullname"`
	Phone    *string `json:"phone"`
	Position *string `json:"position"`
	Email    *string `json:"email"`
}

