package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLogtankResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Logtank        *Logtank `json:"logtank,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
