package invaders

import (
	"github.com/veandco/go-sdl2/sdl"
	"goinvaders/invaders/machine"
	"log"
)

func NewInvaders(cabinet *machine.Cabinet) *Invaders {
	game := &Invaders{
		timer:       0,
		quit:        false,
		initialized: false,
		cabinet:     cabinet,
	}
	if game.windowInit() {
		return game
	} else {
		return nil
	}
}

type Invaders struct {
	quit        bool
	timer       uint32
	initialized bool
	cabinet     *machine.Cabinet

	window   *sdl.Window
	texture  *sdl.Texture
	renderer *sdl.Renderer
}

func (invaders *Invaders) Draw() {
	_ = invaders.renderer.SetDrawColor(0, 0, 255, 255)

	if err := invaders.renderer.Clear(); err != nil {
		sdl.Log("invaders.renderer.Clear: %s", sdl.GetError())
		log.Printf("invaders.renderer.Clear: %s", err.Error())
	}

	if err := invaders.renderer.Copy(invaders.texture, nil, nil); err != nil {
		sdl.Log("invaders.renderer.Copy: %s", sdl.GetError())
		log.Printf("invaders.renderer.Copy: %s", err.Error())
	}
	invaders.renderer.Present()
}

func (invaders *Invaders) PollEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch evType := event.(type) {
		case *sdl.QuitEvent:
			invaders.quit = true
		case *sdl.KeyboardEvent:
			switch event.GetType() {
			case sdl.KEYDOWN:
				invaders.handleKeyDown(evType)
			case sdl.KEYUP:
				invaders.handleKeyUp(evType)
			}
		}
	}
}

func (invaders *Invaders) handleKeyUp(event *sdl.KeyboardEvent) {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		invaders.quit = true
	}
}

func (invaders *Invaders) handleKeyDown(event *sdl.KeyboardEvent) {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		invaders.quit = true
	}
}

func (invaders *Invaders) ShouldQuit() bool {
	return invaders.quit
}

func (invaders *Invaders) Update() {
	cycleCount := uint32(0)
	for cycleCount <= machine.CyclesPerFrame {
		startCycle := invaders.cabinet.CPU.CycleCount

		// grab the next opcode for later to handle space invader special codes
		// if hte opcode is one of the special codes the cpu will treat as a NOP
		//opcode := invaders.cabinet.Memory.Read(invaders.cabinet.CPU.Regs.PC)

		opcode := invaders.cabinet.CPU.Step()
		cycleCount += invaders.cabinet.CPU.CycleCount - startCycle

		// Handle game specific instructions
		if opcode == 0xdb {
			invaders.handleIn()
		} else if opcode == 0xd3 {
			invaders.handleOut()
		}

		// if at a half frame send screen interrupt
		if invaders.cabinet.CPU.CycleCount >= machine.HalfCyclesPerFrame {
			invaders.cabinet.CPU.Interrupt(invaders.cabinet.NextInterrupt)
			invaders.cabinet.CPU.CycleCount -= machine.HalfCyclesPerFrame
			if invaders.cabinet.NextInterrupt == 0x0008 {
				invaders.cabinet.NextInterrupt = 0x0010
			} else {
				invaders.cabinet.NextInterrupt = 0x0008
			}
		}
	}
}

func (invaders *Invaders) GpuUpdate() {
	// one byte of VRAM contains data for 8 pixels
	//for i := uint16(0); i < machine.ScreenWidth*machine.ScreenHeight/8; i++ {
	//	log.Printf("VRAM memory %02x", invaders.cabinet.Memory.Read(machine.VramAddress + i))
	//}

	for i := uint16(0); i < machine.ScreenWidth*machine.ScreenHeight/8; i++ {
		y := i * 8 / machine.ScreenHeight
		baseX := (i * 8) % machine.ScreenHeight
		currentByte := invaders.cabinet.Memory.Read(machine.VramAddress + i)

		for bit := uint16(0); bit < 8; bit++ {
			py := y
			px := baseX + bit
			isLit := ((currentByte >> bit) & 1) == 0x0001

			red := byte(0)
			green := byte(0)
			blue := byte(0)
			if isLit {
				red = 255
			}

			//screen is rotated 90 degrees counter clockwise so compensate the pixels
			tempX := px
			px = py
			py = -tempX + machine.ScreenHeight - 1

			invaders.cabinet.ScreenBuffer[py][px][0] = red
			invaders.cabinet.ScreenBuffer[py][px][1] = green
			invaders.cabinet.ScreenBuffer[py][px][2] = blue
			invaders.cabinet.ScreenBuffer[py][px][3] = 255
		}
	}
	invaders.updateScreen()
}

func (invaders *Invaders) updateScreen() {
	//format, access, width, height, err := invaders.texture.Query()
	//sdl.Log("format %d, access %d, width %d, height %d, err %s", format, access, width, height, err)

	pitch := 4 * machine.ScreenWidth
	var data []byte

	for i := 0; i < machine.ScreenHeight; i++ {
		for j := 0; j < machine.ScreenWidth; j++ {
			//data = append(data, 0, 20, 255, 0)
			data = append(data, invaders.cabinet.ScreenBuffer[i][j][0])
			data = append(data, invaders.cabinet.ScreenBuffer[i][j][1])
			data = append(data, invaders.cabinet.ScreenBuffer[i][j][2])
			data = append(data, invaders.cabinet.ScreenBuffer[i][j][3])
		}
	}
	_ = invaders.texture.Update(nil, data, pitch)
}

func (invaders *Invaders) handleIn() {
	port := invaders.cabinet.CPU.NextByte()
	if port == 0x01 {
		invaders.cabinet.CPU.Regs.A = invaders.cabinet.Port1
	} else if port == 0x02 {
		invaders.cabinet.CPU.Regs.A = invaders.cabinet.Port2
	} else if port == 0x03 {
		shiftVal := uint16(invaders.cabinet.Shift1)<<8 | uint16(invaders.cabinet.Shift0)
		invaders.cabinet.CPU.Regs.A = byte((shiftVal >> (8 - invaders.cabinet.ShiftOffset)) & 0xff)
	}
}

func (invaders *Invaders) handleOut() {
	// port 3 and 5 are used for sound, not currently implemented (missing sound files)
	// port 6 is for a watchdog that has no meaning here

	port := invaders.cabinet.CPU.NextByte()
	if port == 0x02 {
		invaders.cabinet.ShiftOffset = invaders.cabinet.CPU.Regs.A & 0x07
	} else if port == 0x04 {
		invaders.cabinet.Shift0 = invaders.cabinet.Shift1
		invaders.cabinet.Shift0 = invaders.cabinet.CPU.Regs.A
	}
}

func (invaders *Invaders) windowInit() bool {
	var err error
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		sdl.Log("unable to initialise SDL: %s", sdl.GetError())
		return false
	}

	if invaders.window, err = sdl.CreateWindow("Space Invaders",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, machine.ScreenWidth*2, machine.ScreenHeight*2, sdl.WINDOW_RESIZABLE); err != nil {
		sdl.Log("unable to create game window: %s", sdl.GetError())
		return false
	}
	invaders.window.SetMinimumSize(machine.ScreenWidth, machine.ScreenHeight)
	_, _ = sdl.ShowCursor(sdl.DISABLE)

	// create a renderer for the game graphics
	if invaders.renderer, err = sdl.CreateRenderer(
		invaders.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC); err != nil {
		sdl.Log("unable to create render: %s", sdl.GetError())
		return false
	}
	_ = invaders.renderer.SetLogicalSize(machine.ScreenWidth, machine.ScreenHeight)

	rendererInfo, _ := invaders.renderer.GetInfo()
	sdl.Log("using renderer: %s", rendererInfo.Name)

	// create a texture to display
	if invaders.texture, err = invaders.renderer.CreateTexture(
		uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_STREAMING, machine.ScreenWidth, machine.ScreenHeight); err != nil {
		sdl.Log("unable to create texture: %s", sdl.GetError())
		return false
	}
	return true
}
