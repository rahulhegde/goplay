package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"

	grpcproto "github.com/rahulhegde/goplay/gobasic/grpcprotobufplay"
)

func Sqrt(value float64) (string, string) {
	return fmt.Sprint(math.Sqrt(value)), fmt.Sprint(math.Sqrt(value))
}

const STARTER = 100
const DELTA = 0.0001

func NewtonSquareValue(point float64, sqrt int) (return64 float64) {
	return64 = (point*point + float64(sqrt)) / (2 * point)
	return
}

func NewtonSquare(sqrt int) float64 {
	prevSqrtValue := NewtonSquareValue(STARTER, sqrt)
	newSqrtValue := NewtonSquareValue(prevSqrtValue, sqrt)
	for math.Abs(newSqrtValue-prevSqrtValue) > DELTA {
		prevSqrtValue = newSqrtValue
		newSqrtValue = NewtonSquareValue(prevSqrtValue, sqrt)
	}
	return newSqrtValue
}

func TimePlay() {
	location, err := time.LoadLocation("CET")
	if err != nil {
		fmt.Println("LoadLocation failed ", err.Error())
	}

	fmt.Println("Location = ", location.String())
	fmt.Println("time = ", time.Now().In(location))

	timeLOCAL := time.Now()
	fmt.Println("Local Time " + timeLOCAL.Location().String())

	fmt.Println("Location " + timeLOCAL.UTC().Location().String())

	const DateTimeLayoutWithoutTimezone = "2006-01-02T15:04:05"
	timeTZ, _ := time.ParseInLocation(DateTimeLayoutWithoutTimezone, "2017-03-13T15:24:05", time.Local)
	fmt.Println("*** time check-2", timeTZ.String())

	fmt.Println("day = ", time.Now().Day())
	fmt.Println(time.September)

}

func switchUsage(x string) {
	fmt.Println("Entering switch check for OS", runtime.GOOS)

	switch x {
	case "string":
		fmt.Println("switch::string")
	case "sting":
		fmt.Println("switch::sting")
	default:
		fmt.Println("switch::default")
	}
}

func main() {

	fmt.Println("random = ", rand.Intn(100), " bitwise = ", 1<<3)

	var i1 = 3

	for i, j := 0, 20; i < 1; i, j = i+1, j-1 {
		fmt.Printf("random = %d, %d \n", i, j)
	}

	if i1 = 2; i1 > 4 {
		fmt.Print("help...4")

	} else if i1 > 2 {
		fmt.Print("help...2")
	} else {
		fmt.Print("help...1")
	}

	a, b := Sqrt(4)
	fmt.Printf("\nSquare Root1 = %s, %s", a, b)

	fmt.Printf("\nSquare Root via Newton Method = %f\n", NewtonSquare(102))

	switchUsage("sting")

	PointerPlay()

	StructurePlay()

	ArrayPlay()

	ArrayStructurePlay()

	SliceLiteralPlay()

	SliceDefaultPlay()

	D2ArrayPlay()

	RangeForArray()

	// check dynamic creation of 2-d array
	pic.Show(ExerciseSliceCreate2DDynamicImage)

	ReceiverMethodPlay()

	PointerPlay2()

	JSONPlay1()

	JSONPlayTest1()
	JSONPlayTest2()
	JSONPlayTest3()

	HashingPlay()

	InterfacePlay()
	InterfaceTypeAssertionPlay()
	TypeSwitchPlay()
	InterfaceTypeConversion()

	//MapPlay()
	//DecimalPlay()
	//Float64Play()
	//TimePlay()
	//
	//BitwisePlay()
	//
	//GoRoutinePlay()
	//GoChannelPlay()

	grpcproto.GRPCProtoBufPlay()

	BitwiseCheckPlay()

	ShaPlay()

	NewString := ""
	strings.EqualFold(NewString, "hello")
}

type BOOLEAN int

const (
	TRUE  BOOLEAN = 1
	FALSE BOOLEAN = 4
)

func (state *BOOLEAN) setTRUE() {
	*state = *state | TRUE
}

func (state *BOOLEAN) isTRUE() bool {

	status := *state & TRUE
	fmt.Println("isTRUE = ", *state, status)
	return (*state & TRUE) == TRUE
}

func (state *BOOLEAN) isFALSE() bool {
	status := *state & FALSE
	fmt.Println("isFALSE = ", *state, status)
	return (*state & FALSE) == FALSE
}

func (state *BOOLEAN) setFALSE() {
	*state = *state | FALSE
}

func BitwisePlay() {
	state := TRUE
	fmt.Println("state after setting to TRUE: ", state)
	state.setFALSE()
	fmt.Println("state after setting to FALSE: ", state)
	fmt.Println("isFALSE = ", state.isFALSE(), "isTRUE: ", state.isTRUE())

}

func HashingPlay() {
	fmt.Println("*** HashingPlay ***")
	var hashInput1 string = "RahulHegde1"
	var hashInput2 string = "RahulHegde"

	fmt.Println("hashInput", hashInput1)
	hasher := sha256.New()

	hasher.Write([]byte(hashInput1))
	fmt.Println(base64.URLEncoding.EncodeToString(hasher.Sum(nil)), hasher.BlockSize())

	hasher.Reset()
	hasher.Write([]byte(hashInput2))
	fmt.Println(hasher.Sum(nil), base64.URLEncoding.EncodeToString(hasher.Sum(nil)), hex.EncodeToString(hasher.Sum(nil)), hasher.BlockSize())
	fmt.Println("Byte = ", base64.RawStdEncoding.EncodeToString([]byte("255 255")))
}

func PointerPlay2() {
}

type AXIX struct {
	x int
	y int
}

func (x AXIX) PrinterAXIS(a, b int) {
	fmt.Println("PrinterAXIS: x= ", x.x, " y= ", x.y, "a=", a, "b=", b)
}

type FLOAT64 float64

func (f FLOAT64) floater(value float64) {
	fmt.Println("floater=", f, "value=", value)
}

func ReceiverMethodPlay() {
	// call not possible for built-in data type in other package

	FLOAT64(20.02).floater(10.20)

	x := AXIX{1, 2}
	x.PrinterAXIS(11, 22)
	AXIX{1, 10}.PrinterAXIS(5, 50)

}

func MapPlay() {
	type Vertex struct {
		X int
		Y int
	}

	var mapperstub map[int]string = map[int]string{
		1: "name1",
		2: "name2",
	}
	mapperstub[3] = "name3"
	fmt.Println("mapperstub", mapperstub)

	//var mapperstub1 map[int]string
	mapperstub1 := make(map[int]string)
	mapperstub1[1] = "name2"
	value, exist := mapperstub1[1]
	fmt.Println("mapperstub1", mapperstub1, "exist =", exist, "value = ", value)
	delete(mapperstub1, 1)
	value, exist = mapperstub1[1]
	fmt.Println("mapperstub1", mapperstub1, "exist =", exist, "value = ", value)
}

func ExerciseSliceCreate2DDynamicImage(dx, dy int) (arraystub [][]uint8) {
	//dx, dy := 256, 256

	//
	arraystub = make([][]uint8, dx)
	for indexX := 0; indexX < dx; indexX = indexX + 1 {
		arraystub[indexX] = make([]uint8, dy)
		for indexY := 0; indexY < dy; indexY = indexY + 1 {
			arraystub[indexX][indexY] = uint8(indexX + indexX)
		}
	}
	return
}

func RangeForArray() {
	arraystub := []int{1, 2, 3, 4}

	for index, value := range arraystub {
		fmt.Println("Index=", index, "value=", value)
	}
}

func D2ArrayPlay() {
	var array2d [][]string = [][]string{
		{"S11", "S12"},
		{"S21", "S22"},
	}
	fmt.Println("array2d", array2d)

	for index, value := range array2d {
		fmt.Println(".RAHULR.............index", index, "value=", value)

		for indexX, valueX := range value {
			fmt.Println(".RAHUL.............index", index, "indexX", indexX, "valueX=", valueX)
		}

	}

}

func SliceDefaultPlay() {
	var arraystub []int = []int{1, 2, 3, 4, 5}

	slicer := arraystub[:]
	fmt.Println("slicer1=", slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[:3]
	fmt.Println("slicer2=", slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = slicer[:3:3]
	fmt.Println("increase slicer2=", slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[2:]
	fmt.Println("slicer3=", slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[:]
	fmt.Println("slicer4=", slicer, "len=", len(slicer), "cap=", cap(slicer))

	var slicer1 []int
	if slicer1 == nil {
		fmt.Println("slicer1 is nil=", slicer1, " cap=", cap(slicer1), " len=", len(slicer1))
	}
}

func SliceLiteralPlay() {
	slicer := []int{1, 2, 3}
	var arraystub []int = []int{1, 2, 3}
	fmt.Println("slicer=", slicer, "arraystub=", arraystub)
	return
}

func ArrayStructurePlay() {

	type COORDINATE struct {
		X int
		Y int
	}

	// scenario to create array of structure
	var arrcoordinates [3]COORDINATE = [3]COORDINATE{
		COORDINATE{1, 2},
		COORDINATE{3, 5},
		COORDINATE{4, 5},
	}
	fmt.Println("array of coordinates structure =", arrcoordinates)

	slicer1 := arrcoordinates[2:2]
	fmt.Println("array of structured sliced = ", slicer1)

	// scenario where array is constructed first and then slice is created from it.
	slicer := []struct {
		X int
		Y string
	}{
		{1, "hello"},
		{2, "world"},
	}
	fmt.Println("slicer constructed w/o array creation = ", slicer)

	return
}

func Array1PlayCallByValue(arrtmp [5]string) {
	arrtmp[2] = "name20"
	fmt.Println("Array1PlayCallByValue: ", arrtmp)
}

func Array1PlayCallByReference(arrtmp []string) {
	arrtmp[2] = "name20"
	fmt.Println("Array1PlayCallByReference: ", arrtmp)
}

func Array1PlayCallByAddress(arrtmp *string) {
	*arrtmp = "name20"
	fmt.Println("Array1PlayCallByAddress: ", *arrtmp)
}

func ArrayPlay() {
	type COORDINATE struct {
		X int
		Y int
		Z int
	}

	var arr1 [5]int = [5]int{1, 2, 3}
	var arr2 [5]int = [5]int{}
	arr3 := COORDINATE{1, 2, 3}
	arr4 := [5]int{1, 2, 3}

	fmt.Println("arr1=", arr1, "arr2=", arr2, "arr1[2]=", arr1[2], "arr3=", arr3, "arr4=", arr4)

	namelist := [5]string{
		"name1",
		"name2",
		"name3",
	}
	for index := 0; index < len(namelist); index = index + 1 {
		fmt.Println("namelist[", index, "]=", namelist[index])
	}

	fmt.Println("Initial array values", namelist)

	Array1PlayCallByValue(namelist)
	fmt.Println("Call by Value Completed, Back to Initial values =", namelist)

	Array1PlayCallByReference(namelist[0:len(namelist)])
	fmt.Println("Call by Reference Completed, initial value changed due to use of splice=", namelist)

	Array1PlayCallByAddress(&namelist[0])
	fmt.Println("Call by Address Completed, initial value changed due to use of address=", namelist)

	slice := namelist[2:3]
	for index := 0; index < len(slice); index = index + 1 {
		fmt.Println("slice[", index, "]=", slice[index])
	}

	slice[0] = "name22"
	fmt.Println("slice[0] =", slice[0], "namelist[2] =", namelist[2])

	//var arrayed [3]string = [3]string{ "ABC", "EFG", "HIJ"}
	//dstCopy := arrayed[1:2]
	dstCopy := make([]string, 1, 10)

	//dstCopy := arrayed[1:1]	==> length is changed to 0
	//fmt.Println("arrayed =", arrayed, "length = ", len(arrayed), "capacity = ", cap(arrayed))

	fmt.Println("dstCopy =", dstCopy, "length = ", len(dstCopy), "capacity = ", cap(dstCopy))
	fmt.Println("slice =", slice, "length = ", len(slice), "capacity = ", cap(slice))
	length := copy(dstCopy, slice)
	fmt.Println("dstCopy =", dstCopy, "length = ", len(slice), "capacity = ", cap(slice), " copiedLength = ", length)

}

func PointerPlay() {
	var tmp int = 1
	var ptr *int
	fmt.Println("Initial Temporary ", tmp)
	ptr = &tmp
	*ptr = 20
	fmt.Println("Revised Temporary ", tmp, " *ptr ", *ptr)
	return
}

func StructurePlay() {
	type COORDINATE struct {
		X int
		Y int
		Z int
	}

	var (
		tmp1 COORDINATE  = COORDINATE{}
		tmp2 COORDINATE  = COORDINATE{X: 1, Z: 20}
		tmp3 *COORDINATE = &tmp1
		tmp4             = &COORDINATE{1, 2, 3}
	)

	fmt.Println("tmp1=", tmp1, "tmp2=", tmp2, "tmp2::Z=", tmp2.Z, "tmp3=", tmp3, "tmp4=", tmp4, "tmp4::Z=", tmp4.Z)

	// tmp5 := tmp1
	// tmp5.X = 10
	// tmp1.X = 100
	// fmt.Println("tmp5: ", tmp5.X, ", tmp1: ", tmp1.X)
}

func ShaPlay() {
	// stringlist := []string{"h", "e"}
	// bufferfields := &bytes.Buffer{}
	// gob.NewEncoder(bufferfields).Encode(stringlist)
	// shahash := sha256.Sum256([]byte(bufferfields.Bytes()))
	// fmt.Println("shahash256: ", hex.EncodeToString(shahash))

	// stringlist = []string{"h", "", "e"}
	// bufferfields = &bytes.Buffer{}
	// gob.NewEncoder(bufferfields).Encode(stringlist)
	// shahash = sha256.Sum256([]byte(bufferfields.Bytes()))
	// fmt.Println("shahash256: ", hex.EncodeToString(shahash))

	type pair struct {
		value string
	}
	case1Json := pair{value: "hello"}
	case1Bytes, _ := json.Marshal(case1Json)
	case1Sha := sha256.Sum256([]byte(case1Bytes))
	case1string := fmt.Sprintf("%x", case1Sha)
	fmt.Println(case1string)

	case1md5 := md5.New()
	case1md5.Write(case1Bytes)
	case1string = fmt.Sprintf("%x", case1md5.Sum(nil))
	fmt.Println("1-case1string:", case1string)

	case1string = fmt.Sprintf("%x", md5.Sum(case1Bytes))
	fmt.Println("2-case1string:", case1string)

}
