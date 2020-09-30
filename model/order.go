package model

import (
	"errors"
)

type Order struct {
	ID          int64  `db:"id" json:"id"`
	OrderNumber string `db:"orderNumber" json:"orderNumber"`
	Quantity    int    `db:"quantity" json:"quantity"`
	Supplier    string `db:"supplier" json:"supplier"`
	Submitter   int64  `db:"submitter" json:"submitter"`
	CreateAt    string `db:"createAt" json:"createAt"`
	UpdateAt    string `db:"updateAt" json:"updateAt"`
}

// OrderQuery ...
func OrderQuery(orderList *Order, startTime, endTime string, userID int64) error {
	q := `SELECT 
					id, orderNo, modelNumber, quantity, supplier, submitter, createAt, updateAt
				FORM
					order_suppliers
				WHERE
					submitter = ? 
				AND
					create >= ?
				AND
					create <= ?`

	stmt, err := db.Prepare(q)
	if err != nil {
		error := errors.New("order prepare query error: " + err.Error())
		return error
	}

	rows, err := stmt.Query(startTime, endTime, userID)
	if err != nil {
		error := errors.New("order query error: " + err.Error())
		return error
	}

	err = rows.Scan(orderList)
	if err != nil {
		error := errors.New("order query scan error: " + err.Error())
		return error
	}

	return nil
}
