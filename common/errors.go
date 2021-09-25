package common

import (
	"deliverables/common/constants"
)

func HandleError(message string, err error) (bool, constants.PageData) {
	if err != nil {
		return true, constants.PageData{Message: message, Status: "Error"}
	} else {
		return false, constants.PageData{}
	}
}
