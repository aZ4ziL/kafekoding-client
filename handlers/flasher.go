package handlers

var flasher Flasher

type Flasher struct {
	Type    string
	Message string
}

func (f *Flasher) Set(_type, msg string) {
	f.Type = _type
	f.Message = msg
}

func (f *Flasher) Del() {
	f.Type = ""
	f.Message = ""
}
