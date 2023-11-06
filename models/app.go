package models

type HeaderLink struct {
	Title string
	Path  string
}

var PublicLinks = []HeaderLink{
	{
		Title: "Contact",
		Path:  "/contact",
	},
}

var AuthLinks = []HeaderLink{
	{
		Title: "Lists",
		Path:  "/lists",
	},
	{
		Title: "Contact",
		Path:  "/contact",
	},
}
