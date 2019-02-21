package handlers

import (
	"coupons-service-demo/coupon"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func bodyToCoupon(r *http.Request, c *coupon.Coupon) error {
	if r == nil {
		return errors.New("a request is required")
	}
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	if c == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}
	return json.Unmarshal(bd, c)
}

func couponsGetAll(w http.ResponseWriter, r *http.Request) {
	coupons, err := coupon.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"coupons": coupons})
}

func filterCoupons(w http.ResponseWriter, r *http.Request, filters map[string][]string) {
	coupons, err := coupon.Filter(filters)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"coupons": coupons})
}

func couponsGetOne(w http.ResponseWriter, r *http.Request, couponName string) {
	c, err := coupon.One(couponName)
	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"coupon": c})
}

func createCoupon(w http.ResponseWriter, r *http.Request) {
	c := new(coupon.Coupon)
	err := bodyToCoupon(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = c.Save()
	if err != nil {
		if err == coupon.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/coupons/"+c.Name)
	w.WriteHeader(http.StatusCreated)
}

func updateCoupon(w http.ResponseWriter, r *http.Request, couponName string) {
	c, er := coupon.One(couponName)

	if er != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err := bodyToCoupon(r, c)
	if err != nil {

		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save()
	if err != nil {
		if err == coupon.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"coupon": c})
}
