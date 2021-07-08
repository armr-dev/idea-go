package utils

import (
	"encoding/binary"
	"errors"
	"math"
)

/*
	SplitKey gets a 128-bit key (as byte slice) and split it into
	8 16-bit subkeys.

	! Throws an error if the key is not 128-bit

	@param key []byte - The 128-bit key to be split

	@return k1, k2, k3, k4, k5, k6, k7, k8 - The 8 16-bit subkeys
	@return err - Returns an error if ocurred
*/
func SplitKey(key []byte) ([]byte, []byte, []byte, []byte, []byte, []byte, []byte, []byte, error) {
	var k1, k2, k3, k4, k5, k6, k7, k8 []byte
	var err error = nil

	if (len(key) != 8) {
		err = errors.New("the key length should be 128")
	}

	k1 = key[:15]
	k2 = key[16:31]
	k3 = key[32:47]
	k4 = key[48:63]
	k5 = key[64:79]
	k6 = key[80:95]
	k7 = key[96:111]
	k8 = key[112:127]

	return k1, k2, k3, k4, k5, k6, k7, k8, err
}

/*
	AddModulo takes two bytes, add them and return the result modulo 2¹⁶
*/
func AddModulo(a, b byte) byte {
	var auxA, auxB, res uint32
	
	auxA = uint32(a)
	auxB = uint32(b)
	res = (auxA + auxB) % uint32(math.Pow(2, 16))

	return byte(res)
}

/*
	MultiplyModulo takes two bytes, multiply them and return the result modulo 2¹⁶+1
*/
func MultiplyModulo(a, b byte) byte {
	var auxA, auxB, res uint64
	
	exp := uint64(math.Pow(2, 16) + 1)

	auxA = uint64(a)
	auxB = uint64(b)
	res = (auxA * auxB) % exp

	return byte(res)
}

func MergeText(xL, xR uint32) []byte {
	text1 := make([]byte, 4)
	binary.BigEndian.PutUint32(text1, xL)
	text2 := make([]byte, 4)
	binary.BigEndian.PutUint32(text2, xR)
	text := append(text1, text2...)

	return text
}