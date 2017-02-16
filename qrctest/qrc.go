package qrctest  
import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"log"
	"os"
)

func writePng(filename string, img image.Image) {
	file,err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err1 := png.Encode(file,img)
	if err1 != nil {
		log.Fatal(err1)
	}
	file.Close()
	log.Println(file.Name())

}

func QrcChange(url string,width int,height int,result string) {
	base64 := url
	log.Println("Oringinal data:", base64)
	code,err := qr.Encode(base64,qr.L,qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encode data:",code.Content())
	if base64 != code.Content() {
		log.Fatal("data diffres")
	}
	code,err = barcode.Scale(code,width,height)
	if err != nil {
		log.Fatal(err)
	}
	writePng(result,code)
}
