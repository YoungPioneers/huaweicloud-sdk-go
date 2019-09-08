package job

import "github.com/YoungPioneers/huaweicloud-sdk-go"

// Querying the task status based on its job ID.
func GetJobResult(client *gophercloud.ServiceClient, id string) (JobResult, error) {
	var r JobExecResult
	url := jobURL(client, id)

	_, err := client.Get(url, &r.Body, nil)
	if err != nil {
		return JobResult{}, err
	}

	return r.ExtractJobResult()
}
