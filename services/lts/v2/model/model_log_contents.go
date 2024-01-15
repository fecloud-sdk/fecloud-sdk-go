package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LogContents struct {
	Content *string `json:"content,omitempty"`

	LineNum *string `json:"line_num,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}

func (o LogContents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogContents struct{}"
	}

	return strings.Join([]string{"LogContents", string(data)}, " ")
}
