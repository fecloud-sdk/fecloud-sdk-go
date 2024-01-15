package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandInstanceNodeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeResponse", string(data)}, " ")
}
