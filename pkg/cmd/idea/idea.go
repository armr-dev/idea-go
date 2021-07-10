package idea

import "github.com/armr-dev/idea/internal/pkg/utils"

func EncryptBlock(block []byte, key []byte) []byte {
	var steps [14]byte

	for i := 0; i < 6; i++ {
		steps[0] = utils.MultiplyModulo(block[0], key[0])
		steps[1] = utils.AddModulo(block[1], key[1])
		steps[2] = utils.AddModulo(block[2], key[2])
		steps[3] = utils.MultiplyModulo(block[3], key[3])
		steps[4] = steps[0] ^ steps[2]
		steps[5] = steps[1] ^ steps[3]
		steps[6] = steps[4] ^ key[4]
		steps[7] = utils.AddModulo(steps[5], steps[6])
		steps[8] = utils.MultiplyModulo(steps[7], key[5])
		steps[9] = utils.AddModulo(steps[6], steps[8])
		steps[10] = steps[0] ^ steps[8]
		steps[11] = steps[2] ^ steps[8]
		steps[12] = steps[1] ^ steps[9]
		steps[13] = steps[3] ^ steps[9]

		block = []byte{steps[10], steps[12], steps[11], steps[13]}
	}

	block = []byte{steps[10], steps[11], steps[12], steps[13]}

	steps[0] = utils.MultiplyModulo(block[0], key[0])
	steps[1] = utils.MultiplyModulo(block[1], key[1])
	steps[2] = utils.MultiplyModulo(block[2], key[2])
	steps[3] = utils.MultiplyModulo(block[3], key[3])

	return []byte{steps[0], steps[1], steps[2], steps[3]}
}

func DecryptBlock(block []byte, key []byte) []byte {
	
}