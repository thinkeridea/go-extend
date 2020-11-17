// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exsync

import (
	"testing"
	"unsafe"
)

type one int

func (o *one) Increment() {
	*o++
}

func run(t *testing.T, once *Once, o *one, c chan bool) {
	v:=once.Do(func() interface{} {
		o.Increment()
		return o
	}).(*one)

	if *v != 1 {
		t.Errorf("once failed inside run: %d is not 1", v)
	}
	c <- true
}

func TestOnce(t *testing.T) {
	o := new(one)
	once := new(Once)
	c := make(chan bool)
	const N = 10
	for i := 0; i < N; i++ {
		go run(t, once, o, c)
	}
	for i := 0; i < N; i++ {
		<-c
	}
	if *o != 1 {
		t.Errorf("once failed outside run: %d is not 1", *o)
	}
}

func TestOncePanic(t *testing.T) {
	var once Once
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Once.Do did not panic")
			}
		}()

		_ = once.Do(func() interface{}{
			panic("failed")
			return nil
		}).(*one)
	}()

	_ = once.Do(func() interface{}{
		t.Fatalf("Once.Do called twice")
		return nil
	})
}

func BenchmarkOnce(b *testing.B) {
	var once Once
	var o = new(one)
	f := func() interface{}{return o}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = once.Do(f).(*one)
		}
	})
}


func runPointer(t *testing.T, once *OncePointer, o *one, c chan bool) {
	v:=once.Do(func() unsafe.Pointer {
		o.Increment()
		return unsafe.Pointer(o)
	})

	if *(*int)(v) != 1 {
		t.Errorf("once failed inside run: %d is not 1", *(*int)(v))
	}
	c <- true
}

func TestOncePointer(t *testing.T) {
	o := new(one)
	once := new(OncePointer)
	c := make(chan bool)
	const N = 10
	for i := 0; i < N; i++ {
		go runPointer(t, once, o, c)
	}
	for i := 0; i < N; i++ {
		<-c
	}
	if *o != 1 {
		t.Errorf("once failed outside run: %d is not 1", *o)
	}
}

func TestOncePointerPanic(t *testing.T) {
	var once OncePointer
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Once.Do did not panic")
			}
		}()

		once.Do(func() unsafe.Pointer{
			panic("failed")
			return nil
		})
	}()

	_ = once.Do(func() unsafe.Pointer{
		t.Fatalf("Once.Do called twice")
		return nil
	})
}

func BenchmarkOncePointer(b *testing.B) {
	var once OncePointer
	var o = new(one)
	f := func() unsafe.Pointer{
		return unsafe.Pointer(o)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = (*one)(once.Do(f))
		}
	})
}
