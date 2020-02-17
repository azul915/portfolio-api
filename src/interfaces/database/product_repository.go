package database

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/azul915/portfolio-api/src/domain"
)

// ProductRepository は、ProductドメインについてCloudFirestoreとのやり取りを担うRepository
type ProductRepository struct {
	Val interface{}
}

// GetAll は、全ての「products」コレクションを取得する
func (repo *ProductRepository) GetAll() (products domain.Products, err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	// 「products」コレクションをcreated_at/ascで取得
	data := client.Collection("products").OrderBy("created_at", firestore.Asc).Documents(ctx)

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

// Delete は、引数で受け取った値を「domain.DelProduct.Name」として、該当するドキュメントを削除する
func (repo *ProductRepository) Delete(d domain.DelProduct) (err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	_, err = client.Collection("products").Doc(d.Name).Delete(ctx)
	if err != nil {
		return
	}

	defer client.Close()

	return

}
