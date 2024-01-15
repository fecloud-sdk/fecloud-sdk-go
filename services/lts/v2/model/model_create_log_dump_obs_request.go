package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogDumpObsRequest struct {
	ContentType string `json:"Content-Type"`

	Body *CreateLogDumpObsRequestBody `json:"body,omitempty"`
}

func (o CreateLogDumpObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsRequest struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsRequest", string(data)}, " ")
}
