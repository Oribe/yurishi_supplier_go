package server

import "manufacture_supplier_go/model"

// OrderQuery 订单查询
func OrderQuery(startTime, endTime string, userID int64) (model.Order, error) {

	orderList := model.Order{}

	err := model.OrderQuery(&orderList, startTime, endTime, userID)
	if err != nil {
		return orderList, err
	}

	return orderList, nil
}
