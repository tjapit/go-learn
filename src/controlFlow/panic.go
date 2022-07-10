package main

import (
	"net/http"
)

/* SUMMARY
 * Panic
 * - Occur when program cannot continue at all
 *   - Don't use when file can't be opened, unless critical
 *   - Use for unrecoverable events -- cannot obtain TCP port for web server
 * - Function will stop executing
 *   - Deferred functions will still fire
 * - If nothing handles panic, program will exit
 */
func main() {
	/* "panic"
	 * "panic" happens after "defer"
	 * e.g. multiple servers binding same port, normally not allowed
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oh hi"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// trigger by starting the server twice
		panic(err.Error())
	}
}
