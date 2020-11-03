// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package helper

// Must is a generic helper that act just like
// template.Must (go doc template.Must)
func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
