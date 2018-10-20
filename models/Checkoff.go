package models

import "math/big"

type ActiveUserModel struct {
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	CheckInTime  big.Int `json:"checkInTime"`
	CheckOutTime big.Int `json:"checkOutTime"`
}
