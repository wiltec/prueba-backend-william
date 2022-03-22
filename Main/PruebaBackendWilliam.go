package main

import (
	"PruebaBackendWilliam/RestApi"
	"PruebaBackendWilliam/Utils/DependencyInjector"
)

func main() {

	//Se crea las inyecciones de las dependencias
	DependencyInjector.Injector()

	//Se registran los servicios
	RestApi.RegisterControllers()
}
