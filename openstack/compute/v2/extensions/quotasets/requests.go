package quotasets

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

// GetLimits returns limits public data about a previously created QuotaSet.
func GetLimits(client *gophercloud.ServiceClient) GetLimitResult {
	var res GetLimitResult
	_, res.Err = client.Get(getLimitURL(client), &res.Body, nil)
	return res
}

// Get returns public data about a previously created QuotaSet.
func Get(client *gophercloud.ServiceClient, tenantID string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, tenantID), &res.Body, nil)
	return res
}

// GetDetail returns detailed public data about a previously created QuotaSet.
func GetDetail(client *gophercloud.ServiceClient, tenantID string) GetDetailResult {
	var res GetDetailResult
	_, res.Err = client.Get(getDetailURL(client, tenantID), &res.Body, nil)
	return res
}

func GetDefault(client *gophercloud.ServiceClient, ProjectID string) GetDefaultResult {
	var res GetDefaultResult
	_, res.Err = client.Get(getDefaultURL(client, ProjectID), &res.Body, nil)
	return res
}

// Updates the quotas for the given tenantID and returns the new QuotaSet.
func Update(client *gophercloud.ServiceClient, tenantID string, opts UpdateOptsBuilder) (res UpdateResult) {
	reqBody, err := opts.ToComputeQuotaUpdateMap()
	if err != nil {
		res.Err = err
		return
	}

	_, res.Err = client.Put(updateURL(client, tenantID), reqBody, &res.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return res
}

// Resets the quotas for the given tenant to their default values.
func Delete(client *gophercloud.ServiceClient, tenantID string) (res DeleteResult) {
	_, res.Err = client.Delete(deleteURL(client, tenantID), nil)
	return
}

// Options for Updating the quotas of a Tenant.
// All int-values are pointers so they can be nil if they are not needed.
// You can use gopercloud.IntToPointer() for convenience
type UpdateOpts struct {
	// FixedIPs is number of fixed ips alloted this quota_set.
	FixedIPs *int `json:"fixed_ips,omitempty"`

	// FloatingIPs is number of floating ips alloted this quota_set.
	FloatingIPs *int `json:"floating_ips,omitempty"`

	// InjectedFileContentBytes is content bytes allowed for each injected file.
	InjectedFileContentBytes *int `json:"injected_file_content_bytes,omitempty"`

	// InjectedFilePathBytes is allowed bytes for each injected file path.
	InjectedFilePathBytes *int `json:"injected_file_path_bytes,omitempty"`

	// InjectedFiles is injected files allowed for each project.
	InjectedFiles *int `json:"injected_files,omitempty"`

	// KeyPairs is number of ssh keypairs.
	KeyPairs *int `json:"key_pairs,omitempty"`

	// MetadataItems is number of metadata items allowed for each instance.
	MetadataItems *int `json:"metadata_items,omitempty"`

	// RAM is megabytes allowed for each instance.
	RAM *int `json:"ram,omitempty"`

	// SecurityGroupRules is rules allowed for each security group.
	SecurityGroupRules *int `json:"security_group_rules,omitempty"`

	// SecurityGroups security groups allowed for each project.
	SecurityGroups *int `json:"security_groups,omitempty"`

	// Cores is number of instance cores allowed for each project.
	Cores *int `json:"cores,omitempty"`

	// Instances is number of instances allowed for each project.
	Instances *int `json:"instances,omitempty"`

	// Number of ServerGroups allowed for the project.
	ServerGroups *int `json:"server_groups,omitempty"`

	// Max number of Members for each ServerGroup.
	ServerGroupMembers *int `json:"server_group_members,omitempty"`

	// Force will update the quotaset even if the quota has already been used
	// and the reserved quota exceeds the new quota.
	Force bool `json:"force,omitempty"`
}

// UpdateOptsBuilder enables extensins to add parameters to the update request.
type UpdateOptsBuilder interface {
	// Extra specific name to prevent collisions with interfaces for other quotas
	// (e.g. neutron)
	ToComputeQuotaUpdateMap() (map[string]interface{}, error)
}

// ToComputeQuotaUpdateMap builds the update options into a serializable
// format.
func (opts UpdateOpts) ToComputeQuotaUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "quota_set")
}
