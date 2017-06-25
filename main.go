package sonycrapi

var liveviewFrames = make(chan *LiveviewPayload, 60)

// func main() {
// 	DiscoverDevice()
// 	newRequest("startRecMode").Do()
// 	mode, err := GetShootMode()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("Mode:", mode)

// 	go broadcastLiveview()

// 	fs := http.FileServer(http.Dir("client"))
// 	http.Handle("/", fs)
// 	http.HandleFunc("/liveview", liveviewHTTPHandler)

// 	log.Println("Listening...")
// 	http.ListenAndServe(":3000", nil)
// }

// func broadcastLiveview() {
// 	go fillFrameBuffers()

// 	StopLiveview()
// 	lv, err := StartLiveviewWithSize(LiveviewSizes.M)
// 	defer lv.Stop()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	lv.Decode(liveviewFrames)
// }
