package models

type Config1 struct{
	C11 float64
	C12 string
}
type Config3 struct{
	C31 string
}
type Config struct{
	Config2 string
	Config1 Config1
	Config3 Config3
}