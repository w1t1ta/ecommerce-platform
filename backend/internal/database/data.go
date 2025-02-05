// data.go
package data

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	ModelNumber string    `json:"model_number"`
	SKU         string    `json:"sku"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	SellerID    string    `json:"seller_id"`
	ProductType string    `json:"product_type"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	ModelNumber string  `json:"model_number"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	SellerID    string  `json:"seller_id"`
	ProductType string  `json:"product_type"`
	CategoryID  int     `json:"category_id"`
}

type UpdateProduct struct {
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type ProductImage struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	IsPrimary bool      `json:"is_primary"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}

type NewProductImage struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
	SortOrder int    `json:"sort_order"`
}

type UpdateProductImage struct {
	IsPrimary bool `json:"is_primary"`
	SortOrder int  `json:"sort_order"`
}

type ProductQueryParams struct {
	Cursor      string `json:"cursor"`
	Limit       int    `json:"limit"`
	Search      string `json:"search"`
	CategoryID  int    `json:"category_id"`
	SellerID    string `json:"seller_id"`
	Status      string `json:"status"`
	ProductType string `json:"product_type"`
	Sort        string `json:"sort"`
	Order       string `json:"order"`
}

type ProductResponse struct {
	Items      []ProductItem `json:"items"`
	NextCursor string        `json:"next_cursor"`
	Limit      int           `json:"limit"`
}

type ProductItem struct {
	Product
	Categories []Category      `json:"categories"`
	Inventory  Inventory       `json:"inventory"`
	Images     []ProductImage  `json:"images"`
	Options    []ProductOption `json:"options"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Inventory struct {
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductOption struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Values json.RawMessage `json:"values"`
}
type User_local struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Shipping_addresses struct {
	User      string `json:"user_id"`
	Name      string `json:"first_name"`
	Company   string `json:"company_name"`
	Street    string `json:"street_address"`
	Apartment string `json:"apartment_floor"`
	Town      string `json:"town_city"`
	Phone     string `json:"phone_number"`
	Email     string `json:"email"`
}
type Shopping_cart struct {
	User   string `json:"user_id"`
	CartID string `json:"cart_id"`
}

type Shopping_cart_item struct {
	CartID    string  `json:"cart_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Sellers struct {
	Seller_id string `json:"seller_id"`
	Name      string `json:"name"`
	Contact   string `json:"contact_info"`
}

type Categories struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Contact struct {
	ContactID string `json:"contact_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone_number"`
	Message   string `json:"message"`
}

type EcommerceDatabase interface {
	GetProduct(ctx context.Context, id string) (ProductItem, error)
	AddProduct(ctx context.Context, product NewProduct, count int) (Product, error)
	UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error)
	DeleteProduct(ctx context.Context, id string) error
	// I don't know
	GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error)
	// productImage
	GetProductImages(ctx context.Context, productID string) ([]ProductImage, error)
	AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error)
	UpdateProductImage(ctx context.Context, productID, imageID string, update UpdateProductImage) (ProductImage, error)
	DeleteProductImage(ctx context.Context, productID, imageID string) error
	// check
	Close() error
	Ping() error
	Reconnect(connStr string) error
	// user_local
	GetUserLocal(ctx context.Context) ([]User_local, error)
	AddUserLocal(ctx context.Context, name string, email string, password string) (User_local, Shopping_cart, error)
	// shipping_addresses
	GetShippingAddress(ctx context.Context) ([]Shipping_addresses, error)
	UpdateShippingAddress(ctx context.Context, user string, name string, company string, street string, apartment string, town string, phone string, email string) (Shipping_addresses, error)
	// shopping_cart
	GetShoppingCart(ctx context.Context) ([]Shopping_cart, error)
	// shopping_cart_item
	GetAllShoppingCartItem(ctx context.Context) ([]Shopping_cart_item, error)
	GetShoppingCartItem(ctx context.Context, cartID string) ([]Shopping_cart_item, error)
	DeleteShoppingCartItem(ctx context.Context, ProductID string,cartID string) ([]Shopping_cart_item, error)
	UpdateShoppingCartItem(ctx context.Context, cartID string, productID string, quantity int) ([]Shopping_cart_item, error)
	AddCartItem(ctx context.Context, cartID string, productID string, quantity int) (Shopping_cart_item, error)
	// sellers
	GetSellers(ctx context.Context) ([]Sellers, error)
	AddSellers(ctx context.Context, name string, contact string) (Sellers, error)
	UpdateSellerName(ctx context.Context, id string, name string) (Sellers, error)
	UpdateSellerContact(ctx context.Context, id string, contact string) (Sellers, error)
	DeleteSeller(ctx context.Context, id string) error
	// categories
	GetCategory(ctx context.Context) ([]Categories, error)
	AddCategory(ctx context.Context, name string, description string) (Categories, error)
	DeleteCategory(ctx context.Context, ID int) error
	UpdateCategoryName(ctx context.Context, ID int, name string) (Categories, error)
	UpdateCategoryDescription(ctx context.Context, ID int, description string) (Categories, error)
	// contact
	GetContact(ctx context.Context) ([]Contact, error)
	AddContact(ctx context.Context, name string, email string, phone string, message string) (Contact, error)
	// ----------------------------------------------------------------------------------------------------------
	//
	Search(ctx context.Context, word string) ([]ProductItem, error)
	RecommendProduct(ctx context.Context) ([]ProductItem, error)
	GetProductsByCategory(ctx context.Context, categoryID string) ([]ProductItem, error)
	GetProductWithSellers(ctx context.Context, IDSellers string) ([]ProductItem, error)
	GetLatestProductsBySeller(ctx context.Context) ([]ProductItem, error)
}

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(connStr string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &PostgresDatabase{db: db}, nil
}

func (pdb *PostgresDatabase) AddCartItem(ctx context.Context, cartID string, productID string, quantity int) (Shopping_cart_item, error) {
	// First, check if the product exists
	var productPrice float64
	err := pdb.db.QueryRowContext(ctx, `SELECT price FROM products WHERE product_id = $1`, productID).Scan(&productPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return Shopping_cart_item{}, fmt.Errorf("product with ID %s not found", productID)
		}
		return Shopping_cart_item{}, fmt.Errorf("failed to query product price: %v", err)
	}

	// If item does not exist, insert a new item into the cart
	_, err = pdb.db.ExecContext(ctx, `INSERT INTO shopping_cart_items (cart_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`, cartID, productID, quantity, productPrice)
	if err != nil {
		return Shopping_cart_item{}, fmt.Errorf("failed to add item to cartID %s: %v", cartID, err)
	}

	// Return the newly added item
	return Shopping_cart_item{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     productPrice,
	}, nil
}



func (pdb *PostgresDatabase) GetLatestProductsBySeller(ctx context.Context) ([]ProductItem, error) {
	var products []ProductItem
	query := `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, 
		       p.sku, p.price, p.status, p.seller_id, p.product_type, 
		       p.created_at, p.updated_at, c.category_id, c.name AS category_name,
		       i.quantity AS inventory_quantity, pi.image_id, pi.image_url, pi.is_primary, pi.sort_order
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		LEFT JOIN inventory i ON p.product_id = i.product_id
		LEFT JOIN product_images pi ON p.product_id = pi.product_id
		WHERE p.created_at IN (
			SELECT MAX(p2.created_at)
			FROM products p2
			GROUP BY p2.seller_id
		)
	`

	rows, err := pdb.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest products by seller: %v", err)
	}
	defer rows.Close()

	productMap := make(map[string]*ProductItem)

	for rows.Next() {
		var (
			productID, name, description, brand, modelNumber, sku, status, sellerID, productType string
			price                                                                                float64
			inventoryQuantity                                                                    int
			categoryID                                                                           int
			categoryName                                                                         string
			imageID                                                                              sql.NullString
			imageURL                                                                             string
			isPrimary                                                                            bool
			sortOrder                                                                            int
			createdAt, updatedAt                                                                 time.Time
		)

		if err := rows.Scan(&productID, &name, &description, &brand, &modelNumber,
			&sku, &price, &status, &sellerID, &productType,
			&createdAt, &updatedAt,
			&categoryID, &categoryName,
			&inventoryQuantity,
			&imageID, &imageURL, &isPrimary, &sortOrder); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}

		// ตรวจสอบว่ามี product อยู่ใน map หรือยัง ถ้าไม่มีให้สร้างใหม่
		if _, exists := productMap[productID]; !exists {
			productMap[productID] = &ProductItem{
				Product: Product{
					ID:          productID,
					Name:        name,
					Description: description,
					Brand:       brand,
					ModelNumber: modelNumber,
					SKU:         sku,
					Price:       price,
					Status:      status,
					SellerID:    sellerID,
					ProductType: productType,
					CategoryID:  categoryID,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				},
				Categories: []Category{
					{
						ID:   categoryID,
						Name: categoryName,
					},
				},
				Inventory: Inventory{
					Quantity: inventoryQuantity,
				},
				Images: []ProductImage{}, // เริ่มต้นด้วย slice ว่าง
				Options: []ProductOption{}, // เติมข้อมูล options หากต้องการ
			}
		}

		// เพิ่มข้อมูลภาพสินค้าใน ProductItem
		if imageID.Valid {
			productMap[productID].Images = append(productMap[productID].Images, ProductImage{
				ID:        imageID.String,
				ProductID: productID,
				ImageURL:  imageURL,
				IsPrimary: isPrimary,
				SortOrder: sortOrder,
				CreatedAt: createdAt,
			})
		}
	}

	// แปลงจาก map เป็น slice
	for _, product := range productMap {
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetProductWithSellers(ctx context.Context, IDSellers string) ([]ProductItem, error) {
	var products []ProductItem
	// คำสั่ง SQL สำหรับดึงข้อมูลสินค้าที่เชื่อมโยงกับชื่อร้าน
	query := `
        SELECT 
            p.product_id, p.name, p.description, p.brand, p.model_number, 
            p.sku, p.price, p.status, p.seller_id, p.product_type, p.category_id, 
            p.created_at, p.updated_at,
            c.category_id, c.name AS category_name, c.description AS category_description,
            i.quantity AS inventory_quantity,
            pi.image_id, pi.image_url, pi.is_primary, pi.sort_order,
            s.seller_id, s.name AS seller_name, s.contact_info
        FROM 
            products p
        INNER JOIN 
            sellers s ON p.seller_id = s.seller_id
        LEFT JOIN 
            categories c ON p.category_id = c.category_id
        LEFT JOIN 
            inventory i ON p.product_id = i.product_id
        LEFT JOIN 
            product_images pi ON p.product_id = pi.product_id
        WHERE 
            s.seller_id = $1
        ORDER BY 
            p.created_at DESC
    `

	// ดำเนินการ query ข้อมูล
	rows, err := pdb.db.QueryContext(ctx, query, IDSellers)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products with seller name %v: %v", IDSellers, err)
	}
	defer rows.Close()

	// Map สำหรับเก็บข้อมูลสินค้าที่จัดรูปแบบแล้ว
	productMap := make(map[string]*ProductItem)

	// อ่านข้อมูลแต่ละ row
	for rows.Next() {
		var (
			productID, name, description, brand, modelNumber, sku, status, sellerID, productType, categoryName, categoryDescription string
			price                                                                                                                   float64
			categoryID, inventoryQuantity, sortOrder                                                                                int
			imageID                                                                                                                 sql.NullString
			imageURL                                                                                                                string
			isPrimary                                                                                                               bool
			createdAt, updatedAt                                                                                                    time.Time
			sellerName, contactInfo                                                                                                 string
		)

		// อ่านค่าในแต่ละ row
		err := rows.Scan(
			&productID, &name, &description, &brand, &modelNumber, &sku, &price, &status, &sellerID, &productType,
			&categoryID, &createdAt, &updatedAt,
			&categoryID, &categoryName, &categoryDescription,
			&inventoryQuantity,
			&imageID, &imageURL, &isPrimary, &sortOrder,
			&sellerID, &sellerName, &contactInfo,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		// ตรวจสอบว่ามี product อยู่ใน map หรือยัง ถ้าไม่มีให้สร้างใหม่
		if _, exists := productMap[productID]; !exists {
			productMap[productID] = &ProductItem{
				Product: Product{
					ID:          productID,
					Name:        name,
					Description: description,
					Brand:       brand,
					ModelNumber: modelNumber,
					SKU:         sku,
					Price:       price,
					Status:      status,
					SellerID:    sellerID,
					ProductType: productType,
					CategoryID:  categoryID,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				},
				Categories: []Category{ // เปลี่ยนให้เป็น slice ว่างแทน
					{
						ID:   categoryID,
						Name: categoryName,
					},
				},
				Inventory: Inventory{
					Quantity: inventoryQuantity,
				},
				Images: []ProductImage{
					{
						ID:        imageID.String,
						ProductID: productID,
						ImageURL:  imageURL,
						IsPrimary: isPrimary,
						SortOrder: sortOrder,
						CreatedAt: createdAt,
					},
				},
				Options: []ProductOption{},
			}
			products = append(products, *productMap[productID])
		}

		// เพิ่มข้อมูลภาพสินค้าใน ProductItem
		if imageID.Valid {
			productMap[productID].Images = append(productMap[productID].Images, ProductImage{
				ID:        imageID.String,
				ProductID: productID,
				ImageURL:  imageURL,
				IsPrimary: isPrimary,
				SortOrder: sortOrder,
			})
		}
	}

	// ตรวจสอบ error หลังจากวนลูปเสร็จ
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to process rows: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetProductsByCategory(ctx context.Context, categoryID string) ([]ProductItem, error) {
	var products []ProductItem
	query := `
		SELECT 
			p.product_id, p.name, p.description, p.brand, p.model_number, 
			p.sku, p.price, p.status, p.seller_id, p.product_type, 
			p.created_at, p.updated_at,
			c.category_id, c.name AS category_name,
			i.quantity AS inventory_quantity,
			pi.image_id, pi.image_url, pi.is_primary, pi.sort_order
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		LEFT JOIN inventory i ON p.product_id = i.product_id
		LEFT JOIN product_images pi ON p.product_id = pi.product_id
		WHERE p.category_id = $1
	`

	rows, err := pdb.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category: %v", err)
	}
	defer rows.Close()

	// ใช้ map เพื่อเก็บข้อมูลสินค้าตาม product ID
	productMap := make(map[string]*ProductItem)

	for rows.Next() {
		var (
			productID, name, description, brand, modelNumber, sku, status, sellerID, productType string
			price                                                                                float64
			inventoryQuantity                                                                    int
			categoryID                                                                           int
			categoryName                                                                         string
			imageID                                                                              sql.NullString
			imageURL                                                                             string
			isPrimary                                                                            bool
			sortOrder                                                                            int
			createdAt, updatedAt                                                                 time.Time
		)

		if err := rows.Scan(&productID, &name, &description, &brand, &modelNumber,
			&sku, &price, &status, &sellerID, &productType,
			&createdAt, &updatedAt,
			&categoryID, &categoryName,
			&inventoryQuantity,
			&imageID, &imageURL, &isPrimary, &sortOrder); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}

		// ตรวจสอบว่ามี product อยู่ใน map หรือยัง ถ้าไม่มีให้สร้างใหม่
		if _, exists := productMap[productID]; !exists {
			productMap[productID] = &ProductItem{
				Product: Product{
					ID:          productID,
					Name:        name,
					Description: description,
					Brand:       brand,
					ModelNumber: modelNumber,
					SKU:         sku,
					Price:       price,
					Status:      status,
					SellerID:    sellerID,
					ProductType: productType,
					CategoryID:  categoryID,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				},
				Categories: []Category{
					{
						ID:   categoryID,
						Name: categoryName,
					},
				},
				Inventory: Inventory{
					Quantity: inventoryQuantity,
				},
				Images: []ProductImage{{
					ID:        imageID.String,
					ProductID: productID,
					ImageURL:  imageURL,
					IsPrimary: isPrimary,
					SortOrder: sortOrder,
					CreatedAt: createdAt,
				}},
				Options: []ProductOption{}, // เติมข้อมูล options หากต้องการ
			}
		}

	}

	// แปลงจาก map เป็น slice
	for _, product := range productMap {
		products = append(products, *product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) RecommendProduct(ctx context.Context) ([]ProductItem, error) {
	var products []ProductItem
	query := `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
		       p.status, p.seller_id, p.product_type, p.created_at, p.updated_at
		FROM products p
		ORDER BY RANDOM()
		LIMIT 3
	`
	rows, err := pdb.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get random products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductItem
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Status,
			&product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) Search(ctx context.Context, word string) ([]ProductItem, error) {
	var products []ProductItem
	var category Category

	// ปรับคำสั่ง SQL ให้ใช้ ? สำหรับ parameterized query
	query := `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
		       p.status, p.seller_id, p.product_type, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.name ILIKE '%' || $1 || '%'` // ใช้ ILIKE เพื่อค้นหาที่ไม่สนใจตัวพิมพ์

	// ค้นหาผลิตภัณฑ์
	rows, err := pdb.db.QueryContext(ctx, query, "%"+word+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %v", err)
	}
	defer rows.Close()

	// ดึงข้อมูลทุกแถวที่คืนค่ากลับมา
	for rows.Next() {
		var product ProductItem
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Status,
			&product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		product.Categories = []Category{category}

		// ดึงข้อมูล inventory
		err = pdb.db.QueryRowContext(ctx, `
				SELECT quantity, updated_at
				FROM inventory
				WHERE product_id = $1
			`, product.ID).Scan(&product.Inventory.Quantity, &product.Inventory.UpdatedAt)

		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get inventory: %v", err)
		}

		// ดึงข้อมูลรูปภาพ
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetProduct(ctx context.Context, id string) (ProductItem, error) {
	var product ProductItem
	var category Category

	// ดึงข้อมูลหลักของสินค้าและหมวดหมู่
	err := pdb.db.QueryRowContext(ctx, `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
		       p.status, p.seller_id, p.product_type, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.product_id = $1
	`, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Brand,
		&product.ModelNumber, &product.SKU, &product.Price, &product.Status,
		&product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt,
		&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return ProductItem{}, fmt.Errorf("product not found")
		}
		return ProductItem{}, fmt.Errorf("failed to get product: %v", err)
	}

	product.Categories = []Category{category}

	// ดึงข้อมูล inventory
	err = pdb.db.QueryRowContext(ctx, `
		SELECT quantity, updated_at
		FROM inventory
		WHERE product_id = $1
	`, id).Scan(&product.Inventory.Quantity, &product.Inventory.UpdatedAt)

	if err != nil && err != sql.ErrNoRows {
		return ProductItem{}, fmt.Errorf("failed to get inventory: %v", err)
	}

	// ดึงข้อมูลรูปภาพ
	product.Images, err = pdb.GetProductImages(ctx, id)
	if err != nil {
		return ProductItem{}, fmt.Errorf("failed to get product images: %v", err)
	}

	// ดึงข้อมูลตัวเลือก
	product.Options, err = pdb.getProductOptions(ctx, id)
	if err != nil {
		return ProductItem{}, fmt.Errorf("failed to get product options: %v", err)
	}

	return product, nil
}

func (pdb *PostgresDatabase) AddProduct(ctx context.Context, product NewProduct, count int) (Product, error) {
	var createdProduct Product
	err := pdb.db.QueryRowContext(ctx, `
		INSERT INTO products (name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING product_id, name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id, created_at, updated_at
	`,
		product.Name, product.Description, product.Brand, product.ModelNumber, product.SKU, product.Price, product.Status, product.SellerID, product.ProductType, product.CategoryID,
	).Scan(
		&createdProduct.ID, &createdProduct.Name, &createdProduct.Description, &createdProduct.Brand,
		&createdProduct.ModelNumber, &createdProduct.SKU, &createdProduct.Price, &createdProduct.Status,
		&createdProduct.SellerID, &createdProduct.ProductType, &createdProduct.CategoryID,
		&createdProduct.CreatedAt, &createdProduct.UpdatedAt)
	if err != nil {
		return Product{}, fmt.Errorf("failed to add product: %v", err)
	}
	// เพิ่มข้อมูลใน inventory
	_, erro := pdb.db.ExecContext(ctx, `INSERT INTO inventory (product_id, quantity) VALUES ($1, $2)`, createdProduct.ID, count)
	if erro != nil {
		return Product{}, fmt.Errorf("failed to add Inventory: %v", erro)
	}

	return createdProduct, nil
}

func (pdb *PostgresDatabase) UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error) {
	var updatedProduct Product
	err := pdb.db.QueryRowContext(ctx, `
		UPDATE products 
		SET price = $1, status = $2, updated_at = NOW() 
		WHERE product_id = $3 
		RETURNING product_id, name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id, created_at, updated_at
	`,
		update.Price, update.Status, id,
	).Scan(
		&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Description, &updatedProduct.Brand,
		&updatedProduct.ModelNumber, &updatedProduct.SKU, &updatedProduct.Price, &updatedProduct.Status,
		&updatedProduct.SellerID, &updatedProduct.ProductType, &updatedProduct.CategoryID,
		&updatedProduct.CreatedAt, &updatedProduct.UpdatedAt)
	if err != nil {
		return Product{}, fmt.Errorf("failed to update product: %v", err)
	}
	return updatedProduct, nil
}

func (pdb *PostgresDatabase) DeleteProduct(ctx context.Context, id string) error {
	result, err := pdb.db.ExecContext(ctx, "DELETE FROM products WHERE product_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func (pdb *PostgresDatabase) GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error) {
	query := `
        SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
               p.status, p.seller_id, p.product_type, p.created_at, p.updated_at,
               c.category_id, c.name as category_name,
               i.quantity, i.updated_at as inventory_updated_at
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.category_id
        LEFT JOIN inventory i ON p.product_id = i.product_id
        WHERE 1=1`

	args := []interface{}{}
	placeholderCount := 1

	// การจัดการพารามิเตอร์ cursor
	if params.Cursor != "" {
		cursor, err := decodeCursor(params.Cursor)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor: %v", err)
		}
		query += fmt.Sprintf(" AND (p.created_at, p.product_id) > ($%d, $%d)", placeholderCount, placeholderCount+1)
		args = append(args, cursor.CreatedAt, cursor.ProductID)
		placeholderCount += 2
	}

	// การจัดการพารามิเตอร์ search
	if params.Search != "" {
		query += fmt.Sprintf(" AND (p.name ILIKE $%d OR p.description ILIKE $%d)", placeholderCount, placeholderCount+1)
		args = append(args, "%"+params.Search+"%", "%"+params.Search+"%")
		placeholderCount += 2
	}

	// การจัดการพารามิเตอร์ category_id
	if params.CategoryID != 0 {
		query += fmt.Sprintf(" AND p.category_id = $%d", placeholderCount)
		args = append(args, params.CategoryID)
		placeholderCount++
	}

	// การจัดการพารามิเตอร์ seller_id
	if params.SellerID != "" {
		query += fmt.Sprintf(" AND p.seller_id = $%d", placeholderCount)
		args = append(args, params.SellerID)
		placeholderCount++
	}

	// การจัดการพารามิเตอร์ status
	if params.Status != "" {
		query += fmt.Sprintf(" AND p.status = $%d", placeholderCount)
		args = append(args, params.Status)
		placeholderCount++
	}

	// การจัดการพารามิเตอร์ product_type
	if params.ProductType != "" {
		query += fmt.Sprintf(" AND p.product_type = $%d", placeholderCount)
		args = append(args, params.ProductType)
		placeholderCount++
	}

	// การจัดการ ORDER BY ด้วย sort และ order
	sortFields := map[string]string{
		"name":       "p.name",
		"price":      "p.price",
		"created_at": "p.created_at",
	}

	orderDirections := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	sortColumn, ok := sortFields[params.Sort]
	if !ok {
		sortColumn = "p.created_at"
	}

	orderDirection, ok := orderDirections[strings.ToLower(params.Order)]
	if !ok {
		orderDirection = "ASC"
	}

	query += fmt.Sprintf(" ORDER BY %s %s, p.product_id ASC", sortColumn, orderDirection)

	// log.Printf("Query: %s", query)
	// log.Printf("Args: %v", args)

	// ดำเนินการ query และประมวลผลผลลัพธ์
	rows, err := pdb.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()

	var products []ProductItem
	for rows.Next() {
		var product ProductItem
		var category Category
		var inventory Inventory
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Status,
			&product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
			&inventory.Quantity, &inventory.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		product.Categories = []Category{category}
		product.Inventory = inventory

		// การดึงข้อมูลรูปภาพและตัวเลือกของผลิตภัณฑ์
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}

		products = append(products, product)
	}

	return &ProductResponse{Items: products, NextCursor: ""}, nil
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Cursor struct {
	CreatedAt time.Time
	ProductID string
}

func (pdb *PostgresDatabase) GetProductImages(ctx context.Context, productID string) ([]ProductImage, error) {
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT image_id, product_id, image_url, is_primary, sort_order, created_at 
		FROM product_images 
		WHERE product_id = $1 
		ORDER BY sort_order ASC
	`, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product images: %v", err)
	}
	defer rows.Close()

	var images []ProductImage
	for rows.Next() {
		var image ProductImage
		if err := rows.Scan(
			&image.ID, &image.ProductID, &image.ImageURL, &image.IsPrimary,
			&image.SortOrder, &image.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over product images: %v", err)
	}

	return images, nil
}

func (pdb *PostgresDatabase) AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error) {
	var createdImage ProductImage
	err := pdb.db.QueryRowContext(ctx, `
		INSERT INTO product_images (product_id, image_url, is_primary, sort_order) 
		VALUES ($1, $2, $3, $4) 
		RETURNING image_id, product_id, image_url, is_primary, sort_order, created_at
	`,
		productID, image.ImageURL, image.IsPrimary, image.SortOrder,
	).Scan(
		&createdImage.ID, &createdImage.ProductID, &createdImage.ImageURL,
		&createdImage.IsPrimary, &createdImage.SortOrder, &createdImage.CreatedAt)
	if err != nil {
		return ProductImage{}, fmt.Errorf("failed to add product image: %v", err)
	}
	return createdImage, nil
}

func (pdb *PostgresDatabase) UpdateProductImage(ctx context.Context, productID string, imageID string, update UpdateProductImage) (ProductImage, error) {
	var updatedImage ProductImage
	err := pdb.db.QueryRowContext(ctx, `
		UPDATE product_images 
		SET is_primary = $1, sort_order = $2, updated_at = NOW() 
		WHERE product_id = $3 AND image_id = $4 
		RETURNING image_id, product_id, image_url, is_primary, sort_order, created_at
	`,
		update.IsPrimary, update.SortOrder, productID, imageID,
	).Scan(
		&updatedImage.ID, &updatedImage.ProductID, &updatedImage.ImageURL,
		&updatedImage.IsPrimary, &updatedImage.SortOrder, &updatedImage.CreatedAt)
	if err != nil {
		return ProductImage{}, fmt.Errorf("failed to update product image: %v", err)
	}
	return updatedImage, nil
}

func (pdb *PostgresDatabase) DeleteProductImage(ctx context.Context, productID string, imageID string) error {
	result, err := pdb.db.ExecContext(ctx, `
		DELETE FROM product_images 
		WHERE product_id = $1 AND image_id = $2
	`, productID, imageID)
	if err != nil {
		return fmt.Errorf("failed to delete product image: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product image not found")
	}

	return nil
}

func (pdb *PostgresDatabase) Close() error {
	return pdb.db.Close()
}

func (pdb *PostgresDatabase) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return pdb.db.PingContext(ctx)
}

func (pdb *PostgresDatabase) getProductOptions(ctx context.Context, productID string) ([]ProductOption, error) {
	rows, err := pdb.db.QueryContext(ctx, `
        SELECT option_id, name, values
        FROM product_options
        WHERE product_id = $1
    `, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product options: %v", err)
	}
	defer rows.Close()

	var options []ProductOption
	for rows.Next() {
		var option ProductOption
		if err := rows.Scan(&option.ID, &option.Name, &option.Values); err != nil {
			return nil, fmt.Errorf("failed to scan product option: %v", err)
		}
		options = append(options, option)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over product options: %v", err)
	}

	return options, nil
}

func encodeCursor(c Cursor) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s,%s", c.CreatedAt.Format(time.RFC3339Nano), c.ProductID)))
}

func decodeCursor(s string) (Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, err
	}
	parts := strings.Split(string(b), ",")
	if len(parts) != 2 {
		return Cursor{}, fmt.Errorf("invalid cursor format")
	}
	createdAt, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return Cursor{}, err
	}
	return Cursor{CreatedAt: createdAt, ProductID: parts[1]}, nil
}

func (pdb *PostgresDatabase) Reconnect(connStr string) error {
	if pdb.db != nil {
		pdb.db.Close()
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// ตั้งค่า connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	pdb.db = db
	return nil
}

type Store struct {
	db EcommerceDatabase
}

func NewStore(db EcommerceDatabase) *Store {
	return &Store{db: db}
}

func (s *Store) GetProduct(ctx context.Context, id string) (ProductItem, error) {
	return s.db.GetProduct(ctx, id)
}

func (s *Store) AddProduct(ctx context.Context, product NewProduct, count int) (Product, error) {
	return s.db.AddProduct(ctx, product, count)
}

func (s *Store) UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error) {
	return s.db.UpdateProduct(ctx, id, update)
}

func (s *Store) DeleteProduct(ctx context.Context, id string) error {
	return s.db.DeleteProduct(ctx, id)
}

func (s *Store) GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error) {
	return s.db.GetProducts(ctx, params)
}

func (s *Store) GetProductImages(ctx context.Context, productID string) ([]ProductImage, error) {
	return s.db.GetProductImages(ctx, productID)
}

func (s *Store) AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error) {
	return s.db.AddProductImage(ctx, productID, image)
}

func (s *Store) UpdateProductImage(ctx context.Context, productID string, imageID string, update UpdateProductImage) (ProductImage, error) {
	return s.db.UpdateProductImage(ctx, productID, imageID, update)
}

func (s *Store) DeleteProductImage(ctx context.Context, productID string, imageID string) error {
	return s.db.DeleteProductImage(ctx, productID, imageID)
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Ping() error {
	if s.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return s.db.Ping()
}

func (s *Store) Reconnect(connStr string) error {
	return s.db.Reconnect(connStr)
}

// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetUserLocal(ctx context.Context) ([]User_local, error) {

	// Execute the query to retrieve user data
	rows, err := pdb.db.QueryContext(ctx, `SELECT user_id, email, password, username FROM users_local`)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close() // Ensure rows are closed after processing

	// Iterate through the rows
	var users []User_local
	for rows.Next() {
		var User User_local
		if err := rows.Scan(&User.UserID, &User.Email, &User.Password, &User.Username); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}
		users = append(users, User)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return users, nil // Return the slice of users
}

func (pdb *PostgresDatabase) AddUserLocal(ctx context.Context, name string, email string, password string) (User_local, Shopping_cart, error) {
	var user User_local
	var cart Shopping_cart
	var address Shipping_addresses

	// Prepare the SQL statement for inserting a new user and returning user_id
	query := `INSERT INTO users_local (username, email, password) VALUES ($1, $2, $3) RETURNING user_id`
	err := pdb.db.QueryRowContext(ctx, query, name, email, password).Scan(&user.UserID)
	if err != nil {
		return user, cart, fmt.Errorf("failed to add user: %v", err)
	}

	// Initialize a shopping cart for the new user and return cart_id
	cartQuery := `INSERT INTO shopping_cart (user_id) VALUES ($1) RETURNING cart_id`
	err = pdb.db.QueryRowContext(ctx, cartQuery, user.UserID).Scan(&cart.CartID)
	if err != nil {
		return user, cart, fmt.Errorf("failed to create shopping cart: %v, user_id: %s", err, user.UserID)
	}

	ShippingAddressesQuery := `INSERT INTO shipping_addresses (user_id) VALUES ($1) RETURNING address_id`
	err = pdb.db.QueryRowContext(ctx, ShippingAddressesQuery, user.UserID).Scan(&address.User)
	if err != nil {
		return user, cart, fmt.Errorf("failed to create shipping address: %v, user_id: %s", err, user.UserID)
	}
	// Set additional user information
	user.Email = email
	user.Password = password
	user.Username = name // Ensure your struct has a field for name
	cart.User = user.UserID

	return user, cart, nil // Return the user and shopping cart
}

// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetShippingAddress(ctx context.Context) ([]Shipping_addresses, error) {
	var addresses []Shipping_addresses

	// Execute the query to retrieve shipping addresses
	rows, err := pdb.db.QueryContext(ctx, `SELECT user_id, first_name, company_name, street_address, apartment_floor, town_city, phone_number, email FROM shipping_addresses`)
	if err != nil {
		return nil, fmt.Errorf("failed to query shipping addresses: %v", err)
	}
	defer rows.Close() // Ensure rows are closed after processing

	// Iterate through the rows
	for rows.Next() {
		var address Shipping_addresses
		if err := rows.Scan(&address.User, &address.Name, &address.Company, &address.Street, &address.Apartment, &address.Town, &address.Phone, &address.Email); err != nil {
			return nil, fmt.Errorf("failed to scan shipping address row: %v", err)
		}
		addresses = append(addresses, address)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return addresses, nil // Return the slice of shipping addresses
}
func (pdb *PostgresDatabase) UpdateShippingAddress(ctx context.Context, userID string, name string, company string, street string, apartment string, town string, phone string, email string) (Shipping_addresses, error) {
	var updatedAddress Shipping_addresses

	// Prepare the SQL statement for updating the shipping address
	query := `UPDATE shipping_addresses 
			  SET first_name = $1, 
				  company_name = $2, 
				  street_address = $3, 
				  apartment_floor = $4, 
				  town_city = $5, 
				  phone_number = $6, 
				  email = $7 
			  WHERE user_id = $8 
			  RETURNING user_id, first_name, company_name, street_address, apartment_floor, town_city, phone_number, email`

	// Execute the update query
	err := pdb.db.QueryRowContext(ctx, query, name, company, street, apartment, town, phone, email, userID).
		Scan(&updatedAddress.User, &updatedAddress.Name, &updatedAddress.Company, &updatedAddress.Street, &updatedAddress.Apartment, &updatedAddress.Town, &updatedAddress.Phone, &updatedAddress.Email)
	if err != nil {
		return updatedAddress, fmt.Errorf("failed to update shipping address: %v", err)
	}

	return updatedAddress, nil // Return the updated address
}

// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetShoppingCart(ctx context.Context) ([]Shopping_cart, error) {
	var carts []Shopping_cart

	// Execute the query to retrieve all shopping cart items
	rows, err := pdb.db.QueryContext(ctx, `SELECT user_id, cart_id FROM shopping_cart`)
	if err != nil {
		return nil, fmt.Errorf("failed to query shopping cart: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var cart Shopping_cart
		if err := rows.Scan(&cart.User, &cart.CartID); err != nil {
			return nil, fmt.Errorf("failed to scan shopping cart row: %v", err)
		}
		carts = append(carts, cart)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return carts, nil
}

// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetAllShoppingCartItem(ctx context.Context) ([]Shopping_cart_item, error) {
	var items []Shopping_cart_item

	// Query to retrieve all shopping cart items
	rows, err := pdb.db.QueryContext(ctx, `SELECT cart_id, product_id, quantity, price FROM shopping_cart_items`)
	if err != nil {
		return nil, fmt.Errorf("failed to query shopping cart items: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var item Shopping_cart_item
		if err := rows.Scan(&item.CartID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan shopping cart item row: %v", err)
		}
		items = append(items, item)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return items, nil
}

func (pdb *PostgresDatabase) GetShoppingCartItem(ctx context.Context, cartID string) ([]Shopping_cart_item, error) {
	var items []Shopping_cart_item

	// Query to retrieve items for a specific cartID
	rows, err := pdb.db.QueryContext(ctx, `SELECT cart_id, product_id, quantity, price FROM shopping_cart_items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to query shopping cart items for cartID %s: %v", cartID, err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var item Shopping_cart_item
		if err := rows.Scan(&item.CartID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan shopping cart item row: %v", err)
		}
		items = append(items, item)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return items, nil
}
func (pdb *PostgresDatabase) DeleteShoppingCartItem(ctx context.Context, productID string, cartID string) ([]Shopping_cart_item, error) {
	// First, delete the item from the shopping cart based on the productID and cartID
	_, err := pdb.db.ExecContext(ctx, `DELETE FROM shopping_cart_items WHERE product_id = $1 AND cart_id = $2`, productID, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete shopping cart item with productID %s: %v", productID, err)
	}

	// Query to retrieve the remaining items in the shopping cart
	var items []Shopping_cart_item
	rows, err := pdb.db.QueryContext(ctx, `SELECT cart_id, product_id, quantity, price FROM shopping_cart_items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to query remaining shopping cart items: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows to populate the remaining items
	for rows.Next() {
		var item Shopping_cart_item
		if err := rows.Scan(&item.CartID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan shopping cart item row: %v", err)
		}
		items = append(items, item)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return items, nil
}


func (pdb *PostgresDatabase) UpdateShoppingCartItem(ctx context.Context, cartID string, productID string, quantity int) ([]Shopping_cart_item, error) {
	// Update the quantity of the specified item in the shopping cart
	_, err := pdb.db.ExecContext(ctx, `UPDATE shopping_cart_items SET quantity = $1 WHERE cart_id = $2 AND product_id = $3`, quantity, cartID, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to update shopping cart item with cartID %s and productID %s: %v", cartID, productID, err)
	}

	// Retrieve the updated list of items in the cart
	var items []Shopping_cart_item
	rows, err := pdb.db.QueryContext(ctx, `SELECT cart_id, product_id, quantity, price FROM shopping_cart_items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to query updated shopping cart items: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows to populate the list of items
	for rows.Next() {
		var item Shopping_cart_item
		if err := rows.Scan(&item.CartID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan shopping cart item row: %v", err)
		}
		items = append(items, item)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return items, nil
}


// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetSellers(ctx context.Context) ([]Sellers, error) {
	var sellersList []Sellers

	rows, err := pdb.db.QueryContext(ctx, `SELECT seller_id, name, contact_info FROM sellers`)
	if err != nil {
		return nil, fmt.Errorf("failed to query sellers: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var seller Sellers
		if err := rows.Scan(&seller.Seller_id, &seller.Name, &seller.Contact); err != nil {
			return nil, fmt.Errorf("failed to scan seller row: %v", err)
		}
		sellersList = append(sellersList, seller)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return sellersList, nil
}
func (pdb *PostgresDatabase) AddSellers(ctx context.Context, name string, contact string) (Sellers, error) {
	var newSeller Sellers

	// สร้างคำสั่ง SQL สำหรับเพิ่มผู้ขาย
	query := `INSERT INTO sellers (name, contact_info) VALUES ($1, $2) RETURNING seller_id, name, contact_info`
	err := pdb.db.QueryRowContext(ctx, query, name, contact).Scan(&newSeller.Seller_id, &newSeller.Name, &newSeller.Contact)
	if err != nil {
		return newSeller, fmt.Errorf("failed to add seller: %v", err)
	}

	return newSeller, nil
}

func (pdb *PostgresDatabase) UpdateSellerName(ctx context.Context, id string, name string) (Sellers, error) {
	var updatedSeller Sellers

	query := `UPDATE sellers SET name = $1 WHERE seller_id = $2 RETURNING seller_id, name, contact_info`
	err := pdb.db.QueryRowContext(ctx, query, name, id).Scan(&updatedSeller.Seller_id, &updatedSeller.Name, &updatedSeller.Contact)
	if err != nil {
		return updatedSeller, fmt.Errorf("failed to update seller name: %v", err)
	}

	return updatedSeller, nil
}
func (pdb *PostgresDatabase) UpdateSellerContact(ctx context.Context, id string, contact string) (Sellers, error) {
	var updatedSeller Sellers

	query := `UPDATE sellers SET contact_info = $1 WHERE seller_id = $2 RETURNING seller_id, name, contact_info`
	err := pdb.db.QueryRowContext(ctx, query, contact, id).Scan(&updatedSeller.Seller_id, &updatedSeller.Name, &updatedSeller.Contact)
	if err != nil {
		return updatedSeller, fmt.Errorf("failed to update seller contact: %v", err)
	}

	return updatedSeller, nil
}
func (pdb *PostgresDatabase) DeleteSeller(ctx context.Context, id string) error {
	query := `DELETE FROM sellers WHERE seller_id = $1`
	_, err := pdb.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete seller: %v", err)
	}
	return nil
}

// ---------------------------------------------------------------------------------------------
// Retrieve all categories
func (pdb *PostgresDatabase) GetCategory(ctx context.Context) ([]Categories, error) {
	rows, err := pdb.db.QueryContext(ctx, `SELECT category_id, name, description FROM categories`)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %v", err)
	}
	defer rows.Close()

	var categoriesList []Categories
	for rows.Next() {
		var category Categories
		if err := rows.Scan(&category.CategoryID, &category.Name, &category.Description); err != nil {
			return nil, fmt.Errorf("failed to scan category row: %v", err)
		}
		categoriesList = append(categoriesList, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return categoriesList, nil
}

// Add a new category
func (pdb *PostgresDatabase) AddCategory(ctx context.Context, name string, description string) (Categories, error) {
	var category Categories

	query := `INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING category_id`
	err := pdb.db.QueryRowContext(ctx, query, name, description).Scan(&category.CategoryID)
	if err != nil {
		return category, fmt.Errorf("failed to add category: %v", err)
	}

	category.Name = name
	category.Description = description

	return category, nil
}

// Delete a category by ID
func (pdb *PostgresDatabase) DeleteCategory(ctx context.Context, ID int) error {
	_, err := pdb.db.ExecContext(ctx, `DELETE FROM categories WHERE category_id = $1`, ID)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}
	return nil
}

// Update a category's name
func (pdb *PostgresDatabase) UpdateCategoryName(ctx context.Context, ID int, name string) (Categories, error) {
	var category Categories

	query := `UPDATE categories SET name = $1 WHERE category_id = $2 RETURNING category_id, description`
	err := pdb.db.QueryRowContext(ctx, query, name, ID).Scan(&category.CategoryID, &category.Description)
	if err != nil {
		return category, fmt.Errorf("failed to update category name: %v", err)
	}

	category.Name = name

	return category, nil
}

// Update a category's description
func (pdb *PostgresDatabase) UpdateCategoryDescription(ctx context.Context, ID int, description string) (Categories, error) {
	var category Categories

	query := `UPDATE categories SET description = $1 WHERE category_id = $2`
	err := pdb.db.QueryRowContext(ctx, query, description, ID).Scan(&category.CategoryID, &category.Name)
	if err != nil {
		return category, fmt.Errorf("failed to update category description: %v", err)
	}
	category.Description = description

	return category, nil
}

// ---------------------------------------------------------------------------------------------
func (pdb *PostgresDatabase) GetContact(ctx context.Context) ([]Contact, error) {
	var contacts []Contact

	// Execute the query to retrieve contact data
	rows, err := pdb.db.QueryContext(ctx, `SELECT contract_id, name, email, phone_number, message FROM contract`)
	if err != nil {
		return nil, fmt.Errorf("failed to query contacts: %v", err)
	}
	defer rows.Close() // Ensure rows are closed after processing

	// Iterate through the rows
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ContactID, &c.Name, &c.Email, &c.Phone, &c.Message); err != nil {
			return nil, fmt.Errorf("failed to scan contact row: %v", err)
		}
		contacts = append(contacts, c)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return contacts, nil // Return the slice of contacts
}
func (pdb *PostgresDatabase) AddContact(ctx context.Context, name string, email string, phone string, message string) (Contact, error) {
	var c Contact

	// Prepare the SQL statement for inserting a new contact
	query := `INSERT INTO contract (name, email, phone_number, message) VALUES ($1, $2, $3, $4) RETURNING contract_id`
	err := pdb.db.QueryRowContext(ctx, query, name, email, phone, message).Scan(&c.ContactID)
	if err != nil {
		return c, fmt.Errorf("failed to add contact: %v", err)
	}

	// Set additional contact information
	c.Name = name
	c.Email = email
	c.Phone = phone
	c.Message = message

	return c, nil // Return the newly created contact
}

func (s *Store) GetUserLocal(ctx context.Context) ([]User_local, error) {
	return s.db.GetUserLocal(ctx)
}

func (s *Store) AddUserLocal(ctx context.Context, name string, email string, password string) (User_local, Shopping_cart, error) {
	return s.db.AddUserLocal(ctx, name, email, password)
}

func (s *Store) GetShippingAddress(ctx context.Context) ([]Shipping_addresses, error) {
	return s.db.GetShippingAddress(ctx)
}

func (s *Store) UpdateShippingAddress(ctx context.Context, userID string, name string, company string, street string, apartment string, town string, phone string, email string) (Shipping_addresses, error) {
	return s.db.UpdateShippingAddress(ctx, userID, name, company, street, apartment, town, phone, email)
}

func (s *Store) GetShoppingCart(ctx context.Context) ([]Shopping_cart, error) {
	return s.db.GetShoppingCart(ctx)
}
func (s *Store) GetAllShoppingCartItem(ctx context.Context) ([]Shopping_cart_item, error) {
	return s.db.GetAllShoppingCartItem(ctx)
}

func (s *Store) GetShoppingCartItem(ctx context.Context, cartID string) ([]Shopping_cart_item, error) {
	return s.db.GetShoppingCartItem(ctx, cartID)
}

func (s *Store) DeleteShoppingCartItem(ctx context.Context, productID string,cartID string) ([]Shopping_cart_item, error) {
	return s.db.DeleteShoppingCartItem(ctx, productID,cartID)
}

func (s *Store) UpdateShoppingCartItem(ctx context.Context, cartID string, productID string, quantity int) ([]Shopping_cart_item, error) {
	return s.db.UpdateShoppingCartItem(ctx, cartID,productID, quantity)
}

func (s *Store) GetSellers(ctx context.Context) ([]Sellers, error) {
	return s.db.GetSellers(ctx)
}

func (s *Store) AddSellers(ctx context.Context, name string, contact string) (Sellers, error) {
	return s.db.AddSellers(ctx, name, contact)
}

func (s *Store) UpdateSellerName(ctx context.Context, id string, name string) (Sellers, error) {
	return s.db.UpdateSellerName(ctx, id, name)
}

func (s *Store) DeleteSeller(ctx context.Context, id string) error {
	return s.db.DeleteSeller(ctx, id)
}

func (s *Store) GetCategory(ctx context.Context) ([]Categories, error) {
	return s.db.GetCategory(ctx)
}

func (s *Store) AddCategory(ctx context.Context, name string, description string) (Categories, error) {
	return s.db.AddCategory(ctx, name, description)
}

func (s *Store) DeleteCategory(ctx context.Context, ID int) error {
	return s.db.DeleteCategory(ctx, ID)
}

func (s *Store) UpdateCategoryName(ctx context.Context, ID int, name string) (Categories, error) {
	return s.db.UpdateCategoryName(ctx, ID, name)
}

func (s *Store) UpdateCategoryDescription(ctx context.Context, ID int, description string) (Categories, error) {
	return s.db.UpdateCategoryDescription(ctx, ID, description)
}

func (s *Store) GetContact(ctx context.Context) ([]Contact, error) {
	return s.db.GetContact(ctx)
}

func (s *Store) AddContact(ctx context.Context, name string, email string, phone string, message string) (Contact, error) {
	return s.db.AddContact(ctx, name, email, phone, message)
}

func (s *Store) Search(ctx context.Context, word string) ([]ProductItem, error) {
	return s.db.Search(ctx, word)
}

func (s *Store) RecommendProduct(ctx context.Context) ([]ProductItem, error) {
	return s.db.RecommendProduct(ctx)
}

func (s *Store) GetProductsByCategory(ctx context.Context, categoryID string) ([]ProductItem, error) {
	return s.db.GetProductsByCategory(ctx, categoryID)
}

func (s *Store) UpdateSellerContact(ctx context.Context, id string, contact string) (Sellers, error) {
	return s.db.UpdateSellerContact(ctx, id, contact)
}

func (s *Store) GetProductWithSellers(ctx context.Context, IDSellers string) ([]ProductItem, error) {
	return s.db.GetProductWithSellers(ctx, IDSellers)
}

func (s *Store) GetLatestProductsBySeller(ctx context.Context) ([]ProductItem, error) {
	return s.db.GetLatestProductsBySeller(ctx)
}

func (s *Store) AddCartItem(ctx context.Context, cartID string, productID string, quantity int) (Shopping_cart_item, error){
	return s.db.AddCartItem(ctx,cartID,productID,quantity)
}