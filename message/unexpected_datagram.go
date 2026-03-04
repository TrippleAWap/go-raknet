package message

type UnexpectedDatagram struct {
	Data []byte
}

func (pk *UnexpectedDatagram) UnmarshalBinary(data []byte) error {
	pk.Data = data
	return nil
}

func (pk *UnexpectedDatagram) MarshalBinary() (data []byte, err error) {
	return pk.Data, nil
}
