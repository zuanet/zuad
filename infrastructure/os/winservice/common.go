package winservice

import "github.com/zuanet/zuad/infrastructure/config"

// ServiceDescription contains information about a service, needed to administer it
type ServiceDescription struct {
	Name        string
	DisplayName string
	Description string
}

// MainFunc specifies the signature of an application's main function to be able to run as a windows service
type MainFunc func(startedChan chan<- struct{}) error

// WinServiceMain is only invoked on Windows. It detects when zuad is running
// as a service and reacts accordingly.
var WinServiceMain = func(MainFunc, *ServiceDescription, *config.Config) (bool, error) { return false, nil }