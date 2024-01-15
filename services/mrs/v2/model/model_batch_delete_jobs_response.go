package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsResponse", string(data)}, " ")
}
