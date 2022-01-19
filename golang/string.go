package golang

import (
	"bytes"
	"strings"
	"sync"
)

var (
	sPool = sync.Pool{
		New: func() interface{} {
			s :=&strings.Builder{}
			s.Reset()
			return s
		},
	}
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte,0,100)
			return bytes.NewBuffer(buf)
		},
	}
)

func AddStringsByAdd(is ...string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return is[0]
	}

	s := ""
	for i := range is {
		s += is[i]
	}
	return s
}


func AddStringsByStringBufferNoPool(is ...string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return is[0]
	}

	buf := strings.Builder{}

	for i := range is {
		buf.WriteString(is[i])
	}

	return buf.String()
}

func AddStringsByStringBufferPool(is ...string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return is[0]
	}
	buf := sPool.Get().(*strings.Builder)
	for _, i := range is {
		buf.WriteString(i)
	}
	s := buf.String()
	buf.Reset()
	sPool.Put(buf)
	return s
}

func AddStringsByByteBufferPool(is ...string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return is[0]
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(i)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}


func AddStringsByByteBufferNoPool(is ...string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return is[0]
	}

	buf := bytes.NewBuffer(make([]byte,0,100))
	for _, i := range is {
		buf.WriteString(i)
	}
	s := buf.String()
	buf.Reset()
	return s
}

