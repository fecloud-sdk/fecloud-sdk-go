package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankRequestBody struct {
	Logtank *CreateLogtankOption `json:"logtank"`
}

func (o CreateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequestBody", string(data)}, " ")
}
