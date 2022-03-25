package ErrorMessages

const (
	Id_Not_Found     string = "001"
	Null_Field       string = "002"
	Data_Mismatch    string = "003"
	Required_Data    string = "004"
	Incorrect_Format string = "005"
	Exceeded_Length  string = "006"
	Duplicated_Field string = "007"
	Data_Not_Found   string = "008"
	Error_database   string = "009"
	Successfull      string = "999"
)

//ReturnMessage: retorna un texto de acuerdo al código de error pasado
func ReturnMessage(code string) string {
	switch code {
	case Id_Not_Found:
		return "Id not found"
	case Null_Field:
		return "Null field"
	case Data_Mismatch:
		return "Data mismatch"
	case Required_Data:
		return "Required data"
	case Incorrect_Format:
		return "Incorrect format"
	case Exceeded_Length:
		return "Exceeded length"
	case Duplicated_Field:
		return "This field is duplicated"
	case Data_Not_Found:
		return "Data not found"
	case Error_database:
		return "Error in the data base"
	case Successfull:
		return "Sucess"
	default:
		return "Unknow error"
	}
}
