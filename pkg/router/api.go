package router

func RunAPI() {
	router := RouteCustom()

	_ = router.Run(":12001")
}
