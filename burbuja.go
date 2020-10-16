package main

func Burbuja(s []int64)  {
	var aux int
  l := len(arreglo)
  for i:=0;i<l;i++{
    for j:=0;j<l-1;j++{
      if arreglo[j] > arreglo[j+1]{
        aux = arreglo[j]
        arreglo[j]=arreglo[j+1]
        arreglo[j+1]=aux
      }
    }
  }
  fmt.Println(arreglo)
}

func main()  {
	
}