/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

const (
	GRPC ApiType = 0 + iota // support for 1.0
	HTTP
)

type (
	// Bootstrap the door
	Bootstrap struct {
		StaticResources  StaticResources  `yaml:"static_resources" json:"static_resources" mapstructure:"static_resources"`
		DynamicResources DynamicResources `yaml:"dynamic_resources" json:"dynamic_resources" mapstructure:"dynamic_resources"`
		Metric           Metric           `yaml:"metric" json:"metric" mapstructure:"metric"`
	}

	// StaticResources
	StaticResources struct {
		Listeners      []*Listener     `yaml:"listeners" json:"listeners" mapstructure:"listeners"`
		Clusters       []*Cluster      `yaml:"clusters" json:"clusters" mapstructure:"clusters"`
		Adapters       []*Adapter      `yaml:"adapters" json:"adapters" mapstructure:"adapters"`
		ShutdownConfig *ShutdownConfig `yaml:"shutdown_config" json:"shutdown_config" mapstructure:"shutdown_config"`
		PprofConf      PprofConf       `yaml:"pprofConf" json:"pprofConf" mapstructure:"pprofConf"`
	}

	// DynamicResources TODO
	DynamicResources struct {
		adsConfig ApiConfigSourceConfig `yaml:"ads_config" json:"ads_config" mapstructure:"ads_config"`
		cdsConfig ConfigSourceConfig    `yaml:"cds_config" json:"cds_config" mapstructure:"cds_config"`
		ldsConfig ConfigSourceConfig    `yaml:"lds_config" json:"lds_config" mapstructure:"lds_config"`
	}

	// ConfigSourceConfig
	ConfigSourceConfig struct {
		apiConfigSource ApiConfigSourceConfig `yaml:"api_config_source" json:"api_config_source" mapstructure:"api_config_source"`
		// When set, ADS will be used to fetch resources
		ads AggregatedConfigSource `yaml:"ads" json:"ads" mapstructure:"ads"`
	}

	// AggregatedConfigSource
	AggregatedConfigSource struct {
	}

	// API configuration source to fetch xDS api
	ApiConfigSourceConfig struct {
		apiType             ApiType           `yaml:"api_type" json:"api_type" mapstructure:"api_type"`
		transportApiVersion string            `yaml:"transport_api_version" json:"transport_api_version" mapstructure:"transport_api_version"`
		grpcService         GrpcServiceConfig `yaml:"grpc_services" json:"grpc_services" mapstructure:"grpc_services"`
	}

	// GrpcServiceConfig
	GrpcServiceConfig struct {
		clusterName string `yaml:"cluster_name" json:"cluster_name" mapstructure:"cluster_name"`
	}

	// DiscoveryType
	ApiType int32

	// ShutdownConfig how to shutdown server.
	ShutdownConfig struct {
		Timeout      string `default:"60s" yaml:"timeout" json:"timeout,omitempty"`
		StepTimeout  string `default:"10s" yaml:"step_timeout" json:"step_timeout,omitempty"`
		RejectPolicy string `default:"immediacy" yaml:"reject_policy" json:"reject_policy,omitempty"`
	}

	// APIMetaConfig how to find api config, file or etcd etc.
	APIMetaConfig struct {
		Address       string `yaml:"address" json:"address,omitempty"`
		APIConfigPath string `default:"/pixiu/config/api" yaml:"api_config_path" json:"api_config_path,omitempty" mapstructure:"api_config_path"`
	}

	// TimeoutConfig the config of ConnectTimeout and RequestTimeout
	TimeoutConfig struct {
		ConnectTimeoutStr string `default:"5s" yaml:"connect_timeout" json:"connect_timeout,omitempty"` // ConnectTimeout timeout for connect to cluster node
		RequestTimeoutStr string `default:"10s" yaml:"request_timeout" json:"request_timeout,omitempty"`
	}
)

// GetListeners
func (bs *Bootstrap) GetListeners() []*Listener {
	return bs.StaticResources.Listeners
}

func (bs *Bootstrap) GetStaticListeners() []*Listener {
	return bs.StaticResources.Listeners
}

// GetPprof
func (bs *Bootstrap) GetPprof() PprofConf {
	return bs.StaticResources.PprofConf
}

// ExistCluster
func (bs *Bootstrap) ExistCluster(name string) bool {
	if len(bs.StaticResources.Clusters) > 0 {
		for _, v := range bs.StaticResources.Clusters {
			if v.Name == name {
				return true
			}
		}
	}

	return false
}
