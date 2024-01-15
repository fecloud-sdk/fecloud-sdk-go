package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesRequest struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Mode *string `json:"mode,omitempty"`

	DatastoreType *string `json:"datastore_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
