package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateSnatResponse struct {
	SnatRule *PrivateSnat `json:"snat_rule,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatResponse", string(data)}, " ")
}
