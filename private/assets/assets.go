package assets

import _ "embed"

//go:embed monogram.ttf
var MonogramFont []byte

//go:embed clear.mp3
var SoundClear []byte

//go:embed music.mp3
var SoundMusic []byte

//go:embed rotate.mp3
var SoundRotate []byte
