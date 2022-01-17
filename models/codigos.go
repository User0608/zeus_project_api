package models

type CategoriaCodigo struct {
	Index        string       `json:"id" gorm:"primaryKey"`
	Title        string       `json:"title"`
	CodigosItems []CodigoItem `json:"codigos" gorm:"foreignKey:CategoriaCodigoIndex"`
}

type CodigoItem struct {
	Codigo               string `json:"codigo"`
	Detalle              string `json:"detalle"`
	CategoriaCodigoIndex string `json:"-"`
}
