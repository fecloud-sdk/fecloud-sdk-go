package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateDnatResponse struct {
	DnatRule *PrivateDnat `json:"dnat_rule,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatResponse", string(data)}, " ")
}
