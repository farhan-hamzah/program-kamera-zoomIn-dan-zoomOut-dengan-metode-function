// // Soal Kamera Dengan Function
package main
import "fmt"
func zoomIn(c, x, y int )(int, int){
	x = (x)*(c)
	y = y*c
	return x, y
}
func zoomOut(c, x, y int ) (int, int){
	x = x/c
	y = y/c
	x = -x
 	y = -y
 	return x, y
}
func main(){
	var a, b, c1 int
	fmt.Scan(&a, &b)
	fmt.Scan(&c1)
	if c1 > 0 {
		fmt.Print(zoomIn(c1, a, b))
	}else{
		fmt.Print(zoomOut(c1, a, b))
	}
}

