package redisbroker

import (
	"github.com/alash3al/wsify/broker"
)

const name = "redis"

func init() {
	broker.Register(name, &Driver{})
}
