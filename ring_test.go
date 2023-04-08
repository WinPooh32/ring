package ring

import (
	"reflect"
	"testing"
)

func TestRing(t *testing.T) {
	const capacity = 9

	var ringBuf = Make[int](capacity)
	var data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i, v := range data {
		t.Logf("push back %v", v)

		pop, elem := ringBuf.Push(v)
		if pop && i >= capacity && elem != data[i-capacity] {
			t.FailNow()
		}
	}

	t.Logf("ring internal buf: %v", ringBuf.buf)

	var buf = make([]int, capacity)
	n := ringBuf.CopyTo(buf)
	got := buf[:n]

	t.Logf("result buf: %v", got)

	want := []int{8, 9, 10, 11, 12, 13, 14, 15, 16}

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}

	k := 0
	ringBuf.Range(func(i int, v int) {
		if want[i] != v {
			t.Logf("want: %+v, got: %+v", want[i], v)
			t.Fail()
			return
		}

		k++
	})

	if k != len(want) {
		t.Logf("want length: %+v, got: %+v", len(want), k)
	}

	l, r := ringBuf.TwoParts()

	t.Log(l, r)

	buf = append(buf[:0], l...)
	buf = append(buf, r...)
	got = buf

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}

	if front := ringBuf.Front(); front != 8 {
		t.Logf("want: %+v, got: %+v", 8, front)
		t.Fail()
	}

	if back := ringBuf.Back(); back != 16 {
		t.Logf("want: %+v, got: %+v", 16, back)
		t.Fail()
	}

	ringBuf.Reset()

	n = ringBuf.CopyTo(buf[:cap(buf)])
	got = buf[:n]

	t.Logf("result buf: %v", got)

	want = []int{}

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}
}

func TestRing2(t *testing.T) {
	const capacity = 24

	var ringBuf = Make[int](capacity)
	var data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i, v := range data {
		t.Logf("push back %v", v)

		pop, elem := ringBuf.Push(v)
		if pop && i >= capacity && elem != data[i-capacity] {
			t.FailNow()
		}
	}

	t.Logf("ring internal buf: %v", ringBuf.buf)

	var buf = make([]int, capacity)
	n := ringBuf.CopyTo(buf)
	got := buf[:n]

	t.Logf("result buf: %v", got)

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}

	k := 0
	ringBuf.Range(func(i int, v int) {
		if want[i] != v {
			t.Logf("want: %+v, got: %+v", want[i], v)
			t.Fail()
			return
		}

		k++
	})

	if k != len(want) {
		t.Logf("want length: %+v, got: %+v", len(want), k)
	}

	l, r := ringBuf.TwoParts()

	buf = append(buf[:0], l...)
	buf = append(buf, r...)
	got = buf

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}

	if front := ringBuf.Front(); front != 0 {
		t.Logf("want: %+v, got: %+v", 0, front)
		t.Fail()
	}

	if back := ringBuf.Back(); back != 16 {
		t.Logf("want: %+v, got: %+v", 16, back)
		t.Fail()
	}

	ringBuf.Reset()

	n = ringBuf.CopyTo(buf[:cap(buf)])
	got = buf[:n]

	t.Logf("result buf: %v", got)

	want = []int{}

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}
}
