package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesGroupResult struct {
	Id string `json:"id"`

	Status string `json:"status"`

	Volume *Volume `json:"volume"`

	Nodes []ListInstancesNodeResult `json:"nodes"`
}

func (o ListInstancesGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesGroupResult struct{}"
	}

	return strings.Join([]string{"ListInstancesGroupResult", string(data)}, " ")
}
