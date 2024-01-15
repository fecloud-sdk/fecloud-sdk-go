package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobBatchDelete struct {
	JobIdList *[]string `json:"job_id_list,omitempty"`
}

func (o JobBatchDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBatchDelete struct{}"
	}

	return strings.Join([]string{"JobBatchDelete", string(data)}, " ")
}
