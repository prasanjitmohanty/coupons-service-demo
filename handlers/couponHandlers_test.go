package handlers

import (
	"bytes"
	"coupons-service-demo/coupon"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_bodyToCoupon(t *testing.T) {
	valid := &coupon.Coupon{
		Name:      "Save 20 at Tesco",
		Brand:     "Tesco",
		Value:     30,
		CreatedAt: "2019-02-20 10:15:53",
		Expiry:    "2019-03-01 10:15:53",
	}

	js, err := json.Marshal(valid)
	if err != nil {
		t.Errorf("Error marshalling a valid user: %s", err)
		t.FailNow()
	}

	ts := []struct {
		txt string
		r   *http.Request
		u   *coupon.Coupon
		err bool
		exp *coupon.Coupon
	}{
		{
			txt: "nil request",
			err: true,
		},
		{
			txt: "empty request body",
			r:   &http.Request{},
			err: true,
		},
		{
			txt: "empty coupon",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			err: true,
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"name1": "Save 20 at Tesco"`)),
			},
			u:   &coupon.Coupon{},
			err: true,
		},
		{
			txt: "valid request",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(js)),
			},
			u:   &coupon.Coupon{},
			exp: valid,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToCoupon(tc.r, tc.u)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.u, tc.exp) {
			t.Error("Unmarshalled data is different:")
			t.Error(tc.u)
			t.Error(tc.exp)
		}
	}
}
