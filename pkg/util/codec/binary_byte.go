package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
)

// BinaryWrite 将所有的类型 转成byte buffer类型，易于存储
func BinaryWrite(buf *bytes.Buffer, v any) (err error) {
	size := binary.Size(v)
	fmt.Printf("buf size is: %v\n", size)
	if size <= 0 {
		return errors.Wrap(errors.Errorf("encoderPostings binary.Size err, size: %v", size), "binary size error")
	}

	err = binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		err = errors.Wrap(err, "BinaryWrite error")
	}
	return
}

// GobWrite 将所有类型转为 bytes.Buffer 类型
func GobWrite(v any) (buf *bytes.Buffer, err error) {
	if v == nil {
		return buf, errors.Wrap(errors.New("BinaryWrite the value is nil"), "GobWrite error")
	}
	buf = new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err = enc.Encode(v); err != nil {
		err = errors.Wrap(err, "GobWrite encode error")
	}
	return
}

//DecodePosting 解码  return *PostingsList postingslen err
// func DecodePosting(buf []byte)(p *types.)

// BinaryEncoding 二进制编码
func BinaryEncoding(buf *bytes.Buffer, v any) error {
	err := gob.NewEncoder(buf).Encode(v)
	if err != nil {
		err = errors.Wrap(err, "BinaryEncoding error")
	}
	return err
}

// BinaryDecoding 二进制解码
func BinaryDecoding(buf *bytes.Buffer, v any) error {
	err := gob.NewDecoder(buf).Decode(v)
	if err != nil {
		err = errors.Wrap(err, "BinaryDecoding error")
	}
	return err
}
