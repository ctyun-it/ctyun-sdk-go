package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchRebootInstanceApi 重启多台云主机
// https://www.ctyun.cn/document/10026730/10106398
type EcsBatchRebootInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchRebootInstanceApi(client *ctyunsdk.CtyunClient) *EcsBatchRebootInstancesApi {
	return &EcsBatchRebootInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-reboot-instances",
		},
	}
}

func (this *EcsBatchRebootInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchRebootInstanceRequest) (*EcsBatchRebootInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	_, err := builder.WriteJson(&ecsBatchRebootInstanceRealRequest{
		RegionID:       req.RegionID,
		InstanceIDList: req.InstanceIDList,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBatchRebootInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var jobIDListResp []EcsBatchRebootInstanceJobIDListRequest
	for _, info := range realResponse.JobIDList {
		jobIDListResp = append(jobIDListResp, EcsBatchRebootInstanceJobIDListRequest{
			JobID:      info.JobID,
			InstanceID: info.InstanceID,
		})
	}
	return &EcsBatchRebootInstanceResponse{
		JobIDList: jobIDListResp,
	}, nil
}

type ecsBatchRebootInstanceRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type EcsBatchRebootInstanceRealResponse struct {
	JobIDList []struct {
		JobID      string `json:"jobId"`
		InstanceID string `json:"instanceId"`
	} `json:"jobIDList"`
}

type EcsBatchRebootInstanceRequest struct {
	RegionID       string
	InstanceIDList string
}
type EcsBatchRebootInstanceJobIDListRequest struct {
	JobID      string
	InstanceID string
}

type EcsBatchRebootInstanceResponse struct {
	JobIDList []EcsBatchRebootInstanceJobIDListRequest
}
