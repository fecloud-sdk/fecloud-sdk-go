package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateSparkJobRequestBody struct {
	File string `json:"file"`

	ClassName string `json:"className"`

	ClusterName *string `json:"cluster_name,omitempty"`

	Args *[]string `json:"args,omitempty"`

	ScType *string `json:"sc_type,omitempty"`

	Jars *[]string `json:"jars,omitempty"`

	PyFiles *[]string `json:"pyFiles,omitempty"`

	Files *[]string `json:"files,omitempty"`

	Modules *[]string `json:"modules,omitempty"`

	Resources *[]SparkJobResource `json:"resources,omitempty"`

	Groups *[]SparkJobGroup `json:"groups,omitempty"`

	Conf map[string]interface{} `json:"conf,omitempty"`

	Name *string `json:"name,omitempty"`

	DriverMemory *string `json:"driverMemory,omitempty"`

	DriverCores *int32 `json:"driverCores,omitempty"`

	ExecutorMemory *string `json:"executorMemory,omitempty"`

	ExecutorCores *int32 `json:"executorCores,omitempty"`

	NumExecutors *int32 `json:"numExecutors,omitempty"`

	Feature *CreateSparkJobRequestBodyFeature `json:"feature,omitempty"`

	SparkVersion *string `json:"spark_version,omitempty"`

	Queue *string `json:"queue,omitempty"`

	AutoRecovery *bool `json:"auto_recovery,omitempty"`

	MaxRetryTimes *int32 `json:"max_retry_times,omitempty"`

	Image *string `json:"image,omitempty"`

	ObsBucket *string `json:"obs_bucket,omitempty"`

	CatalogName *string `json:"catalog_name,omitempty"`
}

func (o CreateSparkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSparkJobRequestBody", string(data)}, " ")
}

type CreateSparkJobRequestBodyFeature struct {
	value string
}

type CreateSparkJobRequestBodyFeatureEnum struct {
	BASIC  CreateSparkJobRequestBodyFeature
	AI     CreateSparkJobRequestBodyFeature
	CUSTOM CreateSparkJobRequestBodyFeature
}

func GetCreateSparkJobRequestBodyFeatureEnum() CreateSparkJobRequestBodyFeatureEnum {
	return CreateSparkJobRequestBodyFeatureEnum{
		BASIC: CreateSparkJobRequestBodyFeature{
			value: "basic",
		},
		AI: CreateSparkJobRequestBodyFeature{
			value: "ai",
		},
		CUSTOM: CreateSparkJobRequestBodyFeature{
			value: "custom",
		},
	}
}

func (c CreateSparkJobRequestBodyFeature) Value() string {
	return c.value
}

func (c CreateSparkJobRequestBodyFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSparkJobRequestBodyFeature) UnmarshalJSON(b []byte) error {
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
