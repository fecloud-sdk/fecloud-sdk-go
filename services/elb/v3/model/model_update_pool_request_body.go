package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePoolRequestBody This is a auto create Body Object
type UpdatePoolRequestBody struct {
	Pool *UpdatePoolOption `json:"pool"`
}

func (o UpdatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequestBody", string(data)}, " ")
}
