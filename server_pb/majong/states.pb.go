// Code generated by protoc-gen-go. DO NOT EDIT.
// source: states.proto

package majong

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// StateID 状态 ID
type StateID int32

const (
	StateID_state_init            StateID = 0
	StateID_state_xipai           StateID = 1
	StateID_state_fapai           StateID = 2
	StateID_state_huansanzhang    StateID = 3
	StateID_state_dingque         StateID = 4
	StateID_state_chupai          StateID = 5
	StateID_state_angang          StateID = 6
	StateID_state_zimo            StateID = 7
	StateID_state_peng            StateID = 8
	StateID_state_gang            StateID = 9
	StateID_state_hu              StateID = 10
	StateID_state_mopai           StateID = 11
	StateID_state_zixun           StateID = 12
	StateID_state_bugang          StateID = 13
	StateID_state_waitqiangganghu StateID = 14
	StateID_state_qiangganghu     StateID = 15
	StateID_state_chupaiwenxun    StateID = 16
	StateID_state_gameover        StateID = 17
)

var StateID_name = map[int32]string{
	0:  "state_init",
	1:  "state_xipai",
	2:  "state_fapai",
	3:  "state_huansanzhang",
	4:  "state_dingque",
	5:  "state_chupai",
	6:  "state_angang",
	7:  "state_zimo",
	8:  "state_peng",
	9:  "state_gang",
	10: "state_hu",
	11: "state_mopai",
	12: "state_zixun",
	13: "state_bugang",
	14: "state_waitqiangganghu",
	15: "state_qiangganghu",
	16: "state_chupaiwenxun",
	17: "state_gameover",
}
var StateID_value = map[string]int32{
	"state_init":            0,
	"state_xipai":           1,
	"state_fapai":           2,
	"state_huansanzhang":    3,
	"state_dingque":         4,
	"state_chupai":          5,
	"state_angang":          6,
	"state_zimo":            7,
	"state_peng":            8,
	"state_gang":            9,
	"state_hu":              10,
	"state_mopai":           11,
	"state_zixun":           12,
	"state_bugang":          13,
	"state_waitqiangganghu": 14,
	"state_qiangganghu":     15,
	"state_chupaiwenxun":    16,
	"state_gameover":        17,
}

func (x StateID) String() string {
	return proto.EnumName(StateID_name, int32(x))
}
func (StateID) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterEnum("majong.StateID", StateID_name, StateID_value)
}

func init() { proto.RegisterFile("states.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0xc3, 0x40,
	0x0c, 0x45, 0xa1, 0x40, 0x5a, 0xdc, 0x34, 0x75, 0x2c, 0x15, 0x89, 0x2b, 0xb0, 0x60, 0xc3, 0x15,
	0xd8, 0xb0, 0xe6, 0x00, 0x95, 0x0b, 0xc3, 0xc4, 0x48, 0xf1, 0xa4, 0x4d, 0x86, 0x56, 0x39, 0x2a,
	0xa7, 0x41, 0x93, 0xa1, 0x72, 0x97, 0xff, 0xfd, 0xb1, 0xff, 0xf7, 0x40, 0xd9, 0x0f, 0x3c, 0xb8,
	0xfe, 0xb9, 0x3b, 0x84, 0x21, 0x50, 0xd1, 0xf2, 0x77, 0x50, 0xff, 0xf4, 0x3b, 0x83, 0xf9, 0x7b,
	0x32, 0xde, 0x5e, 0xa9, 0x02, 0x98, 0xde, 0x6c, 0x45, 0x65, 0xc0, 0x2b, 0x5a, 0xc3, 0x32, 0xeb,
	0x93, 0x74, 0x2c, 0x78, 0x6d, 0xe0, 0x8b, 0x13, 0x98, 0xd1, 0x03, 0x50, 0x06, 0x4d, 0x64, 0xed,
	0x59, 0xc7, 0x86, 0xd5, 0xe3, 0x0d, 0xd5, 0xb0, 0xca, 0xfc, 0x53, 0xd4, 0xef, 0xa3, 0xc3, 0x5b,
	0xc2, 0xff, 0x02, 0xdb, 0x8f, 0x26, 0xa6, 0xe1, 0x3b, 0x23, 0xac, 0x3e, 0x8d, 0x15, 0x56, 0x60,
	0x94, 0x36, 0xe0, 0xdc, 0x74, 0xe7, 0xd4, 0xe3, 0xc2, 0xf4, 0xf4, 0xfe, 0x9e, 0x4a, 0x58, 0x9c,
	0xe3, 0x11, 0xac, 0x5d, 0x1b, 0x52, 0xc0, 0xd2, 0xc0, 0x28, 0xa7, 0xa8, 0x58, 0x5a, 0xe2, 0x2e,
	0x4e, 0x1b, 0x56, 0xf4, 0x08, 0x9b, 0x4c, 0x8e, 0x2c, 0xc3, 0x5e, 0x58, 0x7d, 0x32, 0x9a, 0x88,
	0x15, 0x6d, 0xa0, 0xce, 0xd6, 0x25, 0x5e, 0xdb, 0xc9, 0xf9, 0x8e, 0xa3, 0xd3, 0xb4, 0x1b, 0x89,
	0xa0, 0x3a, 0x77, 0x6b, 0x5d, 0xf8, 0x71, 0x07, 0xac, 0x77, 0xc5, 0xf4, 0xd7, 0x2f, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0xf6, 0x43, 0x0b, 0x7b, 0x01, 0x00, 0x00,
}
