# config

환경변수 및 설정값을 정의해놓은 패키지.

## common.go

공통으로 사용되는 설정값들을 정의해놓은 폴더.


* `ENVIRONMENT`: 프로그램의 실행 환경으로 (development, staging, production) 으로 구성
    * `development`: 로컬 개발환경
    * `staging`: 테스트 환경
    * `production`: 실제 프로그램 구동 환경
* `JWT_ACCESS_TOKEN_PERIOD`: acess token 유효 기간
* `JWT_REFRESH_TOKEN_PERIOD`: refresh token 유효기간

## key.go 

암호화 비밀키 및 기타 키들을 정의해놓은 폴더.

* `JWT_ACCESS_TOKEN_KEY`: 유저 인증 JWT access token 생성 비밀키
* `JWT_REFRESH_TOKEN_KEY`: 유저 인증 JWT refresh token 생성 비밀키
* `USER_PASSWORD_ENCRYPT_KEY`: 유저 패스워드를 암호화 하기 위한 암호화 키
* `USER_API_ENCRYPT_KEY`: 유저 개인별 API 키를 암호화 하기 위한 암호화 키
