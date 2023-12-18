// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package blockchain

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	prysmexecution "github.com/prysmaticlabs/prysm/v4/beacon-chain/execution"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/interfaces"
	payloadattribute "github.com/prysmaticlabs/prysm/v4/consensus-types/payload-attribute"
	enginev1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"

	"cosmossdk.io/core/header"
	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/types/config"
	v1 "github.com/itsdevbear/bolaris/types/v1"
)

// var defaultLatestValidHash = bytesutil.PadTo([]byte{0xff}, 32)

type ForkChoiceStoreProvider interface {
	ForkChoiceStore(ctx context.Context) v1.ForkChoiceStore
}

type Service struct {
	beaconCfg *config.Beacon
	etherbase common.Address
	logger    log.Logger
	fcsp      ForkChoiceStoreProvider
	engine    execution.EngineCaller
}

func NewService(opts ...Option) *Service {
	s := &Service{}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			s.logger.Error("Failed to apply option", "error", err)
		}
	}

	return s
}

// notifyForkchoiceUpdateArg is the argument for the forkchoice update notification `notifyForkchoiceUpdate`.
type notifyForkchoiceUpdateArg struct {
	headHash []byte
}

func (s *Service) ProcessBlock(ctx context.Context, block header.Info, header interfaces.ExecutionData) (*enginev1.PayloadIDBytes, error) {
	// pid, err := s.notifyForkchoiceUpdate(ctx, uint64(block.Height), &notifyForkchoiceUpdateArg{
	// 	headHash: headHash,
	// })
	// fmt.Println(pid, err)

	_, err := s.validateExecutionOnBlock(ctx, 0, header, nil, [32]byte{})
	if err != nil {
		return nil, err
	}

	// if err != nil {
	return s.notifyForkchoiceUpdate(ctx, uint64(block.Height), &notifyForkchoiceUpdateArg{
		headHash: header.BlockHash(),
	})
}

func (s *Service) notifyForkchoiceUpdate(ctx context.Context, slot uint64, arg *notifyForkchoiceUpdateArg) (*enginev1.PayloadIDBytes, error) {
	// currSafeBlk, err := s.fcsp.ForkChoiceStore(ctx).GetSafeBlockHash()
	// if err != nil {
	// 	return nil, err
	// }

	// currFinalizedBlk, err := s.fcsp.ForkChoiceStore(ctx).GetFinalizedBlockHash()
	// if err != nil {
	// 	return nil, err
	// }

	fcs := &enginev1.ForkchoiceState{
		HeadBlockHash:      arg.headHash,
		SafeBlockHash:      arg.headHash, // currSafeBlk[:],
		FinalizedBlockHash: arg.headHash, // currFinalizedBlk[:],
	}
	// We want to start building the next block as part of this forkchoice update.
	nextSlot := slot + 1 // Cache payload ID for next slot proposer.
	attrs, err := s.getPayloadAttributes(ctx, nextSlot)
	if err != nil {
		return nil, err
	}

	payloadID, lastValidHash, err := s.engine.ForkchoiceUpdated(ctx, fcs, attrs)
	fmt.Println(lastValidHash)
	if err != nil {
		switch err {
		case prysmexecution.ErrAcceptedSyncingPayloadStatus:
			// forkchoiceUpdatedOptimisticNodeCount.Inc()
			// log.WithFields(logrus.Fields{
			// 	"headSlot":                  headBlk.Slot(),
			// 	"headPayloadBlockHash":      fmt.Sprintf("%#x", bytesutil.Trunc(headPayload.BlockHash())),
			// 	"finalizedPayloadBlockHash": fmt.Sprintf("%#x", bytesutil.Trunc(finalizedHash[:])),
			// }).Info("Called fork choice updated with optimistic block")
			return payloadID, nil
		case prysmexecution.ErrInvalidPayloadStatus:
			// // forkchoiceUpdatedInvalidNodeCount.Inc()
			// // headRoot := arg.headRoot
			// if len(lastValidHash) == 0 {
			// 	lastValidHash = defaultLatestValidHash
			// }
			// // // invalidRoots, err := s.cfg.ForkChoiceStore.SetOptimisticToInvalid(ctx, headRoot, headBlk.ParentRoot(), bytesutil.ToBytes32(lastValidHash))
			// if err != nil {
			// 	log.WithError(err).Error("Could not set head root to invalid")
			// 	return nil, nil
			// }
			// if err := s.removeInvalidBlockAndState(ctx, invalidRoots); err != nil {
			// 	log.WithError(err).Error("Could not remove invalid block and state")
			// 	return nil, nil
			// }

			// r, err := s.cfg.ForkChoiceStore.Head(ctx)
			// if err != nil {
			// 	log.WithFields(logrus.Fields{
			// 		"slot":                 headBlk.Slot(),
			// 		"blockRoot":            fmt.Sprintf("%#x", bytesutil.Trunc(headRoot[:])),
			// 		"invalidChildrenCount": len(invalidRoots),
			// 	}).Warn("Pruned invalid blocks, could not update head root")
			// 	return nil, invalidBlock{error: ErrInvalidPayload, root: arg.headRoot, invalidAncestorRoots: invalidRoots}
			// }
			// b, err := s.getBlock(ctx, r)
			// if err != nil {
			// 	log.WithError(err).Error("Could not get head block")
			// 	return nil, nil
			// }
			// st, err := s.cfg.StateGen.StateByRoot(ctx, r)
			// if err != nil {
			// 	log.WithError(err).Error("Could not get head state")
			// 	return nil, nil
			// }
			previousHead := s.fcsp.ForkChoiceStore(ctx).GetLastValidHead()
			pid, err := s.notifyForkchoiceUpdate(ctx, slot, &notifyForkchoiceUpdateArg{
				headHash: previousHead[:],
			})
			if err != nil {
				return nil, err // Returning err because it's recursive here.
			}

			// if err := s.saveHead(ctx, r, b, st); err != nil {
			// 	log.WithError(err).Error("could not save head after pruning invalid blocks")
			// }

			// log.WithFields(logrus.Fields{
			// 	"slot":                 headBlk.Slot(),
			// 	"blockRoot":            fmt.Sprintf("%#x", bytesutil.Trunc(headRoot[:])),
			// 	"invalidChildrenCount": len(invalidRoots),
			// 	"newHeadRoot":          fmt.Sprintf("%#x", bytesutil.Trunc(r[:])),
			// }).Warn("Pruned invalid blocks")
			return pid, errors.New("invalid payload")
			// return pid, invalidBlock{error: ErrInvalidPayload, root: arg.headRoot, invalidAncestorRoots: invalidRoots}

		default:
			// log.WithError(err).Error(ErrUndefinedExecutionEngineError)
			s.logger.Error("undefined execution engine error", "err", err)
			return nil, nil
		}
	}
	// forkchoiceUpdatedValidNodeCount.Inc()
	// if err := s.cfg.ForkChoiceStore.SetOptimisticToValid(ctx, arg.headRoot); err != nil {
	// 	log.WithError(err).Error("Could not set head root to valid")
	// 	return nil, nil
	// }
	// If the forkchoice update call has an attribute, update the proposer payload ID cache.
	// if hasAttr && payloadID != nil {
	// 	var pId [8]byte
	// 	copy(pId[:], payloadID[:])
	// 	log.WithFields(logrus.Fields{
	// 		"blockRoot": fmt.Sprintf("%#x", bytesutil.Trunc(arg.headRoot[:])),
	// 		"headSlot":  headBlk.Slot(),
	// 		"payloadID": fmt.Sprintf("%#x", bytesutil.Trunc(payloadID[:])),
	// 	}).Info("Forkchoice updated with payload attributes for proposal")
	// 	s.cfg.ProposerSlotIndexCache.SetProposerAndPayloadIDs(nextSlot, proposerId, pId, arg.headRoot)
	// } else if hasAttr && payloadID == nil && !features.Get().PrepareAllPayloads {
	// 	log.WithFields(logrus.Fields{
	// 		"blockHash": fmt.Sprintf("%#x", headPayload.BlockHash()),
	// 		"slot":      headBlk.Slot(),
	// 	}).Error("Received nil payload ID on VALID engine response")
	// }
	return payloadID, nil
}

// It returns true if the EL has returned VALID for the block.
func (s *Service) notifyNewPayload(ctx context.Context, preStateVersion int,
	preStateHeader interfaces.ExecutionData, blk interfaces.ReadOnlySignedBeaconBlock) (bool, error) {
	lastValidHash, err := s.engine.NewPayload(ctx, preStateHeader, []common.Hash{}, &common.Hash{} /*empty version hashes and root before Deneb*/)
	return lastValidHash != nil, err
}

// validateExecutionOnBlock notifies the engine of the incoming block execution payload and returns true if the payload is valid.
func (s *Service) validateExecutionOnBlock(ctx context.Context, ver int, header interfaces.ExecutionData, signed interfaces.ReadOnlySignedBeaconBlock, blockRoot [32]byte) (bool, error) {
	isValidPayload, err := s.notifyNewPayload(ctx, ver, header, signed)
	if err != nil {
		return false, err
		// return false, s.handleInvalidExecutionError(ctx, err, blockRoot, signed.Block().ParentRoot())
	}

	fmt.Println("POST NOTIFY NEW PAYLOAD", isValidPayload)
	// if signed.Version() < version.Capella && isValidPayload {
	// 	if err := s.validateMergeTransitionBlock(ctx, ver, header, signed); err != nil {
	// 		return isValidPayload, err
	// 	}
	// }
	return isValidPayload, nil
}

func (s *Service) getPayloadAttributes(ctx context.Context, slot uint64) (payloadattribute.Attributer, error) {
	// TODO: modularize andn make better.
	var random [32]byte
	if _, err := rand.Read(random[:]); err != nil {
		return nil, err
	}

	return payloadattribute.New(&enginev1.PayloadAttributesV2{
		Timestamp:             uint64(time.Now().Unix()),
		SuggestedFeeRecipient: s.etherbase.Bytes(),
		Withdrawals:           nil,
		PrevRandao:            append([]byte{}, random[:]...),
	})
}
