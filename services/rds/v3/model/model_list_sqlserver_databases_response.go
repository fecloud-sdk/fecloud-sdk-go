package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSqlserverDatabasesResponse struct {
	Databases *[]SqlserverDatabaseForDetail `json:"databases,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlserverDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlserverDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlserverDatabasesResponse", string(data)}, " ")
}
