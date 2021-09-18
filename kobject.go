package kson

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/hxoreyer/ktry"
)

type kObject struct {
	value map[string]interface{}
}

func (k *kObject) GetObject(key interface{}) Kson {
	var ret kObject
	ktry.Try(func() {
		r := k.value[key.(string)].(map[string]interface{})
		ret.value = r
	}).CatchAll(func(err error) {
		panic(err)
	}).Finally()
	return &ret
}

func (k *kObject) GetArray(key interface{}) Kson {
	var ret kArray
	ret.value = make([]interface{}, 0)
	ktry.Try(func() {
		r := k.value[key.(string)].([]interface{})
		ret.value = append(ret.value, r...)
	}).CatchAll(func(err error) {
		panic(err)
	}).Finally()
	return &ret
}

func (k *kObject) GetInt(key interface{}) int64 {
	var ret int64
	ktry.Try(func() {
		ret = int64(k.value[key.(string)].(float64))
	}).CatchAll(func(err error) {
		ktry.Try(
			func() {
				ret, _ = strconv.ParseInt(k.value[key.(string)].(string), 10, 64)
			}).CatchAll(func(err error) {
			ret = int64(k.value[key.(string)].(int))
		})
	}).Finally()
	return ret
}

func (k *kObject) GetFloat(key interface{}) float64 {
	var ret float64
	ktry.Try(func() {
		ret = k.value[key.(string)].(float64)
	}).CatchAll(func(err error) {
		ret, _ = strconv.ParseFloat(k.value[key.(string)].(string), 64)
	}).Finally()
	return ret
}

func (k *kObject) GetString(key interface{}) string {
	return k.value[key.(string)].(string)
}

func (k *kObject) GetBool(key interface{}) bool {
	var ret bool
	ktry.Try(func() {
		ret = k.value[key.(string)].(bool)
	}).CatchAll(func(err error) {
		ret, _ = strconv.ParseBool(k.value[key.(string)].(string))
	}).Finally()
	return ret
}

func (k *kObject) Get(key interface{}) interface{} {
	return k.value[key.(string)]
}

func (k *kObject) Length() int {
	return len(k.value)
}

func (k *kObject) Set(key interface{}, val interface{}) Kson {
	ktry.Try(func() {
		k.value[key.(string)] = (val).(Kson).get()
	}).CatchAll(func(err error) {
		k.value[key.(string)] = val
	})
	return k
}

func (k *kObject) Append(val ...interface{}) Kson {
	panic("object can not append")
}

func (k *kObject) SaveAsFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	if err = json.NewEncoder(f).Encode(k.value); err != nil {
		panic(err)
	}
}

func (k *kObject) SaveAsBytes() ([]byte, error) {
	js, err := json.Marshal(k.value)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func (k *kObject) SaveAsMap() map[string]interface{} {
	return k.value
}

func (k *kObject) get() interface{} {
	return k.value
}
