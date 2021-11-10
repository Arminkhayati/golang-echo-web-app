package resmodel


import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"log"
	"github.com/labstack/echo"
)

type UserResp struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	RealUserInfo   realUser
	UserAddress	   address
	UserContact    []contact
	UserProfile	   profile
	jwt.StandardClaims
}

type realUser struct{
	FirstName	   string	 `json:"firstName"`
	LastName	   string	 `json:"lastName"`
	FatherName	   string	 `json:"fatherName"`
	Inic		   string	 `json:"inic"`
	InicFrontImg   uuid.UUID `json:"inicFrontImg"`
	InicBackImg	   uuid.UUID `json:"inicBackImg"`
	confirmed	   string    `json:"confirmed"`
}

type address struct{
	Id             uuid.UUID `json:"id"`
	Zip            string	 `json:"zip"`
	Country        string	 `json:"country"`
	State          string    `json:"state"`
	City		   string	 `json:"city"`
	OtherAdd	   string    `json:"otherAdd"`
}
type contact struct {
	ContactType	   string	 `json:"contactType"`
	ContactValue   string	 `json:"contactValue"`
}

type profile struct {
	Username	   string	 `json:"username"`
	Biography	   string	 `json:"biography"`
		ImgId		   string    `json:"imgId"`
} 


func (usr *UserResp)CreateJwtToken()(token string,err error){
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512,usr)
	token, err = rawToken.SignedString([]byte("ArminRoozbehRezaBackendShireFrontChaiiGhahvehArianChaiiAvordBaCakeAbJooshNadashtim1234462194124628913746912846198246289713462389146298461238974621987461289746128946128746298146284564512874581273452378451238745123874521409650858585901786"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func GetUserID(c echo.Context ) uuid.UUID {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*UserResp)
	userId := claims.ID
	return userId
}
