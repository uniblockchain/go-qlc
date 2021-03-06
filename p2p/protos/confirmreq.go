package protos

import (
	"github.com/gogo/protobuf/proto"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/p2p/protos/pb"
)

type ConfirmReqBlock struct {
	Blk *types.StateBlock
}

// ToProto converts domain ConfirmReqBlock into proto ConfirmReqBlock
func ConfirmReqBlockToProto(confirmReq *ConfirmReqBlock) ([]byte, error) {
	blkData, err := confirmReq.Blk.Serialize()
	if err != nil {
		return nil, err
	}
	blockType := confirmReq.Blk.GetType()
	bpPb := &pb.PublishBlock{
		Blocktype: uint32(blockType),
		Block:     blkData,
	}
	data, err := proto.Marshal(bpPb)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ConfirmReqBlockFromProto parse the data into ConfirmReqBlock message
func ConfirmReqBlockFromProto(data []byte) (*ConfirmReqBlock, error) {
	bp := new(pb.ConfirmReq)
	if err := proto.Unmarshal(data, bp); err != nil {
		return nil, err
	}
	blk := new(types.StateBlock)
	if err := blk.Deserialize(bp.Block); err != nil {
		return nil, err
	}
	confirmReqBlock := &ConfirmReqBlock{
		Blk: blk,
	}
	return confirmReqBlock, nil
}
