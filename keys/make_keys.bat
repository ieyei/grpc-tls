:: gRPC SSL/TLS certificate related file generation

set SERVER_CN=localhost

:: step. 1 RootCA preperations. 
:: 1) make private key for RootCA
:: 2) make self-signed certificate for RootCA
:: 1) RootCA 가 인증서를 발급해줄때 쓸 개인키를 생성한다
:: 2) RootCA 가 발급한 인증서를 검증할 RootCA 자신에 대한, 자신이 서명한 인증서를 만든다. (클라이언트가 사용할 것이다.)

openssl genrsa -aes256 -out rootca.key 2048
openssl req -new -x509 -days 365 -key rootca.key -out rootca.crt -subj "/CN=%SERVER_CN%"

:: step. 2 Server preperations
:: 1) make server private key
:: 2) make "certificate request" which server send to CA to make certificate
:: 3) CA make server certificate
:: 1) 서버용 개인키를 만든다. 
:: 2) certificate request 를 만든다. 이걸 CA 에게 주며 인증서를 만들어 달라고 한다. 
:: 3) CA 는 certificate request 를 받아서 CA 의 개인키를 이용하여 서버의 인증서를 만들어준다. 

openssl genrsa -aes256 -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=%SERVER_CN%"
openssl x509 -req -days 365 -in server.csr -CA rootca.crt -CAkey rootca.key -set_serial 01 -out server.crt

:: step. 3 Server preperations for gRPC
:: 1) change server private key format
:: 1) 서버의 개인키 포맷을 gRPC 가 쓰는 포맷으로 바꿔준다. 
openssl pkcs8 -topk8 -nocrypt -in server.key -out server.pem

