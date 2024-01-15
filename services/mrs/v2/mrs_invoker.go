package v2

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/mrs/v2/model"
)

type BatchDeleteJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) Invoke() (*model.BatchDeleteJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExecuteJobInvoker) Invoke() (*model.CreateExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExecuteJobResponse), nil
	}
}

type ShowAgencyMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyMappingInvoker) Invoke() (*model.ShowAgencyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyMappingResponse), nil
	}
}

type ShowJobExeListNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobExeListNewInvoker) Invoke() (*model.ShowJobExeListNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobExeListNewResponse), nil
	}
}

type ShowSingleJobExeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleJobExeInvoker) Invoke() (*model.ShowSingleJobExeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleJobExeResponse), nil
	}
}

type ShowSqlResultWithJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlResultWithJobInvoker) Invoke() (*model.ShowSqlResultWithJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlResultWithJobResponse), nil
	}
}

type StopJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInvoker) Invoke() (*model.StopJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobResponse), nil
	}
}

type UpdateAgencyMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgencyMappingInvoker) Invoke() (*model.UpdateAgencyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgencyMappingResponse), nil
	}
}

type ShowHdfsFileListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHdfsFileListInvoker) Invoke() (*model.ShowHdfsFileListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHdfsFileListResponse), nil
	}
}

type CancelSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSqlInvoker) Invoke() (*model.CancelSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSqlResponse), nil
	}
}

type ExecuteSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSqlInvoker) Invoke() (*model.ExecuteSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSqlResponse), nil
	}
}

type ShowSqlResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlResultInvoker) Invoke() (*model.ShowSqlResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlResultResponse), nil
	}
}
