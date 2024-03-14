package generate

import (
	"debezium-test-data-generator/model"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func GenerateCustomer(db *gorm.DB, numCustomers int) {
	for i := 0; i < numCustomers; i++ {
		customer := model.Customer{
			FirstName:   gofakeit.FirstName(),
			LastName:    gofakeit.LastName(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Address:     gofakeit.Address().Address,
		}
		db.Create(&customer)
	}
}

func GenerateProduct(db *gorm.DB, numProducts int) {
	for i := 0; i < numProducts; i++ {
		productName := gofakeit.ProductName() // 상품 이름을 생성
		product := model.Product{
			Name:        productName,
			Description: fmt.Sprintf("Introducing the %s: %s", productName, gofakeit.HipsterSentence(4)),
			Price:       gofakeit.Float64Range(10.0, 1000.0),
		}
		db.Create(&product)
	}
}

func GenerateOrders(db *(gorm.DB), customersIdList, productIdList []uint64, numbOrders int) []model.Order {
	src := rand.NewSource(time.Now().UnixNano()) // 새로운 소스 생성
	r := rand.New(src)                           // 새로운 *rand.Rand 인스턴스 생성

	var orders []model.Order
	for i := 0; i < numbOrders; i++ {
		order := model.Order{
			CustomerID: customersIdList[r.Intn(len(customersIdList))], // uint64 타입으로 변경
			ProductID:  productIdList[r.Intn(len(productIdList))],
			OrderDate:  time.Now().Add(-time.Duration(r.Intn(30)) * 24 * time.Hour),
			Quantity:   rand.Intn(10) + 1, // 1 ~ 10 사이의 수량
		}
		db.Create(&order)
		log.Printf("Orderd product.  OrderId: %v \n", order.ID)

	}

	return orders
}
