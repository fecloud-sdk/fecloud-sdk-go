package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeOpsWindowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeOpsWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOpsWindowResponse struct{}"
	}

	return strings.Join([]string{"ChangeOpsWindowResponse", string(data)}, " ")
}
