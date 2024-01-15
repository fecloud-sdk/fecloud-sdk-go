package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDeleteConnectorOrderResponse struct {
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeleteConnectorOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeleteConnectorOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateDeleteConnectorOrderResponse", string(data)}, " ")
}
