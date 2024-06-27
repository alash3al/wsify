WSIFY
======
> a tiny general purpose websocket server that could be used for building real-time apps
> in addition to giving you the power to simply accept/reject any websocket message when using the
> authorizer feature (a webhook that should respond with 200 OK on success and anything else to reject).

Philosophy
==========
- A websocket server should only responsible for transmitting messages in real-time between connected parties.
- It isn't the websocket server responsibility to authorize whether a party is allowed to send certain message or not.
- An authorizer must respond with `200 OK` if a party can send a message.
  - `200 OK` means "Authorized"
  - `5xx` means "the authorizer is down, please close the current connection and reconnect"
  - anything else means "NotAuthorized"
- The client identity should be declared in the query params while connecting to the websocket endpoint and as an argument in the `args` of the message if needed.
- 

Definitions
============
### Message
> is a data structure contains some data to be transmitted.

```json5
{
  // command is a string describes what should be done
  // available commands are:
  // "join": joins a channel (specified in the "args.channel")
  // "leave": leaves a channel (specified in the "args.channel")
  // "broadcast": broadcasts a content (specified in the "args.content") to a channel (in "args.channel")
  "command":  "join",
  "args": {
    "channel": "some_channel"
  }
}
```

### Authorizer
> a webhook that responds with `200 OK` to accept a message sent by a websocket client,
> the authorizer will receive a `POST` request containing a message data structure as described above, but there
> is one special command that isn't described which is "connect", it is fired before accepting the websocket connection
> and it sounds like 
```json5
{
  "command": "connect",
  "args": {
    // array of all http headers sent by the client while trying to open a websocket connection.
    "headers": [/*...*/],
    // an object that contains all available query params sent by the client while trying to open a websocket connection.
    "query": {/*...*/}
  }
}
```
> On all messages from party, wsify will try to authorize it with 'Message' structure inside

Usage
=====
> There is no need to say a lot on how to use this software, just connect using any websocket client to `http://wsify:3000/ws` and start sending messages


Examples
========

#### \> How can a client/device connect to the websocket service?
> by simply connecting to the following endpoint `ws://your.wsify.service:port/ws`


#### \> What is the command used to join a channel named "hello"?
>
```json
{
  "command":  "join",
  "args": {
    "channel": "hello"
  }
}
```

#### \> What is the command used to broadcast a message to the channel "hello"?
```json
{
  "command":  "broadcast",
  "args": {
    "channel": "hello",
    "content": "your message here"
  }
}
```

#### \> How can I publish message to `hello` from another server without connecting to the websocket endpoint?
> Do a post request to `/broadcast` with the following format:
```json5
{
    "channel": "hello",
    "content": "Hello World! from the server"
}
```

### for more info look at `$ ./wsify --help`