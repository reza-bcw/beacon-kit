// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package engineprimitives

import (
	"encoding/json"
	"errors"

	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/uint256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*executableDataDenebMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutableDataDeneb) MarshalJSON() ([]byte, error) {
	type ExecutableDataDeneb struct {
		ParentHash    common.Hash              `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
		FeeRecipient  common.Address           `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
		StateRoot     common.Hash              `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot  common.Hash              `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
		LogsBloom     hexutil.Bytes            `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
		Random        common.Hash              `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
		Number        hexutil.Uint64           `json:"blockNumber"                  gencodec:"required"`
		GasLimit      hexutil.Uint64           `json:"gasLimit"                     gencodec:"required"`
		GasUsed       hexutil.Uint64           `json:"gasUsed"                      gencodec:"required"`
		Timestamp     hexutil.Uint64           `json:"timestamp"                    gencodec:"required"`
		ExtraData     hexutil.Bytes            `json:"extraData"                    gencodec:"required" ssz-max:"32"`
		BaseFeePerGas uint256.LittleEndian     `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
		BlockHash     common.Hash              `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
		Transactions  []hexutil.Bytes          `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
		Withdrawals   []*primitives.Withdrawal `json:"withdrawals"                                      ssz-max:"16"`
		BlobGasUsed   hexutil.Uint64           `json:"blobGasUsed"`
		ExcessBlobGas hexutil.Uint64           `json:"excessBlobGas"`
	}
	var enc ExecutableDataDeneb
	enc.ParentHash = e.ParentHash
	enc.FeeRecipient = e.FeeRecipient
	enc.StateRoot = e.StateRoot
	enc.ReceiptsRoot = e.ReceiptsRoot
	enc.LogsBloom = e.LogsBloom
	enc.Random = e.Random
	enc.Number = hexutil.Uint64(e.Number)
	enc.GasLimit = hexutil.Uint64(e.GasLimit)
	enc.GasUsed = hexutil.Uint64(e.GasUsed)
	enc.Timestamp = hexutil.Uint64(e.Timestamp)
	enc.ExtraData = e.ExtraData
	enc.BaseFeePerGas = e.BaseFeePerGas
	enc.BlockHash = e.BlockHash
	if e.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(e.Transactions))
		for k, v := range e.Transactions {
			enc.Transactions[k] = v
		}
	}
	enc.Withdrawals = e.Withdrawals
	enc.BlobGasUsed = hexutil.Uint64(e.BlobGasUsed)
	enc.ExcessBlobGas = hexutil.Uint64(e.ExcessBlobGas)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutableDataDeneb) UnmarshalJSON(input []byte) error {
	type ExecutableDataDeneb struct {
		ParentHash    *common.Hash             `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
		FeeRecipient  *common.Address          `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
		StateRoot     *common.Hash             `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot  *common.Hash             `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
		LogsBloom     *hexutil.Bytes           `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
		Random        *common.Hash             `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
		Number        *hexutil.Uint64          `json:"blockNumber"                  gencodec:"required"`
		GasLimit      *hexutil.Uint64          `json:"gasLimit"                     gencodec:"required"`
		GasUsed       *hexutil.Uint64          `json:"gasUsed"                      gencodec:"required"`
		Timestamp     *hexutil.Uint64          `json:"timestamp"                    gencodec:"required"`
		ExtraData     *hexutil.Bytes           `json:"extraData"                    gencodec:"required" ssz-max:"32"`
		BaseFeePerGas *uint256.LittleEndian    `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
		BlockHash     *common.Hash             `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
		Transactions  []hexutil.Bytes          `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
		Withdrawals   []*primitives.Withdrawal `json:"withdrawals"                                      ssz-max:"16"`
		BlobGasUsed   *hexutil.Uint64          `json:"blobGasUsed"`
		ExcessBlobGas *hexutil.Uint64          `json:"excessBlobGas"`
	}
	var dec ExecutableDataDeneb
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for ExecutableDataDeneb")
	}
	e.ParentHash = *dec.ParentHash
	if dec.FeeRecipient == nil {
		return errors.New("missing required field 'feeRecipient' for ExecutableDataDeneb")
	}
	e.FeeRecipient = *dec.FeeRecipient
	if dec.StateRoot == nil {
		return errors.New("missing required field 'stateRoot' for ExecutableDataDeneb")
	}
	e.StateRoot = *dec.StateRoot
	if dec.ReceiptsRoot == nil {
		return errors.New("missing required field 'receiptsRoot' for ExecutableDataDeneb")
	}
	e.ReceiptsRoot = *dec.ReceiptsRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutableDataDeneb")
	}
	e.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for ExecutableDataDeneb")
	}
	e.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'blockNumber' for ExecutableDataDeneb")
	}
	e.Number = uint64(*dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for ExecutableDataDeneb")
	}
	e.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for ExecutableDataDeneb")
	}
	e.GasUsed = uint64(*dec.GasUsed)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for ExecutableDataDeneb")
	}
	e.Timestamp = uint64(*dec.Timestamp)
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for ExecutableDataDeneb")
	}
	e.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'baseFeePerGas' for ExecutableDataDeneb")
	}
	e.BaseFeePerGas = *dec.BaseFeePerGas
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for ExecutableDataDeneb")
	}
	e.BlockHash = *dec.BlockHash
	if dec.Transactions == nil {
		return errors.New("missing required field 'transactions' for ExecutableDataDeneb")
	}
	e.Transactions = make([][]byte, len(dec.Transactions))
	for k, v := range dec.Transactions {
		e.Transactions[k] = v
	}
	if dec.Withdrawals != nil {
		e.Withdrawals = dec.Withdrawals
	}
	if dec.BlobGasUsed != nil {
		e.BlobGasUsed = uint64(*dec.BlobGasUsed)
	}
	if dec.ExcessBlobGas != nil {
		e.ExcessBlobGas = uint64(*dec.ExcessBlobGas)
	}
	return nil
}