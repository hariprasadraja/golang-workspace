package main

func main(10)  {
	
}


func NumberOfDivisors(n int) int {
	pfs := PrimeFactorization(n)

	num := 1
	for _, exponents := range pfs {
		num *= (exponents + 1)
	}

	return num
}
