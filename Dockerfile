FROM golang:latest as build

WORKDIR /app

# Cgo mogule dosyalarini kopyalar
COPY go.mod .
COPY go.sum .

RUN go mod download

#bilgisayardaki tum kaynak kodlarini containere a atar
COPY . . 

#kodu derler
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp main.go
 
#yeni bir linux ortami acar
FROM alpine:latest as run


#ilk asamada derlenen programi alir ve yeni ortamin icine kopyalar alpine oldugu icin daha hafiftir
COPY --from=build /myapp /myapp

WORKDIR /app
EXPOSE 3000
CMD ["/myapp"]
   #container baskatildiginda /myapp programi calisir