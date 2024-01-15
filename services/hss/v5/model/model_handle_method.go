package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HandleMethod struct {
}

func (o HandleMethod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleMethod struct{}"
	}

	return strings.Join([]string{"HandleMethod", string(data)}, " ")
}
