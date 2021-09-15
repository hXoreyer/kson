package kson

import (
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
		ret, _ = strconv.ParseInt(k.value[key.(string)].(string), 10, 64)
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
