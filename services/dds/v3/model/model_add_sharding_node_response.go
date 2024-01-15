package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddShardingNodeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddShardingNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeResponse struct{}"
	}

	return strings.Join([]string{"AddShardingNodeResponse", string(data)}, " ")
}
