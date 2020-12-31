package day5

import(
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
)

func StartDay5(){
  data, err := ioutil.ReadFile("./day5/input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
  seats := strings.Split(string(data),"\n")
  ProcessDay5(seats)
}

func ProcessDay5(seats []string) (val int){
  seatID := 0
  topSeatID := 0
  var allSeats []int
  var planeRows []int
  var planeCols = []int {0,1,2,3,4,5,6,7}

  for i := 0; i < 128; i++{
    planeRows = append(planeRows, i)
  }
  
  for _, seatItem := range seats {
    possibleRows := planeRows
    possibleCols := planeCols
    
    for _, seatChar := range seatItem {
    
      if(string(seatChar) == "F"){
        middleKey := len(possibleRows)/2
        possibleRows = possibleRows[0:middleKey]
      }else if(string(seatChar) == "B"){
        middleKey := len(possibleRows)/2
        possibleRows = possibleRows[middleKey:len(possibleRows)]
      }else if(string(seatChar) == "R"){
        middleKey := len(possibleCols)/2
        possibleCols = possibleCols[middleKey:len(possibleCols)]
      }else if(string(seatChar) == "L"){
        middleKey := len(possibleCols)/2
        possibleCols = possibleCols[0:middleKey]
      }
    }
    
    seatID = possibleRows[0] * 8 + possibleCols[0]
    if(seatID > topSeatID){
      topSeatID = seatID
    }
    
    allSeats = append(allSeats, seatID)
  }
  sort.Ints(allSeats)
  fmt.Println(allSeats)
  
  
  for i := 1; i < len(allSeats); i++{
    diff := allSeats[i] - allSeats[i-1]
    if(diff != 1){
      fmt.Println("Missing seat between ", allSeats[i], " and ", allSeats[i-1])
      break
    }
  }
  
  return seatID
}