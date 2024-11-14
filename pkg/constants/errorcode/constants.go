package errorcode

type ErrorDefinition struct {
	Code    int
	Message string
}

var (
	NameEmpty = ErrorDefinition{
		Code:    40001,
		Message: "Name is empty",
	}
	ColorEmpty = ErrorDefinition{
		Code:    40002,
		Message: "Color is Empty",
	}
	SizeEmpty = ErrorDefinition{
		Code:    40003,
		Message: "Size is Empty",
	}
	NotFound = ErrorDefinition{
		Code:    40005,
		Message: "Not Found",
	}
	PriceEmpty = ErrorDefinition{
		Code:    40006,
		Message: "Price is Empty",
	}
	StockEmpty = ErrorDefinition{
		Code:    40007,
		Message: "Stock is Empty",
	}
)
