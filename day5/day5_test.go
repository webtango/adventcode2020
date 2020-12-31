package day5

import (
    "testing"
)
  
func TestValidSeat357(t *testing.T) {
    
    testData := []string {"FBFBBFFRLR"}
    seatValue := ProcessDay5(testData)
    
    if seatValue != 357 {
        t.Fatalf("invalid seat %v", seatValue)
    }
  }

  func TestValidSeat567(t *testing.T) {
    
    testData := []string {"BFFFBBFRRR"}
    seatValue := ProcessDay5(testData)
    
    if seatValue != 567 {
        t.Fatalf("invalid seat %v", seatValue)
    }
  }

func TestValidSeat119(t *testing.T) {
    
    testData := []string {"FFFBBBFRRR"}
    seatValue := ProcessDay5(testData)
    
    if seatValue != 119 {
        t.Fatalf("invalid seat %v", seatValue)
    }
  }

func TestValidSeat820(t *testing.T) {
    
    testData := []string {"BBFFBBFRLL"}
    seatValue := ProcessDay5(testData)
    
    if seatValue != 820 {
        t.Fatalf("invalid seat %v", seatValue)
    }
  }