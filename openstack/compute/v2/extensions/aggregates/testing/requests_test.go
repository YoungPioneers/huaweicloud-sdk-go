package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/compute/v2/extensions/aggregates"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func TestListAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	pages := 0
	err := aggregates.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := aggregates.ExtractAggregates(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 aggregates, got %d", len(actual))
		}
		th.CheckDeepEquals(t, FirstFakeAggregate, actual[0])
		th.CheckDeepEquals(t, SecondFakeAggregate, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestCreateAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	expected := CreatedAggregate

	opts := aggregates.CreateOpts{
		Name:             "name",
		AvailabilityZone: "london",
	}

	actual, err := aggregates.Create(client.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestDeleteAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := aggregates.Delete(client.ServiceClient(), AggregateIDtoDelete).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetAggregates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	expected := SecondFakeAggregate

	actual, err := aggregates.Get(client.ServiceClient(), AggregateIDtoGet).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestUpdateAggregate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	expected := UpdatedAggregate

	opts := aggregates.UpdateOpts{
		Name:             "test-aggregates2",
		AvailabilityZone: "nova2",
	}

	actual, err := aggregates.Update(client.ServiceClient(), expected.ID, opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}
