Stream
Stream
Methods
client.Stream.List(ctx, params) (*SinglePage[
Video
], error)
get/accounts/{account_id}/stream
Lists up to 1000 videos from a single request. For a specific range, refer to the optional parameters.
client.Stream.Get(ctx, identifier, query) (*
Video
, error)
get/accounts/{account_id}/stream/{identifier}
Fetches details for a single video.
client.Stream.New(ctx, params) error
post/accounts/{account_id}/stream
Initiates a video upload using the TUS protocol. On success, the server responds with a status code 201 (created) and includes a location header to indicate where the content should be uploaded. Refer to https://tus.io for protocol details.
client.Stream.Edit(ctx, identifier, params) (*
Video
, error)
post/accounts/{account_id}/stream/{identifier}
Edit details for a single video.
client.Stream.Delete(ctx, identifier, body) error
delete/accounts/{account_id}/stream/{identifier}
Deletes a video and its copies from Cloudflare Stream.
Domain types
typeAllowedOriginsstring
typeVideostruct{…}
Stream
Audio Tracks
Stream.AudioTracks
Methods
client.Stream.AudioTracks.Get(ctx, identifier, query) (*SinglePage[
Audio
], error)
get/accounts/{account_id}/stream/{identifier}/audio
Lists additional audio tracks on a video. Note this API will not return information for audio attached to the video upload.
client.Stream.AudioTracks.Edit(ctx, identifier, audioIdentifier, params) (*
Audio
, error)
patch/accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}
Edits additional audio tracks on a video. Editing the default status of an audio track to true will mark all other audio tracks on the video default status to false.
client.Stream.AudioTracks.Delete(ctx, identifier, audioIdentifier, body) (*string, error)
delete/accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}
Deletes additional audio tracks on a video. Deleting a default audio track is not allowed. You must assign another audio track as default prior to deletion.
client.Stream.AudioTracks.Copy(ctx, identifier, params) (*
Audio
, error)
post/accounts/{account_id}/stream/{identifier}/audio/copy
Adds an additional audio track to a video using the provided audio track URL.
Domain types
typeAudiostruct{…}
Stream
Captions
Stream.Captions
Methods
client.Stream.Captions.Get(ctx, identifier, query) (*SinglePage[
Caption
], error)
get/accounts/{account_id}/stream/{identifier}/captions
Lists the available captions or subtitles for a specific video.
Domain types
typeCaptionstruct{…}
Stream
Captions
Language
Stream.Captions.Language
Methods
client.Stream.Captions.Language.Get(ctx, identifier, language, query) (*
Caption
, error)
get/accounts/{account_id}/stream/{identifier}/captions/{language}
Lists the captions or subtitles for provided language.
client.Stream.Captions.Language.New(ctx, identifier, language, body) (*
Caption
, error)
post/accounts/{account_id}/stream/{identifier}/captions/{language}/generate
Generate captions or subtitles for provided language via AI.
client.Stream.Captions.Language.Update(ctx, identifier, language, params) (*
Caption
, error)
put/accounts/{account_id}/stream/{identifier}/captions/{language}
Uploads the caption or subtitle file to the endpoint for a specific BCP47 language. One caption or subtitle file per language is allowed.
client.Stream.Captions.Language.Delete(ctx, identifier, language, body) (*string, error)
delete/accounts/{account_id}/stream/{identifier}/captions/{language}
Removes the captions or subtitles from a video.
Stream
Captions
Language
Vtt
Stream.Captions.Language.Vtt
Methods
client.Stream.Captions.Language.Vtt.Get(ctx, identifier, language, query) (*string, error)
get/accounts/{account_id}/stream/{identifier}/captions/{language}/vtt
Return WebVTT captions for a provided language.
Stream
Clip
Stream.Clip
Methods
client.Stream.Clip.New(ctx, params) (*
Clip
, error)
post/accounts/{account_id}/stream/clip
Clips a video based on the specified start and end times provided in seconds.
Domain types
typeClipstruct{…}
Stream
Copy
Stream.Copy
Methods
client.Stream.Copy.New(ctx, params) (*
Video
, error)
post/accounts/{account_id}/stream/copy
Uploads a video to Stream from a provided URL.
Stream
Direct Upload
Stream.DirectUpload
Methods
client.Stream.DirectUpload.New(ctx, params) (*
DirectUploadNewResponse
, error)
post/accounts/{account_id}/stream/direct_upload
Creates a direct upload that allows video uploads without an API key.
Stream
Downloads
Stream.Downloads
Methods
client.Stream.Downloads.Get(ctx, identifier, query) (*
DownloadGetResponse
, error)
get/accounts/{account_id}/stream/{identifier}/downloads
Lists the downloads created for a video.
client.Stream.Downloads.New(ctx, identifier, params) (*
DownloadNewResponse
, error)
post/accounts/{account_id}/stream/{identifier}/downloads
Creates a download for a video when a video is ready to view. Use /downloads/{download_type} instead for type-specific downloads. Available types are default and audio.
client.Stream.Downloads.Delete(ctx, identifier, body) (*string, error)
delete/accounts/{account_id}/stream/{identifier}/downloads
Delete the downloads for a video. Use /downloads/{download_type} instead for type-specific downloads. Available types are default and audio.
Stream
Embed
Stream.Embed
Methods
client.Stream.Embed.Get(ctx, identifier, query) (*string, error)
get/accounts/{account_id}/stream/{identifier}/embed
Fetches an HTML code snippet to embed a video in a web page delivered through Cloudflare. On success, returns an HTML fragment for use on web pages to display a video. On failure, returns a JSON response body.
Stream
Keys
Stream.Keys
Methods
client.Stream.Keys.Get(ctx, query) (*SinglePage[
KeyGetResponse
], error)
get/accounts/{account_id}/stream/keys
Lists the video ID and creation date and time when a signing key was created.
client.Stream.Keys.New(ctx, params) (*
Keys
, error)
post/accounts/{account_id}/stream/keys
Creates an RSA private key in PEM and JWK formats. Key files are only displayed once after creation. Keys are created, used, and deleted independently of videos, and every key can sign any video.
client.Stream.Keys.Delete(ctx, identifier, body) (*string, error)
delete/accounts/{account_id}/stream/keys/{identifier}
Deletes signing keys and revokes all signed URLs generated with the key.
Domain types
typeKeysstruct{…}
Stream
Live Inputs
Stream.LiveInputs
Methods
client.Stream.LiveInputs.List(ctx, params) (*
LiveInputListResponse
, error)
get/accounts/{account_id}/stream/live_inputs
Lists the live inputs created for an account. To get the credentials needed to stream to a specific live input, request a single live input.
client.Stream.LiveInputs.Get(ctx, liveInputIdentifier, query) (*
LiveInput
, error)
get/accounts/{account_id}/stream/live_inputs/{live_input_identifier}
Retrieves details of an existing live input.
client.Stream.LiveInputs.New(ctx, params) (*
LiveInput
, error)
post/accounts/{account_id}/stream/live_inputs
Creates a live input, and returns credentials that you or your users can use to stream live video to Cloudflare Stream.
client.Stream.LiveInputs.Update(ctx, liveInputIdentifier, params) (*
LiveInput
, error)
put/accounts/{account_id}/stream/live_inputs/{live_input_identifier}
Updates a specified live input.
client.Stream.LiveInputs.Delete(ctx, liveInputIdentifier, body) error
delete/accounts/{account_id}/stream/live_inputs/{live_input_identifier}
Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls.
Domain types
typeLiveInputstruct{…}
Details about a live input.
Stream
Live Inputs
Outputs
Stream.LiveInputs.Outputs
Methods
client.Stream.LiveInputs.Outputs.List(ctx, liveInputIdentifier, query) (*SinglePage[
Output
], error)
get/accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs
Retrieves all outputs associated with a specified live input.
client.Stream.LiveInputs.Outputs.New(ctx, liveInputIdentifier, params) (*
Output
, error)
post/accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs
Creates a new output that can be used to simulcast or restream live video to other RTMP or SRT destinations. Outputs are always linked to a specific live input — one live input can have many outputs.
client.Stream.LiveInputs.Outputs.Update(ctx, liveInputIdentifier, outputIdentifier, params) (*
Output
, error)
put/accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}
Updates the state of an output.
client.Stream.LiveInputs.Outputs.Delete(ctx, liveInputIdentifier, outputIdentifier, body) error
delete/accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}
Deletes an output and removes it from the associated live input.
Domain types
typeOutputstruct{…}
Stream
Token
Stream.Token
Methods
client.Stream.Token.New(ctx, identifier, params) (*
TokenNewResponse
, error)
post/accounts/{account_id}/stream/{identifier}/token
Creates a signed URL token for a video. If a body is not provided in the request, a token is created with default values.
Stream
Videos
Stream.Videos
Methods
client.Stream.Videos.StorageUsage(ctx, params) (*
VideoStorageUsageResponse
, error)
get/accounts/{account_id}/stream/storage-usage
Returns information about an account's storage use.
Stream
Watermarks
Stream.Watermarks
Methods
client.Stream.Watermarks.List(ctx, query) (*SinglePage[
Watermark
], error)
get/accounts/{account_id}/stream/watermarks
Lists all watermark profiles for an account.
client.Stream.Watermarks.Get(ctx, identifier, query) (*
Watermark
, error)
get/accounts/{account_id}/stream/watermarks/{identifier}
Retrieves details for a single watermark profile.
client.Stream.Watermarks.New(ctx, params) (*
Watermark
, error)
post/accounts/{account_id}/stream/watermarks
Creates watermark profiles using a single HTTP POST multipart/form-data request.
client.Stream.Watermarks.Delete(ctx, identifier, body) (*string, error)
delete/accounts/{account_id}/stream/watermarks/{identifier}
Deletes a watermark profile.
Domain types
typeWatermarkstruct{…}
Stream
Webhooks
Stream.Webhooks
Methods
client.Stream.Webhooks.Get(ctx, query) (*
WebhookGetResponse
, error)
get/accounts/{account_id}/stream/webhook
Retrieves a list of webhooks.
client.Stream.Webhooks.Update(ctx, params) (*
WebhookUpdateResponse
, error)
put/accounts/{account_id}/stream/webhook
Creates a webhook notification.
client.Stream.Webhooks.Delete(ctx, body) (*string, error)
delete/accounts/{account_id}/stream/webhook
Deletes a webhook.