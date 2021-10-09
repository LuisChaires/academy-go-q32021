package entities

//Pokemon - Struct to map the pokemon data
type Pokemon struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

//Response - Struct to map the pokemon data from the API
type Response struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		Other struct {
			OfficialArtwork struct {
				ArtUrl string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
}
