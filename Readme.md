Docker file for GO


func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hello world")
 })
 port := os.Getenv("PORT")
 if port == "" {
  port = "8080"
 }
 log.Fatal(http.ListenAndServe(":"+port, nil))
}

https://medium.com/google-cloud/google-cloud-run-for-go-ec09ddbba111