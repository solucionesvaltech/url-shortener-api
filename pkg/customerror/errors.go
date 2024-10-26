package customerror

func SetupError(domain, message string, callerErr string) *CustomError {
	return NewError(Setup, domain, message, callerErr)
}

func MappingError(domain, message string, callerErr string) *CustomError {
	return NewError(Mapping, domain, message, callerErr)
}

func ParsingError(domain, message string, callerErr string) *CustomError {
	return NewError(Parsing, domain, message, callerErr)
}

func ValidationError(domain, message string, callerErr string) *CustomError {
	return NewError(Validation, domain, message, callerErr)
}

func ProcessingError(domain, message string, callerErr string) *CustomError {
	return NewError(Processing, domain, message, callerErr)
}

func SavingError(domain, message string, callerErr string) *CustomError {
	return NewError(Saving, domain, message, callerErr)
}

func ResponseError(domain, message string, callerErr string) *CustomError {
	return NewError(Response, domain, message, callerErr)
}

func DatabaseError(domain, message string, callerErr string) *CustomError {
	return NewError(Database, domain, message, callerErr)
}

func BusinessError(domain, message string, callerErr string) *CustomError {
	return NewError(Business, domain, message, callerErr)
}
