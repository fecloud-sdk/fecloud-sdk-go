package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListJobInfoResponse Response Object
type ListJobInfoResponse struct {
	Job            *GetJobInfoResponseBodyJob `json:"job,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ListJobInfoResponse", string(data)}, " ")
}
