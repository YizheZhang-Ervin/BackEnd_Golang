func (webhandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	auth := request.Header.Get("Authorization") //获取Basic base64加密后的字段
	if auth == "" {
		writer.Header().Set("WWW-Authenticate", `Basic realm="您必须输入用户名和密码"`)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	auth_list := strings.Split(auth, " ")
	if len(auth_list) == 2 && auth_list[0] == "Basic" {
		res, err := base64.StdEncoding.DecodeString(auth_list[1])
		if err == nil && string(res) == "shenyi:123" { //输入的用户名和密码会用:隔开
			writer.Write([]byte("<h1>web1</h1>"))
			return
		}
	}
	writer.Write([]byte("用户名密码错误"))
}