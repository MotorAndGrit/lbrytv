package responses

import (
	"encoding/json"
	"net/http"

	"github.com/lbryio/lbrytv/internal/errors"
	"github.com/ybbus/jsonrpc"
)

// this is the message to show when authentication info is required but was not provided in the request
// this is NOT the message for when auth info is provided but is not correct
const AuthRequiredErrorMessage = "authentication required"

// AddJSONContentType prepares HTTP response writer for JSON content-type.
func AddJSONContentType(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json; charset=utf-8")
}

func JSONRPCSerialize(r *jsonrpc.RPCResponse) ([]byte, error) {
	var (
		b []byte
		e error
	)
	defer errors.Recover(&e)
	b, err := json.MarshalIndent(r, "", "  ")
	if e != nil {
		return b, e
	} else if err != nil {
		return b, err
	}
	return b, nil
}
