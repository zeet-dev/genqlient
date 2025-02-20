// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// getUserResponse is returned by getUser on success.
type GetUserResponse struct {
	// Lookup a user by login.
	User getUserUser `json:"user"`
}

// GetUser returns GetUserResponse.User, and is useful for accessing the field via an interface.
func (v *GetUserResponse) GetUser() getUserUser { return v.User }

// getViewerResponse is returned by getViewer on success.
type GetViewerResponse struct {
	// The currently authenticated user.
	Viewer getViewerViewerUser `json:"viewer"`
}

// GetViewer returns GetViewerResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetViewerResponse) GetViewer() getViewerViewerUser { return v.Viewer }

// __getUserInput is used internally by genqlient
type __getUserInput struct {
	Login string `json:"Login"`
}

// GetLogin returns __getUserInput.Login, and is useful for accessing the field via an interface.
func (v *__getUserInput) GetLogin() string { return v.Login }

// getUserUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getUserUser struct {
	// The user's public profile name.
	TheirName string `json:"theirName"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
}

// GetTheirName returns getUserUser.TheirName, and is useful for accessing the field via an interface.
func (v *getUserUser) GetTheirName() string { return v.TheirName }

// GetCreatedAt returns getUserUser.CreatedAt, and is useful for accessing the field via an interface.
func (v *getUserUser) GetCreatedAt() time.Time { return v.CreatedAt }

// getViewerViewerUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getViewerViewerUser struct {
	// The user's public profile name.
	MyName string `json:"MyName"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
}

// GetMyName returns getViewerViewerUser.MyName, and is useful for accessing the field via an interface.
func (v *getViewerViewerUser) GetMyName() string { return v.MyName }

// GetCreatedAt returns getViewerViewerUser.CreatedAt, and is useful for accessing the field via an interface.
func (v *getViewerViewerUser) GetCreatedAt() time.Time { return v.CreatedAt }

// The query or mutation executed by getUser.
const getUser_Operation = `
query getUser ($Login: String!) {
	user(login: $Login) {
		theirName: name
		createdAt
	}
}
`

// getUser gets the given user's name from their username.
func GetUserQuery(
	ctx_ context.Context,
	client_ graphql.Client,
	Login string,
) (*GetUserResponse, error) {
	req_ := &graphql.Request{
		OpName: "getUser",
		Query:  getUser_Operation,
		Variables: &__getUserInput{
			Login: Login,
		},
	}
	var err_ error

	var data_ GetUserResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by getViewer.
const getViewer_Operation = `
query getViewer {
	viewer {
		MyName: name
		createdAt
	}
}
`

func GetViewerQuery(
	ctx_ context.Context,
	client_ graphql.Client,
) (*GetViewerResponse, error) {
	req_ := &graphql.Request{
		OpName: "getViewer",
		Query:  getViewer_Operation,
	}
	var err_ error

	var data_ GetViewerResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
