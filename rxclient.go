package rxcleint

import (
	"fmt"
	"github.com/NubeIO/unixclient"
)

type RxClient interface {
	IPValidation(string) *ValidationResponse
}

type rxClient struct {
	rx *unixclient.UnixClient
}

func New() (RxClient, error) {
	client, err := unixclient.NewUnixClient("/tmp/rx-server.sock")
	return &rxClient{
		rx: client,
	}, err

}

func (c *rxClient) IPValidation(ip string) *ValidationResponse {
	validationResponse := &ValidationResponse{}
	_, err := c.rx.Send("validation/ip", ip, 5, &validationResponse, "any")
	return errorResp(validationResponse, err)
}

func errorResp(resp *ValidationResponse, err error) *ValidationResponse {
	if resp == nil {
		resp = &ValidationResponse{}
		resp.ErrorMessage = fmt.Sprintf("reponse was empty")
		resp.IsError = true
		return resp
	}
	if err != nil {
		resp.ErrorMessage = fmt.Sprintf("error sending request: %v", err)
		return resp
	}
	return resp
}

type ValidationResponse struct {
	OkMessage    string `json:"okMessage"`
	Code         string `json:"code"`
	Advice       string `json:"advice,omitempty"` // eg; an exiting entry already contains filed ""
	ErrorMessage string `json:"error,omitempty"`
	IsError      bool   `json:"isError"`
}
