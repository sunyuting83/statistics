package LevelDB

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	leveldb "github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// GetLevel get data
func GetLevel(k string) (data string, has bool) {
	s, err := LevelDB.Get([]byte(k), nil)
	if err != nil {
		return "", false
	}
	if string(s) == "leveldb: not found" {
		return "", false
	}
	return string(s), true
}

// SetLevel set data
func PutLevel(k string, v string) {
	LevelDB.Put([]byte(k), []byte(v), nil)
}

// Delete set data
func DelLevel(k string) {
	LevelDB.Delete([]byte(k), nil)
}

func FindFirst(Prefix string) []string {
	var list []string
	iter := LevelDB.NewIterator(util.BytesPrefix([]byte(Prefix)), nil)
	for iter.Next() {
		value := iter.Value()
		list = append(list, string(value))
		break
	}
	iter.Release()
	if len(list) > 0 {
		return list
	}
	return list
}

func FindAll(Prefix string) []string {
	var list []string
	iter := LevelDB.NewIterator(util.BytesPrefix([]byte(Prefix)), nil)
	for iter.Next() {
		value := iter.Value()
		list = append(list, string(value))
	}
	iter.Release()
	if len(list) > 0 {
		return list
	}
	return list
}
func FindRawAll(Prefix string) []string {
	var list []string
	iter := LevelDB.NewIterator(util.BytesPrefix([]byte(Prefix)), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		data := strings.Join([]string{string(key), string(value)}, "/////")
		list = append(list, string(data))
	}
	iter.Release()
	if len(list) > 0 {
		return list
	}
	return list
}

func BatchPut(personList []string) {
	LeveBatch := new(leveldb.Batch)
	for _, item := range personList {
		Sub := strings.Split(item, "/////")
		key := Sub[0]
		value := Sub[1]
		LeveBatch.Put([]byte(key), []byte(value))
	}
	LevelDB.Write(LeveBatch, nil)
}
func BatchDelete(personList []string) {
	LeveBatch := new(leveldb.Batch)
	for _, item := range personList {
		LeveBatch.Delete([]byte(item))
	}
	LevelDB.Write(LeveBatch, nil)
}

// GetCurrentPath Get Current Path
func GetCurrentPath() (string, error) {
	OS := runtime.GOOS
	LinkPathStr := "/"
	if OS == "windows" {
		LinkPathStr = "\\"
	}
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(path)
	dbPath := strings.Join([]string{dir, "Cache"}, LinkPathStr)
	return dbPath, nil
}
