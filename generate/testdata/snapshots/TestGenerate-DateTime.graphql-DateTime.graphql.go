// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"time"

	"github.com/Khan/genqlient/graphql"
)

// convertTimezoneResponse is returned by convertTimezone on success.
type ConvertTimezoneResponse struct {
	Convert time.Time `json:"convert"`
}

// GetConvert returns ConvertTimezoneResponse.Convert, and is useful for accessing the field via an interface.
func (v *ConvertTimezoneResponse) GetConvert() time.Time { return v.Convert }

// __convertTimezoneInput is used internally by genqlient
type __convertTimezoneInput struct {
	Dt time.Time `json:"dt"`
	Tz string    `json:"tz"`
}

// GetDt returns __convertTimezoneInput.Dt, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetDt() time.Time { return v.Dt }

// GetTz returns __convertTimezoneInput.Tz, and is useful for accessing the field via an interface.
func (v *__convertTimezoneInput) GetTz() string { return v.Tz }

// The query or mutation executed by convertTimezone.
const convertTimezone_Operation = `
query convertTimezone ($dt: DateTime!, $tz: String) {
	convert(dt: $dt, tz: $tz)
}
`

func ConvertTimezoneQuery(
	client_ graphql.Client,
	dt time.Time,
	tz string,
) (*ConvertTimezoneResponse, error) {
	req_ := &graphql.Request{
		OpName: "convertTimezone",
		Query:  convertTimezone_Operation,
		Variables: &__convertTimezoneInput{
			Dt: dt,
			Tz: tz,
		},
	}
	var err_ error

	var data_ ConvertTimezoneResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

