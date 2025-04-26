package target

type Profile struct {
	Name string
	Age int
}

type User struct {
	Profile
	Email string
}

type Product struct {
	ID int
	Title string
}