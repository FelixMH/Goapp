package main

func main() {
	/*
		                SERVIDOR TRADICIONAL CON GO
			// http.Handle("/", http.FileServer(http.Dir("public")))
			// log.Println("Ejecutando server en http://localhost:8080")
			// log.Println(http.ListenAndServe(":8080", nil))
	*/

	/*
	       SERVIDOR TRADICIONAL CON MUX.
	   mux := http.NewServeMux()
	   fs := http.FileServer(http.Dir("public"))
	   mux.Handle("/", fs)
	   log.Println("Ejecutando server en http://localhost:8080")
	   log.Println(http.ListenAndServe(":8080", mux))
	*/
}
