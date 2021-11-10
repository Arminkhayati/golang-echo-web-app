set routes( end points or path) to groups

g.Get("/shit", handlre.shitHandler)


g.Post("/shit", handlre.shitHandler)








if !check {
		// creating jwt token
		token , err := createJwtToken()
		if err != nil{
			log.Println("Error Creating JWT token ", err)
			return c.String(http.StatusInternalServerError,"something went Wrong!")
		}

		return c.JSON(http.StatusOK,
			map[string]string{
				"message": "you are loged in!",
				"token" : token,
			})
	}
