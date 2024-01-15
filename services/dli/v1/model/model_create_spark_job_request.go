package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSparkJobRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *CreateSparkJobRequestBody `json:"body,omitempty"`
}

func (o CreateSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSparkJobRequest", string(data)}, " ")
}
