package model

import (
	"errors"
)

// Order 订单
type Order struct {
	ID          int    `db:"id" json:"id"`
	OrderNumber string `db:"orderNumber" json:"orderNumber"`
	Quantity    int    `db:"quantity" json:"quantity"`
	Supplier    string `db:"supplier" json:"supplier"`
	Submitter   int    `db:"submitter" json:"submitter"`
	CreateAt    string `db:"createAt" json:"createAt"`
	UpdateAt    string `db:"updateAt" json:"updateAt"`
}

// OrderQuery 订单查询
func OrderQuery(orderList *[]Order, startTime, endTime string, userID int) error {
	q := `SELECT 
					id, orderNo, modelNumber, quantity, supplier, submitter, createAt, updateAt 
				FROM 
					order_suppliers 
				WHERE 
					submitter = ? 
				AND 
					createAt >= ? 
				AND 
					createAt <= ?`

	stmt, err := db.Preparex(q)
	if err != nil {
		error := errors.New("order prepare query error: " + err.Error())
		return error
	}

	err = stmt.Select(orderList, startTime, endTime, userID)
	if err != nil {
		error := errors.New("order query error: " + err.Error())
		return error
	}

	return nil
}
