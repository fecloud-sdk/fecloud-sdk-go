package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSqlserverDatabaseExResponse Response Object
type DeleteSqlserverDatabaseExResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSqlserverDatabaseExResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseExResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseExResponse", string(data)}, " ")
}
