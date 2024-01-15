package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckSqlRequest struct {
	Body *CheckSqlRequestBody `json:"body,omitempty"`
}

func (o CheckSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSqlRequest struct{}"
	}

	return strings.Join([]string{"CheckSqlRequest", string(data)}, " ")
}
