package v1

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dli/v1/model"
)

type DliClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDliClient(hcClient *http_client.HcHttpClient) *DliClient {
	return &DliClient{HcClient: hcClient}
}

func DliClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DliClient) AssociateQueueToEnhancedConnection(request *model.AssociateQueueToEnhancedConnectionRequest) (*model.AssociateQueueToEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForAssociateQueueToEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateQueueToEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) AssociateQueueToEnhancedConnectionInvoker(request *model.AssociateQueueToEnhancedConnectionRequest) *AssociateQueueToEnhancedConnectionInvoker {
	requestDef := GenReqDefForAssociateQueueToEnhancedConnection()
	return &AssociateQueueToEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateEnhancedConnection(request *model.CreateEnhancedConnectionRequest) (*model.CreateEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForCreateEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) CreateEnhancedConnectionInvoker(request *model.CreateEnhancedConnectionRequest) *CreateEnhancedConnectionInvoker {
	requestDef := GenReqDefForCreateEnhancedConnection()
	return &CreateEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateGlobalVariable(request *model.CreateGlobalVariableRequest) (*model.CreateGlobalVariableResponse, error) {
	requestDef := GenReqDefForCreateGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalVariableResponse), nil
	}
}

func (c *DliClient) CreateGlobalVariableInvoker(request *model.CreateGlobalVariableRequest) *CreateGlobalVariableInvoker {
	requestDef := GenReqDefForCreateGlobalVariable()
	return &CreateGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateQueue(request *model.CreateQueueRequest) (*model.CreateQueueResponse, error) {
	requestDef := GenReqDefForCreateQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQueueResponse), nil
	}
}

func (c *DliClient) CreateQueueInvoker(request *model.CreateQueueRequest) *CreateQueueInvoker {
	requestDef := GenReqDefForCreateQueue()
	return &CreateQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DeleteEnhancedConnection(request *model.DeleteEnhancedConnectionRequest) (*model.DeleteEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForDeleteEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) DeleteEnhancedConnectionInvoker(request *model.DeleteEnhancedConnectionRequest) *DeleteEnhancedConnectionInvoker {
	requestDef := GenReqDefForDeleteEnhancedConnection()
	return &DeleteEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DeleteGlobalVariable(request *model.DeleteGlobalVariableRequest) (*model.DeleteGlobalVariableResponse, error) {
	requestDef := GenReqDefForDeleteGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalVariableResponse), nil
	}
}

func (c *DliClient) DeleteGlobalVariableInvoker(request *model.DeleteGlobalVariableRequest) *DeleteGlobalVariableInvoker {
	requestDef := GenReqDefForDeleteGlobalVariable()
	return &DeleteGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DeleteQueue(request *model.DeleteQueueRequest) (*model.DeleteQueueResponse, error) {
	requestDef := GenReqDefForDeleteQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueueResponse), nil
	}
}

func (c *DliClient) DeleteQueueInvoker(request *model.DeleteQueueRequest) *DeleteQueueInvoker {
	requestDef := GenReqDefForDeleteQueue()
	return &DeleteQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DisassociateQueueFromEnhancedConnection(request *model.DisassociateQueueFromEnhancedConnectionRequest) (*model.DisassociateQueueFromEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForDisassociateQueueFromEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateQueueFromEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) DisassociateQueueFromEnhancedConnectionInvoker(request *model.DisassociateQueueFromEnhancedConnectionRequest) *DisassociateQueueFromEnhancedConnectionInvoker {
	requestDef := GenReqDefForDisassociateQueueFromEnhancedConnection()
	return &DisassociateQueueFromEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListEnhancedConnections(request *model.ListEnhancedConnectionsRequest) (*model.ListEnhancedConnectionsResponse, error) {
	requestDef := GenReqDefForListEnhancedConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnhancedConnectionsResponse), nil
	}
}

func (c *DliClient) ListEnhancedConnectionsInvoker(request *model.ListEnhancedConnectionsRequest) *ListEnhancedConnectionsInvoker {
	requestDef := GenReqDefForListEnhancedConnections()
	return &ListEnhancedConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListGlobalVariables(request *model.ListGlobalVariablesRequest) (*model.ListGlobalVariablesResponse, error) {
	requestDef := GenReqDefForListGlobalVariables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalVariablesResponse), nil
	}
}

func (c *DliClient) ListGlobalVariablesInvoker(request *model.ListGlobalVariablesRequest) *ListGlobalVariablesInvoker {
	requestDef := GenReqDefForListGlobalVariables()
	return &ListGlobalVariablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListQueues(request *model.ListQueuesRequest) (*model.ListQueuesResponse, error) {
	requestDef := GenReqDefForListQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueuesResponse), nil
	}
}

func (c *DliClient) ListQueuesInvoker(request *model.ListQueuesRequest) *ListQueuesInvoker {
	requestDef := GenReqDefForListQueues()
	return &ListQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowEnhancedConnection(request *model.ShowEnhancedConnectionRequest) (*model.ShowEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForShowEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) ShowEnhancedConnectionInvoker(request *model.ShowEnhancedConnectionRequest) *ShowEnhancedConnectionInvoker {
	requestDef := GenReqDefForShowEnhancedConnection()
	return &ShowEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowEnhancedConnectionPrivilege(request *model.ShowEnhancedConnectionPrivilegeRequest) (*model.ShowEnhancedConnectionPrivilegeResponse, error) {
	requestDef := GenReqDefForShowEnhancedConnectionPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnhancedConnectionPrivilegeResponse), nil
	}
}

func (c *DliClient) ShowEnhancedConnectionPrivilegeInvoker(request *model.ShowEnhancedConnectionPrivilegeRequest) *ShowEnhancedConnectionPrivilegeInvoker {
	requestDef := GenReqDefForShowEnhancedConnectionPrivilege()
	return &ShowEnhancedConnectionPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowQueue(request *model.ShowQueueRequest) (*model.ShowQueueResponse, error) {
	requestDef := GenReqDefForShowQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueueResponse), nil
	}
}

func (c *DliClient) ShowQueueInvoker(request *model.ShowQueueRequest) *ShowQueueInvoker {
	requestDef := GenReqDefForShowQueue()
	return &ShowQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) UpdateEnhancedConnection(request *model.UpdateEnhancedConnectionRequest) (*model.UpdateEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForUpdateEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnhancedConnectionResponse), nil
	}
}

func (c *DliClient) UpdateEnhancedConnectionInvoker(request *model.UpdateEnhancedConnectionRequest) *UpdateEnhancedConnectionInvoker {
	requestDef := GenReqDefForUpdateEnhancedConnection()
	return &UpdateEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) UpdateGlobalVariable(request *model.UpdateGlobalVariableRequest) (*model.UpdateGlobalVariableResponse, error) {
	requestDef := GenReqDefForUpdateGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalVariableResponse), nil
	}
}

func (c *DliClient) UpdateGlobalVariableInvoker(request *model.UpdateGlobalVariableRequest) *UpdateGlobalVariableInvoker {
	requestDef := GenReqDefForUpdateGlobalVariable()
	return &UpdateGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateFlinkJarJob(request *model.CreateFlinkJarJobRequest) (*model.CreateFlinkJarJobResponse, error) {
	requestDef := GenReqDefForCreateFlinkJarJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkJarJobResponse), nil
	}
}

func (c *DliClient) CreateFlinkJarJobInvoker(request *model.CreateFlinkJarJobRequest) *CreateFlinkJarJobInvoker {
	requestDef := GenReqDefForCreateFlinkJarJob()
	return &CreateFlinkJarJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateFlinkSqlJob(request *model.CreateFlinkSqlJobRequest) (*model.CreateFlinkSqlJobResponse, error) {
	requestDef := GenReqDefForCreateFlinkSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSqlJobResponse), nil
	}
}

func (c *DliClient) CreateFlinkSqlJobInvoker(request *model.CreateFlinkSqlJobRequest) *CreateFlinkSqlJobInvoker {
	requestDef := GenReqDefForCreateFlinkSqlJob()
	return &CreateFlinkSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateFlinkSqlJobTemplate(request *model.CreateFlinkSqlJobTemplateRequest) (*model.CreateFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForCreateFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSqlJobTemplateResponse), nil
	}
}

func (c *DliClient) CreateFlinkSqlJobTemplateInvoker(request *model.CreateFlinkSqlJobTemplateRequest) *CreateFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForCreateFlinkSqlJobTemplate()
	return &CreateFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DeleteFlinkJob(request *model.DeleteFlinkJobRequest) (*model.DeleteFlinkJobResponse, error) {
	requestDef := GenReqDefForDeleteFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkJobResponse), nil
	}
}

func (c *DliClient) DeleteFlinkJobInvoker(request *model.DeleteFlinkJobRequest) *DeleteFlinkJobInvoker {
	requestDef := GenReqDefForDeleteFlinkJob()
	return &DeleteFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) DeleteFlinkSqlJobTemplate(request *model.DeleteFlinkSqlJobTemplateRequest) (*model.DeleteFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForDeleteFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkSqlJobTemplateResponse), nil
	}
}

func (c *DliClient) DeleteFlinkSqlJobTemplateInvoker(request *model.DeleteFlinkSqlJobTemplateRequest) *DeleteFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForDeleteFlinkSqlJobTemplate()
	return &DeleteFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListFlinkJobs(request *model.ListFlinkJobsRequest) (*model.ListFlinkJobsResponse, error) {
	requestDef := GenReqDefForListFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobsResponse), nil
	}
}

func (c *DliClient) ListFlinkJobsInvoker(request *model.ListFlinkJobsRequest) *ListFlinkJobsInvoker {
	requestDef := GenReqDefForListFlinkJobs()
	return &ListFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListFlinkSqlJobTemplates(request *model.ListFlinkSqlJobTemplatesRequest) (*model.ListFlinkSqlJobTemplatesResponse, error) {
	requestDef := GenReqDefForListFlinkSqlJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkSqlJobTemplatesResponse), nil
	}
}

func (c *DliClient) ListFlinkSqlJobTemplatesInvoker(request *model.ListFlinkSqlJobTemplatesRequest) *ListFlinkSqlJobTemplatesInvoker {
	requestDef := GenReqDefForListFlinkSqlJobTemplates()
	return &ListFlinkSqlJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowFlinkJob(request *model.ShowFlinkJobRequest) (*model.ShowFlinkJobResponse, error) {
	requestDef := GenReqDefForShowFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobResponse), nil
	}
}

func (c *DliClient) ShowFlinkJobInvoker(request *model.ShowFlinkJobRequest) *ShowFlinkJobInvoker {
	requestDef := GenReqDefForShowFlinkJob()
	return &ShowFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) UpdateFlinkJarJob(request *model.UpdateFlinkJarJobRequest) (*model.UpdateFlinkJarJobResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJarJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJarJobResponse), nil
	}
}

func (c *DliClient) UpdateFlinkJarJobInvoker(request *model.UpdateFlinkJarJobRequest) *UpdateFlinkJarJobInvoker {
	requestDef := GenReqDefForUpdateFlinkJarJob()
	return &UpdateFlinkJarJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) UpdateFlinkSqlJob(request *model.UpdateFlinkSqlJobRequest) (*model.UpdateFlinkSqlJobResponse, error) {
	requestDef := GenReqDefForUpdateFlinkSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkSqlJobResponse), nil
	}
}

func (c *DliClient) UpdateFlinkSqlJobInvoker(request *model.UpdateFlinkSqlJobRequest) *UpdateFlinkSqlJobInvoker {
	requestDef := GenReqDefForUpdateFlinkSqlJob()
	return &UpdateFlinkSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) UpdateFlinkSqlJobTemplate(request *model.UpdateFlinkSqlJobTemplateRequest) (*model.UpdateFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForUpdateFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkSqlJobTemplateResponse), nil
	}
}

func (c *DliClient) UpdateFlinkSqlJobTemplateInvoker(request *model.UpdateFlinkSqlJobTemplateRequest) *UpdateFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForUpdateFlinkSqlJobTemplate()
	return &UpdateFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CancelSparkJob(request *model.CancelSparkJobRequest) (*model.CancelSparkJobResponse, error) {
	requestDef := GenReqDefForCancelSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkJobResponse), nil
	}
}

func (c *DliClient) CancelSparkJobInvoker(request *model.CancelSparkJobRequest) *CancelSparkJobInvoker {
	requestDef := GenReqDefForCancelSparkJob()
	return &CancelSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateSparkJob(request *model.CreateSparkJobRequest) (*model.CreateSparkJobResponse, error) {
	requestDef := GenReqDefForCreateSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSparkJobResponse), nil
	}
}

func (c *DliClient) CreateSparkJobInvoker(request *model.CreateSparkJobRequest) *CreateSparkJobInvoker {
	requestDef := GenReqDefForCreateSparkJob()
	return &CreateSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListSparkJobs(request *model.ListSparkJobsRequest) (*model.ListSparkJobsResponse, error) {
	requestDef := GenReqDefForListSparkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkJobsResponse), nil
	}
}

func (c *DliClient) ListSparkJobsInvoker(request *model.ListSparkJobsRequest) *ListSparkJobsInvoker {
	requestDef := GenReqDefForListSparkJobs()
	return &ListSparkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowSparkJob(request *model.ShowSparkJobRequest) (*model.ShowSparkJobResponse, error) {
	requestDef := GenReqDefForShowSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobResponse), nil
	}
}

func (c *DliClient) ShowSparkJobInvoker(request *model.ShowSparkJobRequest) *ShowSparkJobInvoker {
	requestDef := GenReqDefForShowSparkJob()
	return &ShowSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowSparkJobStatus(request *model.ShowSparkJobStatusRequest) (*model.ShowSparkJobStatusResponse, error) {
	requestDef := GenReqDefForShowSparkJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobStatusResponse), nil
	}
}

func (c *DliClient) ShowSparkJobStatusInvoker(request *model.ShowSparkJobStatusRequest) *ShowSparkJobStatusInvoker {
	requestDef := GenReqDefForShowSparkJobStatus()
	return &ShowSparkJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CancelSqlJob(request *model.CancelSqlJobRequest) (*model.CancelSqlJobResponse, error) {
	requestDef := GenReqDefForCancelSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSqlJobResponse), nil
	}
}

func (c *DliClient) CancelSqlJobInvoker(request *model.CancelSqlJobRequest) *CancelSqlJobInvoker {
	requestDef := GenReqDefForCancelSqlJob()
	return &CancelSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CheckSql(request *model.CheckSqlRequest) (*model.CheckSqlResponse, error) {
	requestDef := GenReqDefForCheckSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckSqlResponse), nil
	}
}

func (c *DliClient) CheckSqlInvoker(request *model.CheckSqlRequest) *CheckSqlInvoker {
	requestDef := GenReqDefForCheckSql()
	return &CheckSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) CreateSqlJob(request *model.CreateSqlJobRequest) (*model.CreateSqlJobResponse, error) {
	requestDef := GenReqDefForCreateSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlJobResponse), nil
	}
}

func (c *DliClient) CreateSqlJobInvoker(request *model.CreateSqlJobRequest) *CreateSqlJobInvoker {
	requestDef := GenReqDefForCreateSqlJob()
	return &CreateSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ListSqlJobs(request *model.ListSqlJobsRequest) (*model.ListSqlJobsResponse, error) {
	requestDef := GenReqDefForListSqlJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlJobsResponse), nil
	}
}

func (c *DliClient) ListSqlJobsInvoker(request *model.ListSqlJobsRequest) *ListSqlJobsInvoker {
	requestDef := GenReqDefForListSqlJobs()
	return &ListSqlJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowSqlJobDetail(request *model.ShowSqlJobDetailRequest) (*model.ShowSqlJobDetailResponse, error) {
	requestDef := GenReqDefForShowSqlJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobDetailResponse), nil
	}
}

func (c *DliClient) ShowSqlJobDetailInvoker(request *model.ShowSqlJobDetailRequest) *ShowSqlJobDetailInvoker {
	requestDef := GenReqDefForShowSqlJobDetail()
	return &ShowSqlJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowSqlJobProgress(request *model.ShowSqlJobProgressRequest) (*model.ShowSqlJobProgressResponse, error) {
	requestDef := GenReqDefForShowSqlJobProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobProgressResponse), nil
	}
}

func (c *DliClient) ShowSqlJobProgressInvoker(request *model.ShowSqlJobProgressRequest) *ShowSqlJobProgressInvoker {
	requestDef := GenReqDefForShowSqlJobProgress()
	return &ShowSqlJobProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DliClient) ShowSqlJobStatus(request *model.ShowSqlJobStatusRequest) (*model.ShowSqlJobStatusResponse, error) {
	requestDef := GenReqDefForShowSqlJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobStatusResponse), nil
	}
}

func (c *DliClient) ShowSqlJobStatusInvoker(request *model.ShowSqlJobStatusRequest) *ShowSqlJobStatusInvoker {
	requestDef := GenReqDefForShowSqlJobStatus()
	return &ShowSqlJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
