package machine

const (
	FPS                 = 60
	ScreenWidth         = 224
	ScreenHeight        = 256
	CyclesPerFrame      = 2000000 / FPS // 2Mhz at 60 fps
	VramAddress         = uint16(0x2400)
	RomFilename         = "roms/invaders"
	HalfCyclesPerFrame  = CyclesPerFrame / 2
	DisassemblyFilename = "roms/invaders.dis"
)
