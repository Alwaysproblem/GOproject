// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/cost_graph.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CostGraphDef struct {
	Node []*CostGraphDef_Node `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
}

func (m *CostGraphDef) Reset()                    { *m = CostGraphDef{} }
func (m *CostGraphDef) String() string            { return proto.CompactTextString(m) }
func (*CostGraphDef) ProtoMessage()               {}
func (*CostGraphDef) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *CostGraphDef) GetNode() []*CostGraphDef_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type CostGraphDef_Node struct {
	// The name of the node. Names are globally unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The device of the node. Can be empty if the node is mapped to the
	// default partition or partitioning hasn't been run yet.
	Device string `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
	// The id of the node. Node ids are only unique inside a partition.
	Id         int32                           `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	InputInfo  []*CostGraphDef_Node_InputInfo  `protobuf:"bytes,4,rep,name=input_info,json=inputInfo" json:"input_info,omitempty"`
	OutputInfo []*CostGraphDef_Node_OutputInfo `protobuf:"bytes,5,rep,name=output_info,json=outputInfo" json:"output_info,omitempty"`
	// Temporary memory used by this node.
	TemporaryMemorySize        int64 `protobuf:"varint,6,opt,name=temporary_memory_size,json=temporaryMemorySize" json:"temporary_memory_size,omitempty"`
	HostTempMemorySize         int64 `protobuf:"varint,10,opt,name=host_temp_memory_size,json=hostTempMemorySize" json:"host_temp_memory_size,omitempty"`
	DeviceTempMemorySize       int64 `protobuf:"varint,11,opt,name=device_temp_memory_size,json=deviceTempMemorySize" json:"device_temp_memory_size,omitempty"`
	HostPersistentMemorySize   int64 `protobuf:"varint,12,opt,name=host_persistent_memory_size,json=hostPersistentMemorySize" json:"host_persistent_memory_size,omitempty"`
	DevicePersistentMemorySize int64 `protobuf:"varint,16,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize" json:"device_persistent_memory_size,omitempty"`
	// Estimate of the computational cost of this node, in microseconds.
	ComputeCost int64 `protobuf:"varint,9,opt,name=compute_cost,json=computeCost" json:"compute_cost,omitempty"`
	// Analytical estimate of the computational cost of this node, in
	// microseconds.
	ComputeTime int64 `protobuf:"varint,14,opt,name=compute_time,json=computeTime" json:"compute_time,omitempty"`
	// Analytical estimate of the memory access cost of this node, in
	// microseconds.
	MemoryTime int64 `protobuf:"varint,15,opt,name=memory_time,json=memoryTime" json:"memory_time,omitempty"`
	// If true, the output is permanent: it can't be discarded, because this
	// node is part of the "final output". Nodes may depend on final nodes.
	IsFinal bool `protobuf:"varint,7,opt,name=is_final,json=isFinal" json:"is_final,omitempty"`
	// Ids of the control inputs for this node.
	ControlInput []int32 `protobuf:"varint,8,rep,packed,name=control_input,json=controlInput" json:"control_input,omitempty"`
}

func (m *CostGraphDef_Node) Reset()                    { *m = CostGraphDef_Node{} }
func (m *CostGraphDef_Node) String() string            { return proto.CompactTextString(m) }
func (*CostGraphDef_Node) ProtoMessage()               {}
func (*CostGraphDef_Node) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *CostGraphDef_Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CostGraphDef_Node) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *CostGraphDef_Node) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CostGraphDef_Node) GetInputInfo() []*CostGraphDef_Node_InputInfo {
	if m != nil {
		return m.InputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetOutputInfo() []*CostGraphDef_Node_OutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetTemporaryMemorySize() int64 {
	if m != nil {
		return m.TemporaryMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetHostTempMemorySize() int64 {
	if m != nil {
		return m.HostTempMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetDeviceTempMemorySize() int64 {
	if m != nil {
		return m.DeviceTempMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetHostPersistentMemorySize() int64 {
	if m != nil {
		return m.HostPersistentMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetDevicePersistentMemorySize() int64 {
	if m != nil {
		return m.DevicePersistentMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetComputeCost() int64 {
	if m != nil {
		return m.ComputeCost
	}
	return 0
}

func (m *CostGraphDef_Node) GetComputeTime() int64 {
	if m != nil {
		return m.ComputeTime
	}
	return 0
}

func (m *CostGraphDef_Node) GetMemoryTime() int64 {
	if m != nil {
		return m.MemoryTime
	}
	return 0
}

func (m *CostGraphDef_Node) GetIsFinal() bool {
	if m != nil {
		return m.IsFinal
	}
	return false
}

func (m *CostGraphDef_Node) GetControlInput() []int32 {
	if m != nil {
		return m.ControlInput
	}
	return nil
}

// Inputs of this node. They must be executed before this node can be
// executed. An input is a particular output of another node, specified
// by the node id and the output index.
type CostGraphDef_Node_InputInfo struct {
	PrecedingNode int32 `protobuf:"varint,1,opt,name=preceding_node,json=precedingNode" json:"preceding_node,omitempty"`
	PrecedingPort int32 `protobuf:"varint,2,opt,name=preceding_port,json=precedingPort" json:"preceding_port,omitempty"`
}

func (m *CostGraphDef_Node_InputInfo) Reset()         { *m = CostGraphDef_Node_InputInfo{} }
func (m *CostGraphDef_Node_InputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_InputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_InputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 0}
}

func (m *CostGraphDef_Node_InputInfo) GetPrecedingNode() int32 {
	if m != nil {
		return m.PrecedingNode
	}
	return 0
}

func (m *CostGraphDef_Node_InputInfo) GetPrecedingPort() int32 {
	if m != nil {
		return m.PrecedingPort
	}
	return 0
}

// Outputs of this node.
type CostGraphDef_Node_OutputInfo struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// If >= 0, the output is an alias of an input. Note that an alias input
	// may itself be an alias. The algorithm will therefore need to follow
	// those pointers.
	AliasInputPort int64             `protobuf:"varint,2,opt,name=alias_input_port,json=aliasInputPort" json:"alias_input_port,omitempty"`
	Shape          *TensorShapeProto `protobuf:"bytes,3,opt,name=shape" json:"shape,omitempty"`
	Dtype          DataType          `protobuf:"varint,4,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (m *CostGraphDef_Node_OutputInfo) Reset()         { *m = CostGraphDef_Node_OutputInfo{} }
func (m *CostGraphDef_Node_OutputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_OutputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_OutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 1}
}

func (m *CostGraphDef_Node_OutputInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetAliasInputPort() int64 {
	if m != nil {
		return m.AliasInputPort
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *CostGraphDef_Node_OutputInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func init() {
	proto.RegisterType((*CostGraphDef)(nil), "tensorflow.CostGraphDef")
	proto.RegisterType((*CostGraphDef_Node)(nil), "tensorflow.CostGraphDef.Node")
	proto.RegisterType((*CostGraphDef_Node_InputInfo)(nil), "tensorflow.CostGraphDef.Node.InputInfo")
	proto.RegisterType((*CostGraphDef_Node_OutputInfo)(nil), "tensorflow.CostGraphDef.Node.OutputInfo")
}

func init() { proto.RegisterFile("github.com/tensorflow/tensorflow/tensorflow/go/core/framework/cost_graph.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe5, 0x26, 0x69, 0x9a, 0x71, 0x9a, 0x56, 0xfb, 0x6f, 0xff, 0x2c, 0x86, 0x0a, 0x03,
	0xaa, 0xb0, 0x2a, 0x94, 0xd2, 0x20, 0x8e, 0x1c, 0x28, 0x55, 0x51, 0x0f, 0x40, 0xe4, 0xe6, 0xc2,
	0xc9, 0x32, 0xf6, 0x3a, 0x59, 0x11, 0x7b, 0x57, 0xbb, 0x1b, 0xaa, 0xf4, 0x79, 0x38, 0xf1, 0x64,
	0x3c, 0x02, 0x47, 0xb4, 0x63, 0xe3, 0x38, 0x40, 0x73, 0x5b, 0xcf, 0xfc, 0xbe, 0x6f, 0x9c, 0xcd,
	0x37, 0x86, 0x13, 0xc3, 0x0a, 0x2d, 0x54, 0x36, 0x17, 0x37, 0xa7, 0x89, 0x50, 0xec, 0x34, 0x53,
	0x71, 0xce, 0x6e, 0x84, 0xfa, 0x72, 0x9a, 0x08, 0x6d, 0xa2, 0xa9, 0x8a, 0xe5, 0x6c, 0x28, 0x95,
	0x30, 0x82, 0xc0, 0x8a, 0xf5, 0x9e, 0xdf, 0xad, 0x2b, 0x3b, 0x91, 0x9e, 0xc5, 0x92, 0x95, 0x4a,
	0xef, 0x78, 0x03, 0xbd, 0x94, 0x4c, 0x97, 0xd8, 0x93, 0x1f, 0x5d, 0xe8, 0xbf, 0x15, 0xda, 0xbc,
	0xb3, 0x43, 0x2f, 0x58, 0x46, 0xce, 0xa0, 0x5d, 0x88, 0x94, 0x51, 0xc7, 0x6f, 0x05, 0xee, 0xe8,
	0x68, 0xb8, 0xb2, 0x19, 0x36, 0xb9, 0xe1, 0x07, 0x91, 0xb2, 0x10, 0x51, 0xef, 0x5b, 0x17, 0xda,
	0xf6, 0x91, 0x10, 0x68, 0x17, 0x71, 0x6e, 0xb5, 0x4e, 0xd0, 0x0b, 0xf1, 0x4c, 0xfe, 0x87, 0xed,
	0x94, 0x7d, 0xe5, 0x09, 0xa3, 0x5b, 0x58, 0xad, 0x9e, 0xc8, 0x00, 0xb6, 0x78, 0x4a, 0x5b, 0xbe,
	0x13, 0x74, 0xc2, 0x2d, 0x9e, 0x92, 0x4b, 0x00, 0x5e, 0xc8, 0x85, 0x89, 0x78, 0x91, 0x09, 0xda,
	0xc6, 0xe9, 0xcf, 0x36, 0x4e, 0x1f, 0x5e, 0x59, 0xfe, 0xaa, 0xc8, 0x44, 0xd8, 0xe3, 0xbf, 0x8f,
	0xe4, 0x0a, 0x5c, 0xb1, 0x30, 0xb5, 0x51, 0x07, 0x8d, 0x82, 0xcd, 0x46, 0x1f, 0x51, 0x80, 0x4e,
	0x20, 0xea, 0x33, 0x19, 0xc1, 0xa1, 0x61, 0xb9, 0x14, 0x2a, 0x56, 0xcb, 0x28, 0x67, 0xb9, 0x50,
	0xcb, 0x48, 0xf3, 0x5b, 0x46, 0xb7, 0x7d, 0x27, 0x68, 0x85, 0xff, 0xd5, 0xcd, 0xf7, 0xd8, 0xbb,
	0xe6, 0xb7, 0x8c, 0x9c, 0xc1, 0xe1, 0xcc, 0xfe, 0x89, 0xb6, 0xb7, 0xa6, 0x01, 0xd4, 0x10, 0xdb,
	0x9c, 0xb0, 0x5c, 0x36, 0x24, 0xaf, 0xe0, 0x5e, 0x79, 0x27, 0x7f, 0x8b, 0x5c, 0x14, 0x1d, 0x94,
	0xed, 0x3f, 0x64, 0xaf, 0xe1, 0x01, 0x4e, 0x92, 0x4c, 0x69, 0xae, 0x0d, 0x2b, 0xcc, 0x9a, 0xb4,
	0x8f, 0x52, 0x6a, 0x91, 0x71, 0x4d, 0x34, 0xe4, 0x6f, 0xe0, 0xa8, 0x9a, 0x7a, 0x87, 0xc1, 0x3e,
	0x1a, 0x78, 0x25, 0xf4, 0x4f, 0x8b, 0xc7, 0xd0, 0x4f, 0x44, 0x2e, 0x17, 0x86, 0x45, 0x36, 0xb8,
	0xb4, 0x87, 0x0a, 0xb7, 0xaa, 0xd9, 0x6b, 0x6e, 0x22, 0x86, 0xe7, 0x8c, 0x0e, 0xd6, 0x90, 0x09,
	0xcf, 0x19, 0x79, 0x04, 0x6e, 0x35, 0x16, 0x89, 0x3d, 0x24, 0xa0, 0x2c, 0x21, 0x70, 0x1f, 0x76,
	0xb8, 0x8e, 0x32, 0x5e, 0xc4, 0x73, 0xda, 0xf5, 0x9d, 0x60, 0x27, 0xec, 0x72, 0x7d, 0x69, 0x1f,
	0xc9, 0x53, 0xd8, 0x4d, 0x44, 0x61, 0x94, 0x98, 0x47, 0x98, 0x00, 0xba, 0xe3, 0xb7, 0x82, 0x4e,
	0xd8, 0xaf, 0x8a, 0x18, 0x10, 0xef, 0x13, 0xf4, 0xea, 0xa4, 0x90, 0x63, 0x18, 0x48, 0xc5, 0x12,
	0x96, 0xf2, 0x62, 0x1a, 0x55, 0x41, 0xb7, 0x11, 0xdc, 0xad, 0xab, 0x98, 0xe4, 0x35, 0x4c, 0x0a,
	0x65, 0x30, 0xbd, 0x4d, 0x6c, 0x2c, 0x94, 0xf1, 0xbe, 0x3b, 0x00, 0xab, 0xf0, 0xd8, 0xfc, 0xe3,
	0xd5, 0x39, 0xf8, 0x1b, 0xf0, 0x4c, 0x02, 0xd8, 0x8f, 0xe7, 0x3c, 0xd6, 0xe5, 0x0b, 0xae, 0xbc,
	0x5a, 0xe1, 0x00, 0xeb, 0xf8, 0x6a, 0xd6, 0x8c, 0x8c, 0xa0, 0x83, 0x0b, 0x8c, 0x4b, 0xe1, 0x8e,
	0x1e, 0x36, 0x33, 0x3b, 0xc1, 0xe3, 0xb5, 0x6d, 0x8f, 0xed, 0xde, 0x86, 0x25, 0x4a, 0x4e, 0xa0,
	0x93, 0xda, 0x75, 0xa6, 0x6d, 0xdf, 0x09, 0x06, 0xa3, 0x83, 0xa6, 0xe6, 0x22, 0x36, 0xf1, 0x64,
	0x29, 0x59, 0x58, 0x22, 0xe7, 0x2f, 0x80, 0x0a, 0x35, 0x6d, 0x12, 0xf5, 0x27, 0xe1, 0x7c, 0xaf,
	0x5e, 0x0a, 0xb4, 0xd7, 0x63, 0xe7, 0xa7, 0xe3, 0x7c, 0xde, 0xc6, 0x6f, 0xc4, 0xcb, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0xa8, 0x1a, 0x17, 0xb2, 0x04, 0x00, 0x00,
}
