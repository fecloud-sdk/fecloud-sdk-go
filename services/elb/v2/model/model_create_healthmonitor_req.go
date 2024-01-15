package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateHealthmonitorReq struct {
	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	MonitorPort *int32 `json:"monitor_port,omitempty"`

	Timeout int32 `json:"timeout"`

	Type CreateHealthmonitorReqType `json:"type"`

	ExpectedCodes *string `json:"expected_codes,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	UrlPath *string `json:"url_path,omitempty"`

	HttpMethod *string `json:"http_method,omitempty"`

	Delay int32 `json:"delay"`

	MaxRetries int32 `json:"max_retries"`

	PoolId string `json:"pool_id"`
}

func (o CreateHealthmonitorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorReq struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorReq", string(data)}, " ")
}

type CreateHealthmonitorReqType struct {
	value string
}

type CreateHealthmonitorReqTypeEnum struct {
	TCP         CreateHealthmonitorReqType
	UDP_CONNECT CreateHealthmonitorReqType
	HTTP        CreateHealthmonitorReqType
}

func GetCreateHealthmonitorReqTypeEnum() CreateHealthmonitorReqTypeEnum {
	return CreateHealthmonitorReqTypeEnum{
		TCP: CreateHealthmonitorReqType{
			value: "TCP",
		},
		UDP_CONNECT: CreateHealthmonitorReqType{
			value: "UDP_CONNECT",
		},
		HTTP: CreateHealthmonitorReqType{
			value: "HTTP",
		},
	}
}

func (c CreateHealthmonitorReqType) Value() string {
	return c.value
}

func (c CreateHealthmonitorReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHealthmonitorReqType) UnmarshalJSON(b []byte) error {
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
