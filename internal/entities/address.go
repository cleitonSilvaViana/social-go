package entities

type Country struct {
	// cca3 is the ISO 3166-1 alpha-3, is a sistem of codes of tree characters used for represently countries.
	CCA3 string // primary key
	
	Name string
}

type State struct {
	ID int
	Name string 
	CountryCCA3 string
}

type City struct {
	Name string
	StateID int
}
