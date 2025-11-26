Memberships
Memberships
Methods
client.Memberships.List(ctx, query) (*V4PagePaginationArray[
Membership
], error)
get/memberships
List memberships of accounts the user can access.
client.Memberships.Get(ctx, membershipID) (*
MembershipGetResponse
, error)
get/memberships/{membership_id}
Get a specific membership.
client.Memberships.Update(ctx, membershipID, body) (*
MembershipUpdateResponse
, error)
put/memberships/{membership_id}
Accept or reject this account invitation.
client.Memberships.Delete(ctx, membershipID) (*
MembershipDeleteResponse
, error)
delete/memberships/{membership_id}
Remove the associated member from an account.
Domain types
typeMembershipstruct{…}