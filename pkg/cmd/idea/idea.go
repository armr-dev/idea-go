package idea

import "github.com/armr-dev/idea/internal/pkg/utils"

func EncryptBlock(block []byte, key []byte) []byte {
	var z [14]byte

	z[0] = utils.MultiplyModulo(block[0], key[0])
	z[1] = utils.AddModulo(block[1], key[1])
	z[2] = utils.AddModulo(block[2],  key[2])
	z[3] = utils.MultiplyModulo(block[3], key[3])
	z[4] = z[0] ^ z[2]
	z[5] = z[1] ^ z[3]
	z[6] = z[4] ^ key[4]
	z[7] = utils.AddModulo(z[5], key[6])
	z[8] = utils.MultiplyModulo(z[7], key[5])
	z[9] = utils.AddModulo(z[6], z[8])
	z[10] = z[0] ^ z[8]
	z[11] = z[2] ^ z[8]
	z[12] = z[1] ^ z[9]
	z[13] = z[3] ^ z[9]


}

/*
11. Step 1 ^ Step 9
12. Step 3 ^ Step 9
13. Step 2 ^ Step 10
14. Step 4 ^ Step 10
*/