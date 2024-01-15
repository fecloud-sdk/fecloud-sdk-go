package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogsResponse struct {
	Logs *[]LogContents `json:"logs,omitempty"`

	Count *int32 `json:"count,omitempty"`

	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
