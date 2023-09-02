package main

import (
	"github.com/Ihpaz/golang-restapi-userfamily/config"
	"github.com/Ihpaz/golang-restapi-userfamily/repository"
)

var (
	DB                                               = config.Init()
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(DB)
)

func main() {

}
