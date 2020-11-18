// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package extime

import (
	"time"
)

// ParseInLocal 将时间解析到 time.Local 本地时区，用于解决使用 time.Parse 因缺失时区信息，默认被设置成 UTC，从而引起一系列时区问题。
func ParseInLocal(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}
