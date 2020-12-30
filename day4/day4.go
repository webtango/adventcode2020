package day4

import (
  "strings"
  "fmt"
  "strconv"
  "regexp"
  "io/ioutil"
  )

var keys = [7]string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} 
func StartDay4(){
  
  data, err := ioutil.ReadFile("./day4/input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
  
  passports := strings.Split(string(data),"\n")
  count := ProcessPassports(passports)
  fmt.Println("Valid Count ", count)
}

func ProcessPassports(passports []string) (count int) {
  invalidPassports := 0
  totalPassports := 0
  totalEntry := ""
  validPassports := 0
  
  for i :=0; i < len(passports); i++ {
    if(len(passports[i]) != 0){
      totalEntry =  totalEntry + " " + passports[i]
    }else{
      totalPassports++
      foundFields := true
      for _, key := range keys {
        foundFields = strings.Contains(totalEntry, key)
        
        if(!foundFields){
          invalidPassports++
          break
        }
      }
      
      if(foundFields){
        if(validatePassport(totalEntry)){
          validPassports++ 
        }
      }
      
      totalEntry = ""
    }
  }
  
  return validPassports
}

func validatePassport(passData string) (val bool){
  passData = strings.TrimSpace(passData)
  
  sets := strings.Split(passData, " ")
  
  allFieldsValid := true
  
  for _, pairs := range sets {
    keyVal := strings.Split(pairs, ":")
    switch keyVal[0]{
      case "byr":
      allFieldsValid = validateRules(keyVal[1], 4, 1920, 2002)
      break
      case "iyr":
      allFieldsValid = validateRules(keyVal[1], 4, 2010, 2020)
      break
      case "eyr":
      allFieldsValid = validateRules(keyVal[1], 4, 2020, 2030)
      break
      case "hgt":
      allFieldsValid = validateHeight(keyVal[1])
      break
      case "hcl":
      allFieldsValid = validateHairColor(keyVal[1])
       break
       case "ecl":
        allFieldsValid = validateEyeColor(keyVal[1])
        break
      case "pid":
        allFieldsValid = validatePID(keyVal[1])
        break
    }
    
    if(!allFieldsValid){
      return false
    }
  }
  
  return true
}

func validatePID(value string) (val bool){
  if(len(value) != 9){
    return false
  }else{
    regexColor := "^[0-9]{9}$"
    re := regexp.MustCompile(regexColor)
    finds := re.FindAllString(value, 1)

    if(len(finds) == 1){
      return true
    }
    fmt.Println("invalid pid ", value)
    return false
  }
}

func validateEyeColor(value string) (val bool){
  searchEyes := [7] string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
  foundColor := false
  for _, eye := range searchEyes {
    if(eye == value){
      foundColor = true
    }
  }
  if(!foundColor){
    fmt.Println("invalid eye ", value)
  }
  return foundColor
}
      
func validateHeight(testVal string) (val bool){
  if(strings.HasSuffix(testVal, "cm")){
    numVal := testVal[:len(testVal)-2]
    num, _ := strconv.ParseInt(numVal, 10, 64)
    if(num >= 150 && num <= 193){
      return true
    }
  }else if(strings.HasSuffix(testVal, "in")){
    numVal := testVal[:len(testVal)-2]
    num, _:= strconv.ParseInt(numVal, 10, 64)
    if(num >= 59 && num <= 76){
      return true
    }
  }
  
  fmt.Println("invalid height ", testVal)
  return false
}

func validateHairColor(value string) (val bool){
  regexColor := "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
  re := regexp.MustCompile(regexColor)
  finds := re.FindAllString(value, 1)

  if(len(finds) == 1){
    return true
  }
  
  fmt.Println("invalid hair ", value)
  return false
}

func validateRules(value string, length int, min int64, max int64) (val bool) {
  if(len(value) != 4){
        return false
      }else{
        num, err := strconv.ParseInt(value, 10, 64)
        
        if(err != nil){
          fmt.Println("invalid date ", value)
          return false
        }else{
          if(num >= min && num <= max){
            return true
          }
        }
      }
  fmt.Println("invalid pid ", value)
  return false
}