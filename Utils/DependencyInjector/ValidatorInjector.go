package DependencyInjector

import (
	"PruebaBackendWilliam/Validators/Implements"

	"github.com/facebookgo/inject"
)

var PlayerLevelValidator Implements.PlayerLevelValidator
var PlayerPaymentsValidator Implements.PlayerPaymentsValidator

//ValidatorInjector: Agrega objetos al contenedor
func ValidatorInjector(container *inject.Graph) error {

	//Agregamos los objetos de tipo controller al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &PlayerLevelValidator},
		&inject.Object{Value: &PlayerPaymentsValidator},
	)

	return errContainer
}
