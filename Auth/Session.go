package auth

import (
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func VerifyJWTAndGetClaim(req *http.Request) (claim *Claims, err error) {
	bearerToken := req.Header.Get("Authorization")
	bearer := strings.Split(bearerToken, " ")
	if len(bearer) != 2 {
		return nil, http.ErrContentLength
	}
	tokenstring := bearer[1]
	claim = &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, err
	})
	if err != nil {
		return claim, err
	}
	if token.Valid {
		return claim, nil
	}
	return claim, jwt.ErrSignatureInvalid
}
func CorsMiddleware(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	fmt.Println("options above")
	if r.Method == http.MethodOptions {
		fmt.Println("in options ")
		return true
	}
	return false

}
func getJWT(user MyDb.User) (JWT, error) {
	var jwtToken JWT
	claims := Claims{
		User_ID:    user.Id,
		Email:      user.Email,
		Role:       user.Role,
		Company_id: user.Company_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 40).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return jwtToken, err
	}
	jwtToken = JWT{
		StatusCode: 200,
		Token:      tokenString,
	}
	return jwtToken, err
}
