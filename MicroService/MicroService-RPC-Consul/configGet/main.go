package configGet

import (
    "github.com/micro/go-micro/v2/config"
    "github.com/micro/go-plugins/config/source/consul/v2"
)

func Get() (result config.Config) {
    address := "0.0.0.0:8500"
    prefix := "/micro/config"
    source := consul.NewSource(
        consul.WithAddress(address),
        consul.WithPrefix(prefix),
        consul.StripPrefix(false),
    )
    result, _ = config.NewConfig()
    _ = result.Load(source)
    return
}
