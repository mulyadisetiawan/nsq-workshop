package ovo

import "github.com/sharring_session/nsq/repository/ovo"

func parseOVOResponse(code string, err error) ovo.Response {
	if err == nil {
		return ovo.Response{
			Code: code,
		}
	}

	return ovo.Response{
		Code:  code,
		Error: err.Error(),
	}
}
