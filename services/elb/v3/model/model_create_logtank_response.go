package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankResponse struct {
	Logtank *Logtank `json:"logtank,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankResponse struct{}"
	}

	return strings.Join([]string{"CreateLogtankResponse", string(data)}, " ")
}
