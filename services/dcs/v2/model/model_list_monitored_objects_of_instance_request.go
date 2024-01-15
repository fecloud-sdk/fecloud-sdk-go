package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListMonitoredObjectsOfInstanceRequest struct {
	InstanceId string `json:"instance_id"`

	DimName string `json:"dim_name"`
}

func (o ListMonitoredObjectsOfInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsOfInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsOfInstanceRequest", string(data)}, " ")
}
