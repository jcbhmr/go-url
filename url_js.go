package url

import (
	"encoding/json"
	"syscall/js"
)

type URL struct {
	value             js.Value
	searchParamsCache *URLSearchParams
}

var url2 = js.Global().Get("URL")

func NewURL(url string, base *string) *URL {
	if base == nil {
		return &URL{url2.New(url), nil}
	} else {
		return &URL{url2.New(url, *base), nil}
	}
}

func URLParse(url string, base *string) *URL {
	if base == nil {
		return &URL{url2.Call("parse", url), nil}
	} else {
		return &URL{url2.Call("parse", url, *base), nil}
	}
}

func URLCanParse(url string, base *string) bool {
	if base == nil {
		return url2.Call("canParse", url).Bool()
	} else {
		return url2.Call("canParse", url, *base).Bool()
	}
}

func (u *URL) String() string {
	return u.Href()
}

func (u *URL) Href() string {
	return u.value.Get("href").String()
}

func (u *URL) SetHref(href string) {
	u.value.Set("href", href)
}

func (u *URL) Origin() string {
	return u.value.Get("origin").String()
}

func (u *URL) Protocol() string {
	return u.value.Get("protocol").String()
}

func (u *URL) SetProtocol(protocol string) {
	u.value.Set("protocol", protocol)
}

func (u *URL) Username() string {
	return u.value.Get("username").String()
}

func (u *URL) SetUsername(username string) {
	u.value.Set("username", username)
}

func (u *URL) Password() string {
	return u.value.Get("password").String()
}

func (u *URL) SetPassword(password string) {
	u.value.Set("password", password)
}

func (u *URL) Host() string {
	return u.value.Get("host").String()
}

func (u *URL) SetHost(host string) {
	u.value.Set("host", host)
}

func (u *URL) Hostname() string {
	return u.value.Get("hostname").String()
}

func (u *URL) SetHostname(hostname string) {
	u.value.Set("hostname", hostname)
}

func (u *URL) Port() string {
	return u.value.Get("port").String()
}

func (u *URL) SetPort(port string) {
	u.value.Set("port", port)
}

func (u *URL) Pathname() string {
	return u.value.Get("pathname").String()
}

func (u *URL) SetPathname(pathname string) {
	u.value.Set("pathname", pathname)
}

func (u *URL) Search() string {
	return u.value.Get("search").String()
}

func (u *URL) SetSearch(search string) {
	u.value.Set("search", search)
}

func (u *URL) SearchParams() *URLSearchParams {
	if u.searchParamsCache == nil {
		u.searchParamsCache = &URLSearchParams{u.value.Get("searchParams")}
	}
	return u.searchParamsCache
}

func (u *URL) Hash() string {
	return u.value.Get("hash").String()
}

func (u *URL) SetHash(hash string) {
	u.value.Set("hash", hash)
}

func (u *URL) ToJSON() string {
	return u.value.Call("toJSON").String()
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.ToJSON())
}
