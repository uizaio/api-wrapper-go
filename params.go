package uiza

import (
	"context"
	"net/http"
	"net/url"

	"github.com/uizaio/api-wrapper-go/form"
)

//
// Public constants
//

// Contains constants for the names of parameters used for pagination in list APIs.
const (
	EndingBefore  = "ending_before"
	StartingAfter = "starting_after"
)

//
// Public types
//

// ExtraValues are extra parameters that are attached to an API request.
// They're implemented as a custom type so that they can have their own
// AppendTo implementation.
type ExtraValues struct {
	url.Values `form:"-"` // See custom AppendTo implementation
}

// AppendTo implements custom form encoding for extra parameter values.
func (v ExtraValues) AppendTo(body *form.Values, keyParts []string) {
	for k, vs := range v.Values {
		for _, v := range vs {
			body.Add(form.FormatKey(append(keyParts, k)), v)
		}
	}
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Total  uint32 `json:"total"`
	Result uint32 `json:"result"`
	Page   uint32 `json:"page"`
	Limit  uint32 `json:"limit"`
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on uiza's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-"`

	EndingBefore *string   `form:"ending_before"`
	Expand       []*string `form:"expand"`
	Limit        *int64    `form:"limit"`

	// Single specifies whether this is a single page iterator. By default,
	// listing through an iterator will automatically grab additional pages as
	// the query progresses. To change this behavior and just load a single
	// page, set this to true.
	Single bool `form:"-"` // Not an API parameter

	StartingAfter *string `form:"starting_after"`
}

func (p *ListParams) GetParams() *Params {
	return p.ToParams()
}

// ListParams is only used to build a set of parameters.
func (p *ListParams) ToParams() *Params {
	return &Params{
		Context: p.Context,
	}
}

// ListParamsContainer is a general interface for which all list parameter
// structs should comply. They achieve this by embedding a ListParams struct
// and inheriting its implementation of this interface.
type ListParamsContainer interface {
	GetListParams() *ListParams
}

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	Metadata map[string]string `form:"metadata"`
	AppID    string            `form:"appId"`
	// key or query the state of the API.
	Context context.Context `form:"-"`
	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `form:"-"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *Params) AddMetadata(key, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// GetParams returns a Params struct (itself). It exists because any structs
// that embed Params will inherit it, and thus implement the ParamsContainer
// interface.
func (p *Params) GetParams() *Params {
	return p
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

// ClientType : Type of Client.
type ClientType string

// Define Client Type
const (
	EntityClientType   ClientType = "Entity"
	StorageClientType  ClientType = "Storage"
	CategoryClientType ClientType = "Category"
	CallbackClientType ClientType = "Callback"
	LiveClientType     ClientType = "Live"
	AnalyticClientType ClientType = "Analytic"
	UserClientType     ClientType = "User"
)
