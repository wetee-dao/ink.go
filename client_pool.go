package ink

import (
	"os"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/wetee-dao/ink.go/util"
)

func (p *ChainClient) Api() *gsrpc.SubstrateAPI {
	p.mu.Lock()
	defer p.mu.Unlock()

	tries := 0
	poolSize := len(p.conns)

	for tries < poolSize {
		conn := p.conns[p.currIndex]
		_, err := conn.RPC.Chain.GetHeaderLatest()
		if err == nil {
			if p.Debug {
				util.LogWithBlue("INK GET API", p.currIndex, "POOL LEN", poolSize)
			}
			p.currIndex = (p.currIndex + 1) % poolSize
			tries++
			return conn
		}

		p.CloseConn(tries)
	}

	util.LogWithRed("All blockchain connections lost")
	os.Exit(0)
	return nil
}

func (p *ChainClient) CloseConn(index int) {
	p.conns[index].Client.Close()
	p.conns = append(p.conns[:index], p.conns[index+1:]...)
}
