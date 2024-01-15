package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPostgresqlDbUserPaginatedRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListPostgresqlDbUserPaginatedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDbUserPaginatedRequest struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDbUserPaginatedRequest", string(data)}, " ")
}
