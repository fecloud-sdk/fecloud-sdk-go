package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// GlanceCreateTagResponse Response Object
type GlanceCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceCreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagResponse", string(data)}, " ")
}
