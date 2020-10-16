package main

func Burbuja(s []int64) {
	var aux int64
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
}

func main() {

}
