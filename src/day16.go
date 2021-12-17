package main

import (
	"fmt"
	"math"
	"strings"
)

var HexaToBinMap = map[string][]int{
	"0": {0, 0, 0, 0},
	"1": {0, 0, 0, 1},
	"2": {0, 0, 1, 0},
	"3": {0, 0, 1, 1},
	"4": {0, 1, 0, 0},
	"5": {0, 1, 0, 1},
	"6": {0, 1, 1, 0},
	"7": {0, 1, 1, 1},
	"8": {1, 0, 0, 0},
	"9": {1, 0, 0, 1},
	"A": {1, 0, 1, 0},
	"B": {1, 0, 1, 1},
	"C": {1, 1, 0, 0},
	"D": {1, 1, 0, 1},
	"E": {1, 1, 1, 0},
	"F": {1, 1, 1, 1},
}

func hexaToBinary(hexa string) []int {
	binArray := make([]int, 0)
	for _, c := range strings.Split(hexa, "") {
		binArray = append(binArray, HexaToBinMap[c]...)
	}
	return binArray
}

type BITSPocket struct {
	version int
	typeId  int
}

func processLiteralValue(literal []int) []int {
	isLastValue := false
	valueStartBit := 0
	decValues := make([]int, 0)
	for !isLastValue {
		if literal[valueStartBit] == 0 {
			isLastValue = true
		}
		decValues = append(decValues, fromBinaryToDec(literal[valueStartBit+1:valueStartBit+5]))
		valueStartBit += 5
	}
	return decValues
}

func fromBinaryToDec(bin []int) int {
	value := 0

	for i := len(bin) - 1; i >= 0; i-- {
		value += int(math.Pow(float64(2), float64(len(bin)-1-i))) * bin[i]
	}
	return value
}

func processPacket(packet []int) ([]int, int, int) {
	fmt.Printf("Packet: %+v\n", packet)
	version := fromBinaryToDec([]int{packet[0], packet[1], packet[2]})
	typeId := fromBinaryToDec([]int{packet[3], packet[4], packet[5]})
	fmt.Printf("version: %+v\n", version)
	fmt.Printf("type id: %+v\n", typeId)
	length := 6
	lastPacketBit := len(packet) - 1
	versionSum := version
	if typeId == 4 {
		//var lastPacketBit int
		for i := 6; i < len(packet); i += 5 {
			if packet[i] == 0 {
				currentNumOfBits := i + 4 + 1
				//if currentNumOfBits%4 != 0 {
				//	lastPacketBit = currentNumOfBits + (4 - currentNumOfBits%4)
				//} else {
				//	lastPacketBit = currentNumOfBits
				//}
				length = currentNumOfBits
				break
			}
		}
		litVal := processLiteralValue(packet[6 : lastPacketBit+1])
		fmt.Printf("Lit val: %+v\n", litVal)
		fmt.Printf("Legnth: %+v\n", length)
		return litVal, length, versionSum
	} else {
		lengthTypeId := packet[6]
		fmt.Printf("length type id: %+v\n", lengthTypeId)
		if lengthTypeId == 1 {
			numOfPackets := fromBinaryToDec(packet[7:18])
			fmt.Printf("Num of packets: %+v\n", numOfPackets)
			lastPacketBit = 18
			length = 18
			litVal := make([]int, 0)
			for i := 1; i <= numOfPackets; i++ {
				v, endBit, vers := processPacket(packet[lastPacketBit:])
				lastPacketBit = endBit + lastPacketBit
				length += endBit
				versionSum += vers
				litVal = append(litVal, v...)
			}

			return litVal, length, versionSum
			//litVal:=make([]int, 0)
		} else {
			bitLength := fromBinaryToDec(packet[7:22])
			fmt.Printf("Bit length: %+v\n", bitLength)
			endOfBlock := 21 + bitLength
			fmt.Printf("End of block: %+v\n", endOfBlock)
			lastPacketBit = 22
			length = 22
			litVal := make([]int, 0)
			for lastPacketBit < endOfBlock {
				v, endBit, vers := processPacket(packet[lastPacketBit : endOfBlock+1])
				lastPacketBit = endBit + lastPacketBit
				length += endBit
				versionSum += vers
				litVal = append(litVal, v...)
			}

			return litVal, length, versionSum

		}

	}
}

func run16_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_16.txt")
	binArray := hexaToBinary(input[0])
	fmt.Printf("Bin array: %+v\n", binArray)
	_, _, versionS := processPacket(binArray)
	fmt.Printf("Version sum: %+v\n", versionS)

}
