package DependencyInjector

import (
	"PruebaBackendWilliam/DataAccess/Repositories/Implements"

	"github.com/facebookgo/inject"
)

var PlayerLevelRepository Implements.PlayerLevelRepository
var HelperRepository Implements.HelperRepository

//RepositoriesInjector: Agrega objetos al contenedor
func RepositoriesInjector(container *inject.Graph) error {

	//Agregamos los objetos de tipo controller al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &PlayerLevelRepository},
		&inject.Object{Value: &HelperRepository},
	)

	return errContainer
}
