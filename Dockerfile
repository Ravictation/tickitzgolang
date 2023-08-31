#! buat image container
FROM golang:1.20.3-alpine AS build


#! buat folder untuk niympan code
WORKDIR /goapptickitz

#! Copy semua file
COPY . .

#! install depedency and build
RUN go mod download
RUN go build -v -o /goapptickitz/tickitzgolang ./cmd/main.go

#! create other images 
# Final stage
FROM alpine:3.14

WORKDIR /goapptickitz

#! copy build file
COPY --from=build /goapptickitz /goapptickitz

ENV PATH="/goapptickitz:${PATH}"

EXPOSE 8082

ENTRYPOINT ["tickitzgolang", "--listen"]

# docker build -t Ravictation/tickitzgolang .