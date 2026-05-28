package dto

type CreateProductRequest struct {
	Brand       string   `form:"brand"       binding:"required"`
	Name        string   `form:"name"        binding:"required"`
	Category    string   `form:"category"    binding:"required,oneof=audio computacion movil smarthome fotografia accesorios"`
	Price       float64  `form:"price"       binding:"required,gt=0"`
	Stock       int      `form:"stock"       binding:"min=0"`
	Condition   string   `form:"condition"   binding:"required,oneof=nuevo reacondicionado segunda_mano"`
	Featured    bool     `form:"featured"`
	Description string   `form:"description"`
	Specs       []string `form:"specs"`
}

type UpdateProductRequest struct {
	Brand       string   `form:"brand"       binding:"required"`
	Name        string   `form:"name"        binding:"required"`
	Category    string   `form:"category"    binding:"required,oneof=audio computacion movil smarthome fotografia accesorios"`
	Price       float64  `form:"price"       binding:"required,gt=0"`
	Stock       int      `form:"stock"       binding:"min=0"`
	Condition   string   `form:"condition"   binding:"required,oneof=nuevo reacondicionado segunda_mano"`
	Featured    bool     `form:"featured"`
	Description string   `form:"description"`
	Specs       []string `form:"specs"`
}
