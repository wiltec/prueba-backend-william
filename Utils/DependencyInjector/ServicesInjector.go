package DependencyInjector

import (
	"PruebaBackendWilliam/Services/Implements"

	"github.com/facebookgo/inject"
)

var PlayerLevelService Implements.PlayerLevelService

//RepositoriesInjector: Agrega objetos al contenedor
func ServicesInjector(container *inject.Graph) error {

	//Agregamos los objetos de tipo controller al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &PlayerLevelService},
	)

	return errContainer
}
