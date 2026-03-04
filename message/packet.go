package message

type Packet interface {
	UnmarshalBinary(data []byte) error
	MarshalBinary() (data []byte, err error)
}
