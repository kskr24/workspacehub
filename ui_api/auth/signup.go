package auth

import "github.com/kskr24/sajha/ui_api/web"

func signup(ctx web.Context){
	var req struct{
		Email string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name`
	}

	if err := ctx.Bind(&req); err != nil{
		
	}
}