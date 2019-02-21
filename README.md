# coupons-service-demo
coupons service demo

## Restful API using GO 
## End points Examples
*localhost:11111/coupons - Get All Coupons (GET)*
* localhost:11111/coupons/?Brand=Tesco&Value=30 - Get Filtered  Coupons (GET)*
* localhost:11111/coupons -Create New Coupon (POST)*
* localhost:11111/coupons/Save 30 at Tesco - Get Single Coupon by Name -(Get)*
* localhost:11111/coupons/Save 30 at Tesco - Update Existing Coupon Details -(PUT)*
## DB
### In memory BoltDB (An embedded key/value database for Go) is used for Data storage using storm (https://github.com/asdine/storm)

Includes Unit tests
