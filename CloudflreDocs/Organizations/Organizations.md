Organizations
Organizations
Methods
client.Organizations.List(ctx, query) (*SinglePage[
Organization
], error)
get/organizations
Retrieve a list of organizations a particular user has access to. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.Get(ctx, organizationID) (*
Organization
, error)
get/organizations/{organization_id}
Retrieve the details of a certain organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.New(ctx, body) (*
Organization
, error)
post/organizations
Create a new organization for a user. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.Update(ctx, organizationID, body) (*
Organization
, error)
put/organizations/{organization_id}
Modify organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Security
API Email + API Key
The previous authorization scheme for interacting with the Cloudflare API, used in conjunction with a Global API key.
Example: X-Auth-Email: user@example.com
The previous authorization scheme for interacting with the Cloudflare API. When possible, use API tokens instead of Global API keys.
Example: X-Auth-Key: 144c9defac04969c7bfad8efaa8ea194
Parameters
organizationIDstring
body
OrganizationUpdateParams
Organizationparam.Field[
Organization
]
References an Organization in the Cloudflare data model.
Returns
Organization
IDstring
CreateTime
Time
(format: date-time)
Meta
OrganizationMeta
Namestring
Parent
OrganizationParent
Optional
Profile
AccountProfile
Optional
Request example Go
package main

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/organizations"
)

func main() {
	client := cloudflare.NewClient(
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	organization, err := client.Organizations.Update(
		context.TODO(),
		"a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8",
		organizations.OrganizationUpdateParams{
			Organization: organizations.OrganizationParam{
				Name: cloudflare.F("name"),
			},
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", organization.ID)
}

200Example
{
  "errors": [],
  "messages": [
    {
      "code": 1000,
      "message": "message",
      "documentation_url": "documentation_url",
      "source": {
        "pointer": "pointer"
      }
    }
  ],
  "result": {
    "id": "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8",
    "create_time": "2019-12-27T18:11:19.117Z",
    "meta": {
      "flags": {
        "account_creation": "account_creation",
        "account_deletion": "account_deletion",
        "account_migration": "account_migration",
        "account_mobility": "account_mobility",
        "sub_org_creation": "sub_org_creation"
      },
      "managed_by": "managed_by"
    },
    "name": "name",
    "parent": {
      "id": "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8",
      "name": "name"
    },
    "profile": {
      "business_address": "business_address",
      "business_email": "business_email",
      "business_name": "business_name",
      "business_phone": "business_phone",
      "external_metadata": "external_metadata"
    }
  },
  "success": true}
client.Organizations.Delete(ctx, organizationID) (*
OrganizationDeleteResponse
, error)
delete/organizations/{organization_id}
Delete an organization. The organization MUST be empty before deleting. It must not contain any sub-organizations, accounts, members or users. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Domain types
typeOrganizationstruct{…}
References an Organization in the Cloudflare data model.
Organizations
Members
Organizations.Members
Methods
client.Organizations.Members.List(ctx, organizationID, query) (*SinglePage[
OrganizationMember
], error)
get/organizations/{organization_id}/members
List memberships for an Organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.Members.Get(ctx, organizationID, memberID) (*
OrganizationMember
, error)
get/organizations/{organization_id}/members/{member_id}
Retrieve a single membership from an Organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.Members.New(ctx, organizationID, body) (*
OrganizationMember
, error)
post/organizations/{organization_id}/members
Create a membership that grants access to a specific Organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.Members.Delete(ctx, organizationID, memberID) error
delete/organizations/{organization_id}/members/{member_id}
Delete a membership to a particular Organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Domain types
typeOrganizationMemberstruct{…}
Organizations
Organization Accounts
Organizations.OrganizationAccounts
Methods
client.Organizations.OrganizationAccounts.Get(ctx, organizationID, query) (*[]
OrganizationAccountGetResponse
, error)
get/organizations/{organization_id}/accounts
Retrieve a list of accounts that belong to a specific organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Domain types
typeOrganizationAccountsstruct{…}
Organizations
Organization Profile
Organizations.OrganizationProfile
Methods
client.Organizations.OrganizationProfile.Get(ctx, organizationID) (*
AccountProfile
, error)
get/organizations/{organization_id}/profile
Get an organizations profile if it exists. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
client.Organizations.OrganizationProfile.Update(ctx, organizationID, body) error
put/organizations/{organization_id}/profile
Modify organization profile. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Domain types
typeOrganizationProfilestruct{…}