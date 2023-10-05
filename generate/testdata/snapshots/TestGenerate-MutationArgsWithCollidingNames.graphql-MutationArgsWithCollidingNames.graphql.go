// Code generated by github.com/codykaup/genqlient, DO NOT EDIT.

package test

import (
	"github.com/codykaup/genqlient/graphql"
	"github.com/codykaup/genqlient/internal/testutil"
)

// MutationArgsWithCollidingNamesResponse is returned by MutationArgsWithCollidingNames on success.
type MutationArgsWithCollidingNamesResponse struct {
	UpdateUser MutationArgsWithCollidingNamesUpdateUser `json:"updateUser"`
}

// GetUpdateUser returns MutationArgsWithCollidingNamesResponse.UpdateUser, and is useful for accessing the field via an interface.
func (v *MutationArgsWithCollidingNamesResponse) GetUpdateUser() MutationArgsWithCollidingNamesUpdateUser {
	return v.UpdateUser
}

// MutationArgsWithCollidingNamesUpdateUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type MutationArgsWithCollidingNamesUpdateUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns MutationArgsWithCollidingNamesUpdateUser.Id, and is useful for accessing the field via an interface.
func (v *MutationArgsWithCollidingNamesUpdateUser) GetId() testutil.ID { return v.Id }

// __MutationArgsWithCollidingNamesInput is used internally by genqlient
type __MutationArgsWithCollidingNamesInput struct {
	Data   string `json:"data"`
	Req    int    `json:"req"`
	Resp   int    `json:"resp"`
	Client string `json:"client"`
}

// GetData returns __MutationArgsWithCollidingNamesInput.Data, and is useful for accessing the field via an interface.
func (v *__MutationArgsWithCollidingNamesInput) GetData() string { return v.Data }

// GetReq returns __MutationArgsWithCollidingNamesInput.Req, and is useful for accessing the field via an interface.
func (v *__MutationArgsWithCollidingNamesInput) GetReq() int { return v.Req }

// GetResp returns __MutationArgsWithCollidingNamesInput.Resp, and is useful for accessing the field via an interface.
func (v *__MutationArgsWithCollidingNamesInput) GetResp() int { return v.Resp }

// GetClient returns __MutationArgsWithCollidingNamesInput.Client, and is useful for accessing the field via an interface.
func (v *__MutationArgsWithCollidingNamesInput) GetClient() string { return v.Client }

// The query or mutation executed by MutationArgsWithCollidingNames.
const MutationArgsWithCollidingNames_Operation = `
mutation MutationArgsWithCollidingNames ($data: String!, $req: Int, $resp: Int, $client: String) {
	updateUser(data: $data, req: $req, resp: $resp, client: $client) {
		id
	}
}
`

func MutationArgsWithCollidingNames(
	client_ graphql.Client,
	data string,
	req int,
	resp int,
	client string,
) (*MutationArgsWithCollidingNamesResponse, error) {
	req_ := &graphql.Request{
		OpName: "MutationArgsWithCollidingNames",
		Query:  MutationArgsWithCollidingNames_Operation,
		Variables: &__MutationArgsWithCollidingNamesInput{
			Data:   data,
			Req:    req,
			Resp:   resp,
			Client: client,
		},
	}
	var err_ error

	var data_ MutationArgsWithCollidingNamesResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

