package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateDatabaseObjectReq struct {
	Jobs []UpdateDatabaseObjectReq `json:"jobs"`
}

func (o BatchUpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateDatabaseObjectReq", string(data)}, " ")
}
