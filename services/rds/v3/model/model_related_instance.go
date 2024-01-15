package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RelatedInstance struct {
	Id string `json:"id"`

	Type string `json:"type"`
}

func (o RelatedInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedInstance struct{}"
	}

	return strings.Join([]string{"RelatedInstance", string(data)}, " ")
}
