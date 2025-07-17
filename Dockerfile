FROM golang:1.20

ENV Dev_DB="root:Admin@123@tcp/AuthData_Test?charset=UTF8&parseTime=True&loc=Local"
ENV Dev_Jwt_Secret="Secret"
ENV Dev_RedisAdder="localhost:6379"
ENV Dev_RedisPassword=""
ENV Dev_Redis=0

ENV Dev_DB=0

RUN go build -o bin . 

ENTRYPOINT [ "/bin" ]