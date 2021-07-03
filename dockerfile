#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:curso-go-web-avancado .
#docker run -p 8080:8080 -d mercurius:curso-go-web-avancado

FROM scratch

ADD curso-go-web-avancado /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/curso-go-web-avancado" ]