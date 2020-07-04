// Copyright 2020 Netflix Inc
// Author: Colin McIntosh (colin@netflix.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configuration

import (
	"crypto/tls"
	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/credentials"
	"os"
	"time"
)

type GatewayConfig struct {
	// gNMI client TLS credentials. Setting this will enable client TLS.
	ClientTLSConfig *tls.Config
	// Enable GNMI Server
	EnableServer bool
	// Logger used by the Gateway code
	Log zerolog.Logger
	// OpenConfig Models 'public' folder location
	OpenConfigDirectory string
	// ServerAddress is the address where other cluster members can reach the gNMI server. The first assigned IP address
	// is used if the parameter is not provided.
	ServerAddress string
	// ServerPort is the TCP port where other cluster members can reach the gNMI server. ServerListenPort is used if the
	// parameter is not provided.
	ServerPort int
	// ServerListenAddress is the interface IP address the gNMI server will listen on.
	ServerListenAddress string
	// ServerListenPort is the TCP port the gNMI server will listen on.
	ServerListenPort int
	// gNMI Server TLS credentials. You must specify either this or both ServerTLSCert & ServerTLSKey if you -EnableServer
	ServerTLSCreds credentials.TransportCredentials
	// gNMI Server TLS Cert
	ServerTLSCert string
	// gNMI Server TLS Key
	ServerTLSKey string
	// Configs for connections to targets
	TargetJSONFile string
	// Interval to reload the target configuration JSON file.
	TargetJSONFileReloadInterval time.Duration
	// Timeout for dialing the target connection
	TargetDialTimeout time.Duration
	// Maximum number of targets that this instance will connect to at once.
	TargetLimit int
	// Updates to be dropped prior to being inserted into the cache
	UpdateRejections [][]*gnmipb.PathElem
	// All of the hosts in your Zookeeper cluster. Setting ZookeeperHosts will enable clustering. -EnableServer is a
	// pre-requisite for clustering to be enabled.
	ZookeeperHosts []string
	// Prefix for the lock path in Zookeeper (e.g. /gnmi/gateway)
	ZookeeperPrefix string
	// Zookeeper timeout time. Minimum is 1 second. Failover time is (ZookeeperTimeout * 2).
	ZookeeperTimeout time.Duration
}

func NewDefaultGatewayConfig() *GatewayConfig {
	config := &GatewayConfig{
		Log: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
	return config
}
