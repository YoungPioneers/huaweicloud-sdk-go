package listeners

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

type LoadBalancerID struct {
	ID string `json:"id"`
}

// Listener is the primary load balancing configuration object that specifies
// the loadbalancer and port on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type Listener struct {
	// The unique ID for the Listener.
	ID string `json:"id"`

	// Owner of the Listener.
	TenantID string `json:"tenant_id"`

	// Human-readable name for the Listener. Does not have to be unique.
	Name string `json:"name"`

	// Human-readable description for the Listener.
	Description string `json:"description"`

	// The protocol to loadbalance. A valid value is TCP, HTTP, or HTTPS.
	Protocol string `json:"protocol"`

	// The port on which to listen to client traffic that is associated with the
	// Loadbalancer. A valid value is from 0 to 65535.
	ProtocolPort int `json:"protocol_port"`

	// The UUID of default pool. Must have compatible protocol with listener.
	DefaultPoolID *string `json:"default_pool_id"`

	// A list of load balancer IDs.
	Loadbalancers []LoadBalancerID `json:"loadbalancers"`

	// The maximum number of connections allowed for the Loadbalancer.
	// Default is -1, meaning no limit.
	ConnLimit int `json:"connection_limit"`

	// The list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef *string `json:"default_tls_container_ref"`

	// The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// Pools are the pools which are part of this listener.
	//Pools []pools.Pool `json:"pools"`

	// tags for Listener
	Tags []string `json:"tags"`
	// client ca cert id
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref"`
	// CreatedAt is the time at which listeners has been created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is the time at which listeners has been created.
	UpdatedAt string `json:"updated_at"`

}

// ListenerPage is the page returned by a pager when traversing over a
// collection of listeners.
type ListenerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of listeners has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r ListenerPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"listeners_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a ListenerPage struct is empty.
func (r ListenerPage) IsEmpty() (bool, error) {
	is, err := ExtractListeners(r)
	return len(is) == 0, err
}

// ExtractListeners accepts a Page struct, specifically a ListenerPage struct,
// and extracts the elements into a slice of Listener structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractListeners(r pagination.Page) ([]Listener, error) {
	var s struct {
		Listeners []Listener `json:"listeners"`
	}
	err := (r.(ListenerPage)).ExtractInto(&s)
	return s.Listeners, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a listener.
func (r commonResult) Extract() (*Listener, error) {
	var s struct {
		Listener *Listener `json:"listener"`
	}
	err := r.ExtractInto(&s)
	return s.Listener, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Listener.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Listener.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
