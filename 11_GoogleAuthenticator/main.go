package main

import (
	"crypto"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sec51/twofactor"
)

func main() {
	otp, err := twofactor.NewTOTP("congluc19297@gmail.com", "Rock Ship", crypto.SHA1, 6)

	b, err := otp.ToBytes()
	
	// fmt.Println(b)
	// fmt.Println("Chuyen ve string")

	// fmt.Println(string(b))
	str := string(b)
	// fmt.Println("Chuyen nguoc ve lai byte")
	_ = str
	// fmt.Println([]byte(str))

	if err != nil {
		log.Println(err)
	}
	// fmt.Println(otp)

	qrBytes, err := otp.QR()
	fmt.Println(otp.Secret())
	fmt.Println(len(otp.Secret()))

	// fmt.Println(qrBytes)
	_ = qrBytes
	if err != nil {
		log.Println(err)
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.FormValue("token")
		fmt.Println(token)

		if token != "" {
			err := otp.Validate(token)
			if err != nil {
				// fmt.Println("Lỗi rồi")
				log.Println(err.Error())
				return
			}
			// fmt.Println("Không có lỗi. Validate thành công")
		}

		// w.Write(`{'test':err}`, qrBytes)
		w.Write(qrBytes)
		// w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, qrBytes)))
	})
	http.ListenAndServe(":1234", r)
	// err = otp.Validate(USER_PROVIDED_TOKEN)
	// if err != nil {
	// 	log.Println(err)
	// }
	// if there is an error, then the authentication failed
	// if it succeeded, then store this information and do not display the QR code ever again.
}
