package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSqlJobRequest struct {
	Body *CreateSqlJobRequestBody `json:"body,omitempty"`
}

func (o CreateSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlJobRequest", string(data)}, " ")
}
