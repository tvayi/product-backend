package models

import "github.com/go-pg/pg/v10"

type Product struct {
    Code string `json:"code"`
    Name string `json:"name"`
    Weight float64 `json:"weight"`
    Description string `json:"description"`
}

func CreateProduct(db *pg.DB, req *Product) (*Product, error) {
    _, err := db.Model(req).Insert()
    if err != nil {
        return nil, err
    }

    product := &Product{}

    err = db.Model(product).
        Where("product.code = ?", req.Code).
        Select()

    return product, err
}

func GetProduct(db *pg.DB, productCode string) (*Product, error) {
    product := &Product{}

    err := db.Model(product).
        Where("product.code = ?", productCode).
        Select()

    return product, err
}

func GetProducts(db *pg.DB) ([]*Product, error) {
    products := make([]*Product, 0)

    err := db.Model(&products).
        Select()

    return products, err
}
