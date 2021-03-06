package DependencyInjector

import (
	"log"

	"github.com/facebookgo/inject"
)

//Injector: función en donde se ejecutará la inyección de dependencias
func Injector() {

	//Definimos un contenedor
	container := inject.Graph{}

	/********DataBase********/
	errContainer := DataBaseInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	/********Contollers********/
	errContainer = ControllerInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	/********Repositories********/
	errContainer = RepositoriesInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	/********Services********/
	errContainer = ServicesInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	/********Validators********/
	errContainer = ValidatorInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	/********Querys********/
	errContainer = QueryInjector(&container)

	//Verificamos si existe algún error
	if errContainer != nil {

		log.Fatal(errContainer)
	}

	//Si existe error poblamos el contenedor
	if err := container.Populate(); err != nil {

		log.Fatal(err)
	}
}
