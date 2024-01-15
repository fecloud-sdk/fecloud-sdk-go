package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowJobExeListNewResponse struct {
	TotalRecord *int32 `json:"total_record,omitempty"`

	JobList        *[]JobQueryBean `json:"job_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowJobExeListNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExeListNewResponse struct{}"
	}

	return strings.Join([]string{"ShowJobExeListNewResponse", string(data)}, " ")
}
