package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateNatTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePrivateNatTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatTagResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatTagResponse", string(data)}, " ")
}
