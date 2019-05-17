package data

import (
	"encoding/json"
	"github.com/MaJloe3Jlo/perx_test_done/util"
	"log"
	"strconv"
)

// Key is a main data structure.
type Key struct {
	Id     int    `json:"id"`
	Key    string `json:"key"`
	Status string `json:"status"`
}

// func CreateKey is create key in Redis.
func CreateKey(k *Key) {
	c := util.RedisConnect()
	defer c.Close()

	data, err := json.Marshal(k)
	if err != nil {
		log.Fatalf("Error create key: %s", err)
	}

	// Save JSON to redis.
	_, err = c.Do("SET", "key:"+strconv.Itoa(k.Id), data)
	if err != nil {
		log.Fatalf("Error saving to Redis with err: %s", err)
	}
}

// func KeyOff is close key in Redis if it "issued".
func KeyOff(keyid int) string {
	ks := GetAllKeys()
	for _, v := range ks {
		if v.Id == keyid && v.Status == "issued" {
			v.Status = "off"
			CreateKey(&v)
			return "Key " + strconv.Itoa(v.Id) + " was closed."
		}
	}
	return "Cannot off this key."
}

// func KeyInfo returns information of key.
func KeyInfo(keyid int) string {
	ks := GetAllKeys()

	for _, v := range ks {
		if v.Id == keyid {
			return "<b>Info about key:</b> status: " + v.Status
		}
	}

	return "<b>Info about key:</b> not found."
}

// func GetKey is return an key if it's "not issued".
func GetKey() (k Key) {
	ks := GetAllKeys()
	for _, v := range ks {
		if v.Status == "not issued" {
			v.Status = "issued"
			CreateKey(&v)
			return v
		}
	}
	return k
}

// func GetAllKeys is returns all keys from Redis.
func GetAllKeys() (ks []Key) {
	c := util.RedisConnect()
	defer c.Close()
	keys, _ := c.Do("KEYS", "key:*")
	for _, k := range keys.([]interface{}) {
		var key Key
		reply, _ := c.Do("GET", k.([]byte))
		if err := json.Unmarshal(reply.([]byte), &key); err != nil {
			panic(err)
		}
		ks = append(ks, key)
	}
	return ks
}

// func GetStatistics is return an information about keys. How many: "issued", "not issued" or "off"
func GetStatistics() string {
	var (
		issued    int
		notissued int
		off       int
	)
	ks := GetAllKeys()
	for _, v := range ks {
		switch v.Status {
		case "issued":
			issued++
		case "not issued":
			notissued++
		case "off":
			off++
		}
	}

	return " issued: " + strconv.Itoa(issued) + " | not issued: " + strconv.Itoa(notissued) + " | off: " + strconv.Itoa(off)
}
