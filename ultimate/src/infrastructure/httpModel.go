package infrastructure

type HttpModel struct {

	Response interface{} `json:"response"`
	Messages string `json:"messages"`
	Status int `json:"status"`
	Type string `json:"type"`

}

func (h *HttpModel) Set(argc *HttpModel) {



}
