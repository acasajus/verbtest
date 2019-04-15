package apigw

import (
	"log"

	capi "github.com/hashicorp/consul/api"
)

type discovery struct {
	consul *consul.Client
}

func newSD(addr string) (*sd, err) {
	consulConfig := capi.DefaultConfig()
	if len(*consulAddr) > 0 {
		consulConfig.Address = *consulAddr
	}
	consulClient, err := capi.NewClient(consulConfig)
	if err != nil {
		logger.Log("err", err)
		return nil, fmt.Error("Cannot connect to consul: %s", err)
	}
	return discovery{consulClient}, nil
}

func (d discovery) getSvcAddr(name string) (string, error) {
	passingOnly := true
	addrs, meta, err := d.consul.Health().Service("name", "latest", passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, err
	}
	log.Printf("Discovered svc %s at %#v", name, addrs)
}

func (d discovery) getAuth() (lol, error) {

}
