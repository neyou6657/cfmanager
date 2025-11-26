Waiting Rooms
WaitingRooms
Methods
client.WaitingRooms.List(ctx, params) (*V4PagePaginationArray[
WaitingRoom
], error)
get/{accounts_or_zones}/{account_or_zone_id}/waiting_rooms
Lists waiting rooms for account or zone.
client.WaitingRooms.Get(ctx, waitingRoomID, query) (*
WaitingRoom
, error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}
Fetches a single configured waiting room.
client.WaitingRooms.New(ctx, params) (*
WaitingRoom
, error)
post/zones/{zone_id}/waiting_rooms
Creates a new waiting room.
client.WaitingRooms.Update(ctx, waitingRoomID, params) (*
WaitingRoom
, error)
put/zones/{zone_id}/waiting_rooms/{waiting_room_id}
Updates a configured waiting room.
client.WaitingRooms.Edit(ctx, waitingRoomID, params) (*
WaitingRoom
, error)
patch/zones/{zone_id}/waiting_rooms/{waiting_room_id}
Patches a configured waiting room.
client.WaitingRooms.Delete(ctx, waitingRoomID, body) (*
WaitingRoomDeleteResponse
, error)
delete/zones/{zone_id}/waiting_rooms/{waiting_room_id}
Deletes a waiting room.
Domain types
typeAdditionalRoutesstruct{…}
typeCookieAttributesstruct{…}
Configures cookie attributes for the waiting room cookie. This encrypted cookie stores a user's status in the waiting room, such as queue position.
typeQuerystruct{…}
typeWaitingRoomstruct{…}
Waiting Rooms
Events
WaitingRooms.Events
Methods
client.WaitingRooms.Events.List(ctx, waitingRoomID, params) (*V4PagePaginationArray[
Event
], error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events
Lists events for a waiting room.
client.WaitingRooms.Events.Get(ctx, waitingRoomID, eventID, query) (*
Event
, error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}
Fetches a single configured event for a waiting room.
client.WaitingRooms.Events.New(ctx, waitingRoomID, params) (*
Event
, error)
post/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events
Only available for the Waiting Room Advanced subscription. Creates an event for a waiting room. An event takes place during a specified period of time, temporarily changing the behavior of a waiting room. While the event is active, some of the properties in the event's configuration may either override or inherit from the waiting room's configuration. Note that events cannot overlap with each other, so only one event can be active at a time.
client.WaitingRooms.Events.Update(ctx, waitingRoomID, eventID, params) (*
Event
, error)
put/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}
Updates a configured event for a waiting room.
client.WaitingRooms.Events.Edit(ctx, waitingRoomID, eventID, params) (*
Event
, error)
patch/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}
Patches a configured event for a waiting room.
client.WaitingRooms.Events.Delete(ctx, waitingRoomID, eventID, body) (*
EventDeleteResponse
, error)
delete/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}
Deletes an event for a waiting room.
Domain types
typeEventstruct{…}
Waiting Rooms
Events
Details
WaitingRooms.Events.Details
Methods
client.WaitingRooms.Events.Details.Get(ctx, waitingRoomID, eventID, query) (*
EventDetailGetResponse
, error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}/details
Previews an event's configuration as if it was active. Inherited fields from the waiting room will be displayed with their current values.
Domain types
typeEventQuerystruct{…}
Waiting Rooms
Page
WaitingRooms.Page
Methods
client.WaitingRooms.Page.Preview(ctx, params) (*
PagePreviewResponse
, error)
post/zones/{zone_id}/waiting_rooms/preview
Creates a waiting room page preview. Upload a custom waiting room page for preview. You will receive a preview URL in the form http://waitingrooms.dev/preview/<uuid>. You can use the following query parameters to change the state of the preview:
force_queue: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website (also known as queueAll).
queue_is_full: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.
queueing_method: The queueing method currently used by the waiting room.
fifo indicates a FIFO queue.
random indicates a Random queue.
passthrough indicates a Passthrough queue. Keep in mind that the waiting room page will only be displayed if force_queue=true or event=prequeueing — for other cases the request will pass through to the origin. For our preview, this will be a fake origin website returning "Welcome". 
reject indicates a Reject queue.
event: Used to preview a waiting room event.
none indicates no event is occurring.
prequeueing indicates that an event is prequeueing (between prequeue_start_time and event_start_time).
started indicates that an event has started (between event_start_time and event_end_time).
shuffle_at_event_start: Boolean indicating if the event will shuffle users in the prequeue when it starts. This can only be set to true if an event is active (event is not none).
For example, you can make a request to http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true 6. waitTime: Non-zero, positive integer indicating the estimated wait time in minutes. The default value is 10 minutes.
For example, you can make a request to http://waitingrooms.dev/preview/<uuid>?waitTime=50 to configure the estimated wait time as 50 minutes.
Waiting Rooms
Rules
WaitingRooms.Rules
Methods
client.WaitingRooms.Rules.Get(ctx, waitingRoomID, query) (*SinglePage[
WaitingRoomRule
], error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
Lists rules for a waiting room.
client.WaitingRooms.Rules.New(ctx, waitingRoomID, params) (*SinglePage[
WaitingRoomRule
], error)
post/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
Only available for the Waiting Room Advanced subscription. Creates a rule for a waiting room.
client.WaitingRooms.Rules.Update(ctx, waitingRoomID, params) (*SinglePage[
WaitingRoomRule
], error)
put/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
Only available for the Waiting Room Advanced subscription. Replaces all rules for a waiting room.
client.WaitingRooms.Rules.Edit(ctx, waitingRoomID, ruleID, params) (*SinglePage[
WaitingRoomRule
], error)
patch/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}
Patches a rule for a waiting room.
client.WaitingRooms.Rules.Delete(ctx, waitingRoomID, ruleID, body) (*SinglePage[
WaitingRoomRule
], error)
delete/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}
Deletes a rule for a waiting room.
Domain types
typeWaitingRoomRulestruct{…}
Waiting Rooms
Settings
WaitingRooms.Settings
Methods
client.WaitingRooms.Settings.Get(ctx, query) (*
SettingGetResponse
, error)
get/zones/{zone_id}/waiting_rooms/settings
Get zone-level Waiting Room settings
client.WaitingRooms.Settings.Update(ctx, params) (*
SettingUpdateResponse
, error)
put/zones/{zone_id}/waiting_rooms/settings
Update zone-level Waiting Room settings
client.WaitingRooms.Settings.Edit(ctx, params) (*
SettingEditResponse
, error)
patch/zones/{zone_id}/waiting_rooms/settings
Patch zone-level Waiting Room settings
Domain types
typeSettingstruct{…}
Waiting Rooms
Statuses
WaitingRooms.Statuses
Methods
client.WaitingRooms.Statuses.Get(ctx, waitingRoomID, query) (*
StatusGetResponse
, error)
get/zones/{zone_id}/waiting_rooms/{waiting_room_id}/status
Fetches the status of a configured waiting room. Response fields include:
status: String indicating the status of the waiting room. The possible status are:
not_queueing indicates that the configured thresholds have not been met and all users are going through to the origin.
queueing indicates that the thresholds have been met and some users are held in the waiting room.
event_prequeueing indicates that an event is active and is currently prequeueing users before it starts.
suspended indicates that the room is suspended.
event_id: String of the current event's id if an event is active, otherwise an empty string.
estimated_queued_users: Integer of the estimated number of users currently waiting in the queue.
estimated_total_active_users: Integer of the estimated number of users currently active on the origin.
max_estimated_time_minutes: Integer of the maximum estimated time currently presented to the users.