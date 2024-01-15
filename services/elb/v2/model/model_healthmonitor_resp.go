package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type HealthmonitorResp struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	AdminStateUp bool `json:"admin_state_up"`

	MonitorPort int32 `json:"monitor_port"`

	Timeout int32 `json:"timeout"`

	Type HealthmonitorRespType `json:"type"`

	ExpectedCodes string `json:"expected_codes"`

	DomainName string `json:"domain_name"`

	UrlPath string `json:"url_path"`

	HttpMethod string `json:"http_method"`

	Delay int32 `json:"delay"`

	MaxRetries int32 `json:"max_retries"`

	Pools []ResourceList `json:"pools"`

	MaxRetriesDown int32 `json:"max_retries_down"`
}

func (o HealthmonitorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthmonitorResp struct{}"
	}

	return strings.Join([]string{"HealthmonitorResp", string(data)}, " ")
}

type HealthmonitorRespType struct {
	value string
}

type HealthmonitorRespTypeEnum struct {
	TCP         HealthmonitorRespType
	UDP_CONNECT HealthmonitorRespType
	HTTP        HealthmonitorRespType
}

func GetHealthmonitorRespTypeEnum() HealthmonitorRespTypeEnum {
	return HealthmonitorRespTypeEnum{
		TCP: HealthmonitorRespType{
			value: "TCP",
		},
		UDP_CONNECT: HealthmonitorRespType{
			value: "UDP_CONNECT",
		},
		HTTP: HealthmonitorRespType{
			value: "HTTP",
		},
	}
}

func (c HealthmonitorRespType) Value() string {
	return c.value
}

func (c HealthmonitorRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthmonitorRespType) UnmarshalJSON(b []byte) error {
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
