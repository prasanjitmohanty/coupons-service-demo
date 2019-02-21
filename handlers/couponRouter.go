package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// CouponsRouter handles the coupons route
func CouponsRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	
	

	if path == "/coupons" {
		
		switch r.Method {
		case http.MethodGet:
			u, _ := url.Parse(r.URL.String())
		queryParams := u.Query()

		if len(queryParams) > 0{
			filterCoupons(w,r,queryParams)

		}else{
			couponsGetAll(w, r)
		}
		return
		case http.MethodPost:
			createCoupon(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	couponName := strings.TrimPrefix(path, "/coupons/")
	fmt.Println(couponName)

	switch r.Method {
	case http.MethodGet:
		couponsGetOne(w, r, couponName)
		return
	case http.MethodPut:
		updateCoupon(w, r, couponName)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
