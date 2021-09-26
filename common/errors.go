package common

import "deliverables/common/constants"

//HandleError - functio to handle the error messages
func HandleError(message string, err error) (bool, constants.PageData) {
	if err != nil {
		return true, constants.PageData{Message: message, Status: constants.InternalError}
	}
	return false, constants.PageData{}

}
