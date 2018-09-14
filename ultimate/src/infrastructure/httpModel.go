package infrastructure

type HttpModel struct {

	Response interface{} `json:"response"`
	Messages string `json:"messages"`
	Status int `json:"status"`
	Content string `json:"content"`

}

func (h *HttpModel) SetResponse(response interface{}) *HttpModel  {
	h.Response = response
	return h
}

func (h *HttpModel) SetMessages(messages string) *HttpModel  {
	h.Messages = messages
	return h
}

func (h *HttpModel) SetStatus(status int) *HttpModel  {
	h.Status = status
	return h
}

func (h *HttpModel) SetType(content string) *HttpModel  {
	h.Content = content
	return h
}

func (h *HttpModel) GetResponse() interface{}  {
	return h.Response
}

func (h *HttpModel) GetMessage() string  {
	return h.Messages
}

func (h *HttpModel) GetStatus() int  {
	return h.Status
}

func (h *HttpModel) GetContentType() string  {
	return h.Content
}

