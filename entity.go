package uiza

// EntitySpecParams defined
type EntitySpecParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

// EntitySpec defined
type EntitySpec struct {
	Data map[string]string `json:"data"`
}
