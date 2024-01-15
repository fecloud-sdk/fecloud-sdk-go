package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogGroupsResponse struct {
	LogGroups      *[]LogGroup `json:"log_groups,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListLogGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListLogGroupsResponse", string(data)}, " ")
}
