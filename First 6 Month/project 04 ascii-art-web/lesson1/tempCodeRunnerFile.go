// server working 
	mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":8080",mux))