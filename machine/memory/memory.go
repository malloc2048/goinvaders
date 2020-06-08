package memory

import (
	"io/ioutil"
	"log"
)

const (
	ROM_SIZE = 0x2000
)

type Memory struct {
	mem [0x10000]byte
}

func (mem *Memory) LoadRom(romFile string) {
	data, err := ioutil.ReadFile(romFile)
	if err != nil {
		panic(err)
	}
	copy(mem.mem[:], data[:])
}

func (mem *Memory) ReadByte(address uint16) uint8 {
	return mem.mem[address]
}

func (mem *Memory) ReadWord(address uint16) uint16 {
	mem.mem[0x0000] = 0xaa
	mem.mem[0x0001] = 0x55
	return uint16(mem.mem[address]) | uint16(mem.mem[address+1])<<8
}

func (mem *Memory) Write(address uint16, value uint8) {
	if address >= ROM_SIZE {
		mem.mem[address] = value
	} else {
		log.Printf("attempt to write to ROM address %04x", address)
	}
}
