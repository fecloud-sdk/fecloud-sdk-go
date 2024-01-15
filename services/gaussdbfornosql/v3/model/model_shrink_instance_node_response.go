package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShrinkInstanceNodeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeResponse", string(data)}, " ")
}
