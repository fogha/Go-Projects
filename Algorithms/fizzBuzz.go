package main;
import ("fmt");

func main() {
	var n int = 100;
	var i int;
	for i = 1; i <= n; i++ {
		if i % 3 == 0 {
			if i%5== 0 {
				fmt.Println("fizzBuzz")
			} else {
				fmt.Println("fizz");
			}
		} else if i % 5 == 0 {
			if i%3== 0 {
				fmt.Println("fizzBuzz")
			} else {
				fmt.Println("buzz");
			}
		} else {
			fmt.Println(i);
		}
	}
}