package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePoolRequest struct {
	PoolId string `json:"pool_id"`
}

func (o DeletePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolRequest struct{}"
	}

	return strings.Join([]string{"DeletePoolRequest", string(data)}, " ")
}
