package reader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// MultipleReader 代表多重读取器的接口
type MultipleReader interface {
	// Reader 用于获取一个可关闭读取器的实例
	// 后者会持有该多重读取器中的数据
	Reader() io.ReadCloser
}

// myMultipleReader 代表多重读取器的实现类型
type myMultipleReader struct {
	data []byte
}

// NewMultipleReader 用于新建并返回一个多重读取器的实例
func NewMultipleReader(reader io.Reader) (MultipleReader, error) {
	var data []byte
	var err error
	if reader != nil {
		// 读取参数读取器多有底层数据，并忽略io.EOF错误。实际上，当碰到io.EOF错误时，该函数就会返回读到的所有数据
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("multiple reader: couldn't create a new one: %s", err)
		}
	} else {
		data = []byte{}
	}
	return &myMultipleReader{
		data: data,
	}, nil
}

func (rr *myMultipleReader) Reader() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(rr.data))
}
