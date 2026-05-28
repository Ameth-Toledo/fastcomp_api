package entities

type Category string

const (
	CategoryAudio       Category = "audio"
	CategoryComputacion Category = "computacion"
	CategoryMovil       Category = "movil"
	CategorySmarthome   Category = "smarthome"
	CategoryFotografia  Category = "fotografia"
	CategoryAccesorios  Category = "accesorios"
)

type Condition string

const (
	ConditionNuevo           Condition = "nuevo"
	ConditionReacondicionado Condition = "reacondicionado"
	ConditionSegundaMano     Condition = "segunda_mano"
)

type Product struct {
	ID          string    `json:"id"`
	Brand       string    `json:"brand"`
	Name        string    `json:"name"`
	Category    Category  `json:"category"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Condition   Condition `json:"condition"`
	Featured    bool      `json:"featured"`
	Description string    `json:"description,omitempty"`
	Specs       []string  `json:"specs,omitempty"`
	ImageURL    string    `json:"imageUrl,omitempty"`
	AddedAt     string    `json:"addedAt"`
}
