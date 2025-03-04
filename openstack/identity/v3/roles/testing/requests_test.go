package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/identity/v3/roles"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func TestListRoles(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListRolesSuccessfully(t)

	count := 0
	err := roles.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := roles.ExtractRoles(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedRolesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}

func TestListRolesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListRolesSuccessfully(t)

	allPages, err := roles.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := roles.ExtractRoles(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedRolesSlice, actual)
	th.AssertEquals(t, ExpectedRolesSlice[1].Extra["description"], "read-only support role")
}

func TestGetRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetRoleSuccessfully(t)

	actual, err := roles.Get(client.ServiceClient(), "9fe1d3").Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondRole, *actual)
}

func TestCreateRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateRoleSuccessfully(t)

	createOpts := roles.CreateOpts{
		Name:     "support",
		DomainID: "1789d1",
		Extra: map[string]interface{}{
			"description": "read-only support role",
		},
	}

	actual, err := roles.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondRole, *actual)
}

func TestUpdateRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateRoleSuccessfully(t)

	updateOpts := roles.UpdateOpts{
		Extra: map[string]interface{}{
			"description": "admin read-only support role",
		},
	}

	actual, err := roles.Update(client.ServiceClient(), "9fe1d3", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SecondRoleUpdated, *actual)
}

func TestDeleteRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteRoleSuccessfully(t)

	res := roles.Delete(client.ServiceClient(), "9fe1d3")
	th.AssertNoErr(t, res.Err)
}

func TestListAssignmentsSinglePage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListRoleAssignmentsSuccessfully(t)

	count := 0
	err := roles.ListAssignments(client.ServiceClient(), roles.ListAssignmentsOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := roles.ExtractRoleAssignments(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedRoleAssignmentsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}

func TestAssign(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssignSuccessfully(t)

	err := roles.Assign(client.ServiceClient(), "{role_id}", roles.AssignOpts{
		UserID:    "{user_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Assign(client.ServiceClient(), "{role_id}", roles.AssignOpts{
		UserID:   "{user_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Assign(client.ServiceClient(), "{role_id}", roles.AssignOpts{
		GroupID:   "{group_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Assign(client.ServiceClient(), "{role_id}", roles.AssignOpts{
		GroupID:  "{group_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnassign(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUnassignSuccessfully(t)

	err := roles.Unassign(client.ServiceClient(), "{role_id}", roles.UnassignOpts{
		UserID:    "{user_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Unassign(client.ServiceClient(), "{role_id}", roles.UnassignOpts{
		UserID:   "{user_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Unassign(client.ServiceClient(), "{role_id}", roles.UnassignOpts{
		GroupID:   "{group_id}",
		ProjectID: "{project_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)

	err = roles.Unassign(client.ServiceClient(), "{role_id}", roles.UnassignOpts{
		GroupID:  "{group_id}",
		DomainID: "{domain_id}",
	}).ExtractErr()
	th.AssertNoErr(t, err)
}
