package database

import (
	"context"

	"github.com/azul915/portfolio-api/src/domain"
)

// ProductRepository は、ProductドメインについてCloudFirestoreとのやり取りを担うRepository
type ProductRepository struct {
	Val interface{}
}

// FindAll は、全ての「products」コレクションを取得する
func (repo *ProductRepository) FindAll() (products domain.Products, err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	data := client.Collection("products").Documents(ctx)

	docs, err := data.GetAll()
	if err != nil {
		return
	}

	products = make(domain.Products, 0)
	for _, doc := range docs {
		p := new(domain.Product)
		mapToStruct(doc.Data(), &p)
		products = append(products, *p)
	}

	defer client.Close()

	return

}
