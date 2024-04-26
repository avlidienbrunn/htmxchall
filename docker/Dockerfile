FROM node:16 as cssbuild
WORKDIR /node
COPY . ./
# Should not be needed?
RUN npm i tailwindcss
RUN npm i tailwind-dracula
RUN npx tailwindcss -c /node/tailwind.config.js -i ./styles.css -o ./styles-build.css 

FROM golang:1.22.1-alpine AS build
WORKDIR /go/src/app
RUN apk add git
#COPY ./backend .
COPY ./ .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
RUN go build -v -o app .

FROM alpine:3.19.1
RUN apk add bash
COPY --from=build /go/src/app/app /go/bin/app
COPY ./public /public
COPY --from=cssbuild /node/styles-build.css /public/assets/styles.css
ENTRYPOINT ["/go/bin/app"]