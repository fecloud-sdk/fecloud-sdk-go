package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddMembersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchAddMembersResponse", string(data)}, " ")
}
