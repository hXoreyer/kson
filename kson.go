package kson

import (
	"encoding/json"
	"os"
)

type Kson interface {
	//Get
	GetObject(key interface{}) Kson
	GetArray(key interface{}) Kson
	GetInt(key interface{}) int64
	GetFloat(key interface{}) float64
	GetString(key interface{}) string
	GetBool(key interface{}) bool
	Get(key interface{}) interface{}
	Length() int
	//Set
	Set(key interface{}, val interface{}) Kson
	//Save
	SaveAsFile(filename string)
	SaveAsBytes() ([]byte, error)
	SaveAsMap() map[string]interface{}
	//????
	get() interface{}
	Append(val ...interface{}) Kson
}

func KparseByFile(fileName string) (Kson, error) {
	f, err := os.Open(fileName)
	defer f.Close()
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

func KparseByBytes(js []byte) (Kson, error) {
	mp := make(map[string]interface{})
	err := json.Unmarshal(js, &mp)
	if err != nil {
		return nil, err
	}
	ret := &kObject{}
	ret.value = mp
	return ret, nil
}

func KparseByMap(mp map[string]interface{}) Kson {
	ret := &kObject{}
	ret.value = mp
	return ret
}

func NewObject() Kson {
	r := &kObject{}
	r.value = make(map[string]interface{})
	return r
}

func NewArray() Kson {
	r := &kArray{}
	r.value = make([]interface{}, 0)
	return r
}
