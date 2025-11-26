Zaraz
Zaraz
Methods
client.Zaraz.Update(ctx, params) (*
Workflow
, error)
put/zones/{zone_id}/settings/zaraz/workflow
Updates Zaraz workflow for a zone.
Domain types
typeButtonTextTranslationstruct{…}
typeNeoEventstruct{…}
Zaraz
Config
Zaraz.Config
Methods
client.Zaraz.Config.Get(ctx, query) (*
Configuration
, error)
get/zones/{zone_id}/settings/zaraz/config
Gets latest Zaraz configuration for a zone. It can be preview or published configuration, whichever was the last updated. Secret variables values will not be included.
client.Zaraz.Config.Update(ctx, params) (*
Configuration
, error)
put/zones/{zone_id}/settings/zaraz/config
Updates Zaraz configuration for a zone.
Domain types
typeConfigurationstruct{…}
Zaraz configuration
Zaraz
Default
Zaraz.Default
Methods
client.Zaraz.Default.Get(ctx, query) (*
Configuration
, error)
get/zones/{zone_id}/settings/zaraz/default
Gets default Zaraz configuration for a zone.
Zaraz
Export
Zaraz.Export
Methods
client.Zaraz.Export.Get(ctx, query) (*
Configuration
, error)
get/zones/{zone_id}/settings/zaraz/export
Exports full current published Zaraz configuration for a zone, secret variables included.
Zaraz
History
Zaraz.History
Methods
client.Zaraz.History.List(ctx, params) (*SinglePage[
HistoryListResponse
], error)
get/zones/{zone_id}/settings/zaraz/history
Lists a history of published Zaraz configuration records for a zone.
client.Zaraz.History.Update(ctx, params) (*
Configuration
, error)
put/zones/{zone_id}/settings/zaraz/history
Restores a historical published Zaraz configuration by ID for a zone.
Zaraz
History
Configs
Zaraz.History.Configs
Methods
client.Zaraz.History.Configs.Get(ctx, params) (*
HistoryConfigGetResponse
, error)
get/zones/{zone_id}/settings/zaraz/history/configs
Gets a history of published Zaraz configurations by ID(s) for a zone.
Zaraz
Publish
Zaraz.Publish
Methods
client.Zaraz.Publish.New(ctx, params) (*string, error)
post/zones/{zone_id}/settings/zaraz/publish
Publish current Zaraz preview configuration for a zone.
Zaraz
Workflow
Zaraz.Workflow
Methods
client.Zaraz.Workflow.Get(ctx, query) (*
Workflow
, error)
get/zones/{zone_id}/settings/zaraz/workflow
Gets Zaraz workflow for a zone.
Domain types
typeWorkflowstring
Zaraz workflow