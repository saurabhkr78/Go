// break,conntine

//nested loop


/*
fmt.Println → prints arguments with spaces automatically.

fmt.Printf → lets you use format specifiers (%d, %s, etc.).

*/
package main
import(
	"fmt"
)


func main(){

for i:=1;i<10;i++{
	for j:=1;j<10;j++{
		fmt.Printf("%d x %d =%d\n",i,j,i*j)
	}
}


for i:=1;i<=10;i++{
	fmt.Printf("%d x %d=%d\n",5,i,5*i);
}

for i:=1;i<=20;i++{
	if i%2==0{
		fmt.Printf("%d\n",i)
	}
}


for i:=1;i<=5;i++{
	for j:=1;j<=i;j++{
		fmt.Print("*")
	}
	fmt.Println()
}

}
