package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	PIX_FMT_RGB24    = C.AV_PIX_FMT_RGB24
)
