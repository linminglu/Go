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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StateID 状态 ID
type StateID int32

const (
	StateID_state_init               StateID = 0
	StateID_state_xipai              StateID = 1
	StateID_state_fapai              StateID = 2
	StateID_state_huansanzhang       StateID = 3
	StateID_state_dingque            StateID = 4
	StateID_state_chupai             StateID = 5
	StateID_state_angang             StateID = 6
	StateID_state_zimo               StateID = 7
	StateID_state_peng               StateID = 8
	StateID_state_gang               StateID = 9
	StateID_state_hu                 StateID = 10
	StateID_state_mopai              StateID = 11
	StateID_state_zixun              StateID = 12
	StateID_state_bugang             StateID = 13
	StateID_state_waitqiangganghu    StateID = 14
	StateID_state_qiangganghu        StateID = 15
	StateID_state_chupaiwenxun       StateID = 16
	StateID_state_gameover           StateID = 17
	StateID_state_gang_settle        StateID = 18
	StateID_state_zimo_settle        StateID = 19
	StateID_state_hu_settle          StateID = 20
	StateID_state_qiangganghu_settle StateID = 21
	StateID_state_gamestart_buhua    StateID = 22
	StateID_state_xingpai_buhua      StateID = 23
	StateID_state_chi                StateID = 24
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
	18: "state_gang_settle",
	19: "state_zimo_settle",
	20: "state_hu_settle",
	21: "state_qiangganghu_settle",
	22: "state_gamestart_buhua",
	23: "state_xingpai_buhua",
	24: "state_chi",
}
var StateID_value = map[string]int32{
	"state_init":               0,
	"state_xipai":              1,
	"state_fapai":              2,
	"state_huansanzhang":       3,
	"state_dingque":            4,
	"state_chupai":             5,
	"state_angang":             6,
	"state_zimo":               7,
	"state_peng":               8,
	"state_gang":               9,
	"state_hu":                 10,
	"state_mopai":              11,
	"state_zixun":              12,
	"state_bugang":             13,
	"state_waitqiangganghu":    14,
	"state_qiangganghu":        15,
	"state_chupaiwenxun":       16,
	"state_gameover":           17,
	"state_gang_settle":        18,
	"state_zimo_settle":        19,
	"state_hu_settle":          20,
	"state_qiangganghu_settle": 21,
	"state_gamestart_buhua":    22,
	"state_xingpai_buhua":      23,
	"state_chi":                24,
}

func (x StateID) String() string {
	return proto.EnumName(StateID_name, int32(x))
}
func (StateID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_states_9f9db09ab52fdd18, []int{0}
}

func init() {
	proto.RegisterEnum("majong.StateID", StateID_name, StateID_value)
}

func init() { proto.RegisterFile("states.proto", fileDescriptor_states_9f9db09ab52fdd18) }

var fileDescriptor_states_9f9db09ab52fdd18 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xff, 0x9f, 0x42, 0x2f, 0xd3, 0x5c, 0x26, 0x13, 0xd2, 0x16, 0x89, 0x27, 0x60, 0xc1,
	0x86, 0x57, 0x60, 0xc3, 0x9a, 0x07, 0x88, 0x5c, 0x30, 0x89, 0x11, 0xb1, 0xd3, 0xc4, 0xa1, 0x55,
	0x1e, 0x1e, 0x61, 0xc7, 0x35, 0x46, 0x62, 0x97, 0xf9, 0xce, 0x5c, 0xce, 0x91, 0x03, 0x51, 0xaf,
	0x99, 0xe6, 0xfd, 0x7d, 0xdb, 0x29, 0xad, 0x68, 0xde, 0xb0, 0x77, 0x25, 0xab, 0xbb, 0xaf, 0x19,
	0x2c, 0x9e, 0xad, 0xf0, 0xf4, 0x48, 0x09, 0xc0, 0xd4, 0x53, 0x0a, 0x29, 0x34, 0xfe, 0xa3, 0x14,
	0xd6, 0xae, 0x3e, 0x89, 0x96, 0x09, 0xfc, 0x1f, 0xc0, 0x1b, 0xb3, 0xe0, 0x82, 0x36, 0x40, 0x0e,
	0xd4, 0x03, 0x93, 0x3d, 0x93, 0x63, 0xcd, 0x64, 0x85, 0x33, 0xca, 0x20, 0x76, 0xfc, 0x55, 0xc8,
	0xea, 0x30, 0x70, 0xbc, 0x24, 0x3c, 0x1b, 0x28, 0x5f, 0xea, 0xc1, 0x0e, 0x5f, 0x05, 0x62, 0x66,
	0xec, 0xd8, 0x3c, 0x18, 0x18, 0x45, 0xa3, 0x70, 0x11, 0xea, 0x96, 0x1b, 0x7d, 0x19, 0xea, 0xa9,
	0x7f, 0x45, 0x11, 0x2c, 0xfd, 0x79, 0x84, 0xe0, 0xae, 0x51, 0xf6, 0xc0, 0x3a, 0x80, 0x51, 0x9c,
	0x06, 0x89, 0x51, 0xb8, 0xb8, 0x1f, 0xa6, 0x0d, 0x31, 0xdd, 0x40, 0xe1, 0xc8, 0x91, 0x09, 0x7d,
	0x10, 0x06, 0x5a, 0xc1, 0xac, 0x4b, 0xa8, 0x80, 0xcc, 0x49, 0xbf, 0x71, 0x1a, 0x22, 0xbb, 0x1c,
	0x47, 0x2e, 0xed, 0x6e, 0x24, 0x82, 0xc4, 0x7b, 0x6b, 0xb8, 0xfa, 0xe4, 0x1d, 0x66, 0x61, 0x85,
	0x9d, 0x2e, 0x7b, 0xae, 0xf5, 0x07, 0x37, 0x9d, 0x3f, 0xd8, 0xc6, 0xf4, 0x38, 0xa7, 0x1c, 0x52,
	0x9f, 0xc6, 0xc3, 0x6b, 0xba, 0x85, 0xdd, 0x1f, 0x17, 0x5e, 0x2d, 0x82, 0x7d, 0x7b, 0xd4, 0x7c,
	0x75, 0xda, 0x44, 0x33, 0x6f, 0x81, 0x1b, 0xda, 0x42, 0xee, 0x1f, 0x4f, 0x56, 0xc6, 0xe8, 0x59,
	0xd8, 0x52, 0x0c, 0x2b, 0x1f, 0x40, 0xe0, 0x6e, 0x3f, 0x9f, 0xfe, 0x87, 0x87, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x8b, 0xde, 0x66, 0x1f, 0x02, 0x00, 0x00,
}
