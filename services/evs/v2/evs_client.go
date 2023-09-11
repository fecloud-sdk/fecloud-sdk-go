package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/evs/v2/model"
)

type EvsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEvsClient(hcClient *http_client.HcHttpClient) *EvsClient {
	return &EvsClient{HcClient: hcClient}
}

func EvsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateVolumeTags 为指定云硬盘批量添加标签
//
// 为指定云硬盘批量添加标签。
//
// 添加标签时，如果云硬盘的标签已存在相同key，则会覆盖已有标签。
// 单个云硬盘最多支持创建10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchCreateVolumeTags(request *model.BatchCreateVolumeTagsRequest) (*model.BatchCreateVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

// BatchCreateVolumeTagsInvoker 为指定云硬盘批量添加标签
func (c *EvsClient) BatchCreateVolumeTagsInvoker(request *model.BatchCreateVolumeTagsRequest) *BatchCreateVolumeTagsInvoker {
	requestDef := GenReqDefForBatchCreateVolumeTags()
	return &BatchCreateVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteVolumeTags 为指定云硬盘批量删除标签
//
// 为指定云硬盘批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchDeleteVolumeTags(request *model.BatchDeleteVolumeTagsRequest) (*model.BatchDeleteVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

// BatchDeleteVolumeTagsInvoker 为指定云硬盘批量删除标签
func (c *EvsClient) BatchDeleteVolumeTagsInvoker(request *model.BatchDeleteVolumeTagsRequest) *BatchDeleteVolumeTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVolumeTags()
	return &BatchDeleteVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderAcceptVolumeTransfer 接受云硬盘过户
//
// 通过云硬盘过户记录ID以及身份认证密钥来接受云硬盘过户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderAcceptVolumeTransfer(request *model.CinderAcceptVolumeTransferRequest) (*model.CinderAcceptVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

// CinderAcceptVolumeTransferInvoker 接受云硬盘过户
func (c *EvsClient) CinderAcceptVolumeTransferInvoker(request *model.CinderAcceptVolumeTransferRequest) *CinderAcceptVolumeTransferInvoker {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()
	return &CinderAcceptVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderCreateVolumeTransfer 创建云硬盘过户
//
// 指定云硬盘来创建云硬盘过户记录，创建成功后，会返回过户记录ID以及身份认证密钥。
// 云硬盘在过户过程中的状态变化如下：创建云硬盘过户后，云硬盘状态由“available”变为“awaiting-transfer”。当云硬盘过户被接收后，云硬盘状态变为“available”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderCreateVolumeTransfer(request *model.CinderCreateVolumeTransferRequest) (*model.CinderCreateVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

// CinderCreateVolumeTransferInvoker 创建云硬盘过户
func (c *EvsClient) CinderCreateVolumeTransferInvoker(request *model.CinderCreateVolumeTransferRequest) *CinderCreateVolumeTransferInvoker {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()
	return &CinderCreateVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderDeleteVolumeTransfer 删除云硬盘过户
//
// 当云硬盘过户未被接受时，您可以删除云硬盘过户记录，接受后则无法执行删除操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderDeleteVolumeTransfer(request *model.CinderDeleteVolumeTransferRequest) (*model.CinderDeleteVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

// CinderDeleteVolumeTransferInvoker 删除云硬盘过户
func (c *EvsClient) CinderDeleteVolumeTransferInvoker(request *model.CinderDeleteVolumeTransferRequest) *CinderDeleteVolumeTransferInvoker {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()
	return &CinderDeleteVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListAvailabilityZones 查询所有的可用分区信息
//
// 查询所有的可用分区信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListAvailabilityZones(request *model.CinderListAvailabilityZonesRequest) (*model.CinderListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForCinderListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

// CinderListAvailabilityZonesInvoker 查询所有的可用分区信息
func (c *EvsClient) CinderListAvailabilityZonesInvoker(request *model.CinderListAvailabilityZonesRequest) *CinderListAvailabilityZonesInvoker {
	requestDef := GenReqDefForCinderListAvailabilityZones()
	return &CinderListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListQuotas 查询租户的详细配额
//
// 查询租户的详细配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListQuotas(request *model.CinderListQuotasRequest) (*model.CinderListQuotasResponse, error) {
	requestDef := GenReqDefForCinderListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListQuotasResponse), nil
	}
}

// CinderListQuotasInvoker 查询租户的详细配额
func (c *EvsClient) CinderListQuotasInvoker(request *model.CinderListQuotasRequest) *CinderListQuotasInvoker {
	requestDef := GenReqDefForCinderListQuotas()
	return &CinderListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListVolumeTransfers 查询云硬盘过户记录列表概要
//
// 查询当前租户下所有云硬盘的过户记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListVolumeTransfers(request *model.CinderListVolumeTransfersRequest) (*model.CinderListVolumeTransfersResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTransfersResponse), nil
	}
}

// CinderListVolumeTransfersInvoker 查询云硬盘过户记录列表概要
func (c *EvsClient) CinderListVolumeTransfersInvoker(request *model.CinderListVolumeTransfersRequest) *CinderListVolumeTransfersInvoker {
	requestDef := GenReqDefForCinderListVolumeTransfers()
	return &CinderListVolumeTransfersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListVolumeTypes 查询云硬盘类型列表
//
// 查询云硬盘类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListVolumeTypes(request *model.CinderListVolumeTypesRequest) (*model.CinderListVolumeTypesResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTypesResponse), nil
	}
}

// CinderListVolumeTypesInvoker 查询云硬盘类型列表
func (c *EvsClient) CinderListVolumeTypesInvoker(request *model.CinderListVolumeTypesRequest) *CinderListVolumeTypesInvoker {
	requestDef := GenReqDefForCinderListVolumeTypes()
	return &CinderListVolumeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderShowVolumeTransfer 查询单个云硬盘过户记录详情
//
// 查询单个云硬盘的过户记录详情，比如过户记录创建时间、ID以及名称等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderShowVolumeTransfer(request *model.CinderShowVolumeTransferRequest) (*model.CinderShowVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderShowVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderShowVolumeTransferResponse), nil
	}
}

// CinderShowVolumeTransferInvoker 查询单个云硬盘过户记录详情
func (c *EvsClient) CinderShowVolumeTransferInvoker(request *model.CinderShowVolumeTransferRequest) *CinderShowVolumeTransferInvoker {
	requestDef := GenReqDefForCinderShowVolumeTransfer()
	return &CinderShowVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumeTags 获取云硬盘资源的所有标签
//
// 获取某个租户的所有云硬盘资源的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumeTags(request *model.ListVolumeTagsRequest) (*model.ListVolumeTagsResponse, error) {
	requestDef := GenReqDefForListVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTagsResponse), nil
	}
}

// ListVolumeTagsInvoker 获取云硬盘资源的所有标签
func (c *EvsClient) ListVolumeTagsInvoker(request *model.ListVolumeTagsRequest) *ListVolumeTagsInvoker {
	requestDef := GenReqDefForListVolumeTags()
	return &ListVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumes 查询所有云硬盘详情
//
// 查询所有云硬盘的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

// ListVolumesInvoker 查询所有云硬盘详情
func (c *EvsClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumesByTags 通过标签查询云硬盘资源实例详情
//
// 通过标签查询云硬盘资源实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumesByTags(request *model.ListVolumesByTagsRequest) (*model.ListVolumesByTagsResponse, error) {
	requestDef := GenReqDefForListVolumesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesByTagsResponse), nil
	}
}

// ListVolumesByTagsInvoker 通过标签查询云硬盘资源实例详情
func (c *EvsClient) ListVolumesByTagsInvoker(request *model.ListVolumesByTagsRequest) *ListVolumesByTagsInvoker {
	requestDef := GenReqDefForListVolumesByTags()
	return &ListVolumesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询job的状态
//
// 查询Job的执行状态。
// 可用于查询创建云硬盘，扩容云硬盘，删除云硬盘等API的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询job的状态
func (c *EvsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolumeTags 查询云硬盘标签
//
// 查询指定云硬盘的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVolumeTags(request *model.ShowVolumeTagsRequest) (*model.ShowVolumeTagsResponse, error) {
	requestDef := GenReqDefForShowVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTagsResponse), nil
	}
}

// ShowVolumeTagsInvoker 查询云硬盘标签
func (c *EvsClient) ShowVolumeTagsInvoker(request *model.ShowVolumeTagsRequest) *ShowVolumeTagsInvoker {
	requestDef := GenReqDefForShowVolumeTags()
	return &ShowVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersions 查询接口版本信息列表
//
// 查询接口版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

// ListVersionsInvoker 查询接口版本信息列表
func (c *EvsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 查询API接口的版本信息
//
// 查询接口的指定版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 查询API接口的版本信息
func (c *EvsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}