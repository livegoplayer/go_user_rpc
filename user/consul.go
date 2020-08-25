package user

import (
	"fmt"
	"time"

	consulapi "github.com/hashicorp/consul/api"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
)

type ConsulRegister struct {
	ConsulAddress                  string // consul address
	ServiceName                    string // service name
	ServiceIP                      string
	Tags                           []string // consul tags
	ServicePort                    int      //service port
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
	EnableTagOverride              bool
}

func NewConsulRegister(consulAddress, serviceName, serviceIP string, servicePort int, tags []string) *ConsulRegister {
	return &ConsulRegister{
		ConsulAddress:                  consulAddress,
		ServiceName:                    serviceName,
		ServiceIP:                      serviceIP,
		Tags:                           tags,
		ServicePort:                    servicePort,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

func (r *ConsulRegister) NewConsulGRPCRegister(c *consulapi.Config) (registar sd.Registrar) {

	consulClient, err := consulapi.NewClient(c)
	if err != nil {
		panic(err)
	}

	// Produce a fake service registration.
	reg := &consulapi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.ServiceName, r.ServiceIP, r.ServicePort),
		Name:    fmt.Sprintf("grpc.health.v1.%v", r.ServiceName),
		Tags:    r.Tags,
		Port:    r.ServicePort,
		Address: r.ServiceIP,
		Check: &consulapi.AgentServiceCheck{
			// 健康检查间隔
			Interval: r.Interval.String(),
			//grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			GRPC: fmt.Sprintf("%v:%v/%v", r.ServiceIP, r.ServicePort, r.ServiceName),
			// 注销时间，相当于过期时间
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}

	//集成到go kit中
	client := consul.NewClient(consulClient)
	p := consul.NewRegistrar(client, reg, log.NewNopLogger())

	return p
}

func NewServerConfig(consulServerAddr string, token string, scheme string, dataCenter string) *consulapi.Config {
	return &consulapi.Config{
		Address:    consulServerAddr,
		Token:      token,
		Scheme:     scheme,
		Datacenter: dataCenter,
	}
}
