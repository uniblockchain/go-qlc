/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package config

import "github.com/qlcchain/go-qlc"

type ConfigV2 struct {
	Version             int    `json:"version"`
	DataDir             string `json:"dataDir"`
	StorageMax          string `json:"storageMax"`
	AutoGenerateReceive bool   `json:"autoGenerateReceive"`
	LogLevel            string `json:"logLevel"` //info,warn,debug
	PerformanceEnabled  bool   `json:"performanceEnabled"`

	RPC *RPCConfigV2 `json:"rpc"`
	P2P *P2PConfigV2 `json:"p2p"`
}

type P2PConfigV2 struct {
	BootNodes []string `json:"bootNode"`
	Listen    string   `json:"listen"`
	//Time in seconds between sync block interval
	SyncInterval int                `json:"syncInterval"`
	Discovery    *DiscoveryConfigV2 `json:"discovery"`
	ID           *IdentityConfigV2  `json:"identity"`
}

type RPCConfigV2 struct {
	Enable bool `json:"rpcEnabled"`
	//Listen string `json:"Listen"`
	HTTPEndpoint     string   `json:"httpEndpoint"`
	HTTPEnabled      bool     `json:"httpEnabled"`
	HTTPCors         []string `json:"httpCors"`
	HttpVirtualHosts []string `json:"httpVirtualHosts"`

	WSEnabled  bool   `json:"webSocketEnabled"`
	WSEndpoint string `json:"webSocketEndpoint"`

	IPCEndpoint   string   `json:"ipcEndpoint"`
	IPCEnabled    bool     `json:"ipcEnabled"`
	PublicModules []string `json:"publicModules"`
}

type DiscoveryConfigV2 struct {
	// Time in seconds between remote discovery rounds
	DiscoveryInterval int `json:"discoveryInterval"`
	//The maximum number of discovered nodes at a time
	Limit       int  `json:"limit"`
	MDNSEnabled bool `json:"mDNSEnabled"`
	// Time in seconds between local discovery rounds
	MDNSInterval int `json:"mDNSInterval"`
}

type IdentityConfigV2 struct {
	PeerID  string `json:"peerId"`
	PrivKey string `json:"privateKey,omitempty"`
}

func DefaultConfigV2(dir string) (*ConfigV2, error) {
	pk, id, err := identityConfig()
	if err != nil {
		return nil, err
	}
	var cfg ConfigV2
	modules := []string{"qlcclassic", "ledger", "account", "net", "util", "wallet", "mintage", "contract", "sms"}
	if goqlc.MAINNET {
		cfg = ConfigV2{
			Version:             2,
			DataDir:             dir,
			StorageMax:          "10GB",
			AutoGenerateReceive: false,
			LogLevel:            "error",
			PerformanceEnabled:  false,
			RPC: &RPCConfigV2{
				Enable:           false,
				HTTPEnabled:      true,
				HTTPEndpoint:     "tcp4://0.0.0.0:9735",
				HTTPCors:         []string{"*"},
				HttpVirtualHosts: []string{},
				WSEnabled:        true,
				WSEndpoint:       "tcp4://0.0.0.0:9736",
				IPCEnabled:       true,
				IPCEndpoint:      defaultIPCEndpoint(),
				PublicModules:    modules,
			},
			P2P: &P2PConfigV2{
				BootNodes: []string{
					"/ip4/47.244.138.61/tcp/9734/ipfs/QmdFSukPUMF3t1JxjvTo14SEEb5JV9JBT6PukGRo6A2g4f",
					"/ip4/47.75.145.146/tcp/9734/ipfs/QmW9ocg4fRjckCMQvRNYGyKxQd6GiutAY4HBRxMrGrZRfc",
				},
				Listen:       "/ip4/0.0.0.0/tcp/9734",
				SyncInterval: 120,
				Discovery: &DiscoveryConfigV2{
					DiscoveryInterval: 10,
					Limit:             20,
					MDNSEnabled:       false,
					MDNSInterval:      30,
				},
				ID: &IdentityConfigV2{id, pk},
			},
		}
	} else {
		cfg = ConfigV2{
			Version:             2,
			DataDir:             dir,
			StorageMax:          "10GB",
			AutoGenerateReceive: false,
			LogLevel:            "error",
			PerformanceEnabled:  false,
			RPC: &RPCConfigV2{
				Enable:           false,
				HTTPEnabled:      true,
				HTTPEndpoint:     "tcp4://0.0.0.0:19735",
				HTTPCors:         []string{"*"},
				HttpVirtualHosts: []string{},
				WSEnabled:        true,
				WSEndpoint:       "tcp4://0.0.0.0:19736",
				IPCEnabled:       true,
				IPCEndpoint:      defaultIPCEndpoint(),
				PublicModules:    modules,
			},
			P2P: &P2PConfigV2{
				BootNodes: []string{
					"/ip4/47.103.40.20/tcp/19734/ipfs/QmdFSukPUMF3t1JxjvTo14SEEb5JV9JBT6PukGRo6A2g4f",
					"/ip4/47.112.112.138/tcp/19734/ipfs/QmW9ocg4fRjckCMQvRNYGyKxQd6GiutAY4HBRxMrGrZRfc",
				},
				Listen:       "/ip4/0.0.0.0/tcp/19734",
				SyncInterval: 120,
				Discovery: &DiscoveryConfigV2{
					DiscoveryInterval: 30,
					Limit:             20,
					MDNSEnabled:       false,
					MDNSInterval:      30,
				},
				ID: &IdentityConfigV2{id, pk},
			},
		}
	}

	return &cfg, nil
}
