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
		Title: "Dashboard",
		Path:  "/dashboard",
	},
	{
		Title: "Contact",
		Path:  "/contact",
	},
}
