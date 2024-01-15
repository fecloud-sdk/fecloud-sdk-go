package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddReadonlyNodeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddReadonlyNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddReadonlyNodeResponse struct{}"
	}

	return strings.Join([]string{"AddReadonlyNodeResponse", string(data)}, " ")
}
