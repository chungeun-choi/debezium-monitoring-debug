package generate

import (
	"debezium-test-data-generator/model"
	"gorm.io/gorm"
)

func RandomCustomerIdMySQL(db *gorm.DB) []uint64 {
	var customers []model.Customer

	// Customer 테이블에서 랜덤하게 10개의 레코드를 선택
	err := db.Order("RAND()").Limit(10).Find(&customers).Error
	if err != nil {
		panic("Failed to query database")
	}

	// ID 값만 추출하여 슬라이스로 저장
	var customerIDs []uint64
	for _, customer := range customers {
		customerIDs = append(customerIDs, customer.ID)
	}

	return customerIDs
}

func RandomProductIdMySQL(db *gorm.DB) []uint64 {
	var products []model.Product

	// Customer 테이블에서 랜덤하게 10개의 레코드를 선택
	err := db.Order("RAND()").Limit(30).Find(&products).Error
	if err != nil {
		panic("Failed to query database")
	}

	// ID 값만 추출하여 슬라이스로 저장
	var productsIDs []uint64
	for _, customer := range products {
		productsIDs = append(productsIDs, customer.ID)
	}

	return productsIDs
}

func RandomCustomerIdPostgreSQL(db *gorm.DB) []uint64 {
	var customers []model.Customer

	// PostgreSQL의 RANDOM() 함수를 사용하여 Customer 테이블에서 랜덤하게 10개의 레코드를 선택
	err := db.Order("RANDOM()").Limit(10).Find(&customers).Error
	if err != nil {
		panic("Failed to query database")
	}

	var customerIDs []uint64
	for _, customer := range customers {
		customerIDs = append(customerIDs, customer.ID)
	}

	return customerIDs
}

func RandomProductIdPostgreSQL(db *gorm.DB) []uint64 {
	var products []model.Product

	// PostgreSQL의 RANDOM() 함수를 사용하여 Product 테이블에서 랜덤하게 30개의 레코드를 선택
	err := db.Order("RANDOM()").Limit(30).Find(&products).Error
	if err != nil {
		panic("Failed to query database")
	}

	var productIDs []uint64
	for _, product := range products {
		productIDs = append(productIDs, product.ID)
	}

	return productIDs
}
