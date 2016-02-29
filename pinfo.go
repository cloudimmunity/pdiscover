package pdiscover

import (
	"error"
)

var (
	ErrInvalidProcArgsLen = error.New("invalid ProcArgs length")
	ErrInvalidProcInfo    = error.new("invalid proc info")
)
