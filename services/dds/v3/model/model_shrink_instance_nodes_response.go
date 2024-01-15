package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShrinkInstanceNodesResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodesResponse", string(data)}, " ")
}
