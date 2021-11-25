FROM golang:1.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV RGE_OWNER_MAIL owner@email.com
ENV RGE_OWNER_PASSWORD Password12345.
ENV RGE_SEEDER FALSE

ENV RGE_TYPE POSTGRES
ENV RGE_USER postgres
ENV RGE_PASSWORD Password12345.
ENV RGE_PORT 5432
ENV RGE_NAME_DB postgres
ENV RGE_HOST localhost

ENV RGE_RSA None
ENV RGE_RSA_PUB None
ENV RGE_COOKIE_KEY MP}mn!=v=xw#fE_Jj{}?PFS%kB;$78hB

ENV RGE_MAIL_USER NONE
ENV RGE_MAIL_PASSWORD NONE
ENV RGE_MAIL_PORT 465
ENV RGE_MAIL_NAME beerparacreer@example.com
ENV RGE_MAIL_HOST smtp.mailtrap.io
ENV RGE_MAIL_CODE_HOST http://localhost
COPY --from=0 /app/public ./public
COPY --from=0 /app/main ./

CMD ["./main"]  


