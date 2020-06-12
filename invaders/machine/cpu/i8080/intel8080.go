package i8080

import (
	"goinvaders/invaders/machine/cpu/i8080/instructions"
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
	"log"
)

func NewCPU(memory *memory.Memory) *Intel8080 {
	cpu := &Intel8080{
		flags: registers.Flags{
			Sign:      false,
			Zero:      false,
			Carry:     false,
			Parity:    false,
			HalfCarry: false,
		},
		Regs: registers.Registers{
			A: 0,
			B: 0, C: 0,
			D: 0, E: 0,
			H: 0, L: 0,
			SP: 0,
			PC: 0,
		},
		CycleCount: 0,
		memory:     memory,

		InterruptEnabled: true,
	}
	return cpu
}

type Intel8080 struct {
	CycleCount uint32
	flags      registers.Flags
	memory     *memory.Memory
	Regs       registers.Registers

	InterruptEnabled bool
}

func (cpu *Intel8080) Interrupt(address uint16) {
	if cpu.InterruptEnabled {
		cpu.InterruptEnabled = true

		cpu.memory.Write(cpu.Regs.SP-1, uint8(cpu.Regs.PC>>8))
		cpu.memory.Write(cpu.Regs.SP-2, uint8(cpu.Regs.PC&0x00ff))
		cpu.Regs.SP -= 2

		cpu.Regs.PC = address
		cpu.CycleCount += 11
	}
}

func (cpu *Intel8080) NextByte() byte {
	value := cpu.memory.Read(cpu.Regs.PC)
	cpu.Regs.PC += 1
	return value
}

func (cpu *Intel8080) debug() {
	switch cpu.Regs.PC {
	case 0x00:
		log.Printf("%04x reset\n", cpu.Regs.PC)
	//case 0x01ef:
	//	log.Printf("%04x DrawShieldPl1\n", cpu.Regs.PC)
	//case 0x021A:
	//	log.Printf("%04x RestoreShields1\n", cpu.Regs.PC)
	//case 0x021e:
	//	log.Printf("%04x CopyShields\n", cpu.Regs.PC)
	//case 0x01cf:
	//	log.Printf("%04x DrawBottomLine\n", cpu.Regs.PC)
	//case 0x14D7:
	//	log.Printf("%04x Return from ClearSmallSprite\n", cpu.Regs.PC)
	//case 0x01d9:
	//	log.Printf("%04x AddDelta\n", cpu.Regs.PC)
	//case 0x01e4:
	//	log.Printf("%04x CopyRAMMirror\n", cpu.Regs.PC)
	//case 0x002d:
	//	log.Printf("%04x Handle bumping credit count\n", cpu.Regs.PC)
	//case 0x005d:
	//	log.Printf("%04x no game is going and there are credits\n", cpu.Regs.PC)
	//case 0x0067:
	//	log.Printf("%04x Mark credit as needing registering\n", cpu.Regs.PC)
	//case 0x006f:
	//	log.Printf("%04x Main game-play timing loop\n", cpu.Regs.PC)
	//case 0x00b1:
	//	log.Printf("%04x InitRack\n", cpu.Regs.PC)
	//case 0x0100:
	//	log.Printf("%04x DrawAlien\n", cpu.Regs.PC)
	//case 0x0141:
	//	log.Printf("%04x CursorNextAlien\n", cpu.Regs.PC)
	//case 0x017a:
	//	log.Printf("%04x GetAlienCoords\n", cpu.Regs.PC)
	//case 0x01a1:
	//	log.Printf("%04x MoveRefAlien\n", cpu.Regs.PC)
	//case 0x01c0:
	//	log.Printf("%04x InitAliens\n", cpu.Regs.PC)
	//case 0x14cb:
	//	log.Printf("%04x ClearSmallSprite\n", cpu.Regs.PC)
	//case 0x0248:
	//	log.Printf("%04x RunGameObjs\n", cpu.Regs.PC)
	//case 0x024b:
	//	log.Printf("%04x RunGameObjs skipping first instruction\n", cpu.Regs.PC)
	//case 0x0b4a:
	//	log.Printf("%04x Play demo\n", cpu.Regs.PC)
	//case 0x08d1:
	//	log.Printf("%04x GetShipsPerCred\n", cpu.Regs.PC)
	//case 0x09d6:
	//	log.Printf("%04x ClearPlayField\n", cpu.Regs.PC)
	//case 0x1a7f:
	//	log.Printf("%04x RemoveShip\n", cpu.Regs.PC)
	//case 0x1618:
	//	log.Printf("%04x PlrFireOrDemo\n", cpu.Regs.PC)
	//case 0x092e:
	//	log.Printf("%04x Get number of ships for active player\n", cpu.Regs.PC)
	//case 0x1611:
	//	log.Printf("%04x GetPlayerDataPtr\n", cpu.Regs.PC)
	//case 0x19e6:
	//	log.Printf("%04x DrawNumShips\n", cpu.Regs.PC)
	//case 0x01f8:
	//	log.Printf("%04x Going to draw 4 shields\n", cpu.Regs.PC)
	//case 0x0205:
	//	log.Printf("%04x Drawing shields RESETTING here somwhere\n", cpu.Regs.PC)
	default:
	}
}

func (cpu *Intel8080) Step() byte {
	cpu.debug()

	opcode := cpu.NextByte()
	cpu.CycleCount += uint32(OpcodesCycles[opcode])

	switch opcode {
	case 0x00:
	case 0x01:
		instructions.LXI(opcode, cpu.memory, &cpu.Regs)
	case 0x02:
		instructions.STAX(opcode, cpu.memory, &cpu.Regs)
	case 0x03:
		instructions.INX(opcode, &cpu.Regs)
	case 0x04:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x05:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x06:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x07:
		instructions.Rotate(opcode, &cpu.Regs, &cpu.flags)
	case 0x09:
		instructions.DAD(opcode, &cpu.Regs, &cpu.flags)
	case 0x08:
	case 0x0a:
		instructions.LDAX(opcode, cpu.memory, &cpu.Regs)
	case 0x0b:
		instructions.DCX(opcode, &cpu.Regs)
	case 0x0c:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x0d:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x0e:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x0f:
		instructions.Rotate(opcode, &cpu.Regs, &cpu.flags)

	case 0x11:
		instructions.LXI(opcode, cpu.memory, &cpu.Regs)
	case 0x12:
		instructions.STAX(opcode, cpu.memory, &cpu.Regs)
	case 0x13:
		instructions.INX(opcode, &cpu.Regs)
	case 0x14:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x15:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x16:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x17:
		instructions.Rotate(opcode, &cpu.Regs, &cpu.flags)
	case 0x19:
		instructions.DAD(opcode, &cpu.Regs, &cpu.flags)
	case 0x1a:
		instructions.LDAX(opcode, cpu.memory, &cpu.Regs)
	case 0x1b:
		instructions.DCX(opcode, &cpu.Regs)
	case 0x1c:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x1d:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x1e:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x1f:
		instructions.Rotate(opcode, &cpu.Regs, &cpu.flags)

	case 0x21:
		instructions.LXI(opcode, cpu.memory, &cpu.Regs)
	case 0x22:
		instructions.SHLD(opcode, cpu.memory, &cpu.Regs)
	case 0x23:
		instructions.INX(opcode, &cpu.Regs)
	case 0x24:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x25:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x26:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x27:
		instructions.DAA(opcode, &cpu.Regs, &cpu.flags)
	case 0x29:
		instructions.DAD(opcode, &cpu.Regs, &cpu.flags)
	case 0x2a:
		instructions.LHLD(opcode, cpu.memory, &cpu.Regs)
	case 0x2b:
		instructions.DCX(opcode, &cpu.Regs)
	case 0x2c:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x2d:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x2e:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x2f:
		instructions.CMA(opcode, &cpu.Regs)

	case 0x31:
		instructions.LXI(opcode, cpu.memory, &cpu.Regs)
	case 0x32:
		instructions.STA(opcode, cpu.memory, &cpu.Regs)
	case 0x33:
		instructions.INX(opcode, &cpu.Regs)
	case 0x34:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x35:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x36:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x37:
		instructions.STC(opcode, &cpu.Regs, &cpu.flags)
	case 0x39:
		instructions.DAD(opcode, &cpu.Regs, &cpu.flags)
	case 0x3a:
		instructions.LDA(opcode, cpu.memory, &cpu.Regs)
	case 0x3b:
		instructions.DCX(opcode, &cpu.Regs)
	case 0x3c:
		instructions.INR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x3d:
		instructions.DCR(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x3e:
		instructions.MVI(opcode, cpu.memory, &cpu.Regs)
	case 0x3f:
		instructions.CMC(opcode, &cpu.Regs, &cpu.flags)

	case 0x40:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x41:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x42:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x43:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x44:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x45:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x46:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x47:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x48:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x49:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4a:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4b:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4c:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4d:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4e:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x4f:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)

	case 0x50:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x51:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x52:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x53:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x54:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x55:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x56:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x57:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x58:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x59:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5a:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5b:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5c:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5d:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5e:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x5f:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)

	case 0x60:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x61:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x62:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x63:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x64:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x65:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x66:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x67:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x68:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x69:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6a:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6b:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6c:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6d:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6e:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x6f:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)

	case 0x70:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x71:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x72:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x73:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x74:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x75:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x76: // HLT
	case 0x77:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x78:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x79:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7a:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7b:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7c:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7d:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7e:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)
	case 0x7f:
		instructions.MOV(opcode, cpu.memory, &cpu.Regs)

	case 0x80:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x81:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x82:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x83:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x84:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x85:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x86:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x87:
		instructions.ADD(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x88:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x89:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8a:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8b:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8c:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8d:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8e:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x8f:
		instructions.ADC(opcode, cpu.memory, &cpu.Regs, &cpu.flags)

	case 0x90:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x91:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x92:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x93:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x94:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x95:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x96:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x97:
		instructions.SUB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x98:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x99:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9a:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9b:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9c:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9d:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9e:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0x9f:
		instructions.SBB(opcode, cpu.memory, &cpu.Regs, &cpu.flags)

	case 0xa0:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa1:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa2:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa3:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa4:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa5:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa6:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa7:
		instructions.ANA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa8:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xa9:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xaa:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xab:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xac:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xad:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xae:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xaf:
		instructions.XRA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)

	case 0xb0:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb1:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb2:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb3:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb4:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb5:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb6:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb7:
		instructions.ORA(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb8:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xb9:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xba:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xbb:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xbc:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xbd:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xbe:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xbf:
		instructions.CMP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)

	case 0xc0:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc1:
		instructions.POP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc2:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc3:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc4:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc5:
		instructions.PUSH(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc6:
		instructions.ADI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc7:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)
	case 0xc8:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xc9:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xca:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xcc:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xcd:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xce:
		instructions.ACI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xcf:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)

	case 0xd0:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd1:
		instructions.POP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd2:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd4:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd5:
		instructions.PUSH(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd6:
		instructions.SUI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xd7:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)
	case 0xd8:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xda:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xdc:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xde:
		instructions.SBI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xdf:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)

	case 0xe0:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe1:
		instructions.POP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe2:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe3:
		instructions.XTHL(opcode, cpu.memory, &cpu.Regs)
	case 0xe4:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe5:
		instructions.PUSH(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe6:
		instructions.ANI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe7:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)
	case 0xe8:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xe9:
		instructions.PCHL(&cpu.Regs)
	case 0xea:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xeb:
		instructions.XCHG(opcode, &cpu.Regs)
	case 0xec:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xee:
		instructions.XRI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xef:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)

	case 0xf0:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf1:
		instructions.POP(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf2:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf3:
		cpu.DI(opcode)
	case 0xf4:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf5:
		instructions.PUSH(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf6:
		instructions.ORI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf7:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)
	case 0xf8:
		instructions.Return(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xf9:
		instructions.SPHL(opcode, &cpu.Regs)
	case 0xfa:
		instructions.Jump(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xfb:
		cpu.EI(opcode)
	case 0xfc:
		instructions.Call(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xfe:
		instructions.CPI(opcode, cpu.memory, &cpu.Regs, &cpu.flags)
	case 0xff:
		instructions.RST(opcode, cpu.memory, &cpu.Regs)

	default:
		// NOP / Unimplemented
	}
	return opcode
}

// consider moving these to instructions
func (cpu *Intel8080) DI(opcode byte) {
	cpu.InterruptEnabled = false
	cpu.Regs.PC += uint16(instructions.OpcodesLength[opcode] - 1)
}

func (cpu *Intel8080) EI(opcode byte) {
	cpu.InterruptEnabled = true
	cpu.Regs.PC += uint16(instructions.OpcodesLength[opcode] - 1)
}
