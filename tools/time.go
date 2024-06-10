package tools

import (
	"github.com/WooRho/rhtool/rhtool_core/rtime"
	"time"
)

const (
	YMD    = "2006-01-02"
	YMDHMS = "2006-01-02 15:04:05"
)

func Time2String(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}
	switch format {
	case YMD:
		return t.Format(YMD)
	case YMDHMS:
		return t.Format(YMDHMS)
	default:
		return t.Format(YMD)
	}
}

func String2Time(str string) time.Time {
	return rtime.Str2Time(str)
}
