package main

import (
	"github.com/Ihpaz/golang-restapi-userfamily/config"
	"github.com/Ihpaz/golang-restapi-userfamily/controller"
	"github.com/Ihpaz/golang-restapi-userfamily/middlewares"
	"github.com/Ihpaz/golang-restapi-userfamily/repository"
	"github.com/Ihpaz/golang-restapi-userfamily/routes"
	"github.com/Ihpaz/golang-restapi-userfamily/service"
)

var (
	DB                                               = config.Init()
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(DB)
	customerService    service.CustomerService       = service.NewPostService(customerRepository)
	customerController controller.CustomerController = controller.NewPostController(customerService)
	httpRouter         routes.Router                 = routes.NewMuxRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/customer", customerController.GetCustomers)
	httpRouter.POST("/customer", middlewares.SetMiddlewareJSON(customerController.AddCustomer))
	httpRouter.PUT("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.UpdateCustomer))
	httpRouter.GET("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.GetCustomer))
	httpRouter.DELETE("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.DeleteCustomer))
	httpRouter.SERVE(port)

}
