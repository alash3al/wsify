package memorybroker

import (
	"github.com/alash3al/wsify/broker"
)

const name = "memory"

func init() {
	broker.Register(name, &Driver{})
}
