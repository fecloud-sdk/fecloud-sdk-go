package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowServerResponse struct {
	Server         *ServerDetail `json:"server,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponse struct{}"
	}

	return strings.Join([]string{"ShowServerResponse", string(data)}, " ")
}
