package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDatabaseReq struct {
	Name string `json:"name"`

	Comment *string `json:"comment,omitempty"`
}

func (o UpdateDatabaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseReq struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseReq", string(data)}, " ")
}
