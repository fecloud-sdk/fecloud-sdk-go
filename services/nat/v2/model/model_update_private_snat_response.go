package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateSnatResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SnatRule       *PrivateSnat `json:"snat_rule,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdatePrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatResponse", string(data)}, " ")
}
