package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LDatastore struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

func (o LDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LDatastore struct{}"
	}

	return strings.Join([]string{"LDatastore", string(data)}, " ")
}
