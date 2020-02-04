package hactar

import "net/http"

type DiskInfoService interface {
	SendDiskInfo(diskInfo DiskInfo) (*http.Response, error)
}

type diskInfoService struct {
	client *Client
}

type DiskInfo struct {
	FreeSpace    string `json:"freeSpace"`
	TakenSpace   string `json:"takenSpace"`
	Node NodeInfo `json:"node"`
}

const (
	DiskInfoPath = "/diskinfo"
)

func (dis *diskInfoService) SendDiskInfo(diskInfo DiskInfo) (*http.Response, error) {
	request, err := dis.client.NewRequest(http.MethodPost, DiskInfoPath, diskInfo)

	if err != nil {
		return nil, err
	}

	response, err := dis.client.Do(request, nil)

	if err != nil {
		return nil, err
	}

	return response, err
}
