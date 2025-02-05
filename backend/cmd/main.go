// main.go

package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"context"
	"log"

	product "backend/internal/database"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := product.NewPostgresDatabase(cfg.GetConnectionString())
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	store := product.NewStore(db)
	h := handlers.NewProductHandlers(store)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := db.Ping(); err != nil {
				log.Printf("Database connection lost: %v", err)
				// พยายามเชื่อมต่อใหม่
				if reconnErr := db.Reconnect(cfg.GetConnectionString()); reconnErr != nil {
					log.Printf("Failed to reconnect: %v", reconnErr)
				} else {
					log.Printf("Successfully reconnected to the database")
				}
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// กำหนดค่า CORS
	configCors := cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // "*" ยอมรับทุกโดเมน
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(configCors))

	r.Use(TimeoutMiddleware(5 * time.Second))

	r.GET("/health", h.HealthCheck)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// Products Routes
		products := v1.Group("/products")
		{
			products.GET("", h.GetProducts)          // ดึงข้อมูลผลิตภัณฑ์ทั้งหมด
			products.POST("", h.AddProduct)          // เพิ่มผลิตภัณฑ์ใหม่
			products.GET("/:id", h.GetProduct)       // ดึงข้อมูลผลิตภัณฑ์ตาม ID
			products.PUT("/:id", h.UpdateProduct)    // อัปเดตผลิตภัณฑ์ตาม ID
			products.DELETE("/:id", h.DeleteProduct) // ลบผลิตภัณฑ์ตาม ID

			products.GET("/seller/:sellerid",h.GetProductWithSellers)
			products.GET("/category/:categoryID", h.GetProductsByCategory)
			products.GET("/search", h.SearchProducts)
			products.GET("/recommend", h.RecommendProduct)
			products.GET("/new",h.GetLatestProductsBySeller)
			// Nested resources - Images
			images := products.Group("/:id/images")
			{
				images.GET("", h.GetProductImages)                // ดึงข้อมูลภาพผลิตภัณฑ์
				images.POST("", h.AddProductImage)                // เพิ่มภาพผลิตภัณฑ์
				images.PUT("/:image_id", h.UpdateProductImage)    // อัปเดตภาพผลิตภัณฑ์
				images.DELETE("/:image_id", h.DeleteProductImage) // ลบภาพผลิตภัณฑ์
			}
		}

		// Categories Routes
		categories := v1.Group("/categories")
		{
			categories.GET("", h.GetCategories)               // ดึงข้อมูลหมวดหมู่ทั้งหมด
			categories.POST("", h.AddCategory)                // เพิ่มหมวดหมู่ใหม่
			categories.PUT("/:id/name", h.UpdateCategoryName) // อัปเดตชื่อหมวดหมู่ตาม ID
			categories.PUT("/:id/description", h.UpdateCategoryDescription)
			categories.DELETE("/:id", h.DeleteCategory) // ลบหมวดหมู่ตาม ID
		}

		// Shipping Addresses Routes
		shippingAddresses := v1.Group("/shipping-addresses")
		{
			shippingAddresses.GET("", h.GetShippingAddresses)      // ดึงข้อมูลที่อยู่จัดส่ง
			shippingAddresses.PUT("/:id", h.UpdateShippingAddress) // อัปเดตที่อยู่จัดส่งตาม ID
		}
		cart := v1.Group("/cart") 
		{
			cart.GET("",h.GetShoppingCart)
			cart.GET("/:cart_id",h.GetShoppingCartItem)
			cart.PUT("/:cart_id/update",h.UpdateShoppingCartItem)
			cart.DELETE("/:cart_id/:product_id",h.DeleteShoppingCartItem)
			cart.POST("/add", h.AddCartItem)
		}
		User := v1.Group("/user")
		{
			User.GET("", h.GetUsers)
			User.POST("", h.AddUser)
		}
		contacts := v1.Group("/contacts")
		{
			contacts.GET("", h.GetContact) // ดึงข้อมูลผู้ติดต่อทั้งหมด
			contacts.POST("", h.AddContact) // เพิ่มผู้ติดต่อใหม่
		}
		sellers := v1.Group("/sellers") 
		{
			sellers.GET("",h.GetSellers)
			sellers.POST("",h.AddSellers)
			sellers.PUT(":id/name",h.UpdateSellerName)
			sellers.PUT(":id/contact",h.UpdateSellerContact)
			sellers.DELETE(":id",h.DeleteSeller)
		}
		// Health Check
		v1.GET("/health", h.HealthCheck) // ตรวจสอบสถานะ API
	}

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
