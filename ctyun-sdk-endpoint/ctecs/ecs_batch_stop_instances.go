package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchStopInstancesApi
type EcsBatchStopInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchStopInstancesApi(client *ctyunsdk.CtyunClient) *EcsBatchStopInstancesApi {
	return &EcsBatchStopInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-stop-instances",
		},
	}
}

func (this *EcsBatchStopInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchStopInstancesRequest) (*EcsBatchStopInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBatchStopInstancesRealRequest{
		RegionID:       req.RegionID,
		InstanceIDList: req.InstanceIDList,
		Force:          req.Force,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBatchStopInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var jobIDList []EcsBatchStopInstancesJobIDListResponse
	for _, res := range realResponse.JobIDList {
		jobIDList = append(jobIDList, EcsBatchStopInstancesJobIDListResponse{
			JobID:      res.JobID,
			InstanceID: res.InstanceID,
		})
	}

	return &EcsBatchStopInstancesResponse{
		JobIDList: jobIDList,
	}, nil
}

type EcsBatchStopInstancesRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
	Force          bool   `json:"force,omitempty"`
}

type EcsBatchStopInstancesRequest struct {
	RegionID       string
	InstanceIDList string
	Force          bool
}

type EcsBatchStopInstancesJobIDListRealResponse struct {
	JobID      string `json:"jobID,omitempty"`
	InstanceID string `json:"instanceID,omitempty"`
}

type EcsBatchStopInstancesRealResponse struct {
	JobIDList []EcsBatchStopInstancesJobIDListRealResponse `json:"jobIDList,omitempty"`
}

type EcsBatchStopInstancesJobIDListResponse struct {
	JobID      string
	InstanceID string
}

type EcsBatchStopInstancesResponse struct {
	JobIDList []EcsBatchStopInstancesJobIDListResponse
}
