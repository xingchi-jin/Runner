package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/harness/runner/task/internalvars"
	"github.com/sirupsen/logrus"
)

// cgi request body is strong typed to support imports.
type requestBody struct {
	Spec    []byte            `json:"spec"`
	Imports map[string]map[string]string `json:"imports,omitempty"`
}

func readBody(req *http.Request) (*requestBody, error) {
	in, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.WithError(err).Error("Read request body failed")
		return nil, err
	}
	reqBody := &requestBody{}
	if err = json.Unmarshal(in, reqBody); err != nil {
		logrus.WithError(err).Error("Unmarshal request body failed")
		return nil, err
	}
	return reqBody, nil
}

func ReadTaskSpec(req *http.Request, a any) error {
	body, err := readBody(req)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body.Spec, a); err != nil {
		return err
	} else {
		return nil
	}
}

func ReadImports(req *http.Request) (map[string]map[string]string, error) {
	if body, err := readBody(req); err != nil {
		return nil, err
	} else {
		return body.Imports, err
	}
}
func NewRequest(ctx context.Context, spec interface{}, imports *internalvars.Imports) (*http.Request, error) {
	specBytes, err := json.Marshal(spec)
	if err != nil {
		logrus.WithError(err).Error("Marshal task spec data failed")
		return nil, err
	}
	body := requestBody{
		Spec:    specBytes,
		Imports: imports.GetAll(),
	}
	payload, err := json.Marshal(body)
	if err != nil {
		logrus.WithError(err).Error("Marshal task request data failed")
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "/", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return req, nil
}
