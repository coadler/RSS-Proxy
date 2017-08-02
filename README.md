note: this was made for [Tatsumaki](https://tatsumaki.xyz) to proxy RSS requests, but in reality it will grab any page and return it so do as you wish ¯\\_(ツ)_/¯

# RSS-Proxy
A simple webserver in Go that can proxy RSS requests through an easy to use API without sacrificing performance. 


# Installation
It is very simple to setup. 
1. Download and install [Go](https://golang.org/) or grab a compiled binary from Thy.  
2. Clone this repository
3. `cd` into the folder the repo was cloned to
4. Run `go build` to compile a binary for your operating system
5. Type `./RSS-Proxy` or run it through your favorite process manager like so: `pm2 start ./RSS-Proxy`
6. The program will indicate when it is ready to serve requests, and you're off to the races!

If you grabbed a compiled binary from Thy, you only need to do step 5

# how do i use this thing
After you've completed the above steps, the binary will begin listening on port 80 on the host device. If you need it to listen on a different port, you can easily change that in the `main.go` file under the variable `port`.

If you're running this on your own PC, type `http://localhost` into your browser to verify that it is running.  

The main route is `/v1/get` (which is poorly named imo but i can't change it now :D)  
To use it, send a POST request to that route with the data
```JSON
{
    "URL": "link to rss page"
}
```
The API will then return the RSS document without exposing the caller's IP address

feel free to copy/use this anywhere :D