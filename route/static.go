package route

import "qkcode/boot/http"

func AddStaticRoute() {
	http.Router.Static("/static", "./Public/")
}
