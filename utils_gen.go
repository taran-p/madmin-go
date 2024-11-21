package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *NodeCommon) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "addr":
			z.Addr, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Addr")
				return
			}
		case "error":
			z.Error, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Error")
				return
			}
			zb0001Mask |= 0x1
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if zb0001Mask != 0x1 {
		if (zb0001Mask & 0x1) == 0 {
			z.Error = ""
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NodeCommon) EncodeMsg(en *msgp.Writer) (err error) {
	// check for omitted fields
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 2 bits */
	_ = zb0001Mask
	if z.Error == "" {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "addr"
	err = en.Append(0xa4, 0x61, 0x64, 0x64, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Addr)
	if err != nil {
		err = msgp.WrapError(err, "Addr")
		return
	}
	if (zb0001Mask & 0x2) == 0 { // if not omitted
		// write "error"
		err = en.Append(0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
		if err != nil {
			return
		}
		err = en.WriteString(z.Error)
		if err != nil {
			err = msgp.WrapError(err, "Error")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NodeCommon) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// check for omitted fields
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 2 bits */
	_ = zb0001Mask
	if z.Error == "" {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "addr"
	o = append(o, 0xa4, 0x61, 0x64, 0x64, 0x72)
	o = msgp.AppendString(o, z.Addr)
	if (zb0001Mask & 0x2) == 0 { // if not omitted
		// string "error"
		o = append(o, 0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
		o = msgp.AppendString(o, z.Error)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NodeCommon) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "addr":
			z.Addr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Addr")
				return
			}
		case "error":
			z.Error, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Error")
				return
			}
			zb0001Mask |= 0x1
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if zb0001Mask != 0x1 {
		if (zb0001Mask & 0x1) == 0 {
			z.Error = ""
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NodeCommon) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Addr) + 6 + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TimedAction) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "count":
			z.Count, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Count")
				return
			}
		case "acc_time_ns":
			z.AccTime, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "AccTime")
				return
			}
		case "bytes":
			z.Bytes, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Bytes")
				return
			}
			zb0001Mask |= 0x1
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if zb0001Mask != 0x1 {
		if (zb0001Mask & 0x1) == 0 {
			z.Bytes = 0
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TimedAction) EncodeMsg(en *msgp.Writer) (err error) {
	// check for omitted fields
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 3 bits */
	_ = zb0001Mask
	if z.Bytes == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "count"
	err = en.Append(0xa5, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Count)
	if err != nil {
		err = msgp.WrapError(err, "Count")
		return
	}
	// write "acc_time_ns"
	err = en.Append(0xab, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.AccTime)
	if err != nil {
		err = msgp.WrapError(err, "AccTime")
		return
	}
	if (zb0001Mask & 0x4) == 0 { // if not omitted
		// write "bytes"
		err = en.Append(0xa5, 0x62, 0x79, 0x74, 0x65, 0x73)
		if err != nil {
			return
		}
		err = en.WriteUint64(z.Bytes)
		if err != nil {
			err = msgp.WrapError(err, "Bytes")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TimedAction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// check for omitted fields
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 3 bits */
	_ = zb0001Mask
	if z.Bytes == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "count"
	o = append(o, 0xa5, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.Count)
	// string "acc_time_ns"
	o = append(o, 0xab, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73)
	o = msgp.AppendUint64(o, z.AccTime)
	if (zb0001Mask & 0x4) == 0 { // if not omitted
		// string "bytes"
		o = append(o, 0xa5, 0x62, 0x79, 0x74, 0x65, 0x73)
		o = msgp.AppendUint64(o, z.Bytes)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TimedAction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "count":
			z.Count, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Count")
				return
			}
		case "acc_time_ns":
			z.AccTime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AccTime")
				return
			}
		case "bytes":
			z.Bytes, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bytes")
				return
			}
			zb0001Mask |= 0x1
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if zb0001Mask != 0x1 {
		if (zb0001Mask & 0x1) == 0 {
			z.Bytes = 0
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TimedAction) Msgsize() (s int) {
	s = 1 + 6 + msgp.Uint64Size + 12 + msgp.Uint64Size + 6 + msgp.Uint64Size
	return
}
