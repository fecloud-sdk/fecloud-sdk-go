package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSqlRequest struct {
	ClusterId string `json:"cluster_id"`

	SqlId string `json:"sql_id"`
}

func (o CancelSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlRequest struct{}"
	}

	return strings.Join([]string{"CancelSqlRequest", string(data)}, " ")
}
