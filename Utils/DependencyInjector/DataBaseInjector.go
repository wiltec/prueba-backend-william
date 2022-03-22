package DependencyInjector

import (
	"PruebaBackendWilliam/Infraestructure/DataBase"

	"github.com/facebookgo/inject"
)

var DbHandler DataBase.DBHandler

//ControllerInjector: Agrega objetos al contenedor
func DataBaseInjector(container *inject.Graph) error {

	//Agregamos los objetos DataBase al contenedor
	errContainer := container.Provide(
		&inject.Object{Value: &DbHandler},
	)

	//Devolvemos el error en caso de existir
	return errContainer
}
