package servers

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

type serverResult struct {
	gophercloud.Result
}

// Extract interprets any serverResult as a Server, if possible.
func (r serverResult) Extract() (*Server, error) {
	var s Server
	err := r.ExtractInto(&s)
	return &s, err
}

func (r serverResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "server")
}

func ExtractServersInto(r pagination.Page, v interface{}) error {
	return r.(ServerPage).Result.ExtractIntoSlicePtr(v, "servers")
}

// CreateResult is the response from a Create operation. Call its Extract
// method to interpret it as a Server.
type CreateResult struct {
	serverResult
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as a Server.
type GetResult struct {
	serverResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Server.
type UpdateResult struct {
	serverResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// RebuildResult is the response from a Rebuild operation. Call its Extract
// method to interpret it as a Server.
type RebuildResult struct {
	serverResult
}

// ActionResult represents the result of server action operations, like reboot.
// Call its ExtractErr method to determine if the action succeeded or failed.
type ActionResult struct {
	gophercloud.ErrResult
}

// RescueResult is the response from a Rescue operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type RescueResult struct {
	ActionResult
}

// CreateImageResult is the response from a CreateImage operation. Call its
// ExtractImageID method to retrieve the ID of the newly created image.
type CreateImageResult struct {
	gophercloud.Result
}

// GetPasswordResult represent the result of a get os-server-password operation.
// Call its ExtractPassword method to retrieve the password.
type GetPasswordResult struct {
	gophercloud.Result
}

// ExtractPassword gets the encrypted password.
// If privateKey != nil the password is decrypted with the private key.
// If privateKey == nil the encrypted password is returned and can be decrypted
// with:
//   echo '<pwd>' | base64 -D | openssl rsautl -decrypt -inkey <private_key>
func (r GetPasswordResult) ExtractPassword(privateKey *rsa.PrivateKey) (string, error) {
	var s struct {
		Password string `json:"password"`
	}
	err := r.ExtractInto(&s)
	if err == nil && privateKey != nil && s.Password != "" {
		return decryptPassword(s.Password, privateKey)
	}
	return s.Password, err
}

func decryptPassword(encryptedPassword string, privateKey *rsa.PrivateKey) (string, error) {
	b64EncryptedPassword := make([]byte, base64.StdEncoding.DecodedLen(len(encryptedPassword)))

	n, err := base64.StdEncoding.Decode(b64EncryptedPassword, []byte(encryptedPassword))
	if err != nil {
		return "", fmt.Errorf("Failed to base64 decode encrypted password: %s", err)
	}
	password, err := rsa.DecryptPKCS1v15(nil, privateKey, b64EncryptedPassword[0:n])
	if err != nil {
		return "", fmt.Errorf("Failed to decrypt password: %s", err)
	}

	return string(password), nil
}

// ExtractImageID gets the ID of the newly created server image from the header.
func (r CreateImageResult) ExtractImageID() (string, error) {
	if r.Err != nil {
		return "", r.Err
	}
	// Get the image id from the header
	u, err := url.ParseRequestURI(r.Header.Get("Location"))
	if err != nil {
		return "", err
	}
	imageID := path.Base(u.Path)
	if imageID == "." || imageID == "/" {
		return "", fmt.Errorf("Failed to parse the ID of newly created image: %s", u)
	}
	return imageID, nil
}

// Extract interprets any RescueResult as an AdminPass, if possible.
func (r RescueResult) Extract() (string, error) {
	var s struct {
		AdminPass string `json:"adminPass"`
	}
	err := r.ExtractInto(&s)
	return s.AdminPass, err
}

// Server represents a server/instance in the OpenStack cloud.
type Server struct {

	//volume attached
	VolumeAttached []map[string]interface{} `json:"os-extended-volumes:volumes_attached"`

	//availbility zone
	AvailbiltyZone string `json:"OS-EXT-AZ:availability_zone"`

	// host
	Host string `json:"OS-EXT-SRV-ATTR:host"`

	// task state
	TaskState string `json:"OS-EXT-STS:task_state"`

	// power state
	PowerState int `json:"OS-EXT-STS:power_state"`

	//  vm state
	VMstate string `json:"OS-EXT-STS:vm_state"`

	// ID uniquely identifies this server amongst all other servers,
	// including those not accessible to the current tenant.
	ID string `json:"id"`

	// TenantID identifies the tenant owning this server resource.
	TenantID string `json:"tenant_id"`

	// UserID uniquely identifies the user account owning the tenant.
	UserID string `json:"user_id"`

	// Name contains the human-readable name for the server.
	Name string `json:"name"`

	// Updated and Created contain ISO-8601 timestamps of when the state of the
	// server last changed, and when it was created.
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`

	// HostID is the host where the server is located in the cloud.
	HostID string `json:"hostid"`

	// Status contains the current operational status of the server,
	// such as IN_PROGRESS or ACTIVE.
	Status string `json:"status"`

	// Progress ranges from 0..100.
	// A request made against the server completes only once Progress reaches 100.
	Progress int `json:"progress"`

	// AccessIPv4 and AccessIPv6 contain the IP addresses of the server,
	// suitable for remote access for administration.
	AccessIPv4 string `json:"accessIPv4"`
	AccessIPv6 string `json:"accessIPv6"`

	// Image refers to a JSON object, which itself indicates the OS image used to
	// deploy the server.
	Image map[string]interface{} `json:"image"`

	// Flavor refers to a JSON object, which itself indicates the hardware
	// configuration of the deployed server.
	Flavor map[string]interface{} `json:"flavor"`

	// Addresses includes a list of all IP addresses assigned to the server,
	// keyed by pool.
	Addresses map[string]interface{} `json:"addresses"`

	// Metadata includes a list of all user-specified key-value pairs attached
	// to the server.
	Metadata map[string]string `json:"metadata"`

	// Links includes HTTP references to the itself, useful for passing along to
	// other APIs that might want a server reference.
	Links []interface{} `json:"links"`

	// KeyName indicates which public key was injected into the server on launch.
	KeyName string `json:"key_name"`

	// AdminPass will generally be empty ("").  However, it will contain the
	// administrative password chosen when provisioning a new server without a
	// set AdminPass setting in the first place.
	// Note that this is the ONLY time this field will be valid.
	AdminPass string `json:"adminPass"`

	// SecurityGroups includes the security groups that this instance has applied
	// to it.
	SecurityGroups []map[string]interface{} `json:"security_groups"`

	// Fault contains failure information about a server.
	Fault Fault `json:"fault"`

	//DCF
	DiskConfig string `json:"OS-DCF:diskConfig"`
	//server tags
	Tags []string `json:"tags"`

	//host status
	HostStatus string `json:"host_status"`
	//host description
	Description string `json:"description"`
	//bool type whether server locked
	Locked *bool `json:"locked"`
	// server hostname
	Hostname string `json:"OS-EXT-SRV-ATTR:hostname"`
	//server kernel_id
	KernelID string `json:"OS-EXT-SRV-ATTR:kernel_id"`
	// server ramdisk_id
	RamdiskID string `json:"OS-EXT-SRV-ATTR:ramdisk_id"`
	//launch_index
	LaunchIndex *int `json:"OS-EXT-SRV-ATTR:launch_index"`
	//user_data
	UserData string `json:"OS-EXT-SRV-ATTR:user_data"`
	//reservation_id
	ReservationID string `json:"OS-EXT-SRV-ATTR:reservation_id"`
	//root_device_name
	RootDeviceName string `json:"OS-EXT-SRV-ATTR:root_device_name"`
	//hypervisor_hostname
	HypervisorHostname string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	//instance_name
	InstanceName string `json:"OS-EXT-SRV-ATTR:instance_name"`

	//instance launched time
	LaunchedAt string `json:"OS-SRV-USG:launched_at"`
	//instance terminated time
	TerminatedAt string `json:"OS-SRV-USG:terminated_at"`
	//Reserved attribute
	ConfigDrive string `json:"config_drive"`
	//Reserved attribute
	EvsOpts int `json:"evs_opts"`
	//Reserved attribute
	HyperThreadAffinity string `json:"hyperThreadAffinity"`
	//Reserved attribute
	NumaOpts int `json:"numa_opts"`
	//Reserved attribute
	VcpuAffinity []int `json:"vcpuAffinity"`
}

//ListBrief接口返回的结构
type ServerBrief struct {
	ID    string
	Name  string
	Links []interface{}
}

type Fault struct {
	Code    int       `json:"code"`
	Created time.Time `json:"created"`
	Details string    `json:"details"`
	Message string    `json:"message"`
}

func (r *Server) UnmarshalJSON(b []byte) error {
	type tmp Server
	var s struct {
		tmp
		Image interface{} `json:"image"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Server(s.tmp)

	switch t := s.Image.(type) {
	case map[string]interface{}:
		r.Image = t
	case string:
		switch t {
		case "":
			r.Image = nil
		}
	}

	return err
}

// ServerPage abstracts the raw results of making a List() request against
// the API. As OpenStack extensions may freely alter the response bodies of
// structures returned to the client, you may only safely access the data
// provided through the ExtractServers call.
type ServerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Server results.
func (r ServerPage) IsEmpty() (bool, error) {
	s, err := ExtractServers(r)
	return len(s) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r ServerPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"servers_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractServers interprets the results of a single page from a List() call,
// producing a slice of Server entities.
func ExtractServers(r pagination.Page) ([]Server, error) {
	var s []Server
	err := ExtractServersInto(r, &s)
	return s, err
}

//对应ListBrief接口
func ExtractServerBriefs(r pagination.Page) ([]ServerBrief, error) {
	var sb []ServerBrief
	err := ExtractServersInto(r, &sb)
	return sb, err
}

// MetadataResult contains the result of a call for (potentially) multiple
// key-value pairs. Call its Extract method to interpret it as a
// map[string]interface.
type MetadataResult struct {
	gophercloud.Result
}

// GetMetadataResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type GetMetadataResult struct {
	MetadataResult
}

type GetConsoleLogResult struct {
	MetadataResult
}

// ResetMetadataResult contains the result of a Reset operation. Call its
// Extract method to interpret it as a map[string]interface.
type ResetMetadataResult struct {
	MetadataResult
}

// UpdateMetadataResult contains the result of an Update operation. Call its
// Extract method to interpret it as a map[string]interface.
type UpdateMetadataResult struct {
	MetadataResult
}

// MetadatumResult contains the result of a call for individual a single
// key-value pair.
type MetadatumResult struct {
	gophercloud.Result
}

// GetMetadatumResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type GetMetadatumResult struct {
	MetadatumResult
}

// CreateMetadatumResult contains the result of a Create operation. Call its
// Extract method to interpret it as a map[string]interface.
type CreateMetadatumResult struct {
	MetadatumResult
}

// DeleteMetadatumResult contains the result of a Delete operation. Call its
// ExtractErr method to determine if the call succeeded or failed.
type DeleteMetadatumResult struct {
	gophercloud.ErrResult
}

// Extract interprets any MetadataResult as a Metadata, if possible.
func (r MetadataResult) Extract() (map[string]string, error) {
	var s struct {
		Metadata map[string]string `json:"metadata"`
	}
	err := r.ExtractInto(&s)
	return s.Metadata, err
}

// Extract interprets any MetadatumResult as a Metadatum, if possible.
func (r MetadatumResult) Extract() (map[string]string, error) {
	var s struct {
		Metadatum map[string]string `json:"meta"`
	}
	err := r.ExtractInto(&s)
	return s.Metadatum, err
}

func (r GetConsoleLogResult) Extract() (string, error) {
	var s struct{
		Output string `json:"output"`
	}
	err := r.ExtractInto(&s)
	return s.Output, err

}

// Address represents an IP address.
type Address struct {
	Version int    `json:"version"`
	Address string `json:"addr"`
}

// AddressPage abstracts the raw results of making a ListAddresses() request
// against the API. As OpenStack extensions may freely alter the response bodies
// of structures returned to the client, you may only safely access the data
// provided through the ExtractAddresses call.
type AddressPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if an AddressPage contains no networks.
func (r AddressPage) IsEmpty() (bool, error) {
	addresses, err := ExtractAddresses(r)
	return len(addresses) == 0, err
}

// ExtractAddresses interprets the results of a single page from a
// ListAddresses() call, producing a map of addresses.
func ExtractAddresses(r pagination.Page) (map[string][]Address, error) {
	var s struct {
		Addresses map[string][]Address `json:"addresses"`
	}
	err := (r.(AddressPage)).ExtractInto(&s)
	return s.Addresses, err
}

// NetworkAddressPage abstracts the raw results of making a
// ListAddressesByNetwork() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures
// returned to the client, you may only safely access the data provided through
// the ExtractAddresses call.
type NetworkAddressPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a NetworkAddressPage contains no addresses.
func (r NetworkAddressPage) IsEmpty() (bool, error) {
	addresses, err := ExtractNetworkAddresses(r)
	return len(addresses) == 0, err
}

// ExtractNetworkAddresses interprets the results of a single page from a
// ListAddressesByNetwork() call, producing a slice of addresses.
func ExtractNetworkAddresses(r pagination.Page) ([]Address, error) {
	var s map[string][]Address
	err := (r.(NetworkAddressPage)).ExtractInto(&s)
	if err != nil {
		return nil, err
	}

	var key string
	for k := range s {
		key = k
	}

	return s[key], err
}

type InstanceActions struct {
	InstanceUUID string `json:"instance_uuid"`
	UserID       string `json:"user_id"`
	StartTime    string `json:"start_time"`
	RequestID    string `json:"request_id"`
	Action       string `json:"action"`
	Message      string `json:"message"`
	ProjectID    string `json:"project_id"`
	UpdatedAt    string `json:"updated_at"`
}

type InstanceAction struct {
	InstanceUUID string   `json:"instance_uuid"`
	UserID       string   `json:"user_id"`
	StartTime    string   `json:"start_time"`
	RequestID    string   `json:"request_id"`
	Action       string   `json:"action"`
	Message      string   `json:"message"`
	ProjectID    string   `json:"project_id"`
	Events       []Events `json:"events"`
}

type Events struct {
	FinishTime string `json:"finish_time"`
	StartTime  string `json:"start_time"`
	Traceback  string `json:"traceback"`
	Event      string `json:"event"`
	Result     string `json:"result"`
}

type InstanceActionResult struct {
	gophercloud.Result
}

func (r InstanceActionResult) ExtractInstanceActionsListResult() (*[]InstanceActions, error) {
	//var ia InstanceActions
	var s struct {
		Ia []InstanceActions `json:"instanceActions"`
	}

	err := r.ExtractInto(&s)
	return &s.Ia, err
}

func (r InstanceActionResult) ExtractInstanceActionsResult() (*InstanceAction, error) {
	var s struct {
		Ia InstanceAction `json:"instanceAction"`
	}

	err := r.ExtractInto(&s)
	return &s.Ia, err
}
