package transactionvalidator

import (
	"github.com/zuanet/zuad/domain/consensus/model"
	"github.com/zuanet/zuad/domain/consensus/model/externalapi"
	"github.com/zuanet/zuad/domain/consensus/utils/txscript"
	"github.com/zuanet/zuad/domain/consensus/vprog" 
	"github.com/zuanet/zuad/util/txmass"
)

const (
	sigCacheSize = 10_000
	// Tambah constant untuk vprog
	maxVProgCodeSize        = 100_000 // 100KB max code size
	maxVProgDataSize        = 50_000  // 50KB max input data
	maxVProgGasLimit        = 1_000_000 // Max gas per transaction
	maxVProgTotalGas        = 10_000_000 // Max gas per block
	supportedVProgVersion   = 1
)

// transactionValidator exposes a set of validation classes, after which
// it's possible to determine whether either a transaction is valid
type transactionValidator struct {
	blockCoinbaseMaturity                   uint64
	databaseContext                         model.DBReader
	pastMedianTimeManager                   model.PastMedianTimeManager
	ghostdagDataStore                       model.GHOSTDAGDataStore
	daaBlocksStore                          model.DAABlocksStore
	enableNonNativeSubnetworks              bool
	maxCoinbasePayloadLength                uint64
	ghostdagK                               externalapi.KType
	coinbasePayloadScriptPublicKeyMaxLength uint8
	sigCache                                *txscript.SigCache
	sigCacheECDSA                           *txscript.SigCacheECDSA
	txMassCalculator                        *txmass.Calculator
	
	// TAMBAH FIELD UNTUK VPROG
	vprogEngine                             *vprog.Engine
	maxVProgCodeSize                        uint64
	maxVProgDataSize                        uint64
	maxVProgGasLimit                        uint64
	maxVProgTotalGas                        uint64
	supportedVProgVersion                   byte
}

// New instantiates a new TransactionValidator
func New(
	blockCoinbaseMaturity uint64,
	enableNonNativeSubnetworks bool,
	maxCoinbasePayloadLength uint64,
	ghostdagK externalapi.KType,
	coinbasePayloadScriptPublicKeyMaxLength uint8,
	databaseContext model.DBReader,
	pastMedianTimeManager model.PastMedianTimeManager,
	ghostdagDataStore model.GHOSTDAGDataStore,
	daaBlocksStore model.DAABlocksStore,
	txMassCalculator *txmass.Calculator,
) model.TransactionValidator {

	// Buat vprog engine
	vprogEngine := vprog.NewEngine(databaseContext)

	return &transactionValidator{
		blockCoinbaseMaturity:                   blockCoinbaseMaturity,
		enableNonNativeSubnetworks:              enableNonNativeSubnetworks,
		maxCoinbasePayloadLength:                maxCoinbasePayloadLength,
		ghostdagK:                               ghostdagK,
		coinbasePayloadScriptPublicKeyMaxLength: coinbasePayloadScriptPublicKeyMaxLength,
		databaseContext:                         databaseContext,
		pastMedianTimeManager:                   pastMedianTimeManager,
		ghostdagDataStore:                       ghostdagDataStore,
		daaBlocksStore:                          daaBlocksStore,
		sigCache:                                txscript.NewSigCache(sigCacheSize),
		sigCacheECDSA:                           txscript.NewSigCacheECDSA(sigCacheSize),
		txMassCalculator:                        txMassCalculator,
		
		// Inisialisasi vprog
		vprogEngine:           vprogEngine,
		maxVProgCodeSize:      maxVProgCodeSize,
		maxVProgDataSize:      maxVProgDataSize,
		maxVProgGasLimit:      maxVProgGasLimit,
		maxVProgTotalGas:      maxVProgTotalGas,
		supportedVProgVersion: supportedVProgVersion,
	}
}
