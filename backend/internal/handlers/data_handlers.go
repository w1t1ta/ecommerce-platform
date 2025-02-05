// data_handlers.go
package handlers

import (
	product "backend/internal/database"
	"encoding/base64"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	store *product.Store
}

func NewProductHandlers(store *product.Store) *ProductHandlers {
	return &ProductHandlers{store: store}
}

func convertTimesToUserTimezone(product *product.ProductItem, loc *time.Location) {
	product.CreatedAt = product.CreatedAt.In(loc)
	product.UpdatedAt = product.UpdatedAt.In(loc)
	product.Inventory.UpdatedAt = product.Inventory.UpdatedAt.In(loc)

	for j := range product.Images {
		product.Images[j].CreatedAt = product.Images[j].CreatedAt.In(loc)
	}
}

func (h *ProductHandlers) GetProducts(c *gin.Context) {
	cursor := c.Query("cursor")
	var decodedCursor string
	var err error
	if cursor != "" {
		decodedCursor, err = decodeCursor(cursor)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor"})
			return
		}
	}

	categoryIDStr := c.Query("category")
	var categoryID int
	if categoryIDStr != "" {
		categoryID, err = strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
	}

	params := product.ProductQueryParams{
		Cursor:      decodedCursor,
		Search:      c.Query("search"),
		CategoryID:  categoryID,
		SellerID:    c.Query("seller_id"),
		Status:      c.Query("status"),
		ProductType: c.Query("product_type"),
		Sort:        c.Query("sort"),
		Order:       c.Query("order"),
	}

	response, err := h.store.GetProducts(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the NextCursor before sending the response
	if response.NextCursor != "" {
		response.NextCursor = encodeCursor(response.NextCursor)
	}

	// ใช้ UTC แทน
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal("ไม่สามารถโหลด timezone ได้:", err)
	}

	for i := range response.Items {
		convertTimesToUserTimezone(&response.Items[i], loc)
	}

	c.JSON(http.StatusOK, response)
}

func encodeCursor(cursor string) string {
	return base64.StdEncoding.EncodeToString([]byte(cursor))
}

func decodeCursor(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (h *ProductHandlers) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.store.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userTimezone := "Asia/Bangkok"
	loc, err := time.LoadLocation(userTimezone)
	if err != nil {
		log.Fatal("ไม่สามารถโหลด timezone ได้:", err)
	}

	convertTimesToUserTimezone(&product, loc)

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandlers) AddProduct(c *gin.Context) {
	var input struct {
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
		Count       int     `json:"inventory"`
	}

	// Bind JSON data to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set values in the product struct from the input
	product := product.NewProduct{
		Name:        input.Name,
		Description: input.Description,
		Brand:       input.Brand,
		ModelNumber: input.ModelNumber,
		SKU:         input.SKU,
		Price:       input.Price,
		Status:      input.Status,
		SellerID:    input.SellerID,
		ProductType: input.ProductType,
		CategoryID:  input.CategoryID,
	}

	// Call the store's AddProduct method
	createdProduct, err := h.store.AddProduct(c.Request.Context(), product, input.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandlers) GetProductImages(c *gin.Context) {
	id := c.Param("id") // รับค่า id เป็น string

	images, err := h.store.GetProductImages(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, images)
}

func (h *ProductHandlers) AddProductImage(c *gin.Context) {
	id := c.Param("id") // รับค่า id เป็น string

	var image product.NewProductImage
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error json": err.Error()})
		return
	}

	createdImage, err := h.store.AddProductImage(c.Request.Context(), id, image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdImage)
}

func (h *ProductHandlers) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var update product.UpdateProduct
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.store.UpdateProduct(c.Request.Context(), id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandlers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := h.store.DeleteProduct(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandlers) UpdateProductImage(c *gin.Context) {
	id := c.Param("id")            // รับค่า id เป็น string
	imageID := c.Param("image_id") // รับค่า imageID เป็น string

	var update product.UpdateProductImage
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedImage, err := h.store.UpdateProductImage(c.Request.Context(), id, imageID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedImage)
}

func (h *ProductHandlers) DeleteProductImage(c *gin.Context) {
	id := c.Param("id")            // รับค่า id เป็น string
	imageID := c.Param("image_id") // รับค่า imageID เป็น string

	if err := h.store.DeleteProductImage(c.Request.Context(), id, imageID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandlers) GetCategories(c *gin.Context) {
	categories, err := h.store.GetCategory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *ProductHandlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func (h *ProductHandlers) GetUsers(c *gin.Context) {
	users, err := h.store.GetUserLocal(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": users})
}

// AddUser เพิ่มผู้ใช้ใหม่
func (h *ProductHandlers) AddUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error json": err.Error()})
		return
	}

	user, cart, err := h.store.AddUserLocal(c.Request.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error server": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user, "shopping_cart": cart})
}

// GetShippingAddress ดึงข้อมูลที่อยู่จัดส่ง
func (h *ProductHandlers) GetShippingAddresses(c *gin.Context) {
	addresses, err := h.store.GetShippingAddress(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

// UpdateShippingAddress อัปเดตที่อยู่จัดส่ง
func (h *ProductHandlers) UpdateShippingAddress(c *gin.Context) {
	id := c.Param("id") // รับค่า id เป็น string

	// Assuming these fields are part of the request body
	var requestBody struct {
		Name      string `json:"name"`
		Company   string `json:"company"`
		Street    string `json:"street"`
		Apartment string `json:"apartment"`
		Town      string `json:"town"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAddress, err := h.store.UpdateShippingAddress(c.Request.Context(), id, requestBody.Name, requestBody.Company, requestBody.Street, requestBody.Apartment, requestBody.Town, requestBody.Phone, requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAddress)
}

// AddCategory เพิ่มหมวดหมู่ใหม่
func (h *ProductHandlers) AddCategory(c *gin.Context) {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := h.store.AddCategory(c.Request.Context(), input.Name, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

// UpdateCategoryName อัปเดตชื่อหมวดหมู่
func (h *ProductHandlers) UpdateCategoryName(c *gin.Context) {
	idStr := c.Param("id") // รับค่า id เป็น string

	// แปลง id จาก string เป็น int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var requestBody struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.store.UpdateCategoryName(c.Request.Context(), id, requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (h *ProductHandlers) UpdateCategoryDescription(c *gin.Context) {
	idStr := c.Param("id") // รับค่า id เป็น string

	// แปลง id จาก string เป็น int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var requestBody struct {
		Description string `json:"description"` // ใช้ field นี้ในการอัปเดตคำอธิบาย
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้งานฟังก์ชัน UpdateCategoryDescription
	updatedCategory, err := h.store.UpdateCategoryDescription(c.Request.Context(), id, requestBody.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งคืนข้อมูลหมวดหมู่ที่ถูกอัปเดต
	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory ลบหมวดหมู่
func (h *ProductHandlers) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id") // รับค่า id เป็น string

	// แปลง id จาก string เป็น int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.store.DeleteCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandlers) GetContact(c *gin.Context) {
	// เรียกใช้ฟังก์ชัน GetContacts เพื่อดึงข้อมูลผู้ติดต่อ
	contacts, err := h.store.GetContact(c.Request.Context())
	if err != nil {
		// หากเกิดข้อผิดพลาด ให้ส่งคืนสถานะ 500 และข้อความข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งคืนข้อมูลผู้ติดต่อในรูปแบบ JSON
	c.JSON(http.StatusOK, contacts)
}

func (h *ProductHandlers) AddContact(c *gin.Context) {
	var requestBody struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Message string `json:"message"`
	}

	// Binding JSON data from the request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call AddContact function with the data from the request
	contact, err := h.store.AddContact(c.Request.Context(), requestBody.Name, requestBody.Email, requestBody.Phone, requestBody.Message)
	if err != nil {
		// If an error occurs, return a 500 status with the error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the newly added contact in JSON format
	c.JSON(http.StatusOK, contact)
}

func (h *ProductHandlers) SearchProducts(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "word query parameter is required"})
		return
	}

	// เรียกใช้ฟังก์ชัน Search จาก PostgresDatabase
	products, err := h.store.Search(c.Request.Context(), word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// รับพารามิเตอร์ sort จาก query string
	sortOrder := c.Query("sort")
	if sortOrder == "asc" || sortOrder == "desc" {
		// แปลง products เป็น []product.Product
		var productSlice []product.Product
		for _, productItem := range products {
			productSlice = append(productSlice, product.Product{
				ID:          productItem.ID,          // แปลง ID
				Name:        productItem.Name,        // แปลง Name
				Description: productItem.Description, // แปลง Description
				Brand:       productItem.Brand,       // แปลง Brand
				ModelNumber: productItem.ModelNumber, // แปลง Model Number
				SKU:         productItem.SKU,         // แปลง SKU
				Price:       productItem.Price,       // แปลง Price
				Status:      productItem.Status,      // แปลง Status
				SellerID:    productItem.SellerID,    // แปลง SellerID
				ProductType: productItem.ProductType, // แปลง ProductType
				CategoryID:  productItem.CategoryID,  // แปลง CategoryID
				CreatedAt:   productItem.CreatedAt,   // แปลง CreatedAt
				UpdatedAt:   productItem.UpdatedAt,   // แปลง UpdatedAt
			})
		}
		SortProductsByPrice(productSlice, sortOrder) // เรียงลำดับตามราคา

		// ส่งข้อมูล products ที่ถูกเรียงลำดับกลับไปยัง client
		c.JSON(http.StatusOK, productSlice)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) RecommendProduct(c *gin.Context) {
	// เรียกใช้ฟังก์ชัน RecommendProduct จาก store เพื่อดึงผลิตภัณฑ์สุ่ม
	products, err := h.store.RecommendProduct(c.Request.Context())
	if err != nil {
		// ส่งกลับข้อผิดพลาดถ้ามี
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งกลับผลิตภัณฑ์ที่แนะนำในรูปแบบ JSON
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetProductsByCategory(c *gin.Context) {
	// รับ categoryID จากพารามิเตอร์ใน URL
	categoryID := c.Param("categoryID")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "categoryID is required"})
		return
	}

	// เรียกใช้ฟังก์ชัน GetProductsByCategory จาก PostgresDatabase
	products, err := h.store.GetProductsByCategory(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// รับพารามิเตอร์ sort จาก query string
	sortOrder := c.Query("sort")
	if sortOrder == "asc" || sortOrder == "desc" {
		// แปลง products เป็น []product.Product
		var productSlice []product.Product
		for _, productItem := range products {
			productSlice = append(productSlice, product.Product{
				ID:          productItem.ID,          // แปลง ID
				Name:        productItem.Name,        // แปลง Name
				Description: productItem.Description, // แปลง Description
				Brand:       productItem.Brand,       // แปลง Brand
				ModelNumber: productItem.ModelNumber, // แปลง Model Number
				SKU:         productItem.SKU,         // แปลง SKU
				Price:       productItem.Price,       // แปลง Price
				Status:      productItem.Status,      // แปลง Status
				SellerID:    productItem.SellerID,    // แปลง SellerID
				ProductType: productItem.ProductType, // แปลง ProductType
				CategoryID:  productItem.CategoryID,  // แปลง CategoryID
				CreatedAt:   productItem.CreatedAt,   // แปลง CreatedAt
				UpdatedAt:   productItem.UpdatedAt,   // แปลง UpdatedAt
			})
		}
		SortProductsByPrice(productSlice, sortOrder) // เรียงลำดับตามราคา

		// ส่งข้อมูล products ที่ถูกเรียงลำดับกลับไปยัง client
		c.JSON(http.StatusOK, productSlice)
		return
	}

	// ถ้าไม่มีการระบุ sort ส่ง products กลับไป
	c.JSON(http.StatusOK, products)
}

func SortProductsByPrice(products []product.Product, sortOrder string) {
	if sortOrder == "desc" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
	} else {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	}
}

func (h *ProductHandlers) GetSellers(c *gin.Context) {
	// เรียกใช้ฟังก์ชัน GetContacts เพื่อดึงข้อมูลผู้ติดต่อ
	Sellers, err := h.store.GetSellers(c.Request.Context())
	if err != nil {
		// หากเกิดข้อผิดพลาด ให้ส่งคืนสถานะ 500 และข้อความข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งคืนข้อมูลผู้ติดต่อในรูปแบบ JSON
	c.JSON(http.StatusOK, Sellers)
}
func (h *ProductHandlers) AddSellers(c *gin.Context) {
	// Struct สำหรับรับข้อมูลจาก request
	var input struct {
		Name    string `json:"name"`
		Contact string `json:"contact_info"`
	}

	// ตรวจสอบและแปลง JSON request เป็น struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้ฟังก์ชัน AddSellers ใน PostgresDatabase เพื่อเพิ่ม seller
	newSeller, err := h.store.AddSellers(c.Request.Context(), input.Name, input.Contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูล seller ที่เพิ่มแล้วกลับไปยัง client
	c.JSON(http.StatusCreated, newSeller)
}

func (h *ProductHandlers) UpdateSellerName(c *gin.Context) {
	// รับ seller_id จากพารามิเตอร์ใน URL
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seller_id is required"})
		return
	}

	// Struct สำหรับรับข้อมูลจาก JSON payload
	var input struct {
		Name string `json:"name"`
	}

	// ตรวจสอบและแปลง JSON request เป็น struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้ฟังก์ชัน UpdateSellerName ใน PostgresDatabase
	updatedSeller, err := h.store.UpdateSellerName(c.Request.Context(), id, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูล seller ที่อัปเดตแล้วกลับไปยัง client
	c.JSON(http.StatusOK, updatedSeller)
}

func (h *ProductHandlers) UpdateSellerContact(c *gin.Context) {
	// รับ seller_id จากพารามิเตอร์ใน URL
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seller_id is required"})
		return
	}

	// Struct สำหรับรับข้อมูลจาก JSON payload
	var input struct {
		Contact string `json:"contact_info""`
	}

	// ตรวจสอบและแปลง JSON request เป็น struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้ฟังก์ชัน UpdateSellerContact ใน PostgresDatabase
	updatedSeller, err := h.store.UpdateSellerContact(c.Request.Context(), id, input.Contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูล seller ที่อัปเดตแล้วกลับไปยัง client
	c.JSON(http.StatusOK, updatedSeller)
}

func (h *ProductHandlers) DeleteSeller(c *gin.Context) {
	// รับ seller_id จากพารามิเตอร์ใน URL
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seller_id is required"})
		return
	}

	// เรียกใช้ฟังก์ชัน DeleteSeller ใน PostgresDatabase
	err := h.store.DeleteSeller(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลการลบสำเร็จกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{"message": "seller deleted successfully"})
}

func (h *ProductHandlers) GetProductWithSellers(c *gin.Context) {
	// ดึง ID ของผู้ขายจาก query parameters
	IDSellers := c.Param("sellerid")
	if IDSellers == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seller_id is required"})
		return
	}

	// เรียกฟังก์ชันเพื่อดึงข้อมูลสินค้า
	ProductSellers, err := h.store.GetProductWithSellers(c.Request.Context(), IDSellers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับไปยัง client
	c.JSON(http.StatusOK, ProductSellers)
}

func (h *ProductHandlers) GetLatestProductsBySeller(c *gin.Context) {
	// เรียกใช้ context จาก Gin
	ctx := c.Request.Context()

	// เรียกฟังก์ชันเพื่อดึงสินค้ามาใหม่ล่าสุดจากแต่ละร้าน
	products, err := h.store.GetLatestProductsBySeller(ctx)
	if err != nil {
		// ส่งข้อผิดพลาดกลับไปยัง client
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลสินค้าที่ดึงมาได้กลับไปยัง client
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetShoppingCartItem(c *gin.Context) {
	cartID := c.Param("cart_id") // รับ cartID จาก URL path

	if cartID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart_id is required"})
		return
	}

	// เรียกใช้ฟังก์ชันจาก PostgresDatabase เพื่อดึงข้อมูลสินค้าจากรถเข็น
	items, err := h.store.GetShoppingCartItem(c.Request.Context(), cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้ในรูปแบบ JSON
	c.JSON(http.StatusOK, items)
}

func (h *ProductHandlers) UpdateShoppingCartItem(c *gin.Context) {
	// รับ cartID จากพารามิเตอร์ URL
	cartID := c.Param("cart_id")
	if cartID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart_id is required"})
		return
	}

	// รับจำนวน (count) จาก JSON body
	var requestBody struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// เรียกใช้ฟังก์ชันเพื่ออัปเดตจำนวนสินค้า
	updatedItems, err := h.store.UpdateShoppingCartItem(c.Request.Context(), cartID, requestBody.ProductID, requestBody.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลที่อัปเดตกลับไปยังผู้ใช้
	c.JSON(http.StatusOK, updatedItems)
}

func (h *ProductHandlers) DeleteShoppingCartItem(c *gin.Context) {
	// รับ productID และ cartID จากพารามิเตอร์ URL
	productID := c.Param("product_id")
	cartID := c.Param("cart_id") // สมมติว่าคุณมี cartID ใน URL

	if productID == "" || cartID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id and cart_id are required"})
		return
	}

	// เรียกใช้ฟังก์ชันเพื่อทำการลบสินค้าออกจากรถเข็น
	remainingItems, err := h.store.DeleteShoppingCartItem(c.Request.Context(), productID, cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลที่เหลืออยู่ในรถเข็นกลับไปยังผู้ใช้
	c.JSON(http.StatusOK, remainingItems)
}

func (h *ProductHandlers) AddCartItem(c *gin.Context) {
	// รับ cartID, productID, และ quantity จาก JSON body
	var request struct {
		CartID    string `json:"cart_id"`
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// เรียกใช้ฟังก์ชัน AddCartItem
	newItem, err := h.store.AddCartItem(c.Request.Context(), request.CartID, request.ProductID, request.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลของรายการที่เพิ่มใหม่กลับไปยังผู้ใช้
	c.JSON(http.StatusOK, newItem)
}

func (h *ProductHandlers) GetShoppingCart(c *gin.Context) {
	// ดึงข้อมูลรายการรถเข็นทั้งหมด
	carts, err := h.store.GetShoppingCart(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์รายการรถเข็นกลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, carts)
}
