package database

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	"portfolio-api/api/domain"
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

// Store は、引数で受け取ったProductについて、新たなドキュメントを追加する
func (repo *ProductRepository) Store(rp domain.ReqProduct) (err error) {

	ctx := context.Background()

	client, err := firebaseInit(ctx)
	if err != nil {
		return
	}

	ap := domain.AddProduct{
		CreatedAt: time.Now(),
		DemoURL:   rp.DemoURL,
		Effort:    rp.DemoURL,
		Feature:   rp.Feature,
		GithubURL: rp.GithubURL,
		Name:      rp.Name,
	}

	// コレクション "products" 作成
	doc := client.Collection("products").Doc(ap.Name)

	_, err = doc.Set(ctx, ap)
	if err != nil {
		return
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
