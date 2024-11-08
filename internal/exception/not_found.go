package exception

import "encoding/json"

type ErrNotFound Err

func NotFound(err interface{}) ErrNotFound {
	var msg string
	switch err.(type) {
	case string:
		msg = err.(string)
	case error:
		msg = err.(error).Error()
	}

	return ErrNotFound{
		ErrorType: TypeErrorNotFound,
		ErrorCode: 404,
		Message:   msg,
	}
}

func (e ErrNotFound) Error() string {
	var msg string
	if IsHttpError {
		payload, _ := json.Marshal(e)
		msg = string(payload)
	}

	return msg
}
