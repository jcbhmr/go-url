package url

import (
	"fmt"
	"iter"
	"syscall/js"
)

type URLSearchParams struct {
	value js.Value
}

var urlSearchParams = js.Global().Get("URLSearchParams")

func NewURLSearchParams(init any) *URLSearchParams {
	switch init := init.(type) {
	case [][]string:
		return &URLSearchParams{urlSearchParams.New(init)}
	case map[string]string:
		return &URLSearchParams{urlSearchParams.New(init)}
	case string:
		return &URLSearchParams{urlSearchParams.New(init)}
	default:
		panic("init not [][]string | map[string]string | string")
	}
}

func (u *URLSearchParams) Size() uint32 {
	return uint32(u.value.Get("size").Int())
}

func (u *URLSearchParams) Append(name string, value string) {
	u.value.Call("append", name, value)
}

func (u *URLSearchParams) Delete(name string) {
	u.value.Call("delete", name)
}

func (u *URLSearchParams) Get(name string) (string, bool) {
	v := u.value.Call("get", name)
	if v.IsUndefined() {
		return "", false
	} else {
		return v.String(), true
	}
}

func (u *URLSearchParams) GetAll(name string) []string {
	v := u.value.Call("getAll", name)
	a := make([]string, v.Length())
	for i := range a {
		a[i] = v.Index(i).String()
	}
	return a
}

func (u *URLSearchParams) Has(name string) bool {
	return u.value.Call("has", name).Bool()
}

func (u *URLSearchParams) Set(name string, value string) {
	u.value.Call("set", name, value)
}

func (u *URLSearchParams) Sort() {
	u.value.Call("sort")
}

func (u *URLSearchParams) All() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		g := u.value.Call("entries")
		defer func() {
			if r := recover(); r != nil {
				g.Call("throw", fmt.Sprint(r))
			}
		}()
		for {
			o := g.Call("next")
			done := o.Get("done").Bool()
			if done {
				break
			}
			value := o.Get("value")
			valueA := value.Index(0).String()
			valueB := value.Index(1).String()
			if !yield(valueA, valueB) {
				g.Call("return")
				break
			}
		}
	}
}

func (u *URLSearchParams) Keys() iter.Seq[string] {
	return func(yield func(string) bool) {
		g := u.value.Call("keys")
		defer func() {
			if r := recover(); r != nil {
				g.Call("throw", fmt.Sprint(r))
			}
		}()
		for {
			o := g.Call("next")
			done := o.Get("done").Bool()
			if done {
				break
			}
			value := o.Get("value").String()
			if !yield(value) {
				g.Call("return")
				break
			}
		}
	}
}

func (u *URLSearchParams) Values() iter.Seq[string] {
	return func(yield func(string) bool) {
		g := u.value.Call("values")
		defer func() {
			if r := recover(); r != nil {
				g.Call("throw", fmt.Sprint(r))
			}
		}()
		for {
			o := g.Call("next")
			done := o.Get("done").Bool()
			if done {
				break
			}
			value := o.Get("value").String()
			if !yield(value) {
				g.Call("return")
				break
			}
		}
	}
}

func (u *URLSearchParams) String() string {
	return u.value.Call("toString").String()
}