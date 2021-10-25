package easegress

import (
	"strconv"
	"strings"
)

type SameSite int32

const (
	DefaultMode SameSite = iota
	LaxMode
	StrictMode
	NoneMode
)

type Cookie struct {
	// TODO mv rawExpires to expires as type date
	name       string
	value      string
	path       string
	domain     string
	rawExpires string
	maxAge     int32
	secure     bool
	httpOnly   bool
	sameSite   SameSite
}

// region getter & setter generated automatically

func (c *Cookie) Name() string {
	return c.name
}

func (c *Cookie) SetName(name string) {
	c.name = name
}

func (c *Cookie) Value() string {
	return c.value
}

func (c *Cookie) SetValue(value string) {
	c.value = value
}

func (c *Cookie) Path() string {
	return c.path
}

func (c *Cookie) SetPath(path string) {
	c.path = path
}

func (c *Cookie) Domain() string {
	return c.domain
}

func (c *Cookie) SetDomain(domain string) {
	c.domain = domain
}

func (c *Cookie) RawExpires() string {
	return c.rawExpires
}

func (c *Cookie) SetRawExpires(rawExpires string) {
	c.rawExpires = rawExpires
}

func (c *Cookie) MaxAge() int32 {
	return c.maxAge
}

func (c *Cookie) SetMaxAge(maxAge int32) {
	c.maxAge = maxAge
}

func (c *Cookie) Secure() bool {
	return c.secure
}

func (c *Cookie) SetSecure(secure bool) {
	c.secure = secure
}

func (c *Cookie) HttpOnly() bool {
	return c.httpOnly
}

func (c *Cookie) SetHttpOnly(httpOnly bool) {
	c.httpOnly = httpOnly
}

func (c *Cookie) SameSite() SameSite {
	return c.sameSite
}

func (c *Cookie) SetSameSite(sameSite SameSite) {
	c.sameSite = sameSite
}

// endregion

func (c *Cookie) marshal() string {
	if len(c.name) <= 0 {
		panic("cookie name must be specified")
	}

	var str = c.name + "=" + c.value

	if len(c.path) > 0 {
		str += "; Path=" + c.path
	}

	if len(c.domain) > 0 {
		str += "; Domain=" + c.domain
	}

	if len(c.rawExpires) > 0 {
		str += "; Expires=" + c.rawExpires
	}

	if c.maxAge > 0 {
		str += "; Max-Age=" + strconv.Itoa(int(c.maxAge))
	} else if c.maxAge < 0 {
		str += "; Max-Age=0"
	}

	if c.secure {
		str += "; Secure"
	}

	if c.httpOnly {
		str += "; HttpOnly"
	}

	switch c.sameSite {
	case NoneMode:
		str += "; SameSite=None"
		break
	case LaxMode:
		str += "; SameSite=Lax"
		break
	case StrictMode:
		str += "; SameSite=Strict"
		break
	}

	return str
}

func (c *Cookie) unmarshal(str string) *Cookie {
	var parts = strings.Split(str, ";")
	var kv = strings.Split(strings.TrimSpace(parts[0]), "=")
	if len(kv) != 2 {
		return nil
	}

	c.name = kv[0]
	c.value = kv[1]

	for i := 1; i < len(parts); i++ {
		kv = strings.Split(strings.TrimSpace(parts[i]), "=")
		if len(kv) != 2 {
			continue
		}

		var k = strings.ToLower(kv[0])
		if k == "path" {
			c.path = kv[1]
		} else if k == "domain" {
			c.domain = kv[1]
		} else if k == "expires" {
			c.rawExpires = kv[1]
		} else if k == "max-age" {
			age, err := strconv.Atoi(kv[1])
			if err != nil {
				panic("max-age in cookie should be int but:" + kv[1])
			}
			c.maxAge = int32(age)
		} else if k == "secure" {
			c.secure = true
		} else if k == "httponly" {
			c.httpOnly = true
		} else if k == "samesite" {
			var v = strings.ToLower(kv[1])
			if v == "lax" {
				c.sameSite = LaxMode
			} else if v == "strict" {
				c.sameSite = StrictMode
			} else if v == "none" {
				c.sameSite = NoneMode
			} else {
				c.sameSite = DefaultMode
			}
		}
	}

	return c
}
