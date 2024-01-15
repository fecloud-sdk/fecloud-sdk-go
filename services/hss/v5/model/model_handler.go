package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Handler struct {
}

func (o Handler) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Handler struct{}"
	}

	return strings.Join([]string{"Handler", string(data)}, " ")
}
