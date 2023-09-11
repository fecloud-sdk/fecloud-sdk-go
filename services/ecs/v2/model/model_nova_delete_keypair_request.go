package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaDeleteKeypairRequest Request Object
type NovaDeleteKeypairRequest struct {

	// 密钥名称。
	KeypairName string `json:"keypair_name"`
}

func (o NovaDeleteKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaDeleteKeypairRequest", string(data)}, " ")
}