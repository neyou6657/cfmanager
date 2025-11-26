Cloudforce One
CloudforceOne
Cloudforce One
Binary Storage
CloudforceOne.BinaryStorage
Methods
client.CloudforceOne.BinaryStorage.Get(ctx, hash, query) error
get/accounts/{account_id}/cloudforce-one/binary/{hash}
Retrieves a file from Binary Storage
client.CloudforceOne.BinaryStorage.New(ctx, params) (*
BinaryStorageNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/binary
Posts a file to Binary Storage
Cloudforce One
Requests
CloudforceOne.Requests
Methods
client.CloudforceOne.Requests.List(ctx, params) (*SinglePage[
ListItem
], error)
post/accounts/{account_id}/cloudforce-one/requests
List Requests
client.CloudforceOne.Requests.Get(ctx, requestID, query) (*
Item
, error)
get/accounts/{account_id}/cloudforce-one/requests/{request_id}
Get a Request
client.CloudforceOne.Requests.New(ctx, params) (*
Item
, error)
post/accounts/{account_id}/cloudforce-one/requests/new
Creating a request adds the request into the Cloudforce One queue for analysis. In addition to the content, a short title, type, priority, and releasability should be provided. If one is not provided, a default will be assigned.
client.CloudforceOne.Requests.Update(ctx, requestID, params) (*
Item
, error)
put/accounts/{account_id}/cloudforce-one/requests/{request_id}
Updating a request alters the request in the Cloudforce One queue. This API may be used to update any attributes of the request after the initial submission. Only fields that you choose to update need to be add to the request body.
client.CloudforceOne.Requests.Delete(ctx, requestID, body) (*
RequestDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/requests/{request_id}
Delete a Request
client.CloudforceOne.Requests.Quota(ctx, query) (*
Quota
, error)
get/accounts/{account_id}/cloudforce-one/requests/quota
Get Request Quota
client.CloudforceOne.Requests.Types(ctx, query) (*SinglePage[string], error)
get/accounts/{account_id}/cloudforce-one/requests/types
Get Request Types
client.CloudforceOne.Requests.Constants(ctx, query) (*
RequestConstants
, error)
get/accounts/{account_id}/cloudforce-one/requests/constants
Get Request Priority, Status, and TLP constants
Domain types
typeItemstruct{…}
typeListItemstruct{…}
typeQuotastruct{…}
typeRequestConstantsstruct{…}
typeRequestTypes[]string
Cloudforce One
Requests
Assets
CloudforceOne.Requests.Assets
Methods
client.CloudforceOne.Requests.Assets.Get(ctx, requestID, assetID, query) (*SinglePage[
RequestAssetGetResponse
], error)
get/accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
Get a Request Asset
client.CloudforceOne.Requests.Assets.New(ctx, requestID, params) (*SinglePage[
RequestAssetNewResponse
], error)
post/accounts/{account_id}/cloudforce-one/requests/{request_id}/asset
List Request Assets
client.CloudforceOne.Requests.Assets.Update(ctx, requestID, assetID, params) (*
RequestAssetUpdateResponse
, error)
put/accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
Update a Request Asset
client.CloudforceOne.Requests.Assets.Delete(ctx, requestID, assetID, body) (*
RequestAssetDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
Delete a Request Asset
Cloudforce One
Requests
Message
CloudforceOne.Requests.Message
Methods
client.CloudforceOne.Requests.Message.Get(ctx, requestID, params) (*SinglePage[
Message
], error)
post/accounts/{account_id}/cloudforce-one/requests/{request_id}/message
List Request Messages
client.CloudforceOne.Requests.Message.New(ctx, requestID, params) (*
Message
, error)
post/accounts/{account_id}/cloudforce-one/requests/{request_id}/message/new
Create a New Request Message
client.CloudforceOne.Requests.Message.Update(ctx, requestID, messageID, params) (*
Message
, error)
put/accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}
Update a Request Message
client.CloudforceOne.Requests.Message.Delete(ctx, requestID, messageID, body) (*
RequestMessageDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}
Delete a Request Message
Domain types
typeMessagestruct{…}
Cloudforce One
Requests
Priority
CloudforceOne.Requests.Priority
Methods
client.CloudforceOne.Requests.Priority.Get(ctx, priorityID, query) (*
Item
, error)
get/accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
Get a Priority Intelligence Requirement
client.CloudforceOne.Requests.Priority.New(ctx, params) (*
Priority
, error)
post/accounts/{account_id}/cloudforce-one/requests/priority/new
Create a New Priority Intelligence Requirement
client.CloudforceOne.Requests.Priority.Update(ctx, priorityID, params) (*
Item
, error)
put/accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
Update a Priority Intelligence Requirement
client.CloudforceOne.Requests.Priority.Delete(ctx, priorityID, body) (*
RequestPriorityDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
Delete a Priority Intelligence Requirement
client.CloudforceOne.Requests.Priority.Quota(ctx, query) (*
Quota
, error)
get/accounts/{account_id}/cloudforce-one/requests/priority/quota
Get Priority Intelligence Requirement Quota
Domain types
typeLabelstring
typePrioritystruct{…}
typePriorityEditstruct{…}
Cloudforce One
Scans
CloudforceOne.Scans
Cloudforce One
Scans
Config
CloudforceOne.Scans.Config
Methods
client.CloudforceOne.Scans.Config.List(ctx, query) (*SinglePage[
ScanConfigListResponse
], error)
get/accounts/{account_id}/cloudforce-one/scans/config
List Scan Configs
client.CloudforceOne.Scans.Config.New(ctx, params) (*
ScanConfigNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/scans/config
Create a new Scan Config
client.CloudforceOne.Scans.Config.Edit(ctx, configID, params) (*
ScanConfigEditResponse
, error)
patch/accounts/{account_id}/cloudforce-one/scans/config/{config_id}
Update an existing Scan Config
client.CloudforceOne.Scans.Config.Delete(ctx, configID, body) (*
ScanConfigDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/scans/config/{config_id}
Delete a Scan Config
Cloudforce One
Scans
Results
CloudforceOne.Scans.Results
Methods
client.CloudforceOne.Scans.Results.Get(ctx, configID, query) (*
ScanResultGetResponse
, error)
get/accounts/{account_id}/cloudforce-one/scans/results/{config_id}
Get the Latest Scan Result
Domain types
typeScanResultstruct{…}
Cloudforce One
Threat Events
CloudforceOne.ThreatEvents
Methods
client.CloudforceOne.ThreatEvents.List(ctx, params) (*[]
ThreatEventListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events
When datasetId is unspecified, events will be listed from the Cloudforce One Threat Events dataset. To list existing datasets (and their IDs), use the List Datasets endpoint). Also, must provide query parameters.
client.CloudforceOne.ThreatEvents.Get(ctx, eventID, query) (*
ThreatEventGetResponse
, error)
Deprecated
get/accounts/{account_id}/cloudforce-one/events/{event_id}
This Method is deprecated. Please use /events/dataset/:dataset_id/events/:event_id instead.
client.CloudforceOne.ThreatEvents.New(ctx, params) (*
ThreatEventNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/events/create
To create a dataset, see the Create Dataset endpoint. When datasetId parameter is unspecified, it will be created in a default dataset named Cloudforce One Threat Events.
client.CloudforceOne.ThreatEvents.Edit(ctx, eventID, params) (*
ThreatEventEditResponse
, error)
patch/accounts/{account_id}/cloudforce-one/events/{event_id}
Updates an event
client.CloudforceOne.ThreatEvents.Delete(ctx, eventID, body) (*
ThreatEventDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/events/{event_id}
The datasetId parameter must be defined. To list existing datasets (and their IDs) in your account, use the List Datasets endpoint.
client.CloudforceOne.ThreatEvents.BulkNew(ctx, params) (*float64, error)
post/accounts/{account_id}/cloudforce-one/events/create/bulk
The datasetId parameter must be defined. To list existing datasets (and their IDs) in your account, use the List Datasets endpoint.
Cloudforce One
Threat Events
Attackers
CloudforceOne.ThreatEvents.Attackers
Methods
client.CloudforceOne.ThreatEvents.Attackers.List(ctx, params) (*
ThreatEventAttackerListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/attackers
Lists attackers across multiple datasets
Cloudforce One
Threat Events
Categories
CloudforceOne.ThreatEvents.Categories
Methods
client.CloudforceOne.ThreatEvents.Categories.List(ctx, params) (*[]
ThreatEventCategoryListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/categories
Lists categories across multiple datasets
client.CloudforceOne.ThreatEvents.Categories.Get(ctx, categoryID, query) (*
ThreatEventCategoryGetResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/categories/{category_id}
Reads a category
client.CloudforceOne.ThreatEvents.Categories.New(ctx, params) (*
ThreatEventCategoryNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/events/categories/create
Creates a new category
client.CloudforceOne.ThreatEvents.Categories.Edit(ctx, categoryID, params) (*
ThreatEventCategoryEditResponse
, error)
patch/accounts/{account_id}/cloudforce-one/events/categories/{category_id}
Updates a category
client.CloudforceOne.ThreatEvents.Categories.Delete(ctx, categoryID, body) (*
ThreatEventCategoryDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/events/categories/{category_id}
Deletes a category
Cloudforce One
Threat Events
Countries
CloudforceOne.ThreatEvents.Countries
Methods
client.CloudforceOne.ThreatEvents.Countries.List(ctx, query) (*[]
ThreatEventCountryListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/countries
Retrieves countries information for all countries
Cloudforce One
Threat Events
Crons
CloudforceOne.ThreatEvents.Crons
Cloudforce One
Threat Events
Datasets
CloudforceOne.ThreatEvents.Datasets
Methods
client.CloudforceOne.ThreatEvents.Datasets.List(ctx, query) (*[]
ThreatEventDatasetListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/dataset
Lists all datasets in an account
client.CloudforceOne.ThreatEvents.Datasets.Get(ctx, datasetID, query) (*
ThreatEventDatasetGetResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}
Reads a dataset
client.CloudforceOne.ThreatEvents.Datasets.New(ctx, params) (*
ThreatEventDatasetNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/events/dataset/create
Creates a dataset
client.CloudforceOne.ThreatEvents.Datasets.Edit(ctx, datasetID, params) (*
ThreatEventDatasetEditResponse
, error)
patch/accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}
Updates an existing dataset
client.CloudforceOne.ThreatEvents.Datasets.Raw(ctx, datasetID, eventID, query) (*
ThreatEventDatasetRawResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/raw/{dataset_id}/{event_id}
Reads data for a raw event
Cloudforce One
Threat Events
Datasets
Health
CloudforceOne.ThreatEvents.Datasets.Health
Cloudforce One
Threat Events
Event Tags
CloudforceOne.ThreatEvents.EventTags
Methods
client.CloudforceOne.ThreatEvents.EventTags.New(ctx, eventID, params) (*
ThreatEventEventTagNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}/create
Adds a tag to an event
client.CloudforceOne.ThreatEvents.EventTags.Delete(ctx, eventID, body) (*
ThreatEventEventTagDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}
Removes a tag from an event
Cloudforce One
Threat Events
Indicator Types
CloudforceOne.ThreatEvents.IndicatorTypes
Methods
client.CloudforceOne.ThreatEvents.IndicatorTypes.List(ctx, query) (*
ThreatEventIndicatorTypeListResponse
, error)
Deprecated
get/accounts/{account_id}/cloudforce-one/events/indicatorTypes
This Method is deprecated. Please use /events/dataset/:dataset_id/indicatorTypes instead.
Cloudforce One
Threat Events
Insights
CloudforceOne.ThreatEvents.Insights
Cloudforce One
Threat Events
Raw
CloudforceOne.ThreatEvents.Raw
Methods
client.CloudforceOne.ThreatEvents.Raw.Get(ctx, eventID, rawID, query) (*
ThreatEventRawGetResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}
Reads data for a raw event
client.CloudforceOne.ThreatEvents.Raw.Edit(ctx, eventID, rawID, params) (*
ThreatEventRawEditResponse
, error)
patch/accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}
Updates a raw event
Cloudforce One
Threat Events
Relate
CloudforceOne.ThreatEvents.Relate
Methods
client.CloudforceOne.ThreatEvents.Relate.Delete(ctx, eventID, body) (*
ThreatEventRelateDeleteResponse
, error)
delete/accounts/{account_id}/cloudforce-one/events/relate/{event_id}
Removes an event reference
Cloudforce One
Threat Events
Tags
CloudforceOne.ThreatEvents.Tags
Methods
client.CloudforceOne.ThreatEvents.Tags.New(ctx, params) (*
ThreatEventTagNewResponse
, error)
post/accounts/{account_id}/cloudforce-one/events/tags/create
Creates a new tag to be used accross threat events.
Cloudforce One
Threat Events
Target Industries
CloudforceOne.ThreatEvents.TargetIndustries
Methods
client.CloudforceOne.ThreatEvents.TargetIndustries.List(ctx, params) (*
ThreatEventTargetIndustryListResponse
, error)
get/accounts/{account_id}/cloudforce-one/events/targetIndustries
Lists target industries across multiple datasets