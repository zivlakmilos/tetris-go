package assets

import (
	"embed"
)

//go:embed monogram.ttf
var MonogramFont []byte

//go:embed Sounds_clear.mp3
var SoundClear embed.FS

//go:embed Sounds_music.mp3
var SoundMusic embed.FS

//go:embed Sounds_rotate.mp3
var SoundRotate embed.FS
