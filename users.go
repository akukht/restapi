package main

type Users struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	TimeZone string `json:"timezone"`
	Token    string `json:"token"`
}

var UsersList = map[string]Users{}

func init() {
	UsersList["1"] = Users{
		Login:    "admin",
		Password: "$2a$14$ZOubd0goKj9Dhfkgd3GsPOZfAHkvGG/0ih8zkSx0.bI1JmbJljSNe", //1
		TimeZone: "Europe/Rome",
	}
	UsersList["2"] = Users{
		Login:    "user1",
		Password: "$2a$14$lTvBiykMlyypDIzsx1bAW.3uTIaImgdW.v10Nd3T6at41R6Jkm3PC", //2
		TimeZone: "Europe/Kiev",
	}
	UsersList["3"] = Users{
		Login:    "user2",
		Password: "$2a$14$lTvBiykMlyypDIzsx1bAW.3uTIaImgdW.v10Nd3T6at41R6Jkm3PC", //2
		TimeZone: "Europe/Kiev",
	}
	UsersList["4"] = Users{
		Login:    "user3",
		Password: "$2a$14$lTvBiykMlyypDIzsx1bAW.3uTIaImgdW.v10Nd3T6at41R6Jkm3PC", //2
		TimeZone: "Pacific/Chuuk",
	}
}
