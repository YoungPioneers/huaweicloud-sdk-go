package extensions

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

// Extension is a struct that represents an OpenStack extension.
type Extension struct {
	Updated     string             `json:"updated"`
	Name        string             `json:"name"`
	Links       []gophercloud.Link `json:"links"`
	Namespace   string             `json:"namespace"`
	Alias       string             `json:"alias"`
	Description string             `json:"description"`
}

// ExtensionPage is the page returned by a pager when traversing over a collection of extensions.
type ExtensionPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether an ExtensionPage struct is empty.
func (r ExtensionPage) IsEmpty() (bool, error) {
	is, err := ExtractExtensions(r)
	return len(is) == 0, err
}

// ExtractExtensions accepts a Page struct, specifically an ExtensionPage
// struct, and extracts the elements into a slice of Extension structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractExtensions(r pagination.Page) ([]Extension, error) {
	var s struct {
		Extensions []Extension `json:"extensions"`
	}
	err := (r.(ExtensionPage)).ExtractInto(&s)
	return s.Extensions, err
}
