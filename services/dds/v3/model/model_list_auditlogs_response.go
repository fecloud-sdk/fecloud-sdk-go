package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAuditlogsResponse struct {
	TotalRecord *int32 `json:"total_record,omitempty"`

	AuditLogs      *[]ListAuditlogsResult `json:"audit_logs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAuditlogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditlogsResponse", string(data)}, " ")
}
