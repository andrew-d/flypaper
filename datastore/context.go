package datastore

import (
	"golang.org/x/net/context"
)

func FromContext(c context.Context) Datastore {
	return c.Value("TKTK").(Datastore)
}
