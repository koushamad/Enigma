package entity

type Chars []byte

type Reflect struct {
	Chars
}

type Reflector interface {
	Save()
}
