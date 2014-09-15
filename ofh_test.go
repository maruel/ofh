// Copyright 2013 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

package ofh

import (
	"net/http"
	"testing"
)

type providerStub struct {
	clients []*http.Client
}

func (p *providerStub) GetClient(scope string, r http.RoundTripper) (*http.Client, error) {
	c := p.clients[0]
	p.clients = p.clients[1:]
	return c, nil
}

func TestMakeOAuth2Settings(t *testing.T) {
	MakeOAuth2Settings()
	MakeInstalledApp()
}
