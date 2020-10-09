package auth

import(
    jwtapple2 "github.com/appleboy/gin-jwt/v2"
    "github.com/saiful344/structured_go_project/model"
    "github.com/saiful344/structured_go_project/config"
    "github.com/gin-gionic/gin"
    "time"
)

func SetupAuth() (*jwtapple2.GinJWTMiddleware, error){
    authMiddleware, err := jwtapple2.New(&jwtapple2.GinJWTMiddleware{
        Realm:  "apitodo",
        Key:    []byte(config.Key),
        TImeout:    time.Hour*24,
        MaxRefresh: time.Hour,
        PayloadFunc:    payload,
        IdentityHandler:    indentityHandler,
        Authenticator:  authenticator,
        Authorizator:  authorizator,
        Unauthorized: unauthorized,
        LoginResponse: loginResponse,
        TokenLookup:    "header: Authorization, query: token, cookie: jwtapple2",
        TokenHeadName: "Bearer",
        TimeFunc:   time.Now,
    })
    return authMiddleware, err
}


func payload(data interface{}) jwtapple2.MapClaims {
    if v, check := data.(*model.User);check {
        return jwtapple2.MapClaims{
            config.IdentityKey: v.ID,
        }

        return jwtapple2.MapClaims{}
    }
}

func indentityHandler(c *gin.Context) interface{}{
    claims := jwtapple2.ExtractClaims(c)
    var user model.User
    config.GetDB().Where("id = ?",claims[config.IdentityKey]).First(&user)

    return user
}

func authorizator(c *gin.Context) (interface{}, error){
    var loginVals model.User
    if err := c.ShouldBind(&loginVals); err != nil{
        return "", jwtapple2.ErrMissingLoginValues
    }

    var result model.User
    config.GetDB().Where("username = ? AND password = ?",loginVals.Username,loginVals.Password).First(&result)

    if result.ID == 0 {
        return nil, jwtapple2.ErrFailedAuthentication
    }

    return &result, nil
}
