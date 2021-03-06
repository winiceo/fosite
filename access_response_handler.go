package fosite

import (
	"net/http"

	"github.com/go-errors/errors"
	"golang.org/x/net/context"
)

func (f *Fosite) NewAccessResponse(ctx context.Context, req *http.Request, requester AccessRequester) (AccessResponder, error) {
	var err error
	var tk TokenEndpointHandler

	response := NewAccessResponse()
	for _, tk = range f.TokenEndpointHandlers {
		if err = tk.PopulateTokenEndpointResponse(ctx, req, requester, response); errors.Is(err, ErrUnknownRequest) {
		} else if err != nil {
			return nil, errors.Wrap(err, 1)
		}
	}

	if response.GetAccessToken() == "" || response.GetTokenType() == "" {
		return nil, ErrServerError
	}

	return response, nil
}
