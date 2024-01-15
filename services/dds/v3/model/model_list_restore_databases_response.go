package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRestoreDatabasesResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Databases      *[]string `json:"databases,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreDatabasesResponse", string(data)}, " ")
}
