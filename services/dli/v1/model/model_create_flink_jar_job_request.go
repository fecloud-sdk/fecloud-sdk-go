package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkJarJobRequest struct {
	Body *CreateFlinkJarJobRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkJarJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobRequest", string(data)}, " ")
}
