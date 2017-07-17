package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/go/docker-machine-driver-idcf"
)

func main() {
	plugin.RegisterDriver(idcf.NewDriver("", ""))
}
