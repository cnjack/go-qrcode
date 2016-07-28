#nightc-qrcode  
a tool website write by Go to create qrcode.  
## Auth  
jack  
## Email
h_7357#qq.com (replace # by @)  
## PowerBy  
BootStrap  
Jquery  
[martini](http://github.com/go-martini/martini)  
[martini-contrib](http://github.com/martini-contrib/render)  
[go-qrcode](http://github.com/skip2/go-qrcode)  
...  
## Design 
![design ](http://github.com/cnjack/go-qrcode/raw/master/doc/index.png?dir=0&filepath=doc%2Findex.png)  
## install  
go get github.com/cnjack/go-qrcode  
cd %GOPATH%/src/github.com/cnjack/go-qrcode(win) cd $GOPATH/src/github.com/cnjack/go-qrcode(mac or linux)  
cd install  
go build main.go  
run main.exe or ./main  
visit http://localhost:8080/(you can change the port at qrcode Line 18)  
## API  
visit http://localhost:port/api?qr_string=[qr_string]&size=[size]
to get the image of the qrcode  
## Demo  
[nightc-qrcode](http://qrcode.nightc.com)  