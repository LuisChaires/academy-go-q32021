package constants

//mover a archivo de conf
const (
	LayoutIndex = "layouts/template.html"
	CvsFile     = "files/commas_file.csv"
)

type Pokemon struct {
	Id       string
	Name     string
	ImageUrl string
}

type Response struct {
	Id      string    `json: "id"`
	Name    string    `json: "name"`
	Sprites []Sprites `json: "sprites"`
}

type Sprites struct {
	Other []Other `json: "other"`
}

type Other struct {
	OfficialArtwork []FrontDefault `json: "official-artwork"`
}

type FrontDefault struct {
	ArtUrl string `json: "font_default"`
}

//
type PageData struct {
	Message string
	Status  string
}
