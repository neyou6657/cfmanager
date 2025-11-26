Images
Images
Images
V1
Images.V1
Methods
client.Images.V1.List(ctx, params) (*V4PagePagination[
V1ListResponse
], error)
Deprecated
get/accounts/{account_id}/images/v1
List up to 100 images with one request. Use the optional parameters below to get a specific range of images.
client.Images.V1.Get(ctx, imageID, query) (*
Image
, error)
get/accounts/{account_id}/images/v1/{image_id}
Fetch details for a single image.
client.Images.V1.New(ctx, params) (*
Image
, error)
post/accounts/{account_id}/images/v1
Upload an image with up to 10 Megabytes using a single HTTP POST (multipart/form-data) request. An image can be uploaded by sending an image file or passing an accessible to an API url.
client.Images.V1.Edit(ctx, imageID, params) (*
Image
, error)
patch/accounts/{account_id}/images/v1/{image_id}
Update image access control. On access control change, all copies of the image are purged from cache.
client.Images.V1.Delete(ctx, imageID, body) (*unknown, error)
delete/accounts/{account_id}/images/v1/{image_id}
Delete an image on Cloudflare Images. On success, all copies of the image are deleted and purged from cache.
Domain types
typeImagestruct{…}
Images
V1
Blobs
Images.V1.Blobs
Methods
client.Images.V1.Blobs.Get(ctx, imageID, query) (*
Response
, error)
get/accounts/{account_id}/images/v1/{image_id}/blob
Fetch base image. For most images this will be the originally uploaded file. For larger images it can be a near-lossless version of the original.
Images
V1
Keys
Images.V1.Keys
Methods
client.Images.V1.Keys.List(ctx, query) (*
V1KeyListResponse
, error)
get/accounts/{account_id}/images/v1/keys
Lists your signing keys. These can be found on your Cloudflare Images dashboard.
client.Images.V1.Keys.Update(ctx, signingKeyName, body) (*
V1KeyUpdateResponse
, error)
put/accounts/{account_id}/images/v1/keys/{signing_key_name}
Create a new signing key with specified name. Returns all keys available.
client.Images.V1.Keys.Delete(ctx, signingKeyName, body) (*
V1KeyDeleteResponse
, error)
delete/accounts/{account_id}/images/v1/keys/{signing_key_name}
Delete signing key with specified name. Returns all keys available. When last key is removed, a new default signing key will be generated.
Domain types
typeKeystruct{…}
Images
V1
Stats
Images.V1.Stats
Methods
client.Images.V1.Stats.Get(ctx, query) (*
Stat
, error)
get/accounts/{account_id}/images/v1/stats
Fetch image statistics details for Cloudflare Images. The returned statistics detail storage usage, including the current image count vs this account's allowance.
Domain types
typeStatstruct{…}
Images
V1
Variants
Images.V1.Variants
Methods
client.Images.V1.Variants.List(ctx, query) (*
Variant
, error)
get/accounts/{account_id}/images/v1/variants
Lists existing variants.
client.Images.V1.Variants.Get(ctx, variantID, query) (*
V1VariantGetResponse
, error)
get/accounts/{account_id}/images/v1/variants/{variant_id}
Fetch details for a single variant.
client.Images.V1.Variants.New(ctx, params) (*
V1VariantNewResponse
, error)
post/accounts/{account_id}/images/v1/variants
Specify variants that allow you to resize images for different use cases.
client.Images.V1.Variants.Edit(ctx, variantID, params) (*
V1VariantEditResponse
, error)
patch/accounts/{account_id}/images/v1/variants/{variant_id}
Updating a variant purges the cache for all images associated with the variant.
client.Images.V1.Variants.Delete(ctx, variantID, body) (*unknown, error)
delete/accounts/{account_id}/images/v1/variants/{variant_id}
Deleting a variant purges the cache for all images associated with the variant.
Domain types
typeVariantstruct{…}
Images
V2
Images.V2
Methods
client.Images.V2.List(ctx, params) (*
V2ListResponse
, error)
get/accounts/{account_id}/images/v2
List up to 10000 images with one request. Use the optional parameters below to get a specific range of images. Endpoint returns continuation_token if more images are present.
Images
V2
Direct Uploads
Images.V2.DirectUploads
Methods
client.Images.V2.DirectUploads.New(ctx, params) (*
V2DirectUploadNewResponse
, error)
post/accounts/{account_id}/images/v2/direct_upload
Direct uploads allow users to upload images without API keys. A common use case are web apps, client-side applications, or mobile devices where users upload content directly to Cloudflare Images. This method creates a draft record for a future image. It returns an upload URL and an image identifier. To verify if the image itself has been uploaded, send an image details request (accounts/:account_identifier/images/v1/:identifier), and check that the draft: true property is not present.