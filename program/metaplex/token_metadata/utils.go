package token_metadata

import (
	"strconv"

	"github.com/portto/solana-go-sdk/common"
)

// The Metadata Account is responsible for storing additional data attached to tokens. As every account in the Token Metadata program, it derives from the Mint Account of the token using a PDA.
func GetTokenMetaPubkey(mint common.PublicKey) (common.PublicKey, error) {
	metadataAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return metadataAccount, nil
}

// The Master Edition account, derived from a Mint Account, is an important component of NFTs because its existence is proof of the Non-Fungibility of the token.
func GetMasterEdition(mint common.PublicKey) (common.PublicKey, error) {
	msaterEdtion, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return msaterEdtion, nil
}

// The Edition account, derived from a Mint Account, represents an NFT that was copied from a Master Edition NFT.
func GetEdition(mint common.PublicKey) (common.PublicKey, error) {
	msaterEdtion, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return msaterEdtion, nil
}

// Edition Marker accounts are used internally by the program to keep track of which editions were printed for a given Master Edition.
func GetEditionMark(mint common.PublicKey, edition uint64) (common.PublicKey, error) {
	editionNumber := edition / EDITION_MARKER_BIT_SIZE
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
			[]byte(strconv.FormatUint(editionNumber, 10)),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

// Token Record accounts are used by Programmable NFTs only. Since Programmable NFTs add another layer on top of tokens, Token Record accounts enable us to attach custom data to token accounts rather than mint accounts.
func GetTokenRecord(mint, tokenAccount common.PublicKey) (common.PublicKey, error) {
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("token_record"),
			tokenAccount.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

// Metadata Delegate Record accounts are used to store multiple delegate authorities for a given Metadata account.
func GetMetadataDelegateRecord(mint, updateAuthority, delegate common.PublicKey, delegateRole MetadataDelegate) (common.PublicKey, error) {
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			uint8ToBytes(uint8(delegateRole)),
			updateAuthority.Bytes(),
			delegate.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

// Collection Authority Record accounts are used internally by the program to keep track of which authorities are allowed to set and/or verify the collection of the token's Metadata account.
func GetCollectionAuthorityRecord(mint, authotity common.PublicKey) (common.PublicKey, error) {
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("collection_authority"),
			authotity.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

// Use Authority Record accounts are used internally by the program to keep track of which authorities are allowed to reduce the uses of the token's Metadata account.
func GetUseAuthorityRecord(mint, authotity common.PublicKey) (common.PublicKey, error) {
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("user"),
			authotity.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

// Convert uint8 to bytes
func uint8ToBytes(u uint8) []byte {
	b := make([]byte, 8)
	b[0] = byte(u)
	return b
}
