package Services

import (
	"context"
	"fmt"
	"strconv"

	"TES/utils"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/time/rate"
)

type UserRequest struct { //封装User请求结构体
	Uid    int `json:"uid"`
	Method string
	Token  string //新加的token字段，用于读取url中的token封装进来再传递给下一层的请求处理
}

type UserResponse struct {
	Result string `json:"result"`
}

//token验证中间件
func CheckTokenMiddleware() endpoint.Middleware { //Middleware type Middleware func(Endpoint) Endpoint
	return func(next endpoint.Endpoint) endpoint.Endpoint { //Endpoint type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest) //通过类型断言获取请求结构体
			uc := UserClaim{}
			//下面的r.Token是在代码DecodeUserRequest那里封装进去的
			getToken, err := jwt.ParseWithClaims(r.Token, &uc, func(token *jwt.Token) (i interface{}, e error) {
				return []byte(secKey), err
			})
			fmt.Println(err, 123)
			if getToken != nil && getToken.Valid { //验证通过
				newCtx := context.WithValue(ctx, "LoginUser", getToken.Claims.(*UserClaim).Uname)
				return next(newCtx, request)
			} else {
				return nil, utils.NewMyError(403, "error token")
			}

			//logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)

		}
	}
}

//日志中间件,每一个service都应该有自己的日志中间件
func UserServiceLogMiddleware(logger log.Logger) endpoint.Middleware { //Middleware type Middleware func(Endpoint) Endpoint
	return func(next endpoint.Endpoint) endpoint.Endpoint { //Endpoint type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest) //通过类型断言获取请求结构体
			logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)
			return next(ctx, request)
		}
	}
}

//加入限流功能中间件
func RateLimit(limit *rate.Limiter) endpoint.Middleware { //Middleware type Middleware func(Endpoint) Endpoint
	return func(next endpoint.Endpoint) endpoint.Endpoint { //Endpoint type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				// return nil, errors.New("too many request")
				// 统一异常处理
				return nil, utils.NewMyError(429, "toot many request") //使用我们自定的错误结构体
			}
			return next(ctx, request)
		}
	}
}

func GenUserEnPoint(userService IUserService) endpoint.Endpoint {
	// return func(ctx context.Context, request interface{}) (response interface{}, err error) {
	// 	r := request.(UserRequest)           //通过类型断言获取请求结构体
	// 	result := userService.GetName(r.Uid) //
	// 	return UserResponse{Result: result}, nil
	// }
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// var logger log.Logger
		// {
		// 	logger = log.NewLogfmtLogger(os.Stdout)
		// 	logger = log.WithPrefix(logger, "mykit", "1.0")
		// 	logger = log.WithPrefix(logger, "time", log.DefaultTimestampUTC) //加上前缀时间
		// 	logger = log.WithPrefix(logger, "caller", log.DefaultCaller)     //加上前缀，日志输出时的文件和第几行代码
		// }

		r := request.(UserRequest) //通过类型断言获取请求结构体
		fmt.Println("当前登录用户为", ctx.Value("LoginUser"))
		result := "nothings"
		if r.Method == "GET" { //通过判断请求方法走不通的处理方法
			//result = userService.GetName(r.Uid)
			result = userService.GetName(r.Uid) + strconv.Itoa(utils.ServicePort)
			//logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)
			//fmt.Println(result)
		} else if r.Method == "DELETE" {
			err := userService.DelUser(r.Uid)
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid为%d的用户已删除", r.Uid)
			}
		}
		return UserResponse{Result: result}, nil
	}
}
