package message

import (
	"encoding/binary"
	"io"
)

type ConnectedPong struct {
	// PingTime is not the Unix Epoch in milliseconds; rather,
	// it represents the duration in milliseconds since the system was started.
	// This is provided by ConnectedPing.PingTime
	PingTime int64
	// PongTime is not the Unix Epoch in milliseconds; rather,
	// it represents the duration in milliseconds since the system was started.
	// You can use this to check the uptime of the server.
	PongTime int64
}

func (pk *ConnectedPong) UnmarshalBinary(data []byte) error {
	if len(data) < 16 {
		return io.ErrUnexpectedEOF
	}
	pk.PingTime = int64(binary.BigEndian.Uint64(data))
	pk.PongTime = int64(binary.BigEndian.Uint64(data[8:]))
	return nil
}

func (pk *ConnectedPong) MarshalBinary() (data []byte, err error) {
	b := make([]byte, 17)
	b[0] = IDConnectedPong
	binary.BigEndian.PutUint64(b[1:], uint64(pk.PingTime))
	binary.BigEndian.PutUint64(b[9:], uint64(pk.PongTime))
	return b, nil
}
