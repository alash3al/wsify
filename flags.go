package main

import (
	"flag"
	"strings"
)

import "github.com/go-redis/redis"

var (
	FLAG_HTTP_ADDR      = flag.String("listen", ":4040", "the http address to listen on")
	FLAG_ALLOWED_ORIGIN = flag.String("allowed-origin", "*", "the allowed websocket origin(s), it accepts a comma separated list of domains, * means anything")
	FLAG_AUTH_BACKEND   = flag.String("auth-webhook", "http://localhost:8000", "the auth endpoint that will validate the response")
	FLAG_REDIS_SERVER   = flag.String("redis", "redis://localhost:6379/1", "the backend server to fallback to")
)

var (
	ALLOWED_ORIGINS = []string{}
	REDIS_CLIENT    *redis.Client
)

func InitFlags() (err error) {
	flag.Parse()
	for _, origin := range strings.Split(*FLAG_ALLOWED_ORIGIN, ",") {
		ALLOWED_ORIGINS = append(ALLOWED_ORIGINS, strings.TrimSpace(origin))
	}
	err, REDIS_CLIENT = RedisConnect(*FLAG_REDIS_SERVER)
	return err
}
