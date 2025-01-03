package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchRebuildInstancesApi 重装多台云主机
// https://www.ctyun.cn/document/10026730/10106438
type EcsBatchRebuildInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchRebuildInstancesApi(client *ctyunsdk.CtyunClient) *EcsBatchRebuildInstancesApi {
	return &EcsBatchRebuildInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-rebuild-instances",
		},
	}
}

func (this *EcsBatchRebuildInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchRebuildInstancesRequest) (*EcsBatchRebuildInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	var rebuildInfoReq []ecsBatchRebuildInstancesRebuildInfoRealRequest
	for _, request := range req.RebuildInfo {
		rebuildInfoReq = append(rebuildInfoReq, ecsBatchRebuildInstancesRebuildInfoRealRequest{
			InstanceID:     request.InstanceID,
			Password:       request.Password,
			ImageID:        request.ImageID,
			UserData:       request.UserData,
			InstanceName:   request.InstanceName,
			MonitorService: request.MonitorService,
		})
	}
	_, err := builder.WriteJson(&ecsBatchRebuildInstancesRealRequest{
		RegionID:    req.RegionID,
		RebuildInfo: rebuildInfoReq,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ecsBatchRebuildInstancesJobIDListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var jobIDListResp []EcsBatchRebuildInstancesJobIDListResponse
	for _, info := range realResponse.JobIDList {
		jobIDListResp = append(jobIDListResp, EcsBatchRebuildInstancesJobIDListResponse{
			JobID:      info.JobID,
			InstanceID: info.InstanceID,
		})
	}
	return &EcsBatchRebuildInstancesResponse{
		JobIDList: jobIDListResp,
	}, nil
}

type ecsBatchRebuildInstancesRebuildInfoRealRequest struct {
	InstanceID     *string `json:"instanceID,omitempty"`
	Password       *string `json:"password,omitempty"`
	ImageID        *string `json:"imageID,omitempty"`
	UserData       *string `json:"userData,omitempty"`
	InstanceName   *string `json:"instanceName,omitempty"`
	MonitorService *bool   `json:"monitorService,omitempty"`
}

type ecsBatchRebuildInstancesRealRequest struct {
	RegionID    *string                                          `json:"regionID"`
	RebuildInfo []ecsBatchRebuildInstancesRebuildInfoRealRequest `json:"rebuildInfo"`
}

type ecsBatchRebuildInstancesJobIDListRealResponse struct {
	JobIDList []struct {
		JobID      string `json:"jobID"`
		InstanceID string `json:"instanceID"`
	} `json:"jobIDList"`
}

type EcsBatchRebuildInstancesRequest struct {
	RegionID    *string
	RebuildInfo []EcsBatchRebuildInstancesRebuildInfoRequest
}

type EcsBatchRebuildInstancesRebuildInfoRequest struct {
	InstanceID     *string
	Password       *string
	ImageID        *string
	UserData       *string
	InstanceName   *string
	MonitorService *bool
}

type EcsBatchRebuildInstancesJobIDListResponse struct {
	JobID      string
	InstanceID string
}

type EcsBatchRebuildInstancesResponse struct {
	JobIDList []EcsBatchRebuildInstancesJobIDListResponse
}
