package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// ZuaMainnetPrivate is the version that is used for
// zua mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var ZuaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// ZuaMainnetPublic is the version that is used for
// zua mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var ZuaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// ZuaTestnetPrivate is the version that is used for
// zua testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var ZuaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// ZuaTestnetPublic is the version that is used for
// zua testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var ZuaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// ZuaDevnetPrivate is the version that is used for
// zua devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var ZuaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// ZuaDevnetPublic is the version that is used for
// zua devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var ZuaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// ZuaSimnetPrivate is the version that is used for
// zua simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var ZuaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// ZuaSimnetPublic is the version that is used for
// zua simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var ZuaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case ZuaMainnetPrivate:
		return ZuaMainnetPublic, nil
	case ZuaTestnetPrivate:
		return ZuaTestnetPublic, nil
	case ZuaDevnetPrivate:
		return ZuaDevnetPublic, nil
	case ZuaSimnetPrivate:
		return ZuaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case ZuaMainnetPrivate:
		return true
	case ZuaTestnetPrivate:
		return true
	case ZuaDevnetPrivate:
		return true
	case ZuaSimnetPrivate:
		return true
	}

	return false
}
