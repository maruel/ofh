// Copyright 2013 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

package aedmz

import (
	"github.com/maruel/ut"
	"net/http"
	"testing"
)

func TestAppIdentity(t *testing.T) {
	app := newAppMock(nil)
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	ut.AssertEqual(t, nil, err)
	c := app.NewContext(req)
	defer CloseRequest(c)

	ut.AssertEqual(t, "Yo", c.AppID())
	ut.AssertEqual(t, "v1", c.AppVersion())
}

func TestConnectivity(t *testing.T) {
	app := newAppMock(nil)
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	ut.AssertEqual(t, nil, err)
	c := app.NewContext(req)
	defer CloseRequest(c)

	r, err := c.HttpClient()
	if r == nil {
		t.Fatal("Expected transport")
	}
	ut.AssertEqual(t, nil, err)

	o, err := c.OAuth2HttpClient("scope")
	if o == nil {
		t.Fatal("Expected transport")
	}
	ut.AssertEqual(t, nil, err)
}
