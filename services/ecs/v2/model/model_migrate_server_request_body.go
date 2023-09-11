package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// MigrateServerRequestBody This is a auto create Body Object
type MigrateServerRequestBody struct {
	Migrate *MigrateServerOption `json:"migrate"`
}

func (o MigrateServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateServerRequestBody", string(data)}, " ")
}
