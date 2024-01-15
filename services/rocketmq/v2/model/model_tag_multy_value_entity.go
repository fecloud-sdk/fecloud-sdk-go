package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagMultyValueEntity struct {
	Key *string `json:"key,omitempty"`

	Values *[]string `json:"values,omitempty"`
}

func (o TagMultyValueEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMultyValueEntity struct{}"
	}

	return strings.Join([]string{"TagMultyValueEntity", string(data)}, " ")
}
