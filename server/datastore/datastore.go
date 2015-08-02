package datastore

type Datastore interface {
	RegionStore
	HostStore
	PortStore
}
