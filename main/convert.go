package main

import "errors"

//convert an int32 to 4 bytes, little-endian
func INTToBytes(i int32) []byte {
	var bytes []byte
	bytes = append(bytes, byte(i))
	bytes = append(bytes, byte(i>>8))
	bytes = append(bytes, byte(i>>16))
	bytes = append(bytes, byte(i>>24))
	return bytes
}

//convert 4 bytes to an int32, little-endian
//error if byte slice length is not 4
func BytesToINT(bytes []byte) (int32, error) {
	if len(bytes) != 4 {
		return 0, errors.New("length of byte slice invalid")
	}
	i := int32(bytes[0]) + int32(bytes[1])<<8 + int32(bytes[2])<<16 + int32(bytes[3])<<24
	return i, nil
}
