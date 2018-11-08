package schema

//go:generate msgp -o=stuff.go -tests=false
type Person struct {
	Name       string `msg:"name"`
	Address    string `msg:"address"`
	Age        int    `msg:"age"`
	Hidden     string `msg:"-"`
	unexported bool
}
