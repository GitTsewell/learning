package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	channel   string
	chaincode string
	key       string

	dbpath string
)

func init() {
	flag.StringVar(&channel, "channel", "mychannel", "Channel name")
	flag.StringVar(&chaincode, "chaincode", "mychaincode", "Chaincode name")
	flag.StringVar(&key, "key", "", "Key to query; empty query all keys")

	flag.StringVar(&dbpath, "dbpath", "/Users/tsewell/.gated/data/application.db", "Path to LevelDB")
}

func readKey(db *leveldb.DB, key string) {
	var b bytes.Buffer
	b.WriteString(channel)
	b.WriteByte(0)
	b.WriteString(chaincode)
	b.WriteByte(0)
	b.WriteString(key)

	value, err := db.Get(b.Bytes(), nil)
	if err != nil {
		fmt.Printf("ERROR: cannot read key[%s], error=[%v]\n", key, err)
		return
	}
	fmt.Printf("Key[%s]=[%s]\n", key, string(value))
}

func readAll(db *leveldb.DB) {
	var b bytes.Buffer
	b.WriteString(channel)
	b.WriteByte(0)
	b.WriteString(chaincode)
	prefix := b.String()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		if strings.HasPrefix(key, prefix) {
			value := string(iter.Value())
			fmt.Printf("Key[%s]=[%s]\n", key, value)
		}
	}
	iter.Release()
	//err := iter.Error()
}

func main() {
	flag.Parse()
	if channel == "" || chaincode == "" || dbpath == "" {
		fmt.Printf("ERROR: Neither of channel, chaincode, key nor dbpath could be empty\n")
		return
	}

	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		fmt.Printf("ERROR: Cannot open LevelDB from [%s], with error=[%v]\n", dbpath, err)
	}
	defer db.Close()

	if key == "" {
		readAll(db)
	} else {
		readKey(db, key)
	}
}
