package dictionary

const (
	UndisclosedError	= "Something went wrong"
	NoError				= "None"
	NotFoundError		= "Not found"
	InvalidParamError	= "Invalid parameter"
)

type Product struct {
	Id			int 		`json:"id"`
	Nama		string 	`json:"nama"`
	Jenis		string 	`json:"jenis"`
	Jumlah	int			`json:"jumlah"`
	Harga 	int			`json:"harga"`
}

type APIResponse struct {
	Data	 interface{}	`json:"data"`
	Detail interface{}	`json:"detail,omitempty"`
	Error	 string				`json:"error"`
}
