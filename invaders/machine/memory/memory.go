package memory

import (
	"io/ioutil"
	"log"
)

const (
	ROM_SIZE = 0x2000
)

type Memory struct {
	Storage [0x10000]byte
}

func (mem *Memory) LoadRom(romFile string) {
	data, err := ioutil.ReadFile(romFile)
	if err != nil {
		panic(err)
	}
	copy(mem.Storage[:], data[:])
}

func (mem *Memory) Read(address uint16) byte {
	return mem.Storage[address]
}

func (mem *Memory) ReadWord(address uint16) uint16 {
	var value uint16
	value = uint16(mem.Storage[address])
	value |= uint16(mem.Storage[address+1]) << 8

	return value
}

func (mem *Memory) Write(address uint16, value byte) {
	if address >= ROM_SIZE {
		mem.Storage[address] = value
	} else {
		log.Printf("attempt to write to ROM address %04x", address)
	}
}
