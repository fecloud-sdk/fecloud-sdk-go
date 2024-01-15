package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmTemplateId struct {
}

func (o AlarmTemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateId struct{}"
	}

	return strings.Join([]string{"AlarmTemplateId", string(data)}, " ")
}
