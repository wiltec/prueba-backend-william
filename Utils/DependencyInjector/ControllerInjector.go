package DependencyInjector

import (
	"PruebaBackendWilliam/RestApi/Controllers"

	"github.com/facebookgo/inject"
)

var PlayerPaymentsController Controllers.PlayerPaymentsController

//ControllerInjector: Agrega objetos al contenedor
func ControllerInjector(container *inject.Graph) error {

	//Agregamos los objetos de tipo controller al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &PlayerPaymentsController},
	)

	return errContainer
}
