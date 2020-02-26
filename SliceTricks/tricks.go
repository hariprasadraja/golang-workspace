package main

func SafeDelete(slice interface{}, i int) {
	// slice = copy(slice[i:], slice[i+1:])
}

func Cut(slice interface{}, from, to int, safe bool) {
	// slice = append(primes[:1], primes[4:]...)

	// if safe {
	// 	slice = copy(slice[i:], slice[j:])
	// 	for k, n := (len(primes) - j + i), len(primes); k < n; k++ {
	// 		primes[k] = 0
	// 	}
	// }
}

func Delete(slice interface{}, index int, preserveOrder bool) {
	// if !preserveOrder {
	// 	slice[i], slice[len(primes)-1], slice = primes[len(primes)-1], 0, primes[:len(primes)-1]
	// }

	// slice := append(slice[:i], slice[i+1:]...)

	// 	copy(primes[i:], primes[i+1:])
	// primes[len(primes)-1], primes = 0, primes[:len(primes)-1]

}

func Filters(slice interface{}, filters ...func()) {

}

func PopFront(slice interface{}) {
	// x, a = slice[0], a[1:]
}

func PopBack() {
	// x, a = a[len(a)-1], a[:len(a)-1]
}

func PushFront() {
	// a = append([]int{a[len(a)-1]}, a...)
}

func PushBack() {}

func RightRotation() {
	// primes = append(primes[len(primes)-rotate:], primes[:len(primes)-rotate]...)

}

func LeftRotation() {
	// primes = append(primes[rotate:], primes[0:rotate]...)
}

func RemoveDuplicates() {
	// sort.Ints(primes)
	// j := 0
	// for i := 1; i < len(primes); i++ {
	// 	if primes[j] == primes[i] {
	// 		continue
	// 	}

	// 	j++
	// 	primes[j] = primes[i]
	// }

	// result := primes[:j+1]
}

func Reverse() {
	// for left, right := 0, len(primes)-1; left < right; left, right = left+1, right-1 {
	// 	primes[left], primes[right] = primes[right], primes[left]
	// }
}
