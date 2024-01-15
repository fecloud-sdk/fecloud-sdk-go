package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ClientData struct {
	Language *string `json:"language,omitempty"`

	Version *string `json:"version,omitempty"`

	ClientId *string `json:"client_id,omitempty"`

	ClientAddr *string `json:"client_addr,omitempty"`

	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}

func (o ClientData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientData struct{}"
	}

	return strings.Join([]string{"ClientData", string(data)}, " ")
}
