package main

func main() {
	n := 150
	records := make([]interface{}, n)
	for i := 0; i < n; i++ {
		records[i] = i
	}
	Test(records)
}

func Test(records []interface{}) {
	for len(records) > 0 {
		record := records
		if len(records) >= 100 {
			record = records[:100]
			records = records[100:]
		} else {
			records = nil
		}
	}
}
