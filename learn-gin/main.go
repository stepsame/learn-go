package main

import "learn-gin/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
