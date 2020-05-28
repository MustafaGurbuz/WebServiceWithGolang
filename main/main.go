package main

import (
	"fmt"
	"github/MustafaGurbuz/webservice/controllers"
	"github/MustafaGurbuz/webservice/models"
	"net/http"
)
func main()  {
	/*if number:=18; number<0 {
		fmt.Println(number," negatiftir...")
	}else if number>10 && number<20{
		fmt.Println(number, " 10 ile 20 arasındadır...")
	}else {
		fmt.Println(number,"Sayı beklenen sınırların dışındadır...")
	}
	fmt.Println("\n")
	toplama := 1
	for ; toplama < 1000; {
		toplama += toplama
	}
	fmt.Println(toplama)
	 */
	//------------ Kullanıcı Ekleme ---------
	u := models.User{
	  	ID:       2,
	  	FirstName: "Mustafa",
	  	LastName: "Gürbüz",
	  }
	  	fmt.Println(u)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	// ------ CONNECT TO LOCALHOST ------
	//-------- go build and webservice.exe ----------

}




