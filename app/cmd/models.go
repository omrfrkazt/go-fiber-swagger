package main

//burada uygulama icerisinde kullanilacak modelleri yazdik.
type(
	UserModel struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	LoginModel struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
