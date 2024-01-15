package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAuditLogResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditLogResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditLogResponse", string(data)}, " ")
}
