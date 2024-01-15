package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateListenerReq struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	ConnectionLimit *int32 `json:"connection_limit,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	InsertHeaders *InsertHeader `json:"insert_headers,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ProtectionStatus *UpdateListenerReqProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdateListenerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerReq struct{}"
	}

	return strings.Join([]string{"UpdateListenerReq", string(data)}, " ")
}

type UpdateListenerReqProtectionStatus struct {
	value string
}

type UpdateListenerReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdateListenerReqProtectionStatus
	CONSOLE_PROTECTION UpdateListenerReqProtectionStatus
}

func GetUpdateListenerReqProtectionStatusEnum() UpdateListenerReqProtectionStatusEnum {
	return UpdateListenerReqProtectionStatusEnum{
		NON_PROTECTION: UpdateListenerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateListenerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateListenerReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdateListenerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerReqProtectionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
