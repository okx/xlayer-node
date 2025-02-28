package apollo

import (
	"github.com/0xPolygonHermez/zkevm-node/jsonrpc"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/apolloconfig/agollo/v4/storage"
)

func (c *Client) loadJsonRPC(value interface{}) {
	dstConf, err := c.unmarshal(value)
	if err != nil {
		log.Fatalf("failed to unmarshal json-rpc config: %v", err)
	}
	// nacos is read from env, so we need to keep and restore it
	nacosConfigRest := c.config.RPC.Nacos
	nacosConfigWebsocket := c.config.RPC.NacosWs

	c.config.RPC = dstConf.RPC

	c.config.RPC.Nacos = nacosConfigRest
	c.config.RPC.NacosWs = nacosConfigWebsocket
	c.config.RPC.DisableAPIs = make([]string, len(dstConf.RPC.DisableAPIs))
	copy(c.config.RPC.DisableAPIs, dstConf.RPC.DisableAPIs)

	log.Infof("loaded json-rpc from apollo config: %+v", value.(string))
}

// fireJsonRPC fires the json-rpc config change
// BatchRequestsEnabled
// BatchRequestsLimit
// GasLimitFactor
// DisableAPIs
func (c *Client) fireJsonRPC(key string, value *storage.ConfigChange) {
	newConf, err := c.unmarshal(value.NewValue)
	if err != nil {
		log.Errorf("failed to unmarshal json-rpc config: %v error: %v", value.NewValue, err)
		return
	}
	log.Infof("apollo json-rpc old config : %+v", value.OldValue.(string))
	log.Infof("apollo json-rpc config changed: %+v", value.NewValue.(string))
	jsonrpc.UpdateConfig(newConf.RPC)
}
