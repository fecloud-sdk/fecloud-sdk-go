package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/mrs/v2/model"
)

type MrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMrsClient(hcClient *http_client.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *MrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

func (c *MrsClient) BatchDeleteJobsInvoker(request *model.BatchDeleteJobsRequest) *BatchDeleteJobsInvoker {
	requestDef := GenReqDefForBatchDeleteJobs()
	return &BatchDeleteJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

func (c *MrsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CreateExecuteJob(request *model.CreateExecuteJobRequest) (*model.CreateExecuteJobResponse, error) {
	requestDef := GenReqDefForCreateExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExecuteJobResponse), nil
	}
}

func (c *MrsClient) CreateExecuteJobInvoker(request *model.CreateExecuteJobRequest) *CreateExecuteJobInvoker {
	requestDef := GenReqDefForCreateExecuteJob()
	return &CreateExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowAgencyMapping(request *model.ShowAgencyMappingRequest) (*model.ShowAgencyMappingResponse, error) {
	requestDef := GenReqDefForShowAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyMappingResponse), nil
	}
}

func (c *MrsClient) ShowAgencyMappingInvoker(request *model.ShowAgencyMappingRequest) *ShowAgencyMappingInvoker {
	requestDef := GenReqDefForShowAgencyMapping()
	return &ShowAgencyMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowJobExeListNew(request *model.ShowJobExeListNewRequest) (*model.ShowJobExeListNewResponse, error) {
	requestDef := GenReqDefForShowJobExeListNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobExeListNewResponse), nil
	}
}

func (c *MrsClient) ShowJobExeListNewInvoker(request *model.ShowJobExeListNewRequest) *ShowJobExeListNewInvoker {
	requestDef := GenReqDefForShowJobExeListNew()
	return &ShowJobExeListNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowSingleJobExe(request *model.ShowSingleJobExeRequest) (*model.ShowSingleJobExeResponse, error) {
	requestDef := GenReqDefForShowSingleJobExe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleJobExeResponse), nil
	}
}

func (c *MrsClient) ShowSingleJobExeInvoker(request *model.ShowSingleJobExeRequest) *ShowSingleJobExeInvoker {
	requestDef := GenReqDefForShowSingleJobExe()
	return &ShowSingleJobExeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowSqlResultWithJob(request *model.ShowSqlResultWithJobRequest) (*model.ShowSqlResultWithJobResponse, error) {
	requestDef := GenReqDefForShowSqlResultWithJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultWithJobResponse), nil
	}
}

func (c *MrsClient) ShowSqlResultWithJobInvoker(request *model.ShowSqlResultWithJobRequest) *ShowSqlResultWithJobInvoker {
	requestDef := GenReqDefForShowSqlResultWithJob()
	return &ShowSqlResultWithJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

func (c *MrsClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) UpdateAgencyMapping(request *model.UpdateAgencyMappingRequest) (*model.UpdateAgencyMappingResponse, error) {
	requestDef := GenReqDefForUpdateAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgencyMappingResponse), nil
	}
}

func (c *MrsClient) UpdateAgencyMappingInvoker(request *model.UpdateAgencyMappingRequest) *UpdateAgencyMappingInvoker {
	requestDef := GenReqDefForUpdateAgencyMapping()
	return &UpdateAgencyMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowHdfsFileList(request *model.ShowHdfsFileListRequest) (*model.ShowHdfsFileListResponse, error) {
	requestDef := GenReqDefForShowHdfsFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHdfsFileListResponse), nil
	}
}

func (c *MrsClient) ShowHdfsFileListInvoker(request *model.ShowHdfsFileListRequest) *ShowHdfsFileListInvoker {
	requestDef := GenReqDefForShowHdfsFileList()
	return &ShowHdfsFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) CancelSql(request *model.CancelSqlRequest) (*model.CancelSqlResponse, error) {
	requestDef := GenReqDefForCancelSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSqlResponse), nil
	}
}

func (c *MrsClient) CancelSqlInvoker(request *model.CancelSqlRequest) *CancelSqlInvoker {
	requestDef := GenReqDefForCancelSql()
	return &CancelSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ExecuteSql(request *model.ExecuteSqlRequest) (*model.ExecuteSqlResponse, error) {
	requestDef := GenReqDefForExecuteSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSqlResponse), nil
	}
}

func (c *MrsClient) ExecuteSqlInvoker(request *model.ExecuteSqlRequest) *ExecuteSqlInvoker {
	requestDef := GenReqDefForExecuteSql()
	return &ExecuteSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MrsClient) ShowSqlResult(request *model.ShowSqlResultRequest) (*model.ShowSqlResultResponse, error) {
	requestDef := GenReqDefForShowSqlResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultResponse), nil
	}
}

func (c *MrsClient) ShowSqlResultInvoker(request *model.ShowSqlResultRequest) *ShowSqlResultInvoker {
	requestDef := GenReqDefForShowSqlResult()
	return &ShowSqlResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
