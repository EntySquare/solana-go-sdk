package token_metadata

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
	InstructionUpdateMetadataAccount
	InstructionDeprecatedCreateMasterEdition                             // deprecated
	InstructionDeprecatedMintNewEditionFromMasterEditionViaPrintingToken // deprecated
	InstructionUpdatePrimarySaleHappenedViaToken
	InstructionDeprecatedSetReservationList    // deprecated
	InstructionDeprecatedCreateReservationList // deprecated
	InstructionSignMetadata
	InstructionDeprecatedMintPrintingTokensViaToken // deprecated
	InstructionDeprecatedMintPrintingTokens         // deprecated
	InstructionCreateMasterEdition
	InstructionMintNewEditionFromMasterEditionViaToken
	InstructionConvertMasterEditionV1ToV2
	InstructionMintNewEditionFromMasterEditionViaVaultProxy
	InstructionPuffMetadata
	InstructionUpdateMetadataAccountV2
	InstructionCreateMetadataAccountV2
	InstructionCreateMasterEditionV3
	InstructionVerifyCollection
	InstructionUtilize
	InstructionApproveUseAuthority
	InstructionRevokeUseAuthority
	InstructionUnverifyCollection
	InstructionApproveCollectionAuthority
	InstructionRevokeCollectionAuthority
	InstructionSetAndVerifyCollection
	InstructionFreezeDelegatedAccount
	InstructionThawDelegatedAccount
	InstructionRemoveCreatorVerification
	InstructionBurnNft
	InstructionVerifySizedCollectionItem
	InstructionUnverifySizedCollectionItem
	InstructionSetAndVerifySizedCollectionItem
	InstructionCreateMetadataAccountV3
	InstructionSetCollectionSize
	InstructionSetTokenStandard
	InstructionUnknown // FIXME: Could not find instruction in the metaplex docs
	InstructionBurnEditionNft
)

type CreateMetadataAccountParam struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	MintData                Data
}

func CreateMetadataAccount(param CreateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        param.MintData,
		IsMutable:   param.IsMutable,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type UpdateMetadataAccountParam struct {
	MetadataAccount     common.PublicKey
	UpdateAuthority     common.PublicKey
	Data                *Data
	NewUpdateAuthority  *common.PublicKey
	PrimarySaleHappened *bool
}

func UpdateMetadataAccount(param UpdateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction         Instruction
		Data                *Data
		NewUpdateAuthority  *common.PublicKey
		PrimarySaleHappened *bool
	}{
		Instruction:         InstructionUpdateMetadataAccount,
		Data:                param.Data,
		NewUpdateAuthority:  param.NewUpdateAuthority,
		PrimarySaleHappened: param.PrimarySaleHappened,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.MetadataAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type UpdateMetadataAccountV2Param struct {
	MetadataAccount     common.PublicKey // Metadata account
	UpdateAuthority     common.PublicKey // Update authority key
	Data                *DataV2
	NewUpdateAuthority  *common.PublicKey
	PrimarySaleHappened *bool
	IsMutable           *bool
}

// This instruction enables us to update parts of the Metadata account.
// Note that some fields have constraints limiting how they can be updated.
// For instance, once the Is Mutable field is set to False, it cannot be changed back to True.
func UpdateMetadataAccountV2(param UpdateMetadataAccountV2Param) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction         Instruction
		Data                *DataV2
		NewUpdateAuthority  *common.PublicKey
		PrimarySaleHappened *bool
		IsMutable           *bool
	}{
		Instruction:         InstructionUpdateMetadataAccountV2,
		Data:                param.Data,
		NewUpdateAuthority:  param.NewUpdateAuthority,
		PrimarySaleHappened: param.PrimarySaleHappened,
		IsMutable:           param.IsMutable,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.MetadataAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMasterEditionParam struct {
	Edition         common.PublicKey
	Mint            common.PublicKey
	UpdateAuthority common.PublicKey
	MintAuthority   common.PublicKey
	Metadata        common.PublicKey
	Payer           common.PublicKey
	MaxSupply       *uint64
}

func CreateMasterEdition(param CreateMasterEditionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		MaxSupply   *uint64
	}{
		Instruction: InstructionCreateMasterEdition,
		MaxSupply:   param.MaxSupply,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type SignMetadataParam struct {
	Metadata common.PublicKey
	Creator  common.PublicKey
}

func SignMetadata(param SignMetadataParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSignMetadata,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Creator,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type MintNewEditionFromMasterEditionViaTokeParam struct {
	NewMetaData                common.PublicKey
	NewEdition                 common.PublicKey
	MasterEdition              common.PublicKey
	NewMint                    common.PublicKey
	EditionMark                common.PublicKey
	NewMintAuthority           common.PublicKey
	Payer                      common.PublicKey
	TokenAccountOwner          common.PublicKey
	TokenAccount               common.PublicKey
	NewMetadataUpdateAuthority common.PublicKey
	MasterMetadata             common.PublicKey
	Edition                    uint64
}

func MintNewEditionFromMasterEditionViaToken(param MintNewEditionFromMasterEditionViaTokeParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Edition     uint64
	}{
		Instruction: InstructionMintNewEditionFromMasterEditionViaToken,
		Edition:     param.Edition,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.NewMetaData,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.EditionMark,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccountOwner,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.NewMetadataUpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MasterMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMetadataAccountV2Param struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	Data                    DataV2
}

func CreateMetadataAccountV2(param CreateMetadataAccountV2Param) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        DataV2
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccountV2,
		Data:        param.Data,
		IsMutable:   param.IsMutable,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// CreateMetadataAccountV3Param is the parameter for CreateMetadataAccountV3
type CreateMetadataAccountV3Param struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	Data                    DataV2
	CollectionDetails       CollectionDetails // (Optional) This optional enum allows us to differentiate Collection NFTs from Regular NFTs whilst adding additional context such as the amount of NFTs that are linked to the Collection NFT.
}

// CreateMetadataAccountV3 instruction creates and initializes a new Metadata account
// for a given Mint account. It is required that the Mint account has been created and
// initialized by the Token Program before executing this instruction.
func CreateMetadataAccountV3(param CreateMetadataAccountV3Param) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction       Instruction
		Data              DataV2
		IsMutable         bool
		CollectionDetails CollectionDetails
	}{
		Instruction:       InstructionCreateMetadataAccountV3,
		Data:              param.Data,
		IsMutable:         param.IsMutable,
		CollectionDetails: param.CollectionDetails,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMasterEditionV3Param struct {
	Edition         common.PublicKey
	Mint            common.PublicKey
	UpdateAuthority common.PublicKey
	MintAuthority   common.PublicKey
	Metadata        common.PublicKey
	Payer           common.PublicKey
	MaxSupply       *uint64
}

func CreateMasterEditionV3(param CreateMasterEditionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		MaxSupply   *uint64
	}{
		Instruction: InstructionCreateMasterEditionV3,
		MaxSupply:   param.MaxSupply,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type BurnNftParam struct {
	Metadata             common.PublicKey  // Metadata (pda of ['metadata', program id, mint id])
	Owner                common.PublicKey  // NFT owner
	Mint                 common.PublicKey  // Mint of the NFT
	TokenAccount         common.PublicKey  // Token account to close
	MasterEditionAccount common.PublicKey  // MasterEdition2 of the NFT
	CollectionMetadata   *common.PublicKey // Metadata of the Collection
}

func BurnNft(param BurnNftParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionBurnNft,
	})
	if err != nil {
		panic(err)
	}

	accounts := []types.AccountMeta{
		{
			PubKey:     param.Metadata,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PubKey:     param.Owner,
			IsSigner:   true,
			IsWritable: true,
		},
		{
			PubKey:     param.Mint,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PubKey:     param.TokenAccount,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PubKey:     param.MasterEditionAccount,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PubKey:     common.TokenProgramID,
			IsSigner:   false,
			IsWritable: false,
		},
	}

	if param.CollectionMetadata != nil && *param.CollectionMetadata != (common.PublicKey{}) {
		accounts = append(accounts, types.AccountMeta{
			PubKey:     *param.CollectionMetadata,
			IsSigner:   false,
			IsWritable: true,
		})
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type BurnEditionNftParam struct {
	Metadata                  common.PublicKey // Metadata (pda of ['metadata', program id, mint id])
	Owner                     common.PublicKey // NFT owner
	PrintEditionMint          common.PublicKey // Mint of the print edition NFT
	MasterEditionMint         common.PublicKey // Mint of the original/master NFT
	PrintEditionTokenAccount  common.PublicKey // Token account the print edition NFT is in
	MasterEditionTokenAccount common.PublicKey // Token account the Master Edition NFT is in
	MasterEditionAccount      common.PublicKey // MasterEdition2 of the original NFT
	PrintEditionAccount       common.PublicKey // Print Edition account of the NFT
	EditionMarkerAccount      common.PublicKey // Edition Marker PDA of the NFT
}

// This instruction enables the owner of a Print Edition NFT to completely burn it:
//   - burning the SPL token and closing the token account
//   - closing the metadata and edition accounts
//   - giving the owner the reclaimed funds from closing these accounts
//
// https://docs.metaplex.com/programs/token-metadata/instructions#burn-an-edition-nft
func BurnEditionNft(param BurnEditionNftParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionBurnEditionNft,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.PrintEditionMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEditionMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.PrintEditionTokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEditionTokenAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MasterEditionAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.PrintEditionAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.EditionMarkerAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type UpdatePrimarySaleHappenedViaTokenParam struct {
	Metadata common.PublicKey // Metadata key (pda of ['metadata', program id, mint id])
	Owner    common.PublicKey // Owner on the token account
	Token    common.PublicKey // Account containing tokens from the metadata's mint
}

// This instruction flips the Primary Sale Happened flag to True,
// indicating that the first sale has happened. Note that this field is indicative
// and is typically used by marketplaces to calculate royalties.
func UpdatePrimarySaleHappenedViaToken(param UpdatePrimarySaleHappenedViaTokenParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUpdatePrimarySaleHappenedViaToken,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Token,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// RemoveCreatorVerificationParams is the params for RemoveCreatorVerification instruction.
type RemoveCreatorVerificationParam struct {
	Metadata common.PublicKey
	Creator  common.PublicKey
}

// This instruction unverifies one creator on the Metadata account.
// As long as the provided Creator account signs the transaction,
// the Verified boolean will be set to False on the appropriate creator
// of the Creators array.
func RemoveCreatorVerification(param RemoveCreatorVerificationParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRemoveCreatorVerification,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Creator,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// VerifyCollectionParam is the params for VerifyCollection instruction.
type VerifyCollectionParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	Payer                          common.PublicKey // Payer of the transaction
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
}

// This instruction verifies the collection of a Metadata account for unsized parent NFTs,
// by setting the Verified boolean to True on the Collection field.
// Calling it on a collection whose parent NFT has a size field will throw an error.
func VerifyCollection(param VerifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionVerifyCollection,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// VerifySizedCollectionItemParam is the params for VerifySizedCollectionItem instruction.
type VerifySizedCollectionItemParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	Payer                          common.PublicKey // Payer of the transaction
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
	CollectionAuthorityRecord      common.PublicKey // Collection Authority Record PDA
}

// This instruction verifies the collection of a Metadata account,
// by setting the Verified boolean to True on the Collection field,
// and increments the size field of the parent NFT.
// Calling it on a collection whose parent NFT does not have a size field will throw an error.
func VerifySizedCollectionItem(param VerifySizedCollectionItemParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionVerifySizedCollectionItem,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// UnverifyCollectionParam is the params for UnverifyCollection instruction.
type UnverifyCollectionParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
	CollectionAuthorityRecord      common.PublicKey // Collection Authority Record PDA
}

// This instruction unverifies the collection of a Metadata account for unsized parent NFTs,
// by setting the Verified boolean to False on the Collection field.
// Calling it on a collection whose parent NFT has a size field will throw an error.
func UnverifyCollection(param UnverifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUnverifyCollection,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// UnverifySizedCollectionItemParam is the params for UnverifySizedCollectionItem instruction.
type UnverifySizedCollectionItemParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	Payer                          common.PublicKey // Payer of the transaction
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
	CollectionAuthorityRecord      common.PublicKey // Collection Authority Record PDA
}

// This instruction unverifies the collection of a Metadata account, by setting
// the Verified boolean to False on the Collection field, and increments
// the size field of the parent NFT. Calling it on a collection whose parent NFT
// does not have a size field will throw an error.
func UnverifySizedCollectionItem(param UnverifySizedCollectionItemParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUnverifySizedCollectionItem,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// SetAndVerifyCollectionParam is the params for SetAndVerifyCollection instruction.
type SetAndVerifyCollectionParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	Payer                          common.PublicKey // Payer of the transaction
	UpdateAuthority                common.PublicKey // Update Authority of Collection NFT and NFT
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
	CollectionAuthorityRecord      common.PublicKey // Collection Authority Record PDA
}

// SetAndVerifyCollection instruction updates the Collection field of a Metadata account using
// the provided Collection Mint account as long as its Collection Authority
// signs the transaction and the parent NFT does not have the collection details
// field populated (unsized).
func SetAndVerifyCollection(param SetAndVerifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSetAndVerifyCollection,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// SetAndVerifySizedCollectionItemParam is the params for SetAndVerifySizedCollectionItem instruction.
type SetAndVerifySizedCollectionItemParam struct {
	Metadata                       common.PublicKey // Metadata account
	CollectionAuthority            common.PublicKey // Collection Update authority
	Payer                          common.PublicKey // Payer of the transaction
	UpdateAuthority                common.PublicKey // Update Authority of Collection NFT and NFT
	CollectionMint                 common.PublicKey // Mint of the collection
	CollectionMetadata             common.PublicKey // Metadata account of the collection
	CollectionMasterEditionAccount common.PublicKey // MasterEdition2 Account of the Collection Token
	CollectionAuthorityRecord      common.PublicKey // Collection Authority Record PDA
}

// SetAndVerifySizedCollectionItem instruction updates the Collection field of a Metadata account
// for sized collections using the provided Collection Mint account as long as its Collection Authority
// signs the transaction and the parent NFT has the collection details field populated (sized).
func SetAndVerifySizedCollectionItem(param SetAndVerifySizedCollectionItemParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSetAndVerifySizedCollectionItem,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// ApproveCollectionAuthorityParam is the params for ApproveCollectionAuthority instruction.
type ApproveCollectionAuthorityParam struct {
	CollectionAuthorityRecord common.PublicKey // Collection Authority Record PDA
	NewCollectionAuthority    common.PublicKey // New Collection Authority
	UpdateAuthority           common.PublicKey // Update Authority of Collection NFT
	Payer                     common.PublicKey // Payer of the transaction
	CollectionMetadata        common.PublicKey // Metadata account of the collection
	CollectionMint            common.PublicKey // Mint of the collection
}

// ApproveCollectionAuthority instruction allows the provided New Collection Authority account
// to update the Collection field of a Metadata account.
func ApproveCollectionAuthority(param ApproveCollectionAuthorityParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionApproveCollectionAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewCollectionAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// RevokeCollectionAuthorityParam is the params for RevokeCollectionAuthority instruction.
type RevokeCollectionAuthorityParam struct {
	CollectionAuthorityRecord common.PublicKey // Collection Authority Record PDA
	DelegateAuthority         common.PublicKey // Delegated Collection Authority
	RevokeAuthority           common.PublicKey // Update Authority, or Delegated Authority, of Collection NFT
	CollectionMetadata        common.PublicKey // Metadata account of the collection
	CollectionMint            common.PublicKey // Mint of the collection
}

// RevokeCollectionAuthority instruction revokes an existing collection authority,
// meaning they will no longer be able to update the Collection field
// of the Metadata account associated with that Mint account.
func RevokeCollectionAuthority(param RevokeCollectionAuthorityParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRevokeCollectionAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.DelegateAuthority,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.RevokeAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// UtilizeParam is the params for Utilize instruction.
type UtilizeParam struct {
	Metadata           common.PublicKey // Metadata account of the NFT
	TokenAccount       common.PublicKey // Token account of the NFT
	Mint               common.PublicKey // Mint of the NFT
	UseAuthority       common.PublicKey // A Use Authority / Can be the current Owner of the NFT
	Owner              common.PublicKey // Owner of the NFT
	UseAuthorityRecord common.PublicKey // Use Authority Record PDA If present the program Assumes a delegated use authority
	Burner             common.PublicKey // Program As Signer (Burner)
	NumberOfUses       uint64           // Number of times the NFT can be used
}

// Utilize instruction reduces the number of uses of a Metadata account.
// This can either be done by the Update Authority of the Metadata account
// or by an approved Use Authority.
func Utilize(param UtilizeParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction  Instruction
		NumberOfUses uint64
	}{
		Instruction:  InstructionUtilize,
		NumberOfUses: param.NumberOfUses,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UseAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   false,
				IsWritable: false,
			},

			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SPLAssociatedTokenAccountProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.UseAuthorityRecord,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Burner, // common.MetaplexTokenMetaProgramID, ???
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// ApproveUseAuthorityParam is the params for ApproveUseAuthority instruction.
type ApproveUseAuthorityParam struct {
	UseAuthorityRecord common.PublicKey // Use Authority Record PDA If present the program Assumes a delegated use authority
	Owner              common.PublicKey // Owner of the NFT
	Payer              common.PublicKey // Payer of the transaction
	User               common.PublicKey // User to be approved
	OwnerTokenAccount  common.PublicKey // Owned Token Account Of Mint
	Metadata           common.PublicKey // Metadata account of the NFT
	Mint               common.PublicKey // Mint of the NFT
	Burner             common.PublicKey // Program As Signer (Burner)
	NumberOfUses       uint64           // Number of times the NFT can be used
}

// ApproveUseAuthority instruction allows the provided User account to utilize a Metadata account.
// The program keeps track of all the use authorities that have been approved via Use Authority Record PDAs.
func ApproveUseAuthority(param ApproveUseAuthorityParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction  Instruction
		NumberOfUses uint64
	}{
		Instruction:  InstructionApproveUseAuthority,
		NumberOfUses: param.NumberOfUses,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.UseAuthorityRecord,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.User,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.OwnerTokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Burner, // common.MetaplexTokenMetaProgramID, ???
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// RevokeUseAuthorityParam is the params for RevokeUseAuthority instruction.
type RevokeUseAuthorityParam struct {
	UseAuthorityRecord common.PublicKey // Use Authority Record PDA If present the program Assumes a delegated use authority
	Owner              common.PublicKey // Owner of the NFT
	User               common.PublicKey // User to be approved
	OwnerTokenAccount  common.PublicKey // Owned Token Account Of Mint
	Mint               common.PublicKey // Mint of the NFT
	Metadata           common.PublicKey // Metadata account of the NFT
}

// ApproveUseAuthority instruction allows the provided User account to utilize a Metadata account.
// The program keeps track of all the use authorities that have been approved via Use Authority Record PDAs.
func RevokeUseAuthority(param RevokeUseAuthorityParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRevokeUseAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.UseAuthorityRecord,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.User,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.OwnerTokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// FreezeDelegatedAccountParam is the params for FreezeDelegatedAccount instruction.
type FreezeDelegatedAccountParam struct {
	Delegate     common.PublicKey // Delegate
	TokenAccount common.PublicKey // Token Account to freeze
	Edition      common.PublicKey // Edition
	Mint         common.PublicKey // Mint of the NFT
}

// FreezeDelegatedAccount instruction freezes a Token account but only
// if you are the Delegate Authority of the Token account. Because Mint Authority
// and Freeze Authority of NFTs are transferred to the Master Edition / Edition PDA,
// this instruction is the only way for a delegate to prevent the owner of an NFT
// to transfer it. This enables a variety of use-cases such as preventing someone
// to sell its NFT whilst being listed in an escrowless marketplace.
func FreezeDelegatedAccount(param FreezeDelegatedAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionFreezeDelegatedAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Delegate,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// ThawDelegatedAccountParam is the params for ThawDelegatedAccount instruction.
type ThawDelegatedAccountParam struct {
	Delegate     common.PublicKey // Delegate
	TokenAccount common.PublicKey // Token Account to freeze
	Edition      common.PublicKey // Edition
	Mint         common.PublicKey // Mint of the NFT
}

// ThawDelegatedAccount instruction reverts the instruction above by unfreezing
// a Token account, only if you are the Delegate Authority of the Token account.
func ThawDelegatedAccount(param ThawDelegatedAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction  Instruction
		NumberOfUses uint64
	}{
		Instruction: InstructionThawDelegatedAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Delegate,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// SetCollectionSizeParam is the params for SetCollectionSize instruction.
type SetCollectionSizeParam struct {
	CollectionMetadata        common.PublicKey // Collection Metadata account
	CollectionAuthority       common.PublicKey // Collection Update Authority
	CollectionMint            common.PublicKey // Mint of the Collection
	CollectionAuthorityRecord common.PublicKey // Collection Authority Record PDA
	Size                      uint64           // Size of the Collection
}

// SetCollectionSize instruction allows the update authority of a collection parent NFT
// to set the size of the collection once in order to allow existing unsized collections
// to be updated to track size. Once a collection is sized it can only be verified and
// unverified by the sized handlers and can't be changed back to unsized.
func SetCollectionSize(param SetCollectionSizeParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Size        uint64
	}{
		Instruction: InstructionSetCollectionSize,
		Size:        param.Size,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

// SetTokenStandardParam is the params for SetTokenStandard instruction.
type SetTokenStandardParam struct {
	Metadata        common.PublicKey // Metadata account
	UpdateAuthority common.PublicKey // Update Authority
	Mint            common.PublicKey // Mint of the NFT
	Edition         common.PublicKey // Edition account of the NFT
}

// SetTokenStandard instruction allows an update authority to pass in a metadata account
// with an optional edition account and then it determines what the correct TokenStandard type
// is and writes it to the metadata.
// See Token Standard for more information:
//   - https://docs.metaplex.com/programs/token-metadata/token-standard
func SetTokenStandard(param SetTokenStandardParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSetTokenStandard,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
