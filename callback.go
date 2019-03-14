package uiza

type CallbackCreateParams struct {
	Params      `form:"*"`
	Url         *string      `form:"url"`
	Method      *HTTPMethod  `form:"method"`
	JsonData    *string      `form:"jsonData"`
	HeadersData *HeadersData `form:"headersData"`
}

type CallbackIDResponse struct {
	Data *CallbackIDData `json:"data"`
}

type CallbackIDData struct {
	ID string `json:"id"`
}

type CallbackIDParams struct {
	Params `form:"*"`
}

type CallbackResponse struct {
	Data *CallbackData `json:"data"`
}

type CallbackData struct {
	ID          string      `json:"id"`
	Url         string      `json:"url"`
	Method      HTTPMethod  `json:"method"`
	JsonData    JsonData    `json:"jsonData"`
	HeadersData HeadersData `json:"headersData"`
	Status      int64       `json:"status"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
}

type CallbackUpdateParams struct {
	Params      `form:"*"`
	ID          *string      `form:"id"`
	Url         *string      `form:"url"`
	Method      *HTTPMethod  `form:"method"`
	JsonData    *JsonData    `form:"jsonData"`
	HeadersData *HeadersData `form:"headersData"`
}

type JsonData struct {
	Text string `json:"text"`
}

type HeadersData struct {
}
