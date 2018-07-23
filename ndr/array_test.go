package ndr

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDimensions(t *testing.T) {
	a := [2][2][2][]SimpleTest{}
	l, ta := parseDimensions(reflect.ValueOf(a))
	assert.Equal(t, 4, len(l), "dimension count not as expected")
	assert.Equal(t, []int{2, 2, 2, 0}, l, "lengths list not as expected")
	assert.Equal(t, "SimpleTest", ta.Name(), "type within array not as expected")
}

func TestMakeSubSlices(t *testing.T) {
	l := []int{2, 5, 3, 1}
	a := new([][][][]uint32)
	v := reflect.ValueOf(a)
	v = v.Elem()
	ty := v.Type()
	s := reflect.MakeSlice(ty, l[0], l[0])
	v.Set(s)
	makeSubSlices(v, l[1:])
	assert.Equal(t, "[[[[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]]] [[[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]] [[0] [0] [0]]]]", fmt.Sprintf("%v", *a))
}

func TestDimensionCountFromTag(t *testing.T) {
	var a StructWithMultiDimensionalConformantSlice
	v := reflect.ValueOf(a)
	d, err := intFromTag(v.Type().Field(0).Tag, "test")
	if err != nil {
		t.Errorf("error getting dimensions from tag: %v", err)
	}
	assert.Equal(t, 3, d, "number of dimensions not as expected")
}

type StructWithArray struct {
	A [4]uint32
}

type StructWithMultiDimArray struct {
	A [2][3][2]uint32
}

type StructWithConformantSlice struct {
	A []uint32 `ndr:"conformant"`
}

type StructWithVaryingSlice struct {
	A []uint32 `ndr:"varying"`
}

type StructWithConformantVaryingSlice struct {
	A []uint32 `ndr:"conformant,varying"`
}

type StructWithMultiDimensionalConformantSlice struct {
	A [][][]uint32 `ndr:"conformant,test:3"`
}

func TestReadUniDimensionalFixedArrary(t *testing.T) {
	hexStr := "01100800cccccccca0040000000000000000020001000000020000000300000004000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithArray)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for i := range a.A {
		assert.Equal(t, uint32(i+1), a.A[i], "Value of index %d not as expected", i)
	}
}

func TestReadMultiDimensionalFixedArrary(t *testing.T) {
	hexStr := "01100800cccccccca004000000000000000002000100000002000000030000000400000005000000060000000700000008000000090000000a0000000b0000000c0000000d0000000e0000000f000000100000001100000012000000130000001400000015000000160000001700000018000000190000001a0000001b0000001c0000001d0000001e0000001f0000002000000021000000220000002300000024000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithMultiDimArray)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	ar := [2][3][2]uint32{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{7, 8},
			{9, 10},
			{11, 12},
		},
	}
	assert.Equal(t, ar, a.A, "multi-dimensional fixed array not as expected")
}

func TestReadUniDimensionalConformantArrary(t *testing.T) {
	hexStr := "01100800cccccccca004000000000000000002000400000001000000020000000300000004000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithConformantSlice)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for i := range a.A {
		assert.Equal(t, uint32(i+1), a.A[i], "Value of index %d not as expected", i)
	}
}

func TestReadMultiDimensionalConformantArrary(t *testing.T) {
	hexStr := "01100800cccccccca004000000000000000002000200000003000000020000000100000002000000030000000400000005000000060000000700000008000000090000000a0000000b0000000c0000000d0000000e0000000f000000100000001100000012000000130000001400000015000000160000001700000018000000190000001a0000001b0000001c0000001d0000001e0000001f0000002000000021000000220000002300000024000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithMultiDimensionalConformantSlice)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	ar := [][][]uint32{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{7, 8},
			{9, 10},
			{11, 12},
		},
	}
	assert.Equal(t, ar, a.A, "multi-dimensional conformant array not as expected")
}

func TestReadUniDimensionalVaryingArrary(t *testing.T) {
	hexStr := "01100800cccccccca004000000000000000002000200000004000000000000000000000001000000020000000300000004000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithVaryingSlice)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for i := range a.A {
		assert.Equal(t, uint32(i+1), a.A[i], "Value of index %d not as expected", i)
	}
}

func TestReadUniDimensionalConformantVaryingArrary(t *testing.T) {
	hexStr := "01100800cccccccca00400000000000000000200060000000200000004000000000000000000000001000000020000000300000004000000"
	b, _ := hex.DecodeString(hexStr)
	a := new(StructWithConformantVaryingSlice)
	dec := NewDecoder(bytes.NewReader(b), 4)
	err := dec.Decode(a)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for i := range a.A {
		assert.Equal(t, uint32(i+1), a.A[i], "Value of index %d not as expected", i)
	}
}
