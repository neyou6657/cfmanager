Queues
Queues
Methods
client.Queues.List(ctx, query) (*SinglePage[
Queue
], error)
get/accounts/{account_id}/queues
Returns the queues owned by an account.
client.Queues.Get(ctx, queueID, query) (*
Queue
, error)
get/accounts/{account_id}/queues/{queue_id}
Get details about a specific queue.
client.Queues.New(ctx, params) (*
Queue
, error)
post/accounts/{account_id}/queues
Create a new queue
client.Queues.Update(ctx, queueID, params) (*
Queue
, error)
put/accounts/{account_id}/queues/{queue_id}
Updates a Queue. Note that this endpoint does not support partial updates. If successful, the Queue's configuration is overwritten with the supplied configuration.
client.Queues.Edit(ctx, queueID, params) (*
Queue
, error)
patch/accounts/{account_id}/queues/{queue_id}
Updates a Queue.
client.Queues.Delete(ctx, queueID, body) (*
QueueDeleteResponse
, error)
delete/accounts/{account_id}/queues/{queue_id}
Deletes a queue
Domain types
typeQueuestruct{…}
Queues
Consumers
Queues.Consumers
Methods
client.Queues.Consumers.List(ctx, queueID, query) (*SinglePage[
Consumer
], error)
get/accounts/{account_id}/queues/{queue_id}/consumers
Returns the consumers for a Queue
client.Queues.Consumers.Get(ctx, queueID, consumerID, query) (*
Consumer
, error)
get/accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
Fetches the consumer for a queue by consumer id
client.Queues.Consumers.New(ctx, queueID, params) (*
Consumer
, error)
post/accounts/{account_id}/queues/{queue_id}/consumers
Creates a new consumer for a Queue
client.Queues.Consumers.Update(ctx, queueID, consumerID, params) (*
Consumer
, error)
put/accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
Updates the consumer for a queue, or creates one if it does not exist.
client.Queues.Consumers.Delete(ctx, queueID, consumerID, body) (*
ConsumerDeleteResponse
, error)
delete/accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
Deletes the consumer for a queue.
Domain types
typeConsumerinterface{…}
Queues
Messages
Queues.Messages
Methods
client.Queues.Messages.Push(ctx, queueID, params) (*
MessagePushResponse
, error)
post/accounts/{account_id}/queues/{queue_id}/messages
Push a message to a Queue
client.Queues.Messages.Ack(ctx, queueID, params) (*
MessageAckResponse
, error)
post/accounts/{account_id}/queues/{queue_id}/messages/ack
Acknowledge + Retry messages from a Queue
client.Queues.Messages.Pull(ctx, queueID, params) (*
MessagePullResponse
, error)
post/accounts/{account_id}/queues/{queue_id}/messages/pull
Pull a batch of messages from a Queue
client.Queues.Messages.BulkPush(ctx, queueID, params) (*
MessageBulkPushResponse
, error)
post/accounts/{account_id}/queues/{queue_id}/messages/batch
Push a batch of message to a Queue
Queues
Purge
Queues.Purge
Methods
client.Queues.Purge.Status(ctx, queueID, query) (*
PurgeStatusResponse
, error)
get/accounts/{account_id}/queues/{queue_id}/purge
Get details about a Queue's purge status.
client.Queues.Purge.Start(ctx, queueID, params) (*
Queue
, error)
post/accounts/{account_id}/queues/{queue_id}/purge
Deletes all messages from the Queue.
Queues
Subscriptions
Queues.Subscriptions
Methods
client.Queues.Subscriptions.List(ctx, params) (*V4PagePaginationArray[
SubscriptionListResponse
], error)
get/accounts/{account_id}/event_subscriptions/subscriptions
Get a paginated list of event subscriptions with optional sorting and filtering
client.Queues.Subscriptions.Get(ctx, subscriptionID, query) (*
SubscriptionGetResponse
, error)
get/accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}
Get details about an existing event subscription
client.Queues.Subscriptions.New(ctx, params) (*
SubscriptionNewResponse
, error)
post/accounts/{account_id}/event_subscriptions/subscriptions
Create a new event subscription for a queue
client.Queues.Subscriptions.Update(ctx, subscriptionID, params) (*
SubscriptionUpdateResponse
, error)
patch/accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}
Update an existing event subscription
client.Queues.Subscriptions.Delete(ctx, subscriptionID, body) (*
SubscriptionDeleteResponse
, error)
delete/accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}
Delete an existing event subscription