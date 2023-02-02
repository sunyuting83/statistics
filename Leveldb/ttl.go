package LevelDB

import (
	"encoding/json"
	"time"

	leveldb_errors "github.com/syndtr/goleveldb/leveldb/errors"
)

type CacheType struct {
	Data    []byte `json:"data"`
	Created int64  `json:"created"`
	Expires int64  `json:"expires"`
}

func Get(key string) ([]byte, error) {

	data, err := LevelDB.Get([]byte(key), nil)

	if err != nil && err != leveldb_errors.ErrNotFound {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	var cache CacheType
	err = json.Unmarshal(data, &cache)

	if err != nil {
		return nil, nil
	}

	secs := time.Now().Unix()

	if cache.Expires > 0 && cache.Expires <= secs {
		LevelDB.Delete([]byte(key), nil)
		return nil, nil
	}

	return cache.Data, nil
}

func Set(key string, value string, expires int64) {
	cache := CacheType{Data: []byte(value), Created: time.Now().Unix(), Expires: 0}

	if expires > 0 {
		cache.Expires = cache.Created + expires
	}
	jsonString, _ := json.Marshal(cache)

	LevelDB.Put([]byte(key), jsonString, nil)
}
