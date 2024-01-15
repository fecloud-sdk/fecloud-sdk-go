package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAuditlogLinksRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ProduceAuditlogLinksRequestBody `json:"body,omitempty"`
}

func (o ListAuditlogLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogLinksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditlogLinksRequest", string(data)}, " ")
}
