# 빌드 스테이지
FROM golang:1.22 AS builder

# 작업 디렉토리 설정
WORKDIR /app

ENV DB_TYPE=$DB_TYPE

# 모듈 파일 복사
COPY go.mod ./
COPY go.sum ./

# 모듈 다운로드
RUN go mod download

# 소스 코드 복사
COPY . .
# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 실행 스테이지
FROM alpine:latest

# SSL 인증서 추가 (필요한 경우)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 빌드 스테이지에서 생성된 바이너리 파일 복사
COPY --from=builder /app/main .
# 실행 명령어
CMD ["./main"]
