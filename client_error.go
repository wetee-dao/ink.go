package client

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func InitErrors(m *types.Metadata) (registry.ErrorRegistry, error) {
	factory := registry.NewFactory()
	return factory.CreateErrorRegistry(m)
}

func (c *ChainClient) GetErrorInfo(index byte, err [4]byte) (*registry.TypeDecoder, error) {
	var errIndex = [4]types.U8{}
	for i := range 4 {
		errIndex[i] = types.NewU8(err[i])
	}

	id := registry.ErrorID{
		ModuleIndex: types.NewU8(index),
		ErrorIndex:  errIndex,
	}

	info, ok := c.ErrorMap[id]
	if !ok {
		return nil, registry.ErrErrorsTypeNotFound
	}
	return info, nil
}
