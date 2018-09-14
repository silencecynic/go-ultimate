package domain

type User struct {
	
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Age  int         `json:"age"`
	Mark int         `json:"mark"`
	Attr interface{} `json:"attr"`
	
}
