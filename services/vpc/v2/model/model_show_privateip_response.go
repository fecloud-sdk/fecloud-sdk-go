package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateipResponse struct {
	Privateip      *Privateip `json:"privateip,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowPrivateipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateipResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateipResponse", string(data)}, " ")
}
