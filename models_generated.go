// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package errx

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type KeyValueT struct {
	Key   string
	Value string
}

func (t *KeyValueT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	keyOffset := builder.CreateString(t.Key)
	valueOffset := builder.CreateString(t.Value)
	KeyValueStart(builder)
	KeyValueAddKey(builder, keyOffset)
	KeyValueAddValue(builder, valueOffset)
	return KeyValueEnd(builder)
}

func (rcv *KeyValue) UnPackTo(t *KeyValueT) {
	t.Key = string(rcv.Key())
	t.Value = string(rcv.Value())
}

func (rcv *KeyValue) UnPack() *KeyValueT {
	if rcv == nil {
		return nil
	}
	t := &KeyValueT{}
	rcv.UnPackTo(t)
	return t
}

type KeyValue struct {
	_tab flatbuffers.Table
}

func GetRootAsKeyValue(buf []byte, offset flatbuffers.UOffsetT) *KeyValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KeyValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *KeyValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KeyValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *KeyValue) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *KeyValue) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func KeyValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KeyValueAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func KeyValueAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func KeyValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type ErrorModelT struct {
	Next   *ErrorModelT
	Text   string
	Detail string
	Stack  []string
	Debug  []*KeyValueT
}

func (t *ErrorModelT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nextOffset := t.Next.Pack(builder)
	textOffset := builder.CreateString(t.Text)
	detailOffset := builder.CreateString(t.Detail)
	stackOffset := flatbuffers.UOffsetT(0)
	if t.Stack != nil {
		stackLength := len(t.Stack)
		stackOffsets := make([]flatbuffers.UOffsetT, stackLength)
		for j := 0; j < stackLength; j++ {
			stackOffsets[j] = builder.CreateString(t.Stack[j])
		}
		ErrorModelStartStackVector(builder, stackLength)
		for j := stackLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(stackOffsets[j])
		}
		stackOffset = builder.EndVector(stackLength)
	}
	debugOffset := flatbuffers.UOffsetT(0)
	if t.Debug != nil {
		debugLength := len(t.Debug)
		debugOffsets := make([]flatbuffers.UOffsetT, debugLength)
		for j := 0; j < debugLength; j++ {
			debugOffsets[j] = t.Debug[j].Pack(builder)
		}
		ErrorModelStartDebugVector(builder, debugLength)
		for j := debugLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(debugOffsets[j])
		}
		debugOffset = builder.EndVector(debugLength)
	}
	ErrorModelStart(builder)
	ErrorModelAddNext(builder, nextOffset)
	ErrorModelAddText(builder, textOffset)
	ErrorModelAddDetail(builder, detailOffset)
	ErrorModelAddStack(builder, stackOffset)
	ErrorModelAddDebug(builder, debugOffset)
	return ErrorModelEnd(builder)
}

func (rcv *ErrorModel) UnPackTo(t *ErrorModelT) {
	t.Next = rcv.Next(nil).UnPack()
	t.Text = string(rcv.Text())
	t.Detail = string(rcv.Detail())
	stackLength := rcv.StackLength()
	t.Stack = make([]string, stackLength)
	for j := 0; j < stackLength; j++ {
		t.Stack[j] = string(rcv.Stack(j))
	}
	debugLength := rcv.DebugLength()
	t.Debug = make([]*KeyValueT, debugLength)
	for j := 0; j < debugLength; j++ {
		x := KeyValue{}
		rcv.Debug(&x, j)
		t.Debug[j] = x.UnPack()
	}
}

func (rcv *ErrorModel) UnPack() *ErrorModelT {
	if rcv == nil {
		return nil
	}
	t := &ErrorModelT{}
	rcv.UnPackTo(t)
	return t
}

type ErrorModel struct {
	_tab flatbuffers.Table
}

func GetRootAsErrorModel(buf []byte, offset flatbuffers.UOffsetT) *ErrorModel {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ErrorModel{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ErrorModel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ErrorModel) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ErrorModel) Next(obj *ErrorModel) *ErrorModel {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ErrorModel)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ErrorModel) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ErrorModel) Detail() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ErrorModel) Stack(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ErrorModel) StackLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ErrorModel) Debug(obj *KeyValue, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ErrorModel) DebugLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ErrorModelStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ErrorModelAddNext(builder *flatbuffers.Builder, next flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(next), 0)
}
func ErrorModelAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(text), 0)
}
func ErrorModelAddDetail(builder *flatbuffers.Builder, detail flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(detail), 0)
}
func ErrorModelAddStack(builder *flatbuffers.Builder, stack flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(stack), 0)
}
func ErrorModelStartStackVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ErrorModelAddDebug(builder *flatbuffers.Builder, debug flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(debug), 0)
}
func ErrorModelStartDebugVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ErrorModelEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
