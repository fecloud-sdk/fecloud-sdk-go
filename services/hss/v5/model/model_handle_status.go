package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HandleStatus struct {
}

func (o HandleStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleStatus struct{}"
	}

	return strings.Join([]string{"HandleStatus", string(data)}, " ")
}
