package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateSnatResponse struct {
	SnatRule *PrivateSnat `json:"snat_rule,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateSnatResponse", string(data)}, " ")
}
