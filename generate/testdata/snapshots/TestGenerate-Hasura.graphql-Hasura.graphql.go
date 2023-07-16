// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

type GetPokemonBoolExp struct {
	And   []*GetPokemonBoolExp `json:"_and"`
	Not   *GetPokemonBoolExp   `json:"_not"`
	Or    []*GetPokemonBoolExp `json:"_or"`
	Level *IntComparisonExp    `json:"level"`
}

// GetAnd returns GetPokemonBoolExp.And, and is useful for accessing the field via an interface.
func (v *GetPokemonBoolExp) GetAnd() []*GetPokemonBoolExp { return v.And }

// GetNot returns GetPokemonBoolExp.Not, and is useful for accessing the field via an interface.
func (v *GetPokemonBoolExp) GetNot() *GetPokemonBoolExp { return v.Not }

// GetOr returns GetPokemonBoolExp.Or, and is useful for accessing the field via an interface.
func (v *GetPokemonBoolExp) GetOr() []*GetPokemonBoolExp { return v.Or }

// GetLevel returns GetPokemonBoolExp.Level, and is useful for accessing the field via an interface.
func (v *GetPokemonBoolExp) GetLevel() *IntComparisonExp { return v.Level }

// GetPokemonResponse is returned by GetPokemon on success.
type GetPokemonResponse struct {
	GetPokemon []*testutil.Pokemon `json:"getPokemon"`
}

// GetGetPokemon returns GetPokemonResponse.GetPokemon, and is useful for accessing the field via an interface.
func (v *GetPokemonResponse) GetGetPokemon() []*testutil.Pokemon { return v.GetPokemon }

type IntComparisonExp struct {
	Eq     *int   `json:"_eq"`
	Gt     *int   `json:"_gt"`
	Gte    *int   `json:"_gte"`
	In     []*int `json:"_in"`
	IsNull *bool  `json:"_isNull"`
	Lt     *int   `json:"_lt"`
	Lte    *int   `json:"_lte"`
	Neq    *int   `json:"_neq"`
	Nin    []*int `json:"_nin"`
}

// GetEq returns IntComparisonExp.Eq, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetEq() *int { return v.Eq }

// GetGt returns IntComparisonExp.Gt, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetGt() *int { return v.Gt }

// GetGte returns IntComparisonExp.Gte, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetGte() *int { return v.Gte }

// GetIn returns IntComparisonExp.In, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetIn() []*int { return v.In }

// GetIsNull returns IntComparisonExp.IsNull, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetIsNull() *bool { return v.IsNull }

// GetLt returns IntComparisonExp.Lt, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetLt() *int { return v.Lt }

// GetLte returns IntComparisonExp.Lte, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetLte() *int { return v.Lte }

// GetNeq returns IntComparisonExp.Neq, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetNeq() *int { return v.Neq }

// GetNin returns IntComparisonExp.Nin, and is useful for accessing the field via an interface.
func (v *IntComparisonExp) GetNin() []*int { return v.Nin }

// __GetPokemonInput is used internally by genqlient
type __GetPokemonInput struct {
	Where *GetPokemonBoolExp `json:"where"`
}

// GetWhere returns __GetPokemonInput.Where, and is useful for accessing the field via an interface.
func (v *__GetPokemonInput) GetWhere() *GetPokemonBoolExp { return v.Where }

// The query or mutation executed by GetPokemon.
const GetPokemon_Operation = `
# @genqlient(pointer: true)
query GetPokemon ($where: getPokemonBoolExp!) {
	getPokemon(where: $where) {
		species
		level
	}
}
`

func GetPokemon(
	client graphql.Client,
	where *GetPokemonBoolExp,
) (*GetPokemonResponse, error) {
	req := &graphql.Request{
		OpName: "GetPokemon",
		Query:  GetPokemon_Operation,
		Variables: &__GetPokemonInput{
			Where: where,
		},
	}
	var err error

	var data GetPokemonResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

