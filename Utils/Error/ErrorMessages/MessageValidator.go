package ErrorMessages

type MessageValidator struct {
	Code        string
	Message     string
	Name        string
	Description string
}

//New : Retorna un MessageValidator de acuerdo los argumentos deben ser ingresados en el sigiente orden:
// code ; Código de mensaje
// name : Nombre del error
// description: Descripción del error
// message : Si el valor no se ingresa su valor por defecto sera ReturnMessage(code)
func New(args ...string) MessageValidator {

	code := ""

	description := ""
	name := ""

	l := len(args)
	if l >= 1 {

		code = args[0]
	}

	if l >= 2 {

		name = args[1]
	}

	if l >= 3 {

		description = args[2]
	}

	message := ReturnMessage(code)
	if l >= 4 {

		message = args[3]
	}

	return MessageValidator{

		Code:        code,
		Message:     message,
		Name:        name,
		Description: description,
	}
}

//NewWithBool : Retorna un booleano true y un MessageValidator si code = Successfull  y false de lo contrario
// el orden de los parametros den de ser :
// code ; Código de mensaje
// name : Nombre del error
// description: Descripción del error
// message : Si el valor no se ingresa su valor por defecto sera ReturnMessage(code)
func NewWithBool(args ...string) (bool, MessageValidator) {

	code := ""

	if len(args) >= 1 {

		code = args[0]
	}

	return (code == Successfull), New(args...)
}
