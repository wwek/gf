// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"

	"github.com/gogf/gf/test/gtest"
)

func Test_ConfigFromMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"address":         ":8199",
			"readTimeout":     "60s",
			"indexFiles":      g.Slice{"index.php", "main.php"},
			"errorLogEnabled": true,
			"cookieMaxAge":    "1y",
		}
		config, err := ghttp.ConfigFromMap(m)
		t.Assert(err, nil)
		d1, _ := time.ParseDuration(gconv.String(m["readTimeout"]))
		d2, _ := time.ParseDuration(gconv.String(m["cookieMaxAge"]))
		t.Assert(config.Address, m["address"])
		t.Assert(config.ReadTimeout, d1)
		t.Assert(config.CookieMaxAge, d2)
		t.Assert(config.IndexFiles, m["indexFiles"])
		t.Assert(config.ErrorLogEnabled, m["errorLogEnabled"])
	})
}

func Test_SetConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"Address": ":8199",
			//"ServerRoot":       "/var/www/MyServerRoot",
			"IndexFiles":       g.Slice{"index.php", "main.php"},
			"AccessLogEnabled": true,
			"ErrorLogEnabled":  true,
			"PProfEnabled":     true,
			"LogPath":          "/var/log/MyServerLog",
			"SessionIdName":    "MySessionId",
			"SessionPath":      "/tmp/MySessionStoragePath",
			"SessionMaxAge":    24 * time.Hour,
		}
		s := g.Server()
		err := s.SetConfigWithMap(m)
		t.Assert(err, nil)
	})
}
