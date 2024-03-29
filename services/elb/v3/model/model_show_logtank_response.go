package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLogtankResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Logtank        *Logtank `json:"logtank,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankResponse struct{}"
	}

	return strings.Join([]string{"ShowLogtankResponse", string(data)}, " ")
}
