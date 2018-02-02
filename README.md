Websocketify (wsify) v1.0
=========================
> Just a tiny, simple and realtime websocket based pub/sub messaging service using redis as backend.

Features
================
- No dependencies except for redis as storage layer !
- Light and Tiny.
- Uses Redis as Pub/Sub backend.
- A client can listen on any redis channel with no hassle.
- You can handle the user authentication using a simple `webhook`.
- You can set a message to only be sent to certain users.

How it Works?
===============
- `Wsify` implements a `http` server.
- That http server translates the request `path` i.e `/some/channel/` to a redis channel `some/channel`.
- After the client requests `ws://wsify.dev/some/channel/` the `wsify` server will send a request to a `webhook` to just authenticate the request.
- The `webhook` will recieve the `Authorization` header `Bearer XXXXXXX` and a query param `?channel=some/channel` that will contain the requested channel.
- In the `webhook`, you do your own logic to tell `wsify` that this client is authorized or not by simply returns a status code `200` in case of success, or anything else to say that this user isn't authorized.
- To publish a message you need to `redis-cli> PUBLISH some/channel '{"payload": "hi", "to": []}'` this will publish the payload "hi" to all the subscribers on that channel `some/channel`.
- To only send it to certain users, you will need to specify the target(s) tokens in the `to` field `redis-cli> PUBLISH some/channel '{"payload": "hi", "to": ["Bearer XXXXXX"]}'.`

Installation
==============

- **Docker ?** > `docker run --network host alash3al/wsify -listen :8080 -auth-webhook "http://localhost/auth.php"`   
- **Binary ?** > goto the [releases](https://github.com/alash3al/wsify/releases) page and download yours.
- **From Source ?** > `go get -u github.com/alash3al/wsify`

Author
=============
This project has been created by [Mohamed Al Ashaal](http://github.com/alash3al) a Crazy Gopher ^^!

Contribution
=============
- Fork the Repo
- Create a feature branch
- Push your changes to the created branch
- Create a pull request.

License
=============
Wsify is open-sourced software licensed under the [MIT License](LICENSE).
