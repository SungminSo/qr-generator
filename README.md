# QR Code Generator
## QR 코드 생성 및 검증
### usage

port : 3506

#### for Server Run :
```
go run main.go
```

#### for Test Run :
```
go test .
```

#### for make docker :
``` 
make docker
```

<br />

#### Generate QR Code 
URI : ``` /qr/generate ```

Method : ``` POST ```

API Request 
 - Content
 
API Response
 - http status code
 - QR Code(256 x 256 size)
 - QR Content
 - QR Token(6-digit)
    
<br />

#### Verify QR Code
URI : ``` /qr/verify/:qrToken ```

Method : ``` GET ```

API Request
 - QRToken
 
API Response
 - http status code 
 - result message