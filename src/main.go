package main
import (
	"fmt"
	"github.com/magicon-top/go-pkg/dll"
	)

func main() {	// Теперь мы вызываем функцию через имя пакета: dll.Dll
	_, err := dll.Dll("asm-Splash-dll.dll", "ShowSplash", 3000)
	if err != nil { fmt.Println("Ошибка при вызове DLL:", err); return}
}
