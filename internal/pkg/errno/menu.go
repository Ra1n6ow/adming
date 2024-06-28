package errno

var (
	ErrMenuAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.MenuAlreadyExist", Message: "Menu already exist."}

	ErrMenuConstraintFail = &Errno{HTTP: 400, Code: "FailedOperation.MenuConstraintFail", Message: "Menu foreign key constraint fails."}

	ErrMenuNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.MenuNotFound", Message: "Menu was not found."}
)
