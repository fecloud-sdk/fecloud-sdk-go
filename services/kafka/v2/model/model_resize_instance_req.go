package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceReq struct {
	NewSpecCode *string `json:"new_spec_code,omitempty"`

	NewStorageSpace *int32 `json:"new_storage_space,omitempty"`

	OperType *string `json:"oper_type,omitempty"`

	NewBrokerNum *int32 `json:"new_broker_num,omitempty"`

	NewProductId *string `json:"new_product_id,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o ResizeInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeInstanceReq", string(data)}, " ")
}
