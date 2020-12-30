package day4

import (
    "testing"
)
  
  func TestValidBirthYearMin(t *testing.T) {
    
    validCount := validateRules("1920", 4, 1920, 2002)
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidBirthYearBelowMin(t *testing.T) {
    
    validCount := validateRules("1900", 4, 1920, 2002)
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidBirthYearMax(t *testing.T) {
    
    validCount := validateRules("2002", 4, 1920, 2002)
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidBirthYearAboveMax(t *testing.T) {
    
    validCount := validateRules("2003", 4, 1920, 2002)
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

  func TestValidHeightMinCm(t *testing.T) {
    
    validCount := validateHeight("150cm")
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightBelowMinCm(t *testing.T) {
    
    validCount := validateHeight("100cm")
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightMaxCm(t *testing.T) {
    
    validCount := validateHeight("193cm")
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightAboveMaxCm(t *testing.T) {
    
    validCount := validateHeight("200cm")
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

  func TestValidHeightMinIn(t *testing.T) {
    
    validCount := validateHeight("59in")
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightBelowMinIn(t *testing.T) {
    
    validCount := validateHeight("50in")
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightMaxIn(t *testing.T) {
    
    validCount := validateHeight("76in")
    
    if !validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightAboveMaxIn(t *testing.T) {
    
    validCount := validateHeight("100in")
    
    if validCount {
        t.Fatalf("Count does not match")
    }
  }

func TestValidHeightRandomString(t *testing.T) {

  validCount := validateHeight("vghkcghj")

  if validCount {
      t.Fatalf("Count does not match")
  }
}

func TestInValidHairColor(t *testing.T) {

  validCount := validateHairColor("vghkcghj")

  if validCount {
      t.Fatalf("Count does not match")
  }
}

func TestValidHairColor(t *testing.T) {

  validCount := validateHairColor("#888785")

  if !validCount {
      t.Fatalf("Count does not match")
  }
}

func TestValidEyeColor(t *testing.T) {

  searchEyes := [7] string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
  
  for _, eye := range searchEyes {
  validCount := validateEyeColor(eye)

    if !validCount {
        t.Fatalf("Count does not match %s", eye)
    }
  }
}

func TestInValidEyeColor(t *testing.T) {

  validCount := validateEyeColor("vghjcghj")

  if validCount {
      t.Fatalf("Count does not match")
  }
}

func TestInValidPID(t *testing.T) {

  validCount := validatePID("456745674567")

  if validCount {
      t.Fatalf("Count does not match")
  }
}

func TestValidPID(t *testing.T) {

  validCount := validatePID("123456789")

  if !validCount {
      t.Fatalf("Count does not match")
  }
}

func TestValidPIDLeadingZero(t *testing.T) {

  validCount := validatePID("023456789")

  if !validCount {
      t.Fatalf("Count does not match")
  }
}