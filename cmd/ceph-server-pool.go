package cmd

import (
	"sync"
	"github.com/ceph/go-ceph/rados"
)

type cephServerPools struct {
	poolMetaMutex sync.RWMutex
	poolMeta      poolMeta
}

func newCephServerPools() error {
	con, err := rados.NewConnWithClusterAndUser("ceph", "test")
	return err
}
