package header

import "project/lib/errors"

type Header struct {
	Errno int    `json:"errno"`
	Error string `json:"error"`
}

func (h *Header) SetError(errno int) {
	h.Errno = errno
	if h.Errno != errors.ERROR_NONE {
		if err, ok := errors.AllErrors[h.Errno]; ok {
			h.Error = err
		}
	}
}
