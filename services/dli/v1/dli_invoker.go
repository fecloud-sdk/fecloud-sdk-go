package v1

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dli/v1/model"
)

type AssociateQueueToEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateQueueToEnhancedConnectionInvoker) Invoke() (*model.AssociateQueueToEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateQueueToEnhancedConnectionResponse), nil
	}
}

type CreateEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnhancedConnectionInvoker) Invoke() (*model.CreateEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnhancedConnectionResponse), nil
	}
}

type CreateGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalVariableInvoker) Invoke() (*model.CreateGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalVariableResponse), nil
	}
}

type CreateQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQueueInvoker) Invoke() (*model.CreateQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueueResponse), nil
	}
}

type DeleteEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnhancedConnectionInvoker) Invoke() (*model.DeleteEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnhancedConnectionResponse), nil
	}
}

type DeleteGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalVariableInvoker) Invoke() (*model.DeleteGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalVariableResponse), nil
	}
}

type DeleteQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueueInvoker) Invoke() (*model.DeleteQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueueResponse), nil
	}
}

type DisassociateQueueFromEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateQueueFromEnhancedConnectionInvoker) Invoke() (*model.DisassociateQueueFromEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateQueueFromEnhancedConnectionResponse), nil
	}
}

type ListEnhancedConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnhancedConnectionsInvoker) Invoke() (*model.ListEnhancedConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnhancedConnectionsResponse), nil
	}
}

type ListGlobalVariablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalVariablesInvoker) Invoke() (*model.ListGlobalVariablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalVariablesResponse), nil
	}
}

type ListQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuesInvoker) Invoke() (*model.ListQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuesResponse), nil
	}
}

type ShowEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnhancedConnectionInvoker) Invoke() (*model.ShowEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnhancedConnectionResponse), nil
	}
}

type ShowEnhancedConnectionPrivilegeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnhancedConnectionPrivilegeInvoker) Invoke() (*model.ShowEnhancedConnectionPrivilegeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnhancedConnectionPrivilegeResponse), nil
	}
}

type ShowQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueInvoker) Invoke() (*model.ShowQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueResponse), nil
	}
}

type UpdateEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnhancedConnectionInvoker) Invoke() (*model.UpdateEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnhancedConnectionResponse), nil
	}
}

type UpdateGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalVariableInvoker) Invoke() (*model.UpdateGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalVariableResponse), nil
	}
}

type CreateFlinkJarJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkJarJobInvoker) Invoke() (*model.CreateFlinkJarJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkJarJobResponse), nil
	}
}

type CreateFlinkSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSqlJobInvoker) Invoke() (*model.CreateFlinkSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSqlJobResponse), nil
	}
}

type CreateFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSqlJobTemplateInvoker) Invoke() (*model.CreateFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSqlJobTemplateResponse), nil
	}
}

type DeleteFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkJobInvoker) Invoke() (*model.DeleteFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkJobResponse), nil
	}
}

type DeleteFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkSqlJobTemplateInvoker) Invoke() (*model.DeleteFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkSqlJobTemplateResponse), nil
	}
}

type ListFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobsInvoker) Invoke() (*model.ListFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobsResponse), nil
	}
}

type ListFlinkSqlJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkSqlJobTemplatesInvoker) Invoke() (*model.ListFlinkSqlJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkSqlJobTemplatesResponse), nil
	}
}

type ShowFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobInvoker) Invoke() (*model.ShowFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobResponse), nil
	}
}

type UpdateFlinkJarJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkJarJobInvoker) Invoke() (*model.UpdateFlinkJarJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkJarJobResponse), nil
	}
}

type UpdateFlinkSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkSqlJobInvoker) Invoke() (*model.UpdateFlinkSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkSqlJobResponse), nil
	}
}

type UpdateFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkSqlJobTemplateInvoker) Invoke() (*model.UpdateFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkSqlJobTemplateResponse), nil
	}
}

type CancelSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkJobInvoker) Invoke() (*model.CancelSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkJobResponse), nil
	}
}

type CreateSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSparkJobInvoker) Invoke() (*model.CreateSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSparkJobResponse), nil
	}
}

type ListSparkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkJobsInvoker) Invoke() (*model.ListSparkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkJobsResponse), nil
	}
}

type ShowSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobInvoker) Invoke() (*model.ShowSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobResponse), nil
	}
}

type ShowSparkJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobStatusInvoker) Invoke() (*model.ShowSparkJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobStatusResponse), nil
	}
}

type CancelSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSqlJobInvoker) Invoke() (*model.CancelSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSqlJobResponse), nil
	}
}

type CheckSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckSqlInvoker) Invoke() (*model.CheckSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckSqlResponse), nil
	}
}

type CreateSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlJobInvoker) Invoke() (*model.CreateSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlJobResponse), nil
	}
}

type ListSqlJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlJobsInvoker) Invoke() (*model.ListSqlJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlJobsResponse), nil
	}
}

type ShowSqlJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobDetailInvoker) Invoke() (*model.ShowSqlJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobDetailResponse), nil
	}
}

type ShowSqlJobProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobProgressInvoker) Invoke() (*model.ShowSqlJobProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobProgressResponse), nil
	}
}

type ShowSqlJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobStatusInvoker) Invoke() (*model.ShowSqlJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobStatusResponse), nil
	}
}
