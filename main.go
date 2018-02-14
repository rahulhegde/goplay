package main

import "fmt"
import (
	"stringutil"
	"math"
	"math/rand"
	"runtime"
	"time"
)

import (
	"golang.org/x/tour/pic"
	"encoding/json"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)


func Sqrt (value float64) (string, string) {
	return fmt.Sprint(math.Sqrt(value)), fmt.Sprint(math.Sqrt(value))
}

const STARTER=100
const DELTA = 0.0001

func NewtonSquareValue(point float64, sqrt int) ( return64 float64) {
	return64 = (point * point + float64(sqrt)) / (2 * point)
	return
}

func NewtonSquare (sqrt int) (float64) {
	prevSqrtValue := NewtonSquareValue(STARTER, sqrt)
	newSqrtValue := NewtonSquareValue(prevSqrtValue, sqrt)
	for math.Abs(newSqrtValue - prevSqrtValue) > DELTA {
		prevSqrtValue = newSqrtValue;
		newSqrtValue = NewtonSquareValue(prevSqrtValue, sqrt)
	}
	return newSqrtValue;
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



	fmt.Println("day = ", time.Now().Day());
	fmt.Println(time.September)

}

func switchUsage(x string) {
	fmt.Println("Entering switch check for OS", runtime.GOOS);

	switch x {
	case "string":
		fmt.Println("switch::string");
	case "sting":
		fmt.Println("switch::sting")
	default:
		fmt.Println("switch::default")
	}
}



func main() {
	fmt.Println("%q", stringutil.Reverse("hello, world rahul") )
	fmt.Println("random = ", rand.Intn(100))


	var i1 = 3

	for i,j := 0,20; i < 1; i,j = i+1,j-1 {
		fmt.Printf("random = %d, %d \n", i, j)
	}

	if i1 = 2; i1 > 4 {
		fmt.Print("help...4")

	} else if i1 > 2 {
		fmt.Print("help...2")
	} else {
		fmt.Print("help...1")
	}

	a, b := Sqrt(4);
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
	//InterfaceTypeAssertionPlay()
	//TypeSwitchPlay()

	//MapPlay()
	//DecimalPlay()
	//Float64Play()
	//TimePlay()
	//
	//BitwisePlay()
	//
	//GoRoutinePlay()
	//GoChannelPlay()

	BitwiseCheckPlay()
}


type BOOLEAN int

const (
	TRUE BOOLEAN = 1
	FALSE BOOLEAN = 4
)

func (state *BOOLEAN) setTRUE () {
	*state = *state | TRUE
}

func (state *BOOLEAN) isTRUE () bool {

	status := *state & TRUE
	fmt.Println("isTRUE = ", *state, status)
	return (*state & TRUE)==TRUE
}

func (state *BOOLEAN) isFALSE () bool {
	status := *state & FALSE
	fmt.Println("isFALSE = ", *state, status)
	return (*state & FALSE)==FALSE
}

func (state *BOOLEAN) setFALSE ()  {
	*state = *state | FALSE
}

func BitwisePlay() {
	var state BOOLEAN
	state = TRUE
	fmt.Println("state after setting to TRUE: ", state)
	state.setFALSE()
	fmt.Println("state after setting to FALSE: ", state)
	fmt.Println("isFALSE = ", state.isFALSE(), "isTRUE: ", state.isTRUE())

}

type marble struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Size       int 	  `json:"size"`
	Owner      string `json:"owner"`
}

func JSONPlayTest3() {

/*	"JSONPlayTest3: Unmarshalling to Generic Interface however using NewDecoder/UseNumber API \n" +
		"This will ensure float-pointing representation done by the JSON package for integer/json-number \n" +
	            "at the time of internal storage is reverted to original format."*/
	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`
	jsonResult := make(map[string]interface{})

	decoder := json.NewDecoder(bytes.NewBuffer([]byte(myJSON)))
	decoder.UseNumber()
	err := decoder.Decode(&jsonResult);
	if err != nil && jsonResult["data"] == nil {
		fmt.Println("Nothing found")
	} else {

		fmt.Println("Unmarshalled to Genric Structure using NewDecoder API: ", jsonResult["data"])
	}
}

func JSONPlayTest2()  {
	/*"JSONPlayTest3: Unmarshalling to Generic Interface however using Unmarshall API \n" +
		"This unmarshalling causes problem when used with Generic Interface like below for integer \n" +
		"digits size >= 7. Integer value is represented as with E notation and this will fail when " +
		"unmarshalling to integer");
	*/
	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`

	jsonResult := make(map[string]interface{})
	json.Unmarshal([]byte(myJSON), &jsonResult)

	if jsonResult["data"] == nil {
		fmt.Println("Nothing found")
	} else {

		fmt.Println("Unmarshalled to Generic Interface: ", jsonResult["data"])
	}
}


func JSONPlayTest1() {
/*
	Marshalling to a known structure is not a problem only to generic interface.
 */

	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`

	marble1 := make(map[string]marble)
	err := json.Unmarshal([]byte(myJSON), &marble1)
	if err != nil {
		fmt.Println(err)
	}

	marblesBytes, err := json.Marshal(marble1)
	if err != nil {
		fmt.Println(err)
	}

	marble2 := make(map[string]marble)
	err = json.Unmarshal(marblesBytes, &marble2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshalled to know structure: ", marble2["data"])

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

func JSONPlay1() {
	type App struct {
		Target1 struct 	{
			Id    string    `json:"id"`
			Title string 	`json:"title"`
		} `json: "Target1"`
		Target2 struct {
				Test    string   `json:"CurrencyIso"`
			Title string    `json:"title"`
		} `json: "Target2"`
	}

	data := []byte(`
		{
			"Target1":
				{ "id":"1000", "title":"Bookish"},
			"Target2":
				{ "CurrencyISO":"10010", "title":"Bookish"}
		}
	`)

	var app App
	_ = json.Unmarshal(data, &app)
	fmt.Print("App: ", app)

	var out bytes.Buffer
	byteData, _  := json.Marshal(app.Target2)

	json.Indent(&out, byteData, "", "-")
	fmt.Println("Marshalled ", out.String())
}

func PointerPlay2() {
}


type AXIX struct {
	x int;
	y int;
}
func  (x AXIX) PrinterAXIS(a,b int)  {
	fmt.Println("PrinterAXIS: x= ", x.x, " y= ", x.y, "a=",a, "b=", b)
}

type FLOAT64 float64

func (f FLOAT64) floater (value float64) {
	fmt.Println("floater=", f, "value=", value)
}



func ReceiverMethodPlay() {
	stringutil.AXIS{}.Printer();
	// call not possible for built-in data type in other package
	//stringutil.STRFLOAT64.Printer(12.2)

	FLOAT64(20.02).floater(10.20)

	x := AXIX{1, 2}
	x.PrinterAXIS(11,22)
	AXIX{1,10}.PrinterAXIS(5,50)

}

func MapPlay() () {
	type Vertex struct {
		X int
		Y int
	}

	var mapperstub map[int]string = map[int]string {
		1 : "name1",
		2 : "name2",
	}
	mapperstub[3] = "name3"
	fmt.Println("mapperstub", mapperstub)

	//var mapperstub1 map[int]string
	mapperstub1 := make (map[int]string)
	mapperstub1[1] = "name2"
	value, exist := mapperstub1[1];
	fmt.Println("mapperstub1", mapperstub1, "exist =", exist, "value = ",value)
	delete(mapperstub1, 1)
	value, exist = mapperstub1[1];
	fmt.Println("mapperstub1", mapperstub1, "exist =", exist, "value = ",value)
}

func ExerciseSliceCreate2DDynamicImage(dx, dy int) (arraystub [][]uint8) {
	//dx, dy := 256, 256

	//
	arraystub = make ([][]uint8, dx)
	for indexX :=0; indexX < dx; indexX = indexX + 1 {
		arraystub[indexX] = make ([]uint8, dy)
		for indexY := 0; indexY < dy; indexY = indexY + 1 {
			arraystub [indexX][indexY] = uint8(indexX + indexX)
		}
	}
	return
}

func RangeForArray() {
	arraystub := []int {1,2,3,4}

	for index, value := range arraystub {
		fmt.Println("Index=", index, "value=", value)
	}
}

func D2ArrayPlay() {
	var array2d [][]string = [][]string {
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
	var arraystub []int = []int {1,2,3,4,5}

	slicer := arraystub[:]
	fmt.Println("slicer1=",slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[:3]
	fmt.Println("slicer2=",slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[2:]
	fmt.Println("slicer3=",slicer, "len=", len(slicer), "cap=", cap(slicer))

	slicer = arraystub[:]
	fmt.Println("slicer4=",slicer, "len=", len(slicer), "cap=", cap(slicer))


	var slicer1 []int
	if slicer1 == nil {
		fmt.Println("slicer1 is nil=", slicer1, " cap=", cap(slicer1), " len=", len(slicer1))
	}
}

func SliceLiteralPlay() () {
	slicer := []int {1,2,3}
	var arraystub []int = []int{1,2,3}
	fmt.Println("slicer=", slicer, "arraystub=", arraystub)
	return
}

func ArrayStructurePlay() () {

	type COORDINATE struct {
		X int
		Y int
	}

	// scenario to create array of structure
	var arrcoordinates [3]COORDINATE = [3]COORDINATE {
		COORDINATE{1,2},
		COORDINATE{3,5},
		COORDINATE{4,5},
	}
	fmt.Println("array of coordinates structure =", arrcoordinates);

	slicer1 := arrcoordinates[2:2]
	fmt.Println("array of structured sliced = ", slicer1);

	// scenario where array is constructed first and then slice is created from it.
	slicer := []struct {
		X int
		Y string
	}{
		{1, "hello"},
		{2,"world"},
	}
	fmt.Println("slicer constructed w/o array creation = ", slicer)


	return;
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

	var arr1 [5]int = [5]int {1,2,3}
	var arr2 [5]int = [5]int {}
	arr3 := COORDINATE{1,2,3}
	arr4 := [5]int{1,2,3}

	fmt.Println("arr1=", arr1, "arr2=", arr2, "arr1[2]=", arr1[2],"arr3=", arr3, "arr4=", arr4)


	namelist := [5]string {
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
		tmp1 COORDINATE = COORDINATE{}
		tmp2 COORDINATE = COORDINATE{X:1, Z:20}
		tmp3 *COORDINATE = &tmp1
		tmp4 = &COORDINATE{1,2,3}
	)
	fmt.Println("tmp1=", tmp1, "tmp2=", tmp2, "tmp2::Z=", tmp2.Z, "tmp3=", tmp3, "tmp4=", tmp4, "tmp4::Z=", tmp4.Z)
}