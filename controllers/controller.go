package controllers

import (
	"build-api/database"
	"build-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddOrder(ctx *gin.Context) {
	var order AddOrderRequest

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var itemsInput []models.Items
	for _, value := range order.Items {
		itemsInput = append(itemsInput, models.Items{
			ItemCode:    value.ItemCode,
			Description: value.Description,
			Quantity:    value.Quantity,
		})
	}
	orderInput := models.Order{
		CustomerName: order.CustomerName,
		OrderAt:      order.OrderAt,
		Items:        itemsInput,
	}

	var db = database.GetDB()

	if err := db.Create(&orderInput).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error create order",
			"data":    "",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "success",
		"data":    order,
	},
	)
}

func GetAllData(ctx *gin.Context) {
	var db = database.GetDB()

	var orders []models.Order

	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error get all orders",
			"data":    "",
		})
		return
	}

	var ordersRes []OrderResponse
	for _, orderValue := range orders {
		var itemsResp []ItemsResponse

		for _, itemValue := range orderValue.Items {
			itemsResp = append(itemsResp, ItemsResponse{
				ID:          itemValue.ID,
				ItemCode:    itemValue.ItemCode,
				OrderId:     itemValue.OrderId,
				Description: itemValue.Description,
				Quantity:    itemValue.Quantity,
				CreatedAt:   itemValue.CreatedAt,
				UpdatedAt:   itemValue.UpdatedAt,
			})
		}

		ordersRes = append(ordersRes, OrderResponse{
			ID:           orderValue.ID,
			CustomerName: orderValue.CustomerName,
			Items:        itemsResp,
			CreatedAt:    orderValue.CreatedAt,
			UpdatedAt:    orderValue.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    ordersRes,
	},
	)
}

func DeleteOrder(ctx *gin.Context) {
	var db = database.GetDB()

	orderId := ctx.Param("orderId")

	db.Where("order_id = ?", orderId).Delete(&models.Items{})

	if err := db.Where("id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
			"data":    "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	},
	)
}

func DeleteItem(ctx *gin.Context) {
	var db = database.GetDB()

	orderId := ctx.Param("orderId")
	itemId := ctx.Param("itemId")

	if err := db.Where("id = ? AND order_id = ?", itemId, orderId).Delete(&models.Items{}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
			"data":    "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	},
	)
}

func UpdateOrder(ctx *gin.Context) {
	var db = database.GetDB()

	var orderData models.Order
	orderId := ctx.Param("orderId")

	if err := db.Preload("Items").First(&orderData, orderId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var order AddOrderRequest

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var items []models.Items
	itemExist := make(map[string]bool)

	for i := 0; i < len(order.Items); i++ {

		// if there is duplicate item code in body request, input the last data duplicate data from request
		if itemExist[order.Items[i].ItemCode] {
			for k := 0; k < len(items); k++ {
				if items[k].ItemCode == order.Items[i].ItemCode {
					items[k].Description = order.Items[i].Description
					items[k].Quantity = order.Items[i].Quantity
				}
			}
			break
		}

		// mapping items data
		for j := 0; j < len(orderData.Items); j++ {

			// if items is already exist in database, update it
			if order.Items[i].ItemCode == orderData.Items[j].ItemCode {
				items = append(items, models.Items{
					ID:          orderData.Items[j].ID,
					ItemCode:    order.Items[i].ItemCode,
					Description: order.Items[i].Description,
					Quantity:    order.Items[i].Quantity,
					OrderId:     orderData.ID,
				})

				itemExist[order.Items[i].ItemCode] = true
				break
			}

			// if items in not exist, add it
			if j == len(orderData.Items)-1 {
				items = append(items, models.Items{
					ItemCode:    order.Items[i].ItemCode,
					Description: order.Items[i].Description,
					Quantity:    order.Items[i].Quantity,
					OrderId:     orderData.ID,
				})
				itemExist[order.Items[i].ItemCode] = true
			}
		}
	}

	// Change order data
	orderData.CustomerName = order.CustomerName
	orderData.OrderAt = order.OrderAt

	// update order
	if err := db.Save(&orderData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update orders",
			"data":    "",
		})
		return
	}

	// update items
	if err := db.Save(&items).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update items",
			"data":    "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    order,
	},
	)
}
