// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// ListOfListsOfListsResponse is returned by ListOfListsOfLists on success.
type ListOfListsOfListsResponse struct {
	ListOfListsOfLists [][][]string `json:"listOfListsOfLists"`
}

// GetListOfListsOfLists returns ListOfListsOfListsResponse.ListOfListsOfLists, and is useful for accessing the field via an interface.
func (v *ListOfListsOfListsResponse) GetListOfListsOfLists() [][][]string {
	return v.ListOfListsOfLists
}

// The query or mutation executed by ListOfListsOfLists.
const ListOfListsOfLists_Operation = `
query ListOfListsOfLists {
	listOfListsOfLists
}
`

func ListOfListsOfListsQuery(
	client_ graphql.Client,
) (*ListOfListsOfListsResponse, error) {
	req_ := &graphql.Request{
		OpName: "ListOfListsOfLists",
		Query:  ListOfListsOfLists_Operation,
	}
	var err_ error

	var data_ ListOfListsOfListsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

