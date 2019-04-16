// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FloatingPoint struct {
	_tab flatbuffers.Table
}

func GetRootAsFloatingPoint(buf []byte, offset flatbuffers.UOffsetT) *FloatingPoint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FloatingPoint{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FloatingPoint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FloatingPoint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FloatingPoint) Precision() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FloatingPoint) MutatePrecision(n int16) bool {
	return rcv._tab.MutateInt16Slot(4, n)
}

func FloatingPointStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FloatingPointAddPrecision(builder *flatbuffers.Builder, precision int16) {
	builder.PrependInt16Slot(0, precision, 0)
}
func FloatingPointEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
