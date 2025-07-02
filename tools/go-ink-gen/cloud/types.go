package cloud

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wetee-dao/ink.go/util"
)

type Service struct { // Enum
	Tcp        *uint16 // 0
	Udp        *uint16 // 1
	Http       *uint16 // 2
	Https      *uint16 // 3
	ProjectTcp *uint16 // 4
	ProjectUdp *uint16 // 5
}

func (ty Service) Encode(encoder scale.Encoder) (err error) {
	if ty.Tcp != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Tcp)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Udp != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Udp)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Http != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Http)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Https != nil {
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Https)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.ProjectTcp != nil {
		err = encoder.PushByte(4)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.ProjectTcp)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.ProjectUdp != nil {
		err = encoder.PushByte(5)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.ProjectUdp)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Service) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0: // Inline
		ty.Tcp = new(uint16)
		err = decoder.Decode(ty.Tcp)
		if err != nil {
			return err
		}
		return
	case 1: // Inline
		ty.Udp = new(uint16)
		err = decoder.Decode(ty.Udp)
		if err != nil {
			return err
		}
		return
	case 2: // Inline
		ty.Http = new(uint16)
		err = decoder.Decode(ty.Http)
		if err != nil {
			return err
		}
		return
	case 3: // Inline
		ty.Https = new(uint16)
		err = decoder.Decode(ty.Https)
		if err != nil {
			return err
		}
		return
	case 4: // Inline
		ty.ProjectTcp = new(uint16)
		err = decoder.Decode(ty.ProjectTcp)
		if err != nil {
			return err
		}
		return
	case 5: // Inline
		ty.ProjectUdp = new(uint16)
		err = decoder.Decode(ty.ProjectUdp)
		if err != nil {
			return err
		}
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Disk struct { // Composite
	Path DiskClass
	Size uint32
}
type DiskClass struct { // Enum
	SSD *[]byte // 0
}

func (ty DiskClass) Encode(encoder scale.Encoder) (err error) {
	if ty.SSD != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.SSD)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *DiskClass) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0: // Inline
		ty.SSD = new([]byte)
		err = decoder.Decode(ty.SSD)
		if err != nil {
			return err
		}
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Container struct { // Composite
	Image   []byte
	Command Command
	Port    []Service
	Cr      Cr
}
type Command struct { // Enum
	SH   *[]byte // 0
	BASH *[]byte // 1
	ZSH  *[]byte // 2
	NONE *bool   // 3
}

func (ty Command) Encode(encoder scale.Encoder) (err error) {
	if ty.SH != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.SH)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.BASH != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.BASH)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.ZSH != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.ZSH)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.NONE != nil {
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Command) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0: // Inline
		ty.SH = new([]byte)
		err = decoder.Decode(ty.SH)
		if err != nil {
			return err
		}
		return
	case 1: // Inline
		ty.BASH = new([]byte)
		err = decoder.Decode(ty.BASH)
		if err != nil {
			return err
		}
		return
	case 2: // Inline
		ty.ZSH = new([]byte)
		err = decoder.Decode(ty.ZSH)
		if err != nil {
			return err
		}
		return
	case 3: // Base
		t := true
		ty.NONE = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Cr struct { // Composite
	Cpu  uint32
	Mem  uint32
	Disk []Disk
	Gpu  uint32
}
type Pod struct { // Composite
	Id            uint64
	Creator       types.H160
	ContractId    types.H160
	StartBlock    uint32
	Name          []byte
	TemplateId    util.Option[types.U128]
	Image         []byte
	Meta          []byte
	Command       Command
	Port          []Service
	Cr            Cr
	SideContainer []Container
	TeeVersion    TEEVersion
}
type TEEVersion struct { // Enum
	SGX *bool // 0
	CVM *bool // 1
}

func (ty TEEVersion) Encode(encoder scale.Encoder) (err error) {
	if ty.SGX != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.CVM != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *TEEVersion) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0: // Base
		t := true
		ty.SGX = &t
		return
	case 1: // Base
		t := true
		ty.CVM = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Error struct { // Enum
	SetCodeFailed          *bool // 0
	MustCallByMainContract *bool // 1
}

func (ty Error) Encode(encoder scale.Encoder) (err error) {
	if ty.SetCodeFailed != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.MustCallByMainContract != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Error) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0: // Base
		t := true
		ty.SetCodeFailed = &t
		return
	case 1: // Base
		t := true
		ty.MustCallByMainContract = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}
func (ty *Error) Error() string {
	if ty.SetCodeFailed != nil {
		return "SetCodeFailed"
	}

	if ty.MustCallByMainContract != nil {
		return "MustCallByMainContract"
	}
	return "Unknown"
}

type Tuple_48 struct { // Tuple
	F0 uint32
	F1 Pod
}
