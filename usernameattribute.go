package stun

import (
  "bytes"
  "encoding/binary"
  "errors"
)

type UsernameAttribute struct {
  Username string
}

func (h *UsernameAttribute) Type() (AttributeType) {
  return Username
}

func (h *UsernameAttribute) Encode(_ *Message) ([]byte, error) {
  buf := new(bytes.Buffer)
  err := binary.Write(buf, binary.BigEndian, attributeHeader(Attribute(h)))
  err = binary.Write(buf, binary.BigEndian, h.Username)

  if err != nil {
    return nil, err
  }
  return buf.Bytes(), nil
}

func (h *UsernameAttribute) Decode(data []byte, length uint16, _ *Header) (error) {
  if uint16(len(data)) < length {
    return errors.New("Truncated Username Attribute")
  }
  h.Username = string(data[0:length])
  return nil
}

func (h *UsernameAttribute) Length() (uint16) {
  return uint16(len(h.Username))
}
