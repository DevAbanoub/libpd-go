// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 22 Sep 2016 23:18:48 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package core

/*
#cgo LDFLAGS: -lpd
#include "z_libpd.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

func (x PrintHook) PassRef() (ref *C.t_libpd_printhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if printHookE41F15CDFunc == nil {
		printHookE41F15CDFunc = x
	}
	return (*C.t_libpd_printhook)(C.t_libpd_printhook_e41f15cd), nil
}

func (x PrintHook) PassValue() (ref C.t_libpd_printhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if printHookE41F15CDFunc == nil {
		printHookE41F15CDFunc = x
	}
	return (C.t_libpd_printhook)(C.t_libpd_printhook_e41f15cd), nil
}

//export printHookE41F15CD
func printHookE41F15CD(crecv *C.char) {
	if printHookE41F15CDFunc != nil {
		recve41f15cd := packPCharString(crecv)
		printHookE41F15CDFunc(recve41f15cd)
		return
	}
	panic("callback func has not been set (race?)")
}

var printHookE41F15CDFunc PrintHook

func (x BangHook) PassRef() (ref *C.t_libpd_banghook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if bangHook7EE06EB6Func == nil {
		bangHook7EE06EB6Func = x
	}
	return (*C.t_libpd_banghook)(C.t_libpd_banghook_7ee06eb6), nil
}

func (x BangHook) PassValue() (ref C.t_libpd_banghook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if bangHook7EE06EB6Func == nil {
		bangHook7EE06EB6Func = x
	}
	return (C.t_libpd_banghook)(C.t_libpd_banghook_7ee06eb6), nil
}

//export bangHook7EE06EB6
func bangHook7EE06EB6(crecv *C.char) {
	if bangHook7EE06EB6Func != nil {
		recv7ee06eb6 := packPCharString(crecv)
		bangHook7EE06EB6Func(recv7ee06eb6)
		return
	}
	panic("callback func has not been set (race?)")
}

var bangHook7EE06EB6Func BangHook

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

func (x FloatHook) PassRef() (ref *C.t_libpd_floathook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if floatHookC0728AC0Func == nil {
		floatHookC0728AC0Func = x
	}
	return (*C.t_libpd_floathook)(C.t_libpd_floathook_c0728ac0), nil
}

func (x FloatHook) PassValue() (ref C.t_libpd_floathook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if floatHookC0728AC0Func == nil {
		floatHookC0728AC0Func = x
	}
	return (C.t_libpd_floathook)(C.t_libpd_floathook_c0728ac0), nil
}

//export floatHookC0728AC0
func floatHookC0728AC0(crecv *C.char, cx C.float) {
	if floatHookC0728AC0Func != nil {
		recvc0728ac0 := packPCharString(crecv)
		xc0728ac0 := (float32)(cx)
		floatHookC0728AC0Func(recvc0728ac0, xc0728ac0)
		return
	}
	panic("callback func has not been set (race?)")
}

var floatHookC0728AC0Func FloatHook

func (x SymbolHook) PassRef() (ref *C.t_libpd_symbolhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if symbolHookA2F11E30Func == nil {
		symbolHookA2F11E30Func = x
	}
	return (*C.t_libpd_symbolhook)(C.t_libpd_symbolhook_a2f11e30), nil
}

func (x SymbolHook) PassValue() (ref C.t_libpd_symbolhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if symbolHookA2F11E30Func == nil {
		symbolHookA2F11E30Func = x
	}
	return (C.t_libpd_symbolhook)(C.t_libpd_symbolhook_a2f11e30), nil
}

//export symbolHookA2F11E30
func symbolHookA2F11E30(crecv *C.char, csym *C.char) {
	if symbolHookA2F11E30Func != nil {
		recva2f11e30 := packPCharString(crecv)
		syma2f11e30 := packPCharString(csym)
		symbolHookA2F11E30Func(recva2f11e30, syma2f11e30)
		return
	}
	panic("callback func has not been set (race?)")
}

var symbolHookA2F11E30Func SymbolHook

func (x ListHook) PassRef() (ref *C.t_libpd_listhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listHookC08DCEF4Func == nil {
		listHookC08DCEF4Func = x
	}
	return (*C.t_libpd_listhook)(C.t_libpd_listhook_c08dcef4), nil
}

func (x ListHook) PassValue() (ref C.t_libpd_listhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listHookC08DCEF4Func == nil {
		listHookC08DCEF4Func = x
	}
	return (C.t_libpd_listhook)(C.t_libpd_listhook_c08dcef4), nil
}

//export listHookC08DCEF4
func listHookC08DCEF4(crecv *C.char, cargc C.int, cargv *C.t_atom) {
	if listHookC08DCEF4Func != nil {
		recvc08dcef4 := packPCharString(crecv)
		argcc08dcef4 := (int32)(cargc)
		argvc08dcef4 := NewAtomRef(unsafe.Pointer(cargv))
		listHookC08DCEF4Func(recvc08dcef4, argcc08dcef4, argvc08dcef4)
		return
	}
	panic("callback func has not been set (race?)")
}

var listHookC08DCEF4Func ListHook

func (x MessageHook) PassRef() (ref *C.t_libpd_messagehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if messageHookCB0280Func == nil {
		messageHookCB0280Func = x
	}
	return (*C.t_libpd_messagehook)(C.t_libpd_messagehook_cb0280), nil
}

func (x MessageHook) PassValue() (ref C.t_libpd_messagehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if messageHookCB0280Func == nil {
		messageHookCB0280Func = x
	}
	return (C.t_libpd_messagehook)(C.t_libpd_messagehook_cb0280), nil
}

//export messageHookCB0280
func messageHookCB0280(crecv *C.char, cmsg *C.char, cargc C.int, cargv *C.t_atom) {
	if messageHookCB0280Func != nil {
		recvcb0280 := packPCharString(crecv)
		msgcb0280 := packPCharString(cmsg)
		argccb0280 := (int32)(cargc)
		argvcb0280 := NewAtomRef(unsafe.Pointer(cargv))
		messageHookCB0280Func(recvcb0280, msgcb0280, argccb0280, argvcb0280)
		return
	}
	panic("callback func has not been set (race?)")
}

var messageHookCB0280Func MessageHook

func (x NoteOnHook) PassRef() (ref *C.t_libpd_noteonhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if noteOnHook73E8687DFunc == nil {
		noteOnHook73E8687DFunc = x
	}
	return (*C.t_libpd_noteonhook)(C.t_libpd_noteonhook_73e8687d), nil
}

func (x NoteOnHook) PassValue() (ref C.t_libpd_noteonhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if noteOnHook73E8687DFunc == nil {
		noteOnHook73E8687DFunc = x
	}
	return (C.t_libpd_noteonhook)(C.t_libpd_noteonhook_73e8687d), nil
}

//export noteOnHook73E8687D
func noteOnHook73E8687D(cchannel C.int, cpitch C.int, cvelocity C.int) {
	if noteOnHook73E8687DFunc != nil {
		channel73e8687d := (int32)(cchannel)
		pitch73e8687d := (int32)(cpitch)
		velocity73e8687d := (int32)(cvelocity)
		noteOnHook73E8687DFunc(channel73e8687d, pitch73e8687d, velocity73e8687d)
		return
	}
	panic("callback func has not been set (race?)")
}

var noteOnHook73E8687DFunc NoteOnHook

func (x ControlChangeHook) PassRef() (ref *C.t_libpd_controlchangehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if controlChangeHook7248F877Func == nil {
		controlChangeHook7248F877Func = x
	}
	return (*C.t_libpd_controlchangehook)(C.t_libpd_controlchangehook_7248f877), nil
}

func (x ControlChangeHook) PassValue() (ref C.t_libpd_controlchangehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if controlChangeHook7248F877Func == nil {
		controlChangeHook7248F877Func = x
	}
	return (C.t_libpd_controlchangehook)(C.t_libpd_controlchangehook_7248f877), nil
}

//export controlChangeHook7248F877
func controlChangeHook7248F877(cchannel C.int, ccontroller C.int, cvalue C.int) {
	if controlChangeHook7248F877Func != nil {
		channel7248f877 := (int32)(cchannel)
		controller7248f877 := (int32)(ccontroller)
		value7248f877 := (int32)(cvalue)
		controlChangeHook7248F877Func(channel7248f877, controller7248f877, value7248f877)
		return
	}
	panic("callback func has not been set (race?)")
}

var controlChangeHook7248F877Func ControlChangeHook

func (x ProgramChangeHook) PassRef() (ref *C.t_libpd_programchangehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if programChangeHook527988C8Func == nil {
		programChangeHook527988C8Func = x
	}
	return (*C.t_libpd_programchangehook)(C.t_libpd_programchangehook_527988c8), nil
}

func (x ProgramChangeHook) PassValue() (ref C.t_libpd_programchangehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if programChangeHook527988C8Func == nil {
		programChangeHook527988C8Func = x
	}
	return (C.t_libpd_programchangehook)(C.t_libpd_programchangehook_527988c8), nil
}

//export programChangeHook527988C8
func programChangeHook527988C8(cchannel C.int, cvalue C.int) {
	if programChangeHook527988C8Func != nil {
		channel527988c8 := (int32)(cchannel)
		value527988c8 := (int32)(cvalue)
		programChangeHook527988C8Func(channel527988c8, value527988c8)
		return
	}
	panic("callback func has not been set (race?)")
}

var programChangeHook527988C8Func ProgramChangeHook

func (x PitchbendHook) PassRef() (ref *C.t_libpd_pitchbendhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if pitchbendHook8BF1C1F3Func == nil {
		pitchbendHook8BF1C1F3Func = x
	}
	return (*C.t_libpd_pitchbendhook)(C.t_libpd_pitchbendhook_8bf1c1f3), nil
}

func (x PitchbendHook) PassValue() (ref C.t_libpd_pitchbendhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if pitchbendHook8BF1C1F3Func == nil {
		pitchbendHook8BF1C1F3Func = x
	}
	return (C.t_libpd_pitchbendhook)(C.t_libpd_pitchbendhook_8bf1c1f3), nil
}

//export pitchbendHook8BF1C1F3
func pitchbendHook8BF1C1F3(cchannel C.int, cvalue C.int) {
	if pitchbendHook8BF1C1F3Func != nil {
		channel8bf1c1f3 := (int32)(cchannel)
		value8bf1c1f3 := (int32)(cvalue)
		pitchbendHook8BF1C1F3Func(channel8bf1c1f3, value8bf1c1f3)
		return
	}
	panic("callback func has not been set (race?)")
}

var pitchbendHook8BF1C1F3Func PitchbendHook

func (x AftertouchHook) PassRef() (ref *C.t_libpd_aftertouchhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if aftertouchHookE7D4F475Func == nil {
		aftertouchHookE7D4F475Func = x
	}
	return (*C.t_libpd_aftertouchhook)(C.t_libpd_aftertouchhook_e7d4f475), nil
}

func (x AftertouchHook) PassValue() (ref C.t_libpd_aftertouchhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if aftertouchHookE7D4F475Func == nil {
		aftertouchHookE7D4F475Func = x
	}
	return (C.t_libpd_aftertouchhook)(C.t_libpd_aftertouchhook_e7d4f475), nil
}

//export aftertouchHookE7D4F475
func aftertouchHookE7D4F475(cchannel C.int, cvalue C.int) {
	if aftertouchHookE7D4F475Func != nil {
		channele7d4f475 := (int32)(cchannel)
		valuee7d4f475 := (int32)(cvalue)
		aftertouchHookE7D4F475Func(channele7d4f475, valuee7d4f475)
		return
	}
	panic("callback func has not been set (race?)")
}

var aftertouchHookE7D4F475Func AftertouchHook

func (x PolyAftertouchHook) PassRef() (ref *C.t_libpd_polyaftertouchhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if polyAftertouchHookC3CFEAF7Func == nil {
		polyAftertouchHookC3CFEAF7Func = x
	}
	return (*C.t_libpd_polyaftertouchhook)(C.t_libpd_polyaftertouchhook_c3cfeaf7), nil
}

func (x PolyAftertouchHook) PassValue() (ref C.t_libpd_polyaftertouchhook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if polyAftertouchHookC3CFEAF7Func == nil {
		polyAftertouchHookC3CFEAF7Func = x
	}
	return (C.t_libpd_polyaftertouchhook)(C.t_libpd_polyaftertouchhook_c3cfeaf7), nil
}

//export polyAftertouchHookC3CFEAF7
func polyAftertouchHookC3CFEAF7(cchannel C.int, cpitch C.int, cvalue C.int) {
	if polyAftertouchHookC3CFEAF7Func != nil {
		channelc3cfeaf7 := (int32)(cchannel)
		pitchc3cfeaf7 := (int32)(cpitch)
		valuec3cfeaf7 := (int32)(cvalue)
		polyAftertouchHookC3CFEAF7Func(channelc3cfeaf7, pitchc3cfeaf7, valuec3cfeaf7)
		return
	}
	panic("callback func has not been set (race?)")
}

var polyAftertouchHookC3CFEAF7Func PolyAftertouchHook

func (x MIDIByteHook) PassRef() (ref *C.t_libpd_midibytehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if mIDIByteHook6393D784Func == nil {
		mIDIByteHook6393D784Func = x
	}
	return (*C.t_libpd_midibytehook)(C.t_libpd_midibytehook_6393d784), nil
}

func (x MIDIByteHook) PassValue() (ref C.t_libpd_midibytehook, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if mIDIByteHook6393D784Func == nil {
		mIDIByteHook6393D784Func = x
	}
	return (C.t_libpd_midibytehook)(C.t_libpd_midibytehook_6393d784), nil
}

//export mIDIByteHook6393D784
func mIDIByteHook6393D784(cport C.int, cbyte C.int) {
	if mIDIByteHook6393D784Func != nil {
		port6393d784 := (int32)(cport)
		byte6393d784 := (int32)(cbyte)
		mIDIByteHook6393D784Func(port6393d784, byte6393d784)
		return
	}
	panic("callback func has not been set (race?)")
}

var mIDIByteHook6393D784Func MIDIByteHook

// Ref returns a reference to C object as it is.
func (x *PdInstance) Ref() *C.struct__pdinstance {
	if x == nil {
		return nil
	}
	return (*C.struct__pdinstance)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *PdInstance) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewPdInstanceRef converts the C object reference into a raw struct reference without wrapping.
func NewPdInstanceRef(ref unsafe.Pointer) *PdInstance {
	return (*PdInstance)(ref)
}

// NewPdInstance allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewPdInstance() *PdInstance {
	return (*PdInstance)(allocStruct_PdInstanceMemory(1))
}

// allocStruct_PdInstanceMemory allocates memory for type C.struct__pdinstance in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_PdInstanceMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_PdInstanceValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStruct_PdInstanceValue = unsafe.Sizeof([1]C.struct__pdinstance{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *PdInstance) PassRef() *C.struct__pdinstance {
	if x == nil {
		x = (*PdInstance)(allocStruct_PdInstanceMemory(1))
	}
	return (*C.struct__pdinstance)(unsafe.Pointer(x))
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
