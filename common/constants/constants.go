package constants

const (
	LayoutIndex   = "layouts/template.html"
	CvsFile       = "files/commas_file.csv"
	InternalError = "500"
	NotFound      = "404"
	Success       = "200"
)

//Pokemon - Struct to map the pokemon data
type Pokemon struct {
	ID       string
	Name     string
	ImageUrl string
}

//Response - Struct to map the pokemon data from the API
type Response struct {
	ID      string    `json: "id"`
	Name    string    `json: "name"`
	Sprites []Sprites `json: "sprites"`
}

//Sprites - Struct to map pokemon's sprites
type Sprites struct {
	Other []Other `json: "other"`
}

//Other - Struct to map pokemon´s artwork
type Other struct {
	OfficialArtwork []FrontDefault `json: "official-artwork"`
}

//FrontDefault - Struct to map pokemon's art url
type FrontDefault struct {
	ArtUrl string `json: "font_default"`
}

//PageData - Struct to send data to the view
type PageData struct {
	Message string
	Status  string
}
