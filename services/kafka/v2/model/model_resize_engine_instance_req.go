package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeEngineInstanceReq struct {
	OperType string `json:"oper_type"`

	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	NewProductId *string `json:"new_product_id,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o ResizeEngineInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceReq", string(data)}, " ")
}
