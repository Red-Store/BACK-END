package data

import (
	"MyEcommerce/features/order"
	"MyEcommerce/utils/externalapi"
	cd "MyEcommerce/features/cart/data"
	"errors"

	"gorm.io/gorm"
)

type orderQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mi externalapi.MidtransInterface) order.OrderDataInterface {
	return &orderQuery{
		db:              db,
		paymentMidtrans: mi,
	}
}

// InsertOrder implements order.OrderDataInterface.
func (repo *orderQuery) InsertOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore) (*order.OrderCore, error) {
	payment, errPay := repo.paymentMidtrans.NewOrderPayment(inputOrder)
	if errPay != nil {
		return nil, errPay
	}

	orderModel := CoreToModelOrder(inputOrder)
	orderModel.UserID = uint(userIdLogin)

	orderModel.PaymentType = payment.PaymentType
	orderModel.Status = payment.Status
	orderModel.VaNumber = payment.VaNumber

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	inputOrder.ID = orderModel.ID

	for _, cartId := range cartIds {
		orderItem := OrderItem{
			OrderID: orderModel.ID,
			CartID:  cartId,
		}

		if err := repo.db.Create(&orderItem).Error; err != nil {
			return nil, err
		}
		
		var cart  cd.Cart

		if err := repo.db.Preload("Product").Where("id = ?", cartId).First(&cart).Error; err != nil {
			return nil, err
		}

		cart.Product.Stock -= cart.Quantity

		if err := repo.db.Save(&cart.Product).Error; err != nil {
			return nil, err
	}

	}

	return payment, nil
}

// SelectOrderUser implements order.OrderDataInterface.
func (repo *orderQuery) SelectOrderUser(userIdLogin int) ([]order.OrderItemCore, error) {
	var orderItems []OrderItem
	err := repo.db.Joins("Order").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.User").Where("user_id = ?", userIdLogin).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	orderItemCores := make([]order.OrderItemCore, len(orderItems))
	for i, item := range orderItems {
		orderItemCores[i] = item.ModelToCoreOrderItemUser()
	}

	return orderItemCores, nil
}

// SelectOrderAdmin implements order.OrderDataInterface.
func (repo *orderQuery) SelectOrderAdmin(page, limit int) ([]order.OrderItemCore, error) {
	var orderItems []OrderItem
	err := repo.db.Joins("Order").Preload("Cart").Preload("Cart.Product").Limit(limit).Offset((page - 1) * limit).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	orderItemCores := make([]order.OrderItemCore, len(orderItems))
	for i, item := range orderItems {
		orderItemCores[i] = item.ModelToCoreOrderItemAdmin()
	}

	return orderItemCores, nil
}

// SelectOrderAdmin implements order.OrderDataInterface.
func (repo *orderQuery) CancleOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "cancelled" {
		repo.paymentMidtrans.CancelOrderPayment(orderId)
	}

	dataGorm := CoreToModelOrderCancle(orderCore)
	tx := repo.db.Model(&Order{}).Where("id = ? AND user_id = ?", orderId, userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Update implements user.UserDataInterface.
func (repo *orderQuery) WebhoocksData(reqNotif order.OrderCore) error {
	dataGorm := CoreToModel(reqNotif)
	tx := repo.db.Model(&Order{}).Where("id = ?", reqNotif.ID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
