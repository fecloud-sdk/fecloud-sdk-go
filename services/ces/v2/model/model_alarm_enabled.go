package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmEnabled struct {
}

func (o AlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabled struct{}"
	}

	return strings.Join([]string{"AlarmEnabled", string(data)}, " ")
}
