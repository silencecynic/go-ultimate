package domain

type (

	User struct {
	  ID int `json:"id"`
	  Name string `json:"name"`
	  Attr interface{} `json:"attr"`
	  Lock string `json:"lock"`
	  Age int `json:"age"`
	}

)
