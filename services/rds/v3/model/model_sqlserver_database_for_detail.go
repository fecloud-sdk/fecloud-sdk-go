package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SqlserverDatabaseForDetail struct {
	Name string `json:"name"`

	CharacterSet string `json:"character_set"`

	State string `json:"state"`
}

func (o SqlserverDatabaseForDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverDatabaseForDetail struct{}"
	}

	return strings.Join([]string{"SqlserverDatabaseForDetail", string(data)}, " ")
}
