package snapshots

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	VolumeID    string                 `json:"volume_id" required:"true"`
	Description string                 `json:"display_description,omitempty"`
	Force       bool                   `json:"force,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Name        string                 `json:"display_name,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "snapshot")
}

// Create will create a new Snapshot based on the values in CreateOpts. To
// extract the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot
// object from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts hold options for listing Snapshots. It is passed to the
// snapshots.List function.
type ListOpts struct {
	Name     string `q:"display_name"`
	Status   string `q:"status"`
	VolumeID string `q:"volume_id"`
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Snapshots optionally limited by the conditions provided in
// ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

// UpdateMetadataOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateMetadataOptsBuilder interface {
	ToSnapshotUpdateMetadataMap() (map[string]interface{}, error)
}

// UpdateMetadataOpts contain options for updating an existing Snapshot. This
// object is passed to the snapshots.Update function. For more information
// about the parameters, see the Snapshot object.
type UpdateMetadataOpts struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ToSnapshotUpdateMetadataMap assembles a request body based on the contents of
// an UpdateMetadataOpts.
func (opts UpdateMetadataOpts) ToSnapshotUpdateMetadataMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// UpdateMetadata will update the Snapshot with provided information. To
// extract the updated Snapshot from the response, call the ExtractMetadata
// method on the UpdateMetadataResult.
func UpdateMetadata(client *gophercloud.ServiceClient, id string, opts UpdateMetadataOptsBuilder) (r UpdateMetadataResult) {
	b, err := opts.ToSnapshotUpdateMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateMetadataURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// IDFromName is a convienience function that returns a snapshot's ID given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractSnapshots(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		//return "", gophercloud.ErrResourceNotFound{Name: name, ResourceType: "snapshot"}

		message := fmt.Sprintf(gophercloud.CE_ResourceNotFoundMessage, "snapshot", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_ResourceNotFoundCode, message)
		return "", err
	case 1:
		return id, nil
	default:
		//return "", gophercloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "snapshot"}

		message := fmt.Sprintf(gophercloud.CE_MultipleResourcesFoundMessage, count, "snapshot", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MultipleResourcesFoundCode, message)
		return "", err
	}
}
