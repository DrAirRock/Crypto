package DES

import (
	"bytes"
	"log"
	"strconv"
)

// rotate a string leftward by a specified value
func rotateLeft(str string, shiftVal int) string {
	retStr := str[shiftVal:]
	retStr = retStr + str[0:shiftVal]

	return retStr
}

// Permute a string given a permutation array
func permuteString(oldString string, permuteArray []int) string {
	tmpRunes := []string{}

	for _, v := range permuteArray {
		tmpRunes = append(tmpRunes, oldString[v-1:v])
	}

	var buffer bytes.Buffer

	for i := 0; i < len(tmpRunes); i++ {
		buffer.WriteString(tmpRunes[i])
	}

	resultString := buffer.String()

	return resultString
}

// GenSubkeys Generate subkeys from an original private key (In form of "0x[0-F]^16")
func GenSubkeys(key string) []string {
	decimal, err := strconv.ParseUint(key, 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	binString := strconv.FormatUint(decimal, 2)

	for len(binString) != 64 {
		binString = "0" + binString
	}

	pc1 := []int{57, 49, 41, 33, 25, 17, 9, 1, 58, 50, 42, 34, 26, 18, 10, 2, 59, 51, 43, 35, 27, 19, 11, 3, 60, 52, 44, 36, 63, 55, 47, 39, 31, 23, 15, 7, 62, 54, 46, 38, 30, 22, 14, 6, 61, 53, 45, 37, 29, 21, 13, 5, 28, 20, 12, 4}

	kplusString := permuteString(binString, pc1)

	c0, d0 := kplusString[0:28], kplusString[28:56]

	cS := []string{}
	dS := []string{}

	cS = append(cS, c0)
	dS = append(dS, d0)

	shifts := []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	for i, v := range shifts {
		cS = append(cS, rotateLeft(cS[i], v))
		dS = append(dS, rotateLeft(dS[i], v))
	}

	kS := []string{}

	for i := 1; i <= 16; i++ {
		kS = append(kS, cS[i]+dS[i])
	}

	pc2 := []int{14, 17, 11, 24, 1, 5, 3, 28, 15, 6, 21, 10, 23, 19, 12, 4, 26, 8, 16, 7, 27, 20, 13, 2, 41, 52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48, 44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32}

	for i := 0; i < 16; i++ {
		kS[i] = permuteString(kS[i], pc2)
	}

	return kS
}

// do an XOR operation on two bitstrings
func xorString(a, b string) string {
	aprime, _ := strconv.ParseInt(a, 2, 64)
	bprime, _ := strconv.ParseInt(b, 2, 64)

	combined := aprime ^ bprime

	resString := strconv.FormatInt(combined, 2)

	for len(resString)%16 != 0 {
		resString = "0" + resString
	}

	return resString
}

// S-boxes computation
func sBox(idx int, str string) string {
	sboxes := [][][]int{
		[][]int{
			[]int{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
			[]int{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
			[]int{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
			[]int{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
		},
		[][]int{
			[]int{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
			[]int{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
			[]int{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
			[]int{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
		},
		[][]int{
			[]int{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
			[]int{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
			[]int{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
			[]int{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
		},
		[][]int{
			[]int{7, 12, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
			[]int{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
			[]int{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
			[]int{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
		},
		[][]int{
			[]int{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
			[]int{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
			[]int{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
			[]int{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
		},
		[][]int{
			[]int{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
			[]int{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
			[]int{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
			[]int{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
		},
		[][]int{
			[]int{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
			[]int{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
			[]int{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
			[]int{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
		},
		[][]int{
			[]int{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
			[]int{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
			[]int{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
			[]int{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
		},
	}

	i := str[0:1] + str[5:6]
	j := str[1:5]

	iInt, _ := strconv.ParseInt(i, 2, 64)
	jInt, _ := strconv.ParseInt(j, 2, 64)

	resInt := sboxes[idx][iInt][jInt]

	resStr := strconv.FormatInt(int64(resInt), 2)

	for len(resStr) != 4 {
		resStr = "0" + resStr
	}

	return resStr
}

// The fiestel cipher
func fiestel(str, key string) string {
	E := []int{32, 1, 2, 3, 4, 5, 4, 5, 6, 7, 8, 9, 8, 9, 10, 11, 12, 13, 12, 13, 14, 15, 16, 17, 16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1}

	newStr := permuteString(str, E)

	KE := xorString(key, newStr)

	bS := []string{}

	for i := 0; i < 8; i++ {
		bS = append(bS, KE[i*6:i*6+6])
	}

	nBS := []string{}

	for i := 0; i < 8; i++ {
		nBS = append(nBS, sBox(i, bS[i]))
	}

	sBoxxedString := nBS[0] + nBS[1] + nBS[2] + nBS[3] + nBS[4] + nBS[5] + nBS[6] + nBS[7]

	P := []int{16, 7, 20, 21, 29, 12, 28, 17, 1, 15, 23, 26, 5, 18, 31, 10, 2, 8, 24, 14, 32, 27, 3, 9, 19, 13, 30, 6, 22, 11, 4, 25}

	result := permuteString(sBoxxedString, P)

	return result
}

// Des run DES on the input using subkeys
func Des(message string, subkeys []string, decrypt bool) string {
	decimal, err := strconv.ParseUint(message, 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	binString := strconv.FormatUint(decimal, 2)
	for len(binString) != 64 {
		binString = "0" + binString
	}

	IP := []int{58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20, 12, 4, 62, 54, 46, 38, 30, 22, 14, 6, 64, 56, 48, 40, 32, 24, 16, 8, 57, 49, 41, 33, 25, 17, 9, 1, 59, 51, 43, 35, 27, 19, 11, 3, 61, 53, 45, 37, 29, 21, 13, 5, 63, 55, 47, 39, 31, 23, 15, 7}

	mIP := permuteString(binString, IP)

	l0, r0 := mIP[0:32], mIP[32:64]

	oldL, oldR := l0, r0
	newL, newR := "", ""

	for i := 0; i < 16; i++ {
		newL = oldR
		if !decrypt {
			newR = xorString(oldL, fiestel(oldR, subkeys[i]))
		} else {
			newR = xorString(oldL, fiestel(oldR, subkeys[15-i]))
		}

		oldL = newL
		oldR = newR
	}

	finalStr := newR + newL

	IPInv := []int{40, 8, 48, 16, 56, 24, 64, 32, 39, 7, 47, 15, 55, 23, 63, 31, 38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29, 36, 4, 44, 12, 52, 20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27, 34, 2, 42, 10, 50, 18, 58, 26, 33, 1, 41, 9, 49, 17, 57, 25}

	finalStr = permuteString(finalStr, IPInv)

	return finalStr
}
