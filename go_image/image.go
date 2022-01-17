package go_image

import (
	"fmt"
	"image/jpeg"
	"os"
)

func ImageMain(){
	fp,err:=os.OpenFile("./abc.jpeg",os.O_RDONLY, 6)
	if err!=nil{
		fmt.Println(err)
	}

	//img, err := jpeg.Decode(fp) //解码
	//if err != nil {
	//	fmt.Println(err)
	//}

	c,err:=jpeg.DecodeConfig(fp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("width = ", c.Width)
	fmt.Println("height = ", c.Height)
}