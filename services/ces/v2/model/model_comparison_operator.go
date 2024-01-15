package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ComparisonOperator struct {
}

func (o ComparisonOperator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComparisonOperator struct{}"
	}

	return strings.Join([]string{"ComparisonOperator", string(data)}, " ")
}
