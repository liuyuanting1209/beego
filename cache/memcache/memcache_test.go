// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memcache

import (
	_ "github.com/bradfitz/gomemcache/memcache"

	"strconv"
	"testing"
	"time"

	"github.com/liuyuanting1209/beego/cache"
)

func TestMemcacheCache(t *testing.T) {
	bm, err := cache.NewCache("memcache", `{"conn": "127.0.0.1:11211"}`)
	if err != nil {
		t.Error("init err")
	}
	timeoutDuration := 10 * time.Second
	if err = bm.Put("liuyuanting1209", "1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("liuyuanting1209") {
		t.Error("check err")
	}

	time.Sleep(11 * time.Second)

	if bm.IsExist("liuyuanting1209") {
		t.Error("check err")
	}
	if err = bm.Put("liuyuanting1209", "1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if v, err := strconv.Atoi(string(bm.Get("liuyuanting1209").([]byte))); err != nil || v != 1 {
		t.Error("get err")
	}

	if err = bm.Incr("liuyuanting1209"); err != nil {
		t.Error("Incr Error", err)
	}

	if v, err := strconv.Atoi(string(bm.Get("liuyuanting1209").([]byte))); err != nil || v != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("liuyuanting1209"); err != nil {
		t.Error("Decr Error", err)
	}

	if v, err := strconv.Atoi(string(bm.Get("liuyuanting1209").([]byte))); err != nil || v != 1 {
		t.Error("get err")
	}
	bm.Delete("liuyuanting1209")
	if bm.IsExist("liuyuanting1209") {
		t.Error("delete err")
	}

	//test string
	if err = bm.Put("liuyuanting1209", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("liuyuanting1209") {
		t.Error("check err")
	}

	if v := bm.Get("liuyuanting1209").([]byte); string(v) != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err = bm.Put("liuyuanting12091", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("liuyuanting12091") {
		t.Error("check err")
	}

	vv := bm.GetMulti([]string{"liuyuanting1209", "liuyuanting12091"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if string(vv[0].([]byte)) != "author" && string(vv[0].([]byte)) != "author1" {
		t.Error("GetMulti ERROR")
	}
	if string(vv[1].([]byte)) != "author1" && string(vv[1].([]byte)) != "author" {
		t.Error("GetMulti ERROR")
	}

	// test clear all
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}
}
