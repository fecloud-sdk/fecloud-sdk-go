package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupDatabase struct {
	Name string `json:"name"`
}

func (o BackupDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDatabase struct{}"
	}

	return strings.Join([]string{"BackupDatabase", string(data)}, " ")
}
