package config

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

const AelConfigFile = ".aelconfig.json"

type Node struct {
	IP   net.IP `json:"address"`
	Name string `json:"name"`
}

type TrackedFile struct {
	Path string `json:"path"`
}

type AelConfig struct {
	CreatedAt int64         `json:"created_at"`
	Nodes     []Node        `json:"nodes"`
	Tracked   []TrackedFile `json:"tracked_files"`
}

func EmptyAelConfig() AelConfig {
	return AelConfig{
		CreatedAt: time.Now().UnixNano(),
		Nodes:     []Node{},
		Tracked:   []TrackedFile{},
	}
}

func DoesConfigExist() bool {
	_, err := os.Stat(AelConfigFile)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func (config *AelConfig) FindNodeByName(name string) (node *Node, index int) {
	for i, node := range config.Nodes {
		if node.Name == name {
			return &node, i
		}
	}

	return nil, -1
}

func (config *AelConfig) FindNodeByAddress(address net.IP) (node *Node, index int) {
	for i, node := range config.Nodes {
		if node.IP.Equal(address) {
			return &node, i
		}
	}

	return nil, -1
}

func (config *AelConfig) StoreConfiguration() error {
	file, err := os.OpenFile(AelConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return cli.Exit("Could not create configuration file: "+err.Error(), 1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(config)
	if err != nil {
		return cli.Exit("Could not write configuration file: "+err.Error(), 1)
	}

	return nil
}

func StoreEmptyConfiguration() error {
	cfg := EmptyAelConfig()
	return cfg.StoreConfiguration()
}

func ReadConfiguration(optional bool) (error, *AelConfig) {
	var file *os.File
	var config AelConfig

	_, err := os.Stat(AelConfigFile)
	if err != nil {
		goto fail
	}

	file, err = os.Open(AelConfigFile)
	if err != nil {
		goto fail
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		goto fail
	}

	return nil, &config

fail:
	if optional {
		return nil, nil
	}

	return cli.Exit("Could not read or parse configuration file: "+err.Error(), 1), nil
}
