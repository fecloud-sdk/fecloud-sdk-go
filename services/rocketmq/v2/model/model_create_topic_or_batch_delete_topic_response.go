package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicOrBatchDeleteTopicResponse struct {
	Id *string `json:"id,omitempty"`

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTopicOrBatchDeleteTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicResponse", string(data)}, " ")
}
