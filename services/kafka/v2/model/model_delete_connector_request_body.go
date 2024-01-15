package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteConnectorRequestBody struct {
	OrderId *string `json:"order_id,omitempty"`

	CsbInstanceProductId *string `json:"csb_instance_product_id,omitempty"`
}

func (o DeleteConnectorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectorRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteConnectorRequestBody", string(data)}, " ")
}
