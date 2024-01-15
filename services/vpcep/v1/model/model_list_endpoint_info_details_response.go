package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListEndpointInfoDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceType *ListEndpointInfoDetailsResponseServiceType `json:"service_type,omitempty"`

	Status *ListEndpointInfoDetailsResponseStatus `json:"status,omitempty"`

	ActiveStatus *[]string `json:"active_status,omitempty"`

	EnableStatus *ListEndpointInfoDetailsResponseEnableStatus `json:"enable_status,omitempty"`

	SpecificationName *string `json:"specification_name,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	DnsNames *[]string `json:"dns_names,omitempty"`

	Ip *string `json:"ip,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Error *QueryError `json:"error,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Routetables *[]string `json:"routetables,omitempty"`

	Description *string `json:"description,omitempty"`

	PolicyStatement *[]string `json:"policy_statement,omitempty"`

	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ListEndpointInfoDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointInfoDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointInfoDetailsResponse", string(data)}, " ")
}

type ListEndpointInfoDetailsResponseServiceType struct {
	value string
}

type ListEndpointInfoDetailsResponseServiceTypeEnum struct {
	INTERFACE ListEndpointInfoDetailsResponseServiceType
	GATEWAY   ListEndpointInfoDetailsResponseServiceType
}

func GetListEndpointInfoDetailsResponseServiceTypeEnum() ListEndpointInfoDetailsResponseServiceTypeEnum {
	return ListEndpointInfoDetailsResponseServiceTypeEnum{
		INTERFACE: ListEndpointInfoDetailsResponseServiceType{
			value: "interface",
		},
		GATEWAY: ListEndpointInfoDetailsResponseServiceType{
			value: "gateway",
		},
	}
}

func (c ListEndpointInfoDetailsResponseServiceType) Value() string {
	return c.value
}

func (c ListEndpointInfoDetailsResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointInfoDetailsResponseServiceType) UnmarshalJSON(b []byte) error {
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

type ListEndpointInfoDetailsResponseStatus struct {
	value string
}

type ListEndpointInfoDetailsResponseStatusEnum struct {
	PENDING_ACCEPTANCE ListEndpointInfoDetailsResponseStatus
	CREATING           ListEndpointInfoDetailsResponseStatus
	ACCEPTED           ListEndpointInfoDetailsResponseStatus
	REJECTED           ListEndpointInfoDetailsResponseStatus
	FAILED             ListEndpointInfoDetailsResponseStatus
	DELETING           ListEndpointInfoDetailsResponseStatus
}

func GetListEndpointInfoDetailsResponseStatusEnum() ListEndpointInfoDetailsResponseStatusEnum {
	return ListEndpointInfoDetailsResponseStatusEnum{
		PENDING_ACCEPTANCE: ListEndpointInfoDetailsResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: ListEndpointInfoDetailsResponseStatus{
			value: "creating",
		},
		ACCEPTED: ListEndpointInfoDetailsResponseStatus{
			value: "accepted",
		},
		REJECTED: ListEndpointInfoDetailsResponseStatus{
			value: "rejected",
		},
		FAILED: ListEndpointInfoDetailsResponseStatus{
			value: "failed",
		},
		DELETING: ListEndpointInfoDetailsResponseStatus{
			value: "deleting",
		},
	}
}

func (c ListEndpointInfoDetailsResponseStatus) Value() string {
	return c.value
}

func (c ListEndpointInfoDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointInfoDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ListEndpointInfoDetailsResponseEnableStatus struct {
	value string
}

type ListEndpointInfoDetailsResponseEnableStatusEnum struct {
	ENABLE  ListEndpointInfoDetailsResponseEnableStatus
	DISABLE ListEndpointInfoDetailsResponseEnableStatus
}

func GetListEndpointInfoDetailsResponseEnableStatusEnum() ListEndpointInfoDetailsResponseEnableStatusEnum {
	return ListEndpointInfoDetailsResponseEnableStatusEnum{
		ENABLE: ListEndpointInfoDetailsResponseEnableStatus{
			value: "enable",
		},
		DISABLE: ListEndpointInfoDetailsResponseEnableStatus{
			value: "disable",
		},
	}
}

func (c ListEndpointInfoDetailsResponseEnableStatus) Value() string {
	return c.value
}

func (c ListEndpointInfoDetailsResponseEnableStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointInfoDetailsResponseEnableStatus) UnmarshalJSON(b []byte) error {
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
