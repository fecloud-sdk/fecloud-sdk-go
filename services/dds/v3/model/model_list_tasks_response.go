package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTasksResponse struct {
	Jobs *[]JobInfo `json:"jobs,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
