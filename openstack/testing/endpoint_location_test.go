package testing

import (
	"strings"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack"
	tokens2 "github.com/huaweicloud/huaweicloud-sdk-go/openstack/identity/v2/tokens"
	tokens3 "github.com/huaweicloud/huaweicloud-sdk-go/openstack/identity/v3/tokens"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
)

// Service catalog fixtures take too much vertical space!
var catalog2 = tokens2.ServiceCatalog{
	Entries: []tokens2.CatalogEntry{
		tokens2.CatalogEntry{
			Type: "same",
			Name: "same",
			Endpoints: []tokens2.Endpoint{
				tokens2.Endpoint{
					Region:      "same",
					PublicURL:   "https://public.correct.com/",
					InternalURL: "https://internal.correct.com/",
					AdminURL:    "https://admin.correct.com/",
				},
				tokens2.Endpoint{
					Region:    "different",
					PublicURL: "https://badregion.com/",
				},
			},
		},
		tokens2.CatalogEntry{
			Type: "same",
			Name: "different",
			Endpoints: []tokens2.Endpoint{
				tokens2.Endpoint{
					Region:    "same",
					PublicURL: "https://badname.com/",
				},
				tokens2.Endpoint{
					Region:    "different",
					PublicURL: "https://badname.com/+badregion",
				},
			},
		},
		tokens2.CatalogEntry{
			Type: "different",
			Name: "different",
			Endpoints: []tokens2.Endpoint{
				tokens2.Endpoint{
					Region:    "same",
					PublicURL: "https://badtype.com/+badname",
				},
				tokens2.Endpoint{
					Region:    "different",
					PublicURL: "https://badtype.com/+badregion+badname",
				},
			},
		},
	},
}

func TestV2EndpointExact(t *testing.T) {
	expectedURLs := map[gophercloud.Availability]string{
		gophercloud.AvailabilityPublic:   "https://public.correct.com/",
		gophercloud.AvailabilityAdmin:    "https://admin.correct.com/",
		gophercloud.AvailabilityInternal: "https://internal.correct.com/",
	}

	for availability, expected := range expectedURLs {
		actual, err := openstack.V2EndpointURL(&catalog2, gophercloud.EndpointOpts{
			Type:         "same",
			Name:         "same",
			Region:       "same",
			Availability: availability,
		})
		th.AssertNoErr(t, err)
		th.CheckEquals(t, expected, actual)
	}
}

func TestV2EndpointNone(t *testing.T) {
	_, actual := openstack.V2EndpointURL(&catalog2, gophercloud.EndpointOpts{
		Type:         "nope",
		Availability: gophercloud.AvailabilityPublic,
	})
	expected := gophercloud.NewSystemCommonError(gophercloud.CE_NoEndPointInCatalogCode, gophercloud.CE_NoEndPointInCatalogMessage)
	th.CheckEquals(t, expected.Error(), actual.Error())
}

func TestV2EndpointMultiple(t *testing.T) {
	_, err := openstack.V2EndpointURL(&catalog2, gophercloud.EndpointOpts{
		Type:         "same",
		Region:       "same",
		Availability: gophercloud.AvailabilityPublic,
	})
	if !strings.HasPrefix(err.Error(), "Discovered 2 matching endpoints:") {
		t.Errorf("Received unexpected error: %v", err)
	}
}

func TestV2EndpointBadAvailability(t *testing.T) {
	_, err := openstack.V2EndpointURL(&catalog2, gophercloud.EndpointOpts{
		Type:         "same",
		Name:         "same",
		Region:       "same",
		Availability: "wat",
	})
	th.CheckEquals(t, "Unexpected availability in endpoint query: wat", err.Error())
}

var catalog3 = tokens3.ServiceCatalog{
	Entries: []tokens3.CatalogEntry{
		tokens3.CatalogEntry{
			Type: "same",
			Name: "same",
			Endpoints: []tokens3.Endpoint{
				tokens3.Endpoint{
					ID:        "1",
					Region:    "same",
					Interface: "public",
					URL:       "https://public.correct.com/",
				},
				tokens3.Endpoint{
					ID:        "2",
					Region:    "same",
					Interface: "admin",
					URL:       "https://admin.correct.com/",
				},
				tokens3.Endpoint{
					ID:        "3",
					Region:    "same",
					Interface: "internal",
					URL:       "https://internal.correct.com/",
				},
				tokens3.Endpoint{
					ID:        "4",
					Region:    "different",
					Interface: "public",
					URL:       "https://badregion.com/",
				},
			},
		},
		tokens3.CatalogEntry{
			Type: "same",
			Name: "different",
			Endpoints: []tokens3.Endpoint{
				tokens3.Endpoint{
					ID:        "5",
					Region:    "same",
					Interface: "public",
					URL:       "https://badname.com/",
				},
				tokens3.Endpoint{
					ID:        "6",
					Region:    "different",
					Interface: "public",
					URL:       "https://badname.com/+badregion",
				},
			},
		},
		tokens3.CatalogEntry{
			Type: "different",
			Name: "different",
			Endpoints: []tokens3.Endpoint{
				tokens3.Endpoint{
					ID:        "7",
					Region:    "same",
					Interface: "public",
					URL:       "https://badtype.com/+badname",
				},
				tokens3.Endpoint{
					ID:        "8",
					Region:    "different",
					Interface: "public",
					URL:       "https://badtype.com/+badregion+badname",
				},
			},
		},
	},
}

func TestV3EndpointExact(t *testing.T) {
	expectedURLs := map[gophercloud.Availability]string{
		gophercloud.AvailabilityPublic:   "https://public.correct.com/",
		gophercloud.AvailabilityAdmin:    "https://admin.correct.com/",
		gophercloud.AvailabilityInternal: "https://internal.correct.com/",
	}

	for availability, expected := range expectedURLs {
		actual, err := openstack.V3EndpointURL(&catalog3, gophercloud.EndpointOpts{
			Type:         "same",
			Name:         "same",
			Region:       "same",
			Availability: availability,
		},nil)
		th.AssertNoErr(t, err)
		th.CheckEquals(t, expected, actual)
	}
}

func TestV3EndpointNone(t *testing.T) {
	_, actual := openstack.V3EndpointURL(&catalog3, gophercloud.EndpointOpts{
		Type:         "nope",
		Availability: gophercloud.AvailabilityPublic,
	},nil)
	expected := gophercloud.NewSystemCommonError(gophercloud.CE_NoEndPointInCatalogCode, gophercloud.CE_NoEndPointInCatalogMessage)
	th.CheckEquals(t, expected.Error(), actual.Error())
}

func TestV3EndpointMultiple(t *testing.T) {
	_, err := openstack.V3EndpointURL(&catalog3, gophercloud.EndpointOpts{
		Type:         "same",
		Region:       "same",
		Availability: gophercloud.AvailabilityPublic,
	},nil)
	if !strings.HasPrefix(err.Error(), "Discovered 2 matching endpoints:") {
		t.Errorf("Received unexpected error: %v", err)
	}
}

func TestV3EndpointBadAvailability(t *testing.T) {
	_, err := openstack.V3EndpointURL(&catalog3, gophercloud.EndpointOpts{
		Type:         "same",
		Name:         "same",
		Region:       "same",
		Availability: "wat",
	},nil)
	th.CheckEquals(t, "Unexpected availability in endpoint query: wat", err.Error())
}
