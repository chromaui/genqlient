// Code generated by github.com/codykaup/genqlient, DO NOT EDIT.

package test

import (
	"github.com/codykaup/genqlient/graphql"
	"github.com/codykaup/genqlient/internal/testutil"
)

// SimpleQueryResponse is returned by SimpleQuery on success.
type SimpleQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryUser `json:"user"`
}

// GetUser returns SimpleQueryResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryResponse) GetUser() SimpleQueryUser { return v.User }

// SimpleQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns SimpleQueryUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryUser) GetId() testutil.ID { return v.Id }

// The query or mutation executed by SimpleQuery.
const SimpleQuery_Operation = `
query SimpleQuery {
	user {
		id
	}
}
`

func SimpleQuery(
	client_ graphql.Client,
) (*SimpleQueryResponse, error) {
	req_ := &graphql.Request{
		OpName: "SimpleQuery",
		Query:  SimpleQuery_Operation,
	}
	var err_ error

	var data_ SimpleQueryResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

