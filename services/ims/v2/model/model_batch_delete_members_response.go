package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteMembersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersResponse", string(data)}, " ")
}
