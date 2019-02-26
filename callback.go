package uiza

type CallbackCreateParams struct {
	Params      `form:"*"`
	Url         *string     `form:"url"`
	Method      *HttpMethod `form:"method"`
	JsonData    *string     `form:"jsonData"`
	HeadersData *string     `form:"headersData"`
}

type CallbackIDResponse struct {
	Data *CallbackIDData `json:"data"`
}

type CallbackIDData struct {
	ID string `json:"id"`
}

type CallbackIDParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type CallbackResponse struct {
	Data *CallbackData `json:"data"`
}

type CallbackData struct {
	ID          string      `json:"id"`
	Url         string      `json:"url"`
	Method      HttpMethod  `json:"method"`
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
	Method      *HttpMethod  `form:"method"`
	JsonData    *JsonData    `form:"jsonData"`
	HeadersData *HeadersData `form:"headersData"`
}

type JsonData struct {
	Text string `json:"text"`
}

type HeadersData struct {
}
