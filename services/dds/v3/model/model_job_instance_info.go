package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobInstanceInfo struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

func (o JobInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstanceInfo struct{}"
	}

	return strings.Join([]string{"JobInstanceInfo", string(data)}, " ")
}
