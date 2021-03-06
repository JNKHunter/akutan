// Copyright 2019 eBay Inc.
// Primary authors: Simon Fell, Diego Ongaro,
//                  Raymond Kroeker, and Sathish Kandasamy.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discoveryfactory

import (
	"context"
	"net"

	"github.com/ebay/akutan/config"
	"github.com/ebay/akutan/discovery"
)

func init() {
	impls["static"] = &discoveryImpl{setUp: setUpStatic}
}

func setUpStatic() (locatorFactory, error) {
	return newStaticLocator, nil
}

func newStaticLocator(ctx context.Context, cfg *config.Locator) (discovery.Locator, error) {
	endpoints := make([]*discovery.Endpoint, len(cfg.Addresses))
	for i := range cfg.Addresses {
		host, port, err := net.SplitHostPort(cfg.Addresses[i])
		if err != nil {
			return nil, err
		}
		endpoints[i] = &discovery.Endpoint{
			Network: "tcp",
			Host:    host,
			Port:    port,
		}
	}
	return discovery.NewStaticLocator(endpoints), nil
}
