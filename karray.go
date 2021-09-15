package kson

import (
	"strconv"

	"github.com/hxoreyer/ktry"
)

type kArray struct {
	value []interface{}
}

func (k *kArray) GetObject(key interface{}) Kson {
	var ret kObject
	ktry.Try(func() {
		r := k.value[key.(int)].(map[string]interface{})
		ret.value = r
	}).CatchAll(func(err error) {
		panic(err)
	}).Finally()
	return &ret
}

func (k *kArray) GetArray(key interface{}) Kson {
	return k
}

func (k *kArray) GetInt(key interface{}) int64 {
	var ret int64
	ktry.Try(func() {
		ret = int64(k.value[key.(int)].(float64))
	}).CatchAll(func(err error) {
		ret, _ = strconv.ParseInt(k.value[key.(int)].(string), 10, 64)
	}).Finally()
	return ret
}

func (k *kArray) GetFloat(key interface{}) float64 {
	var ret float64
	ktry.Try(func() {
		ret = k.value[key.(int)].(float64)
	}).CatchAll(func(err error) {
		ret, _ = strconv.ParseFloat(k.value[key.(int)].(string), 64)
	}).Finally()
	return ret
}

func (k *kArray) GetString(key interface{}) string {
	return k.value[key.(int)].(string)
}

func (k *kArray) GetBool(key interface{}) bool {
	var ret bool
	ktry.Try(func() {
		ret = k.value[key.(int)].(bool)
	}).CatchAll(func(err error) {
		ret, _ = strconv.ParseBool(k.value[key.(int)].(string))
	}).Finally()
	return ret
}

func (k *kArray) Get(key interface{}) interface{} {
	return k.value[key.(int)]
}

func (k *kArray) Length() int {
	return len(k.value)
}
