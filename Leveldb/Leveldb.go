package LevelDB

import (
	"fmt"
	"sync"

	leveldb "github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// LevelDB a
type DB struct {
	sync.RWMutex
	// contains filtered or unexported fields
}

var (
	LevelDB *leveldb.DB
	Errdb   error
)

func init() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	DbPath, _ := GetCurrentPath()
	LevelDB, Errdb = leveldb.OpenFile(DbPath, o)
	if Errdb != nil {
		fmt.Println(Errdb)
	}
}
