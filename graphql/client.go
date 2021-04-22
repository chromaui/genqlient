package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Client is the interface that the generate code calls into to actually make
// requests.
//
// Unstable: This interface is likely to change before v1.0, see #19.
type Client interface {
	// MakeRequest must make a request to the client's GraphQL API.
	//
	// ctx is the context that should be used to make this request.  If context
	// is disabled in the genqlient settings, this will be set to
	// context.Background().
	//
	// query is the literal string representing the GraphQL query, e.g.
	// `query myQuery { myField }`.  variables contains the GraphQL variables
	// to be sent along with the query, or may be nil if there are none.
	// Typically, GraphQL APIs will accept a JSON payload of the form
	//	{"query": "query myQuery { ... }", "variables": {...}}`
	// but MakeRequest may use some other transport, handle extensions, or set
	// other parameters, if it wishes.
	//
	// retval is a pointer to the struct representing the query result, e.g.
	// new(myQueryResponse).  Typically, GraphQL APIs will return a JSON
	// payload of the form
	//	{"data": {...}, "errors": {...}}
	// and retval is designed so that `data` will json-unmarshal into `retval`.
	// (Errors are returned.) But again, MakeRequest may customize this.
	MakeRequest(
		ctx context.Context,
		opName string,
		query string,
		retval interface{},
		variables map[string]interface{},
	) error
}

type client struct {
	endpoint   string
	method     string
	httpClient *http.Client
}

// NewClient returns a Client which makes requests to the given endpoint,
// suitable for most users.
//
// The client makes POST requests to the given GraphQL endpoint using standard
// GraphQL HTTP-over-JSON transport.  It will use the given http client, or
// http.DefaultClient if a nil client is passed.
//
// The typical method of adding authentication headers is to wrap the client's
// Transport to add those headers.  See example/caller.go for an example.
func NewClient(endpoint string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &client{endpoint, http.MethodPost, httpClient}
}

type payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
	// OpName is only required if there are multiple queries in the document,
	// but we set it unconditionally, because that's easier.
	OpName string `json:"operationName"`
}

type response struct {
	Data   interface{}   `json:"data"`
	Errors gqlerror.List `json:"errors"`
}

func (c *client) MakeRequest(ctx context.Context, opName string, query string, retval interface{}, variables map[string]interface{}) error {
	body, err := json.Marshal(payload{
		Query:     query,
		Variables: variables,
		OpName:    opName,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		c.method,
		c.endpoint,
		bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if ctx != nil {
		req = req.WithContext(ctx)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			respBody = []byte("<unreadable>")
		}
		return fmt.Errorf("returned error %v: %s", resp.Status, respBody)
	}

	var dataAndErrors response
	dataAndErrors.Data = retval
	err = json.NewDecoder(resp.Body).Decode(&dataAndErrors)
	if err != nil {
		return err
	}

	if len(dataAndErrors.Errors) > 0 {
		return dataAndErrors.Errors
	}
	return nil
}
