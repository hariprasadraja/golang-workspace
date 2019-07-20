// file system cache implementation in Golang
package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var currentDir, _ = os.Getwd()

const (
	cache_dir = "/tmp"        // Cache directory
	expire    = 8 * time.Hour // Hours to keep the cache
)

func Set(k *Key, data interface{}) error {
	t1 := time.Now()
	err := k.validate()
	if err != nil {
		return err
	}

	val, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	key := regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(k.String(), "")
	clean(key)

	file := "filecache." + key + "." + strconv.FormatInt(time.Now().Add(expire).Unix(), 10)
	fpath := filepath.Join(currentDir+cache_dir, file)

	var fmutex sync.RWMutex
	fmutex.Lock()
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err

	}

	defer fp.Close()
	if _, err = fp.Write(val); err != nil {
		return err
	}

	defer fmutex.Unlock()
	t2 := time.Now()
	log.Println("Set filecache: ", key, " : ", t2.Sub(t1))

	return nil
}

func Get(k *Key, dst interface{}) error {
	t := time.Now()
	key := regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(k.String(), "")
	pattern := filepath.Join(currentDir+cache_dir, "filecache."+key+".*")
	files, err := filepath.Glob(pattern)
	if len(files) < 1 || err != nil {
		return errors.New("filecache: no cache file found")
	}

	if _, err = os.Stat(files[0]); err != nil {
		return err
	}

	fp, err := os.OpenFile(files[0], os.O_RDONLY, 0400)
	if err != nil {
		return err
	}
	defer fp.Close()

	var serialized []byte
	buf := make([]byte, 1024)
	for {
		var n int
		n, err = fp.Read(buf)
		serialized = append(serialized, buf[0:n]...)
		if err != nil || err == io.EOF {
			break
		}
	}
	if err = json.Unmarshal(serialized, &dst); err != nil {
		return err
	}

	for _, file := range files {
		exptime, err := strconv.ParseInt(strings.Split(file, ".")[2], 10, 64)
		if err != nil {
			return err
		}

		if exptime < time.Now().Unix() {
			if _, err = os.Stat(file); err == nil {
				os.Remove(file)
			}
		}
	}

	log.Println("Get filecache: ", key, " : ", time.Since(t))
	return nil
}

func clean(key string) {
	pattern := filepath.Join(currentDir+cache_dir, "filecache."+key+".*")
	log.Print(pattern)
	files, _ := filepath.Glob(pattern)
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			os.Remove(file)
		}
	}
}

type Key struct {
	Collection string
	Id         string
}

// Stringify the Key
func (k *Key) String() string {
	data = k.Collection + "_" + k.Id.Hex() + data
	return data
}

func (k *Key) validate() error {
	return nil
}

func main() {
	var customer = map[string]interface{}{
		"name":            "customer 1",
		"firstName":       "",
		"lastName":        "",
		"gender":          "",
		"imageAvailable":  false,
		"imageVersion":    0,
		"gateCode":        ""
	}

	key := Key{
		Collection: "customer", // Set DB collection name
		Id:         "5936895978362e42ac5e02ca",
	}

	err := Set(&key, &customer)
	if err != nil {
		log.Print("error:", err)
	}

	var data interface{}
	err = Get(&key, &data)
	if err != nil {
		log.Print("Error", err)
	}

	log.Print(data)

}
