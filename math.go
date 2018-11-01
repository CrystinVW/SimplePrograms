package main

import (
  "fmt"
  "math
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	
	return lim
}

func main() {
  fmt.Println("This is a simple example of a for loop!")
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println("This is the sum: ", sum)
  
  
  sum := 1
  fmt.Println("This is a while loop!")
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum) 
  
  fmt.Println("This calls for the square root of a number from function sqrt")
  fmt.Println(sqrt(2), sqrt(-4))
  
  fmt.Println("This calls for the power of a number from function pow")
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
  
  fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

