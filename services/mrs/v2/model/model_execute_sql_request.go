package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExecuteSqlRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *SqlExecutionReq `json:"body,omitempty"`
}

func (o ExecuteSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSqlRequest struct{}"
	}

	return strings.Join([]string{"ExecuteSqlRequest", string(data)}, " ")
}
