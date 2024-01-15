package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MongoUpdateReplSetV3RequestBody struct {
	Name string `json:"name"`
}

func (o MongoUpdateReplSetV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MongoUpdateReplSetV3RequestBody struct{}"
	}

	return strings.Join([]string{"MongoUpdateReplSetV3RequestBody", string(data)}, " ")
}
