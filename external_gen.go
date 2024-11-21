package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *cpuTimesStat) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "cpu":
			z.CPU, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "CPU")
				return
			}
		case "user":
			z.User, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "User")
				return
			}
		case "system":
			z.System, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "System")
				return
			}
		case "idle":
			z.Idle, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Idle")
				return
			}
		case "nice":
			z.Nice, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Nice")
				return
			}
		case "iowait":
			z.Iowait, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Iowait")
				return
			}
		case "irq":
			z.Irq, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Irq")
				return
			}
		case "softirq":
			z.Softirq, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Softirq")
				return
			}
		case "steal":
			z.Steal, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Steal")
				return
			}
		case "guest":
			z.Guest, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Guest")
				return
			}
		case "guestNice":
			z.GuestNice, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "GuestNice")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *cpuTimesStat) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "cpu"
	err = en.Append(0x8b, 0xa3, 0x63, 0x70, 0x75)
	if err != nil {
		return
	}
	err = en.WriteString(z.CPU)
	if err != nil {
		err = msgp.WrapError(err, "CPU")
		return
	}
	// write "user"
	err = en.Append(0xa4, 0x75, 0x73, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.User)
	if err != nil {
		err = msgp.WrapError(err, "User")
		return
	}
	// write "system"
	err = en.Append(0xa6, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.System)
	if err != nil {
		err = msgp.WrapError(err, "System")
		return
	}
	// write "idle"
	err = en.Append(0xa4, 0x69, 0x64, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Idle)
	if err != nil {
		err = msgp.WrapError(err, "Idle")
		return
	}
	// write "nice"
	err = en.Append(0xa4, 0x6e, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Nice)
	if err != nil {
		err = msgp.WrapError(err, "Nice")
		return
	}
	// write "iowait"
	err = en.Append(0xa6, 0x69, 0x6f, 0x77, 0x61, 0x69, 0x74)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Iowait)
	if err != nil {
		err = msgp.WrapError(err, "Iowait")
		return
	}
	// write "irq"
	err = en.Append(0xa3, 0x69, 0x72, 0x71)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Irq)
	if err != nil {
		err = msgp.WrapError(err, "Irq")
		return
	}
	// write "softirq"
	err = en.Append(0xa7, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x72, 0x71)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Softirq)
	if err != nil {
		err = msgp.WrapError(err, "Softirq")
		return
	}
	// write "steal"
	err = en.Append(0xa5, 0x73, 0x74, 0x65, 0x61, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Steal)
	if err != nil {
		err = msgp.WrapError(err, "Steal")
		return
	}
	// write "guest"
	err = en.Append(0xa5, 0x67, 0x75, 0x65, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Guest)
	if err != nil {
		err = msgp.WrapError(err, "Guest")
		return
	}
	// write "guestNice"
	err = en.Append(0xa9, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.GuestNice)
	if err != nil {
		err = msgp.WrapError(err, "GuestNice")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *cpuTimesStat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "cpu"
	o = append(o, 0x8b, 0xa3, 0x63, 0x70, 0x75)
	o = msgp.AppendString(o, z.CPU)
	// string "user"
	o = append(o, 0xa4, 0x75, 0x73, 0x65, 0x72)
	o = msgp.AppendFloat64(o, z.User)
	// string "system"
	o = append(o, 0xa6, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d)
	o = msgp.AppendFloat64(o, z.System)
	// string "idle"
	o = append(o, 0xa4, 0x69, 0x64, 0x6c, 0x65)
	o = msgp.AppendFloat64(o, z.Idle)
	// string "nice"
	o = append(o, 0xa4, 0x6e, 0x69, 0x63, 0x65)
	o = msgp.AppendFloat64(o, z.Nice)
	// string "iowait"
	o = append(o, 0xa6, 0x69, 0x6f, 0x77, 0x61, 0x69, 0x74)
	o = msgp.AppendFloat64(o, z.Iowait)
	// string "irq"
	o = append(o, 0xa3, 0x69, 0x72, 0x71)
	o = msgp.AppendFloat64(o, z.Irq)
	// string "softirq"
	o = append(o, 0xa7, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x72, 0x71)
	o = msgp.AppendFloat64(o, z.Softirq)
	// string "steal"
	o = append(o, 0xa5, 0x73, 0x74, 0x65, 0x61, 0x6c)
	o = msgp.AppendFloat64(o, z.Steal)
	// string "guest"
	o = append(o, 0xa5, 0x67, 0x75, 0x65, 0x73, 0x74)
	o = msgp.AppendFloat64(o, z.Guest)
	// string "guestNice"
	o = append(o, 0xa9, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x69, 0x63, 0x65)
	o = msgp.AppendFloat64(o, z.GuestNice)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *cpuTimesStat) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "cpu":
			z.CPU, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CPU")
				return
			}
		case "user":
			z.User, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "User")
				return
			}
		case "system":
			z.System, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "System")
				return
			}
		case "idle":
			z.Idle, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Idle")
				return
			}
		case "nice":
			z.Nice, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Nice")
				return
			}
		case "iowait":
			z.Iowait, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Iowait")
				return
			}
		case "irq":
			z.Irq, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Irq")
				return
			}
		case "softirq":
			z.Softirq, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Softirq")
				return
			}
		case "steal":
			z.Steal, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Steal")
				return
			}
		case "guest":
			z.Guest, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Guest")
				return
			}
		case "guestNice":
			z.GuestNice, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "GuestNice")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *cpuTimesStat) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.CPU) + 5 + msgp.Float64Size + 7 + msgp.Float64Size + 5 + msgp.Float64Size + 5 + msgp.Float64Size + 7 + msgp.Float64Size + 4 + msgp.Float64Size + 8 + msgp.Float64Size + 6 + msgp.Float64Size + 6 + msgp.Float64Size + 10 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *loadAvgStat) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "load1":
			z.Load1, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Load1")
				return
			}
		case "load5":
			z.Load5, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Load5")
				return
			}
		case "load15":
			z.Load15, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Load15")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z loadAvgStat) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "load1"
	err = en.Append(0x83, 0xa5, 0x6c, 0x6f, 0x61, 0x64, 0x31)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Load1)
	if err != nil {
		err = msgp.WrapError(err, "Load1")
		return
	}
	// write "load5"
	err = en.Append(0xa5, 0x6c, 0x6f, 0x61, 0x64, 0x35)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Load5)
	if err != nil {
		err = msgp.WrapError(err, "Load5")
		return
	}
	// write "load15"
	err = en.Append(0xa6, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Load15)
	if err != nil {
		err = msgp.WrapError(err, "Load15")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z loadAvgStat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "load1"
	o = append(o, 0x83, 0xa5, 0x6c, 0x6f, 0x61, 0x64, 0x31)
	o = msgp.AppendFloat64(o, z.Load1)
	// string "load5"
	o = append(o, 0xa5, 0x6c, 0x6f, 0x61, 0x64, 0x35)
	o = msgp.AppendFloat64(o, z.Load5)
	// string "load15"
	o = append(o, 0xa6, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35)
	o = msgp.AppendFloat64(o, z.Load15)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *loadAvgStat) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "load1":
			z.Load1, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Load1")
				return
			}
		case "load5":
			z.Load5, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Load5")
				return
			}
		case "load15":
			z.Load15, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Load15")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z loadAvgStat) Msgsize() (s int) {
	s = 1 + 6 + msgp.Float64Size + 6 + msgp.Float64Size + 7 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *procfsNetDevLine) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "rx_bytes":
			z.RxBytes, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxBytes")
				return
			}
		case "rx_packets":
			z.RxPackets, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxPackets")
				return
			}
		case "rx_errors":
			z.RxErrors, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxErrors")
				return
			}
		case "rx_dropped":
			z.RxDropped, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxDropped")
				return
			}
		case "rx_fifo":
			z.RxFIFO, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxFIFO")
				return
			}
		case "rx_frame":
			z.RxFrame, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxFrame")
				return
			}
		case "rx_compressed":
			z.RxCompressed, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxCompressed")
				return
			}
		case "rx_multicast":
			z.RxMulticast, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "RxMulticast")
				return
			}
		case "tx_bytes":
			z.TxBytes, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxBytes")
				return
			}
		case "tx_packets":
			z.TxPackets, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxPackets")
				return
			}
		case "tx_errors":
			z.TxErrors, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxErrors")
				return
			}
		case "tx_dropped":
			z.TxDropped, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxDropped")
				return
			}
		case "tx_fifo":
			z.TxFIFO, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxFIFO")
				return
			}
		case "tx_collisions":
			z.TxCollisions, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxCollisions")
				return
			}
		case "tx_carrier":
			z.TxCarrier, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxCarrier")
				return
			}
		case "tx_compressed":
			z.TxCompressed, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TxCompressed")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *procfsNetDevLine) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 17
	// write "name"
	err = en.Append(0xde, 0x0, 0x11, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "rx_bytes"
	err = en.Append(0xa8, 0x72, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxBytes)
	if err != nil {
		err = msgp.WrapError(err, "RxBytes")
		return
	}
	// write "rx_packets"
	err = en.Append(0xaa, 0x72, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxPackets)
	if err != nil {
		err = msgp.WrapError(err, "RxPackets")
		return
	}
	// write "rx_errors"
	err = en.Append(0xa9, 0x72, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxErrors)
	if err != nil {
		err = msgp.WrapError(err, "RxErrors")
		return
	}
	// write "rx_dropped"
	err = en.Append(0xaa, 0x72, 0x78, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxDropped)
	if err != nil {
		err = msgp.WrapError(err, "RxDropped")
		return
	}
	// write "rx_fifo"
	err = en.Append(0xa7, 0x72, 0x78, 0x5f, 0x66, 0x69, 0x66, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxFIFO)
	if err != nil {
		err = msgp.WrapError(err, "RxFIFO")
		return
	}
	// write "rx_frame"
	err = en.Append(0xa8, 0x72, 0x78, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxFrame)
	if err != nil {
		err = msgp.WrapError(err, "RxFrame")
		return
	}
	// write "rx_compressed"
	err = en.Append(0xad, 0x72, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxCompressed)
	if err != nil {
		err = msgp.WrapError(err, "RxCompressed")
		return
	}
	// write "rx_multicast"
	err = en.Append(0xac, 0x72, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.RxMulticast)
	if err != nil {
		err = msgp.WrapError(err, "RxMulticast")
		return
	}
	// write "tx_bytes"
	err = en.Append(0xa8, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxBytes)
	if err != nil {
		err = msgp.WrapError(err, "TxBytes")
		return
	}
	// write "tx_packets"
	err = en.Append(0xaa, 0x74, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxPackets)
	if err != nil {
		err = msgp.WrapError(err, "TxPackets")
		return
	}
	// write "tx_errors"
	err = en.Append(0xa9, 0x74, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxErrors)
	if err != nil {
		err = msgp.WrapError(err, "TxErrors")
		return
	}
	// write "tx_dropped"
	err = en.Append(0xaa, 0x74, 0x78, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxDropped)
	if err != nil {
		err = msgp.WrapError(err, "TxDropped")
		return
	}
	// write "tx_fifo"
	err = en.Append(0xa7, 0x74, 0x78, 0x5f, 0x66, 0x69, 0x66, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxFIFO)
	if err != nil {
		err = msgp.WrapError(err, "TxFIFO")
		return
	}
	// write "tx_collisions"
	err = en.Append(0xad, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxCollisions)
	if err != nil {
		err = msgp.WrapError(err, "TxCollisions")
		return
	}
	// write "tx_carrier"
	err = en.Append(0xaa, 0x74, 0x78, 0x5f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxCarrier)
	if err != nil {
		err = msgp.WrapError(err, "TxCarrier")
		return
	}
	// write "tx_compressed"
	err = en.Append(0xad, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TxCompressed)
	if err != nil {
		err = msgp.WrapError(err, "TxCompressed")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *procfsNetDevLine) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 17
	// string "name"
	o = append(o, 0xde, 0x0, 0x11, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "rx_bytes"
	o = append(o, 0xa8, 0x72, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	o = msgp.AppendUint64(o, z.RxBytes)
	// string "rx_packets"
	o = append(o, 0xaa, 0x72, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73)
	o = msgp.AppendUint64(o, z.RxPackets)
	// string "rx_errors"
	o = append(o, 0xa9, 0x72, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73)
	o = msgp.AppendUint64(o, z.RxErrors)
	// string "rx_dropped"
	o = append(o, 0xaa, 0x72, 0x78, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64)
	o = msgp.AppendUint64(o, z.RxDropped)
	// string "rx_fifo"
	o = append(o, 0xa7, 0x72, 0x78, 0x5f, 0x66, 0x69, 0x66, 0x6f)
	o = msgp.AppendUint64(o, z.RxFIFO)
	// string "rx_frame"
	o = append(o, 0xa8, 0x72, 0x78, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65)
	o = msgp.AppendUint64(o, z.RxFrame)
	// string "rx_compressed"
	o = append(o, 0xad, 0x72, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64)
	o = msgp.AppendUint64(o, z.RxCompressed)
	// string "rx_multicast"
	o = append(o, 0xac, 0x72, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74)
	o = msgp.AppendUint64(o, z.RxMulticast)
	// string "tx_bytes"
	o = append(o, 0xa8, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	o = msgp.AppendUint64(o, z.TxBytes)
	// string "tx_packets"
	o = append(o, 0xaa, 0x74, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73)
	o = msgp.AppendUint64(o, z.TxPackets)
	// string "tx_errors"
	o = append(o, 0xa9, 0x74, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73)
	o = msgp.AppendUint64(o, z.TxErrors)
	// string "tx_dropped"
	o = append(o, 0xaa, 0x74, 0x78, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64)
	o = msgp.AppendUint64(o, z.TxDropped)
	// string "tx_fifo"
	o = append(o, 0xa7, 0x74, 0x78, 0x5f, 0x66, 0x69, 0x66, 0x6f)
	o = msgp.AppendUint64(o, z.TxFIFO)
	// string "tx_collisions"
	o = append(o, 0xad, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendUint64(o, z.TxCollisions)
	// string "tx_carrier"
	o = append(o, 0xaa, 0x74, 0x78, 0x5f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72)
	o = msgp.AppendUint64(o, z.TxCarrier)
	// string "tx_compressed"
	o = append(o, 0xad, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64)
	o = msgp.AppendUint64(o, z.TxCompressed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *procfsNetDevLine) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "rx_bytes":
			z.RxBytes, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxBytes")
				return
			}
		case "rx_packets":
			z.RxPackets, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxPackets")
				return
			}
		case "rx_errors":
			z.RxErrors, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxErrors")
				return
			}
		case "rx_dropped":
			z.RxDropped, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxDropped")
				return
			}
		case "rx_fifo":
			z.RxFIFO, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxFIFO")
				return
			}
		case "rx_frame":
			z.RxFrame, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxFrame")
				return
			}
		case "rx_compressed":
			z.RxCompressed, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxCompressed")
				return
			}
		case "rx_multicast":
			z.RxMulticast, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RxMulticast")
				return
			}
		case "tx_bytes":
			z.TxBytes, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxBytes")
				return
			}
		case "tx_packets":
			z.TxPackets, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxPackets")
				return
			}
		case "tx_errors":
			z.TxErrors, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxErrors")
				return
			}
		case "tx_dropped":
			z.TxDropped, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxDropped")
				return
			}
		case "tx_fifo":
			z.TxFIFO, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxFIFO")
				return
			}
		case "tx_collisions":
			z.TxCollisions, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxCollisions")
				return
			}
		case "tx_carrier":
			z.TxCarrier, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxCarrier")
				return
			}
		case "tx_compressed":
			z.TxCompressed, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxCompressed")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *procfsNetDevLine) Msgsize() (s int) {
	s = 3 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.Uint64Size + 11 + msgp.Uint64Size + 10 + msgp.Uint64Size + 11 + msgp.Uint64Size + 8 + msgp.Uint64Size + 9 + msgp.Uint64Size + 14 + msgp.Uint64Size + 13 + msgp.Uint64Size + 9 + msgp.Uint64Size + 11 + msgp.Uint64Size + 10 + msgp.Uint64Size + 11 + msgp.Uint64Size + 8 + msgp.Uint64Size + 14 + msgp.Uint64Size + 11 + msgp.Uint64Size + 14 + msgp.Uint64Size
	return
}
