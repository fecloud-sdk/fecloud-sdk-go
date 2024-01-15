package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSqlJobRequestBody struct {
	Sql string `json:"sql"`

	Currentdb *string `json:"currentdb,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	Conf *[]string `json:"conf,omitempty"`

	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o CreateSqlJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlJobRequestBody", string(data)}, " ")
}
