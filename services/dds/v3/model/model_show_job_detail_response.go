package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowJobDetailResponse struct {
	Job            *JobDetail `json:"job,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
