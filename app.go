package main
 
import (
	"net/http"
	"./server"
)

func main(){
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	http.Handle(server.SetDirToUrl, http.StripPrefix(server.SetDirToUrl, http.FileServer(http.Dir(server.SetDir))))

	http.HandleFunc("/treefs", 	server.ScanDirTree)
	http.HandleFunc("/readfile", 	server.ReadFile)
	http.HandleFunc("/writefile", 	server.WriteFile)	
	http.HandleFunc("/getconfig", 	server.GetConfig)	
	
    print("已建立服务：http://localhost") 
	http.ListenAndServe(":80", nil)
}


