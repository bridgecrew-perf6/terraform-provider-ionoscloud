//go:build compute || all || user

package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccUserBasic(t *testing.T) {
	var user ionoscloud.User

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckUserDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists(UserResource+"."+UserTestResource, &user),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "first_name", UserTestResource),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "last_name", UserTestResource),
					resource.TestCheckResourceAttrSet(UserResource+"."+UserTestResource, "email"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "administrator", "true"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "force_sec_auth", "true"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "active", "true"),
				),
			},
			{
				Config: testAccDataSourceUserMatchId,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "first_name", UserResource+"."+UserTestResource, "first_name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "last_name", UserResource+"."+UserTestResource, "last_name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "email", UserResource+"."+UserTestResource, "email"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "administrator", UserResource+"."+UserTestResource, "administrator"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "force_sec_auth", UserResource+"."+UserTestResource, "force_sec_auth"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "sec_auth_active", UserResource+"."+UserTestResource, "sec_auth_active"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "s3_canonical_user_id", UserResource+"."+UserTestResource, "s3_canonical_user_id"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceById, "active", UserResource+"."+UserTestResource, "active"),
				),
			},
			{
				Config: testAccDataSourceUserMatchEmail,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "first_name", UserResource+"."+UserTestResource, "first_name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "last_name", UserResource+"."+UserTestResource, "last_name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "email", UserResource+"."+UserTestResource, "email"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "administrator", UserResource+"."+UserTestResource, "administrator"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "force_sec_auth", UserResource+"."+UserTestResource, "force_sec_auth"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "sec_auth_active", UserResource+"."+UserTestResource, "sec_auth_active"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "s3_canonical_user_id", UserResource+"."+UserTestResource, "s3_canonical_user_id"),
					resource.TestCheckResourceAttrPair(DataSource+"."+UserResource+"."+UserDataSourceByName, "active", UserResource+"."+UserTestResource, "active"),
				),
			},
			{
				Config:      testAccDataSourceUserWrongEmail,
				ExpectError: regexp.MustCompile(`no user found with the specified criteria: email`),
			},
			{
				Config: testAccCheckUserConfigUpdateForceSec,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists(UserResource+"."+UserTestResource, &user),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "first_name", UserTestResource),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "last_name", UserTestResource),
					resource.TestCheckResourceAttrSet(UserResource+"."+UserTestResource, "email"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "administrator", "true"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "force_sec_auth", "false"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "active", "true"),
				),
			},
			{
				Config: testAccCheckUserConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "first_name", UpdatedResources),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "last_name", UpdatedResources),
					resource.TestCheckResourceAttrSet(UserResource+"."+UserTestResource, "email"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "administrator", "false"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "force_sec_auth", "false"),
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "active", "false")),
			},
			{
				Config: testAccCheckUserMultipleGroups,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "group_ids.#", "3"),
					resource.TestCheckResourceAttr(DataSource+"."+UserResource+"."+UserDataSourceById, "groups.#", "3"),
					resource.TestCheckTypeSetElemNestedAttrs(DataSource+"."+UserResource+"."+UserDataSourceById, "groups.*", map[string]string{
						"name": "group1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(DataSource+"."+UserResource+"."+UserDataSourceById, "groups.*", map[string]string{
						"name": "group2",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(DataSource+"."+UserResource+"."+UserDataSourceById, "groups.*", map[string]string{
						"name": "group3",
					})),
			},
			{
				Config: testAccCheckUserRemoveAllGroups,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(UserResource+"."+UserTestResource, "group_ids.#", "0")),
			},
		},
	})
}

func testAccCheckUserDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(SdkBundle).CloudApiClient

	ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Delete)
	if cancel != nil {
		defer cancel()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != UserResource {
			continue
		}
		_, apiResponse, err := client.UserManagementApi.UmUsersFindById(ctx, rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			if !httpNotFound(apiResponse) {
				return fmt.Errorf("user still exists %s - an error occurred while checking it %s", rs.Primary.ID, err)
			}
		} else {
			return fmt.Errorf("user still exists %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckUserExists(n string, user *ionoscloud.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(SdkBundle).CloudApiClient
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckUserExists: Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)

		if cancel != nil {
			defer cancel()
		}

		foundUser, apiResponse, err := client.UserManagementApi.UmUsersFindById(ctx, rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			return fmt.Errorf("error occured while fetching User: %s %s", rs.Primary.ID, err)
		}
		if *foundUser.Id != rs.Primary.ID {
			return fmt.Errorf("record not found")
		}

		user = &foundUser

		return nil
	}
}

var testAccCheckUserConfigBasic = `
resource ` + UserResource + ` ` + UserTestResource + ` {
  first_name = "` + UserTestResource + `"
  last_name = "` + UserTestResource + `"
  email = "` + utils.GenerateEmail() + `"
  password = "abc123-321CBA"
  administrator = true
  force_sec_auth= true
  active  = true
}`

var testAccCheckUserConfigUpdateForceSec = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name = "` + UserTestResource + `"
 last_name = "` + UserTestResource + `"
 email = "` + utils.GenerateEmail() + `"
 password = "abc123-321CBA"
 administrator = true
 force_sec_auth= false
 active  = true
}`

var testAccCheckUserConfigUpdate = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name = "` + UpdatedResources + `"
 last_name = "` + UpdatedResources + `"
 email = "` + utils.GenerateEmail() + `"
 password = "abc123-321CBAupdated"
 administrator = false
 force_sec_auth= false
 active  = false
}`

var testAccCheckUserMultipleGroups = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name 	= "` + UpdatedResources + `"
 last_name 		= "` + UpdatedResources + `"
 email 			= "` + utils.GenerateEmail() + `"
 password 		= "abc123-321CBAupdated"
 administrator  = false
 force_sec_auth = false
 active  		= false
 group_ids 		= [ ionoscloud_group.group1.id, ionoscloud_group.group2.id, ionoscloud_group.group3.id]
}

resource "ionoscloud_group" "group1" {
  name = "group1"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group2" {
  name = "group2"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group3" {
  name = "group3"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

data ` + UserResource + ` ` + UserDataSourceById + ` {
	id = ionoscloud_user.` + UserTestResource + `.id
}
`

var testAccCheckUserMultipleGroups1Element = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name 	= "` + UpdatedResources + `"
 last_name 		= "` + UpdatedResources + `"
 email 			= "` + utils.GenerateEmail() + `"
 password 		= "abc123-321CBAupdated"
 administrator  = false
 force_sec_auth = false
 active  		= false
 group_ids 		= [ ionoscloud_group.group1.id]
}

resource "ionoscloud_group" "group1" {
  name = "group1"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group2" {
  name = "group2"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group3" {
  name = "group3"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

data ` + UserResource + ` ` + UserDataSourceById + ` {
	id = ionoscloud_user.` + UserTestResource + `.id
}
`

var testAccCheckUserRemoveAllGroups = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name 	= "` + UpdatedResources + `"
 last_name 		= "` + UpdatedResources + `"
 email 			= "` + utils.GenerateEmail() + `"
 password 		= "abc123-321CBAupdated"
 administrator  = false
 force_sec_auth = false
 active  		= false
}

resource "ionoscloud_group" "group1" {
  name = "group1"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group2" {
  name = "group2"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group3" {
  name = "group3"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

data ` + UserResource + ` ` + UserDataSourceById + ` {
	id = ionoscloud_user.` + UserTestResource + `.id
}
`

var testAccCheckUserWrongGroupId = `
resource ` + UserResource + ` ` + UserTestResource + ` {
 first_name 	= "` + UpdatedResources + `"
 last_name 		= "` + UpdatedResources + `"
 email 			= "` + utils.GenerateEmail() + `"
 password 		= "abc123-321CBAupdated"
 administrator  = false
 force_sec_auth = false
 active  		= false
 group_ids = ["notAnId"]
}

resource "ionoscloud_group" "group1" {
  name = "group1"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group2" {
  name = "group2"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
  create_k8s_cluster = true
}
resource "ionoscloud_group" "group3" {
  name = "group3"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

data ` + UserResource + ` ` + UserDataSourceById + ` {
	id = ionoscloud_user.` + UserTestResource + `.id
}
`

var testAccDataSourceUserMatchId = testAccCheckUserConfigBasic + `
data ` + UserResource + ` ` + UserDataSourceById + ` {
  id			= ` + UserResource + `.` + UserTestResource + `.id
}
`

var testAccDataSourceUserMatchEmail = testAccCheckUserConfigBasic + `
data ` + UserResource + ` ` + UserDataSourceByName + ` {
  email			= ` + UserResource + `.` + UserTestResource + `.email
}
`

var testAccDataSourceUserWrongEmail = testAccCheckUserConfigBasic + `
data ` + UserResource + ` ` + UserDataSourceByName + ` {
  email			= "wrong_email"
}
`
