package main

import "encoding/binary"


/*实现一个从 socket connection 中解码出 goim 协议的解码器。*/

type GoimData struct{
	PackageLength uint32
	HeaderLength uint16
	Version uint16
	Operation uint32
	Sequence uint32
	Body  string
}

func goimDecoder(data []byte) *GoimData{


	packetLen := binary.BigEndian.Uint32(data[:4])
	headerLen := binary.BigEndian.Uint16(data[4:6])
	version := binary.BigEndian.Uint16(data[6:8])
	operation := binary.BigEndian.Uint32(data[8:12])
	sequence := binary.BigEndian.Uint32(data[12:16])
	body := string(data[16:])

	return &GoimData{PackageLength: packetLen,
					HeaderLength: headerLen,
					Version:version,
					Operation:operation,
					Sequence: sequence,
					Body: body,
				}
}