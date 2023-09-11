package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ProjectFlavorLimit
type ProjectFlavorLimit struct {
}

func (o ProjectFlavorLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectFlavorLimit struct{}"
	}

	return strings.Join([]string{"ProjectFlavorLimit", string(data)}, " ")
}
