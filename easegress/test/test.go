//go:build test
// +build test

package test

import (
	"errors"

	"github.com/xmh19936688/easegress-go-sdk/easegress/util"
)

func RunTest() error {
	if err := testParseTime(); err != nil {
		return err
	}

	if err := testParseInt(); err != nil {
		return err
	}

	if err := testParseUint(); err != nil {
		return err
	}

	if err := testAtoi(); err != nil {
		return err
	}

	if err := testParseFloat(); err != nil {
		return err
	}

	return nil
}

func testParseTime() error {
	cases := []struct {
		String  string
		Layout  string
		Seconds int64
	}{
		{
			String:  "2021-11-06 21:29:36",
			Layout:  "2006-01-02 15:04:05",
			Seconds: int64(1636205376),
		},
		{
			String:  "2021-11-06T21:29:36",
			Layout:  "2006-01-02T15:04:05",
			Seconds: int64(1636205376),
		},
	}

	for _, c := range cases {
		res, err := util.ParseTime(c.Layout, c.String)
		if err != nil {
			return err
		}
		if res.Unix() != c.Seconds {
			return errors.New("ParseTime failed")
		}
	}

	return nil
}

func testParseInt() error {
	cases := []struct {
		String string
		Base   int
		Bit    int
		Value  int64
	}{
		{
			String: "123",
			Base:   10,
			Bit:    64,
			Value:  123,
		},
	}

	for _, c := range cases {
		res, err := util.ParseInt(c.String, c.Base, c.Bit)
		if err != nil {
			return err
		}
		if res != c.Value {
			return errors.New("ParseInt failed")
		}
	}

	return nil
}

func testParseUint() error {
	cases := []struct {
		String string
		Base   int
		Bit    int
		Value  uint64
	}{
		{
			String: "123",
			Base:   10,
			Bit:    64,
			Value:  123,
		},
	}

	for _, c := range cases {
		res, err := util.ParseUint(c.String, c.Base, c.Bit)
		if err != nil {
			return err
		}
		if res != c.Value {
			return errors.New("ParseInt failed")
		}
	}

	return nil
}

func testAtoi() error {
	cases := []struct {
		String string
		Value  int
	}{
		{
			String: "123",
			Value:  123,
		},
	}

	for _, c := range cases {
		res, err := util.Atoi(c.String)
		if err != nil {
			return err
		}
		if res != c.Value {
			return errors.New("Atoi failed")
		}
	}

	return nil
}

func testParseFloat() error {
	cases := []struct {
		String string
		Bit    int
		Value  float64
	}{
		{
			String: "123.456",
			Bit:    64,
			Value:  123.456,
		},
	}

	for _, c := range cases {
		res, err := util.ParseFloat(c.String, c.Bit)
		if err != nil {
			return err
		}
		if res != c.Value {
			return errors.New("ParseFloat failed")
		}
	}

	return nil
}
