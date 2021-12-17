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

func processLiteralValue(literal []int) int {
	isLastValue := false
	valueStartBit := 0
	binVal := make([]int, 0)
	for !isLastValue {
		if literal[valueStartBit] == 0 {
			isLastValue = true
		}
		binVal = append(binVal, literal[valueStartBit+1:valueStartBit+5]...)
		valueStartBit += 5
	}
	return fromBinaryToDec(binVal)
}

func fromBinaryToDec(bin []int) int {
	value := 0

	for i := len(bin) - 1; i >= 0; i-- {
		value += int(math.Pow(float64(2), float64(len(bin)-1-i))) * bin[i]
	}
	return value
}

func processPacket(packet []int) (int, int, int) {
	fmt.Printf("Packet: %+v\n", packet)
	version := fromBinaryToDec([]int{packet[0], packet[1], packet[2]})
	typeId := fromBinaryToDec([]int{packet[3], packet[4], packet[5]})
	fmt.Printf("version: %+v\n", version)
	fmt.Printf("type id: %+v\n", typeId)
	length := 6
	lastPacketBit := len(packet) - 1
	versionSum := version
	litVal := make([]int, 0)
	if typeId == 4 {
		for i := 6; i < len(packet); i += 5 {
			if packet[i] == 0 {
				currentNumOfBits := i + 4 + 1
				length = currentNumOfBits
				break
			}
		}
		val := processLiteralValue(packet[6 : lastPacketBit+1])
		fmt.Printf("Lit val: %+v\n", litVal)
		fmt.Printf("Legnth: %+v\n", length)
		return val, length, versionSum
	} else {
		var num int
		lengthTypeId := packet[6]
		fmt.Printf("length type id: %+v\n", lengthTypeId)
		if lengthTypeId == 1 {
			numOfPackets := fromBinaryToDec(packet[7:18])
			fmt.Printf("Num of packets: %+v\n", numOfPackets)
			lastPacketBit = 18
			length = 18
			for i := 1; i <= numOfPackets; i++ {
				v, endBit, vers := processPacket(packet[lastPacketBit:])
				lastPacketBit = endBit + lastPacketBit
				length += endBit
				versionSum += vers
				litVal = append(litVal, v)
			}
		} else {
			bitLength := fromBinaryToDec(packet[7:22])
			fmt.Printf("Bit length: %+v\n", bitLength)
			endOfBlock := 21 + bitLength
			fmt.Printf("End of block: %+v\n", endOfBlock)
			lastPacketBit = 22
			length = 22
			for lastPacketBit < endOfBlock {
				v, endBit, vers := processPacket(packet[lastPacketBit : endOfBlock+1])
				lastPacketBit = endBit + lastPacketBit
				length += endBit
				versionSum += vers
				litVal = append(litVal, v)
			}
		}

		switch typeId {
		case 0:
			num = sumPac(litVal)
			break
		case 1:
			num = prosPac(litVal)
			break
		case 2:
			num = minPac(litVal)
			break
		case 3:
			num = maxPac(litVal)
			break
		case 5:
			if litVal[0] > litVal[1] {
				num = 1
			} else {
				num = 0
			}
			break
		case 6:
			if litVal[0] < litVal[1] {
				num = 1
			} else {
				num = 0
			}
			break
		case 7:
			if litVal[0] == litVal[1] {
				num = 1
			} else {
				num = 0
			}
			break
		}

		return num, length, versionSum
	}
}

func sumPac(pacets []int) int {
	sum := 0
	for _, v := range pacets {
		sum += v
	}
	return sum
}

func prosPac(pacets []int) int {
	num := 1
	for _, v := range pacets {
		num *= v
	}
	return num
}

func minPac(pacets []int) int {
	min := MAX_INT
	for _, v := range pacets {
		if min > v {
			min = v
		}
	}
	return min
}

func maxPac(pacets []int) int {
	max := -1
	for _, v := range pacets {
		if max < v {
			max = v
		}
	}
	return max
}

func run16_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_16.txt")
	binArray := hexaToBinary(input[0])
	fmt.Printf("Bin array: %+v\n", binArray)
	_, _, versionS := processPacket(binArray)
	fmt.Printf("Version sum: %+v\n", versionS)

}

func run16_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_16.txt")
	binArray := hexaToBinary(input[0])
	fmt.Printf("Bin array: %+v\n", binArray)
	num, _, _ := processPacket(binArray)
	fmt.Printf("Num: %+v\n", num)
}
