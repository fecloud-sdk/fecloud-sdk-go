package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ImportImageQuickRequest Request Object
type ImportImageQuickRequest struct {
	Body *QuickImportImageByFileRequestBody `json:"body,omitempty"`
}

func (o ImportImageQuickRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickRequest struct{}"
	}

	return strings.Join([]string{"ImportImageQuickRequest", string(data)}, " ")
}
