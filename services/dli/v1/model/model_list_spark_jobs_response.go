package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSparkJobsResponse struct {
	From *int32 `json:"from,omitempty"`

	Total *int32 `json:"total,omitempty"`

	Sessions *[]SparkJobInfo `json:"sessions,omitempty"`

	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSparkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListSparkJobsResponse", string(data)}, " ")
}
