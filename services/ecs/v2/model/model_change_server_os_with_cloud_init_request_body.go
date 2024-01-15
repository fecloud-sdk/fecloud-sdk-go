package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithCloudInitRequestBody struct {
	OsChange *ChangeServerOsWithCloudInitOption `json:"os-change"`
}

func (o ChangeServerOsWithCloudInitRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitRequestBody", string(data)}, " ")
}
