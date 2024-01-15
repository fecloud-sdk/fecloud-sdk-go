package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppliedInstancesResponse struct {
	Instances      *[]ApplicableInstancesInfo `json:"instances,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListAppliedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAppliedInstancesResponse", string(data)}, " ")
}
