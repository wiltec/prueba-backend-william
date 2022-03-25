package DependencyInjector

import (
	"PruebaBackendWilliam/DataAccess/Queries/Implements"

	"github.com/facebookgo/inject"
)

var PlayerLevelQuery Implements.PlayerLevelQuery

//QueryInjector: Agrega objetos al contenedor
func QueryInjector(container *inject.Graph) error {

	//Agregamos los objetos de tipo controller al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &PlayerLevelQuery},
	)

	return errContainer
}
