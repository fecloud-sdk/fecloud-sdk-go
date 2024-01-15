package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MigrateAzRequestBody struct {
	TargetAzs string `json:"target_azs"`
}

func (o MigrateAzRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateAzRequestBody", string(data)}, " ")
}
