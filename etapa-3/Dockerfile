FROM golang:1.14.3-alpine AS build

WORKDIR /src
COPY  . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch as bin

COPY --from=build /src .

EXPOSE 8000
CMD [ "./main" ]