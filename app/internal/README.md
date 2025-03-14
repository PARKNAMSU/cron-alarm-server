# internal

프로젝트 내부에서 사용되는 핵심 애플리케이션 로직이 포함되어 있는 패키지

## router

</br>

## controller

### [user controller](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/controller/user_controller/user_controller.go)

유저 정보 가져오기,추가,업데이트,인증,삭제 등 유저관련 요청을 처리

</br>

## usecase

### [user usecase](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/usecase/user_usecase/user_usecase.go)

유저관련 비즈니스 로적을 처리

</br>

## repository

### [user repository](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/repository/user_repository/repository.go) 

<details>

<summary>유저 DB데이터를 처리(데이터 추가, 업데이트, 삭제, 조회)</summary>

* GetUser
    * 유저 전체 정보 가져오기 (select)
* CreateUser
    * 유저 index 테이블 유저 생성 (insert)
* SetUserLoginData
    * 유저 로그인 정보 생성 및 업데이트 (duplicate)
* SetUserOauth
    * 유저 oauth 정보 생성 및 업데이트 (duplicate)
* SetUserInformation
    * 유저 상세정보 생성 및 업데이트 (duplicate)
* Authorization
    * 유저 인증 (update)
* SetUserRefreshToken
    * 유저 갱신 토큰 생성 및 업데이트 (duplicate)
* DeleteUser
    * 유저 삭제상태로 업데이트 (update)
* GetUserApiKey
    * 유저 api key 가져오기 (select)
* GetRefreshToken
    * 유저 갱신 토큰 가져오기 (select)
* SetUserApiKey
    * 유저 api key 생성 및 업데이트 (duplicate)

</details>

</br>

### [log repository](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/repository/log_repository/repository.go)

</br>

### [stat repository](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/repository/stat_repository/repository.go)


## entity

### user entity

</br>

## [di](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/di/inject.go) 
의존성 주입을 활용하여 
`repository` , `usecase` , `controller`, `middleware` 객체를 중앙에서 초기화 후 필요한 외부 패키지에서 사용

</br>

## middleware

인증 및 검증 등 API 호출 전 수행되어야 하는 미들웨어 목록

### [UserValidation](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/middleware/middleware.go)

JWT 토큰을 검증하여 사용자 인증을 수행하는 미들웨어.
이를 통해 유효한 사용자만 API 요청을 수행할 수 있도록 보호

1. 요청 헤더에서 필요한 정보를 추출
    * `access-token`: 접근 JWT 토큰
    * `refresh-token`: 갱신 JWT 토큰 
2. JWT 토큰 검증
    * 접근토큰 검증후 성공시 다음 로직 진행
    * 실패 시 전달받은 갱신토큰이 존재하는지 DB 검색 
    * 해당 토큰이 유효한경우 접근토큰 갱신 진행
3. 이후의 로직에서 사용할 수 있게 사용자 데이터를 컨텍스트에 저장

### [APIKeyValidation](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/middleware/middleware.go)

외부 사이트에서 해당 서버의 오픈 API 를 사용할 수 있는 API KEY를 검증

1. 요청 헤더에서 필요한 정보를 추출
    * `x-api-key`: 암호화된 API 키
2. API 키 복호화
3. 복호화 한 API 키를 이용하여 해당 키가 존재하는 키 여부를 확인

### [BodyParsor](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/middleware/middleware.go)

byte 형태로 구성되어 있는 request body 데이터를 이후 로직에서 편리하게 사용 가능할 수 있게 map 타입으로 변경 

### [BodyValidator](https://github.com/PARKNAMSU/cron-alarm-server/blob/main/app/internal/middleware/middleware.go)

body key 와 data type 을 넘겨받아 해당 데이터 존재 여부와 타입 검증 진행
