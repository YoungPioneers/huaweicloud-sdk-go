package recordsets

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRecordSetListQuery() (string, error)
}

type ListByZoneOptsBuilder interface {
	ToRecordSetListByZoneQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	// Integer value for the limit of values to return.
	Limit  int `q:"limit"`
	Offset int `q:"offset"`
	// UUID of the recordset at which you want to set a marker.
	Marker   string `q:"marker"`
	Name     string `q:"name"`
	SortDir  string `q:"sort_dir"`
	SortKey  string `q:"sort_key"`
	Status   string `q:"status"`
	Type     string `q:"type"`
	Id       string `q:"id"`
	ZoneType string `q:"zone_type"`
	Records  string `q:"records"`
}

// ToRecordSetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRecordSetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type ListByZoneOpts struct {
	// Integer value for the limit of values to return.
	Limit  int `q:"limit"`
	Offset int `q:"offset"`
	// UUID of the recordset at which you want to set a marker.
	Marker  string `q:"marker"`
	SortDir string `q:"sort_dir"`
	SortKey string `q:"sort_key"`
}

// ToRecordSetListByZoneQuery formats a ListOpts into a query string.
func (opts ListByZoneOpts) ToRecordSetListByZoneQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListByZone implements the recordset list request.
func ListByZone(client *gophercloud.ServiceClient, zoneID string, opts ListByZoneOptsBuilder) pagination.Pager {
	url := baseURL(client, zoneID)
	if opts != nil {
		query, err := opts.ToRecordSetListByZoneQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RecordSetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// List implements the recordset list request.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToRecordSetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RecordSetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get implements the recordset Get request.
func Get(client *gophercloud.ServiceClient, zoneID string, rrsetID string) (r GetResult) {
	_, r.Err = client.Get(rrsetURL(client, zoneID, rrsetID), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToRecordSetCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies the base attributes that may be used to create a
// RecordSet.
type CreateOpts struct {
	// Name is the name of the RecordSet.
	Name string `json:"name" required:"true"`

	// Description is a description of the RecordSet.
	Description string `json:"description,omitempty"`

	// Records are the DNS records of the RecordSet.
	Records []string `json:"records,omitempty"`

	// TTL is the time to live of the RecordSet.
	TTL int `json:"ttl,omitempty"`

	// Type is the RRTYPE of the RecordSet.
	Type string `json:"type,omitempty"`
}

// ToRecordSetCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToRecordSetCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create creates a recordset in a given zone.
func Create(client *gophercloud.ServiceClient, zoneID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRecordSetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client, zoneID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToRecordSetUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the base attributes that may be updated on an existing
// RecordSet.
type UpdateOpts struct {
	// Description is a description of the RecordSet.
	Description string `json:"description,omitempty"`

	// TTL is the time to live of the RecordSet.
	TTL int `json:"ttl,omitempty"`

	// Records are the DNS records of the RecordSet.
	Records []string `json:"records,omitempty"`
}

// ToRecordSetUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToRecordSetUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.TTL > 0 {
		b["ttl"] = opts.TTL
	} else {
		b["ttl"] = nil
	}

	return b, nil
}

// Update updates a recordset in a given zone
func Update(client *gophercloud.ServiceClient, zoneID string, rrsetID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRecordSetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(rrsetURL(client, zoneID, rrsetID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete removes an existing RecordSet.
func Delete(client *gophercloud.ServiceClient, zoneID string, rrsetID string) (r DeleteResult) {
	_, r.Err = client.Delete(rrsetURL(client, zoneID, rrsetID), &gophercloud.RequestOpts{
		OkCodes: []int{202}, JSONResponse: &r.Body,
	})
	return
}
