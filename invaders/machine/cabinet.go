package machine

import (
	"goinvaders/invaders/machine/cpu/i8080"
	"goinvaders/invaders/machine/memory"
)

func NewCabinet() *Cabinet {
	cabinet := &Cabinet{
		Port1:         0x08,
		Port2:         0x03,
		Shift0:        0x00,
		Shift1:        0x00,
		ShiftOffset:   0x00,
		NextInterrupt: 0x0008,
	}
	cabinet.CPU = i8080.NewCPU(&cabinet.Memory)
	cabinet.ScreenBuffer = [ScreenHeight][ScreenWidth][4]byte{}
	//cabinet.ScreenBuffer = make([]byte, ScreenWidth * ScreenHeight * 4)
	cabinet.Memory.LoadRom(RomFilename)
	return cabinet
}

type Cabinet struct {
	Memory memory.Memory
	CPU    *i8080.Intel8080

	NextInterrupt uint16
	ScreenBuffer  [ScreenHeight][ScreenWidth][4]byte
	//ScreenBuffer  []byte

	// space invaders specific ports and shift registers
	Port1, Port2                byte
	Shift0, Shift1, ShiftOffset byte
}
