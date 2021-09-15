package kson

import (
	"encoding/json"
	"os"
)

type Kson interface {
	GetObject(key interface{}) Kson
	GetArray(key interface{}) Kson
	GetInt(key interface{}) int64
	GetFloat(key interface{}) float64
	GetString(key interface{}) string
	GetBool(key interface{}) bool
	Get(key interface{}) interface{}
	Length() int
}

func KparseByFile(fileName string) (Kson, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]interface{})
	if err = json.NewDecoder(f).Decode(&mp); err != nil {
		return nil, err
	}
	ret := &kObject{}
	ret.value = mp
	return ret, nil
}

func KparseByString(js string) (Kson, error) {
	mp := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &mp)
	if err != nil {
		return nil, err
	}
	ret := &kObject{}
	ret.value = mp
	return ret, nil
}
