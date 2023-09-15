package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobInstanceInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`
}

func (o JobInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstanceInfo struct{}"
	}

	return strings.Join([]string{"JobInstanceInfo", string(data)}, " ")
}
