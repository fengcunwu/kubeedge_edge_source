package main

import  (
     "fmt"
     "github.com/go-resty/resty/v2"
     "os"
 )

func main(){
    client := resty.New()
    resp, err := client.R().EnableTrace().Get("http://192.168.226.142:9100/metrics")
    
    f,err := os.Create("file.txt")
    defer f.Close()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        _,err=f.Write(resp.Body())
    }
    defer f.Close()

  //  fmt.Println("hello")
  //  resp_post, err_post := client.R().SetFile("my_file", "/home/kubeedgenode/workspace/GoSpace/go_src/src/file.txt").Post("http://192.168.226.139:9091/metrics/job/some_job/instance/some_instance")
  //  if err_post != nil {
  //      fmt.Println(err_post.Error())
  //  } else {
  //      fmt.Println(resp_post)
  //  }
 }


