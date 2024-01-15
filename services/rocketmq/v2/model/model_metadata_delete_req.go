package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MetadataDeleteReq struct {
	TaskIds []string `json:"taskIds"`
}

func (o MetadataDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataDeleteReq struct{}"
	}

	return strings.Join([]string{"MetadataDeleteReq", string(data)}, " ")
}
