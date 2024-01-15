package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteMigrationTaskResponse struct {
	TaskIdList     *[]string `json:"task_id_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigrationTaskResponse", string(data)}, " ")
}
