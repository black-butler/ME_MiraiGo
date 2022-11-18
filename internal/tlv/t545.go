package tlv

import "github.com/black-butler/ME_MiraiGo/binary"

func T545(imei []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x545)
		w.WriteBytesShort(imei)
	})
}
