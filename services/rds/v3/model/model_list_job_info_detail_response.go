package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListJobInfoDetailResponse struct {
	Jobs *GetTaskDetailListRspJobs `json:"jobs,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobInfoDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoDetailResponse struct{}"
	}

	return strings.Join([]string{"ListJobInfoDetailResponse", string(data)}, " ")
}
