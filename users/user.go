package users

import (
	"fmt"
	"github/rigel-developer/golang-course/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.Usuario)
	u.AddUser(1, "Rigel", time.Now(), true)
	fmt.Println(u)

}
