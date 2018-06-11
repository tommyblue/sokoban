package utils

import (
	"bufio"
	"path"
	"runtime"
)

func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func GetRelativePath(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("error")
	}
	return path.Join(path.Dir(filename), filepath)
}
