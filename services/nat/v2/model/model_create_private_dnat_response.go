package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateDnatResponse struct {
	DnatRule *PrivateDnat `json:"dnat_rule,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateDnatResponse", string(data)}, " ")
}
