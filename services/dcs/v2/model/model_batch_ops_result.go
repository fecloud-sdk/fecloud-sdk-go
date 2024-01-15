package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchOpsResult struct {
	Result *string `json:"result,omitempty"`

	Instance *string `json:"instance,omitempty"`
}

func (o BatchOpsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOpsResult struct{}"
	}

	return strings.Join([]string{"BatchOpsResult", string(data)}, " ")
}
