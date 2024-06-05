package broker

import (
	"fmt"
	"sync"
)

var (
	driversMap   = map[string]Driver{}
	driversMutex = &sync.RWMutex{}
)

func Register(name string, driver Driver) {
	driversMutex.Lock()
	defer driversMutex.Unlock()

	if _, found := driversMap[name]; found {
		panic(fmt.Errorf("specified driver (%s) already registered", name))
	}

	driversMap[name] = driver
}

func Connect(name string, dsn string) (Driver, error) {
	driversMutex.RLock()
	defer driversMutex.RUnlock()

	driver, found := driversMap[name]
	if !found {
		return nil, fmt.Errorf("specified driver (%s) not registered", name)
	}

	if err := driver.Connect(dsn); err != nil {
		return nil, err
	}

	return driver, nil
}
