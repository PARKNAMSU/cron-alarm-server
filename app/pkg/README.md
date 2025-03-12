# pkg

외부에서 사용 가능한 공통적으로 사용될 유틸리티 패키지

## database

## tool

프로젝트에서 사용할 툴 목록

### mail tool

smtp 이메일 전송 툴

* encodeBase64 - 문자열 인코딩 함수
    * parameters
        * `input` (string) : 인코딩할 문자열
* readTemplate - 이메일 전송 템플릿 생성
    * parameters 
        * `title` (string) : 이메일 제목 
        * `body` (string) : 전송할 내용
* SendMail - 이메일 전송 함수
    * parameters
        * `to` (string) : 전송할 이메일 주소
        * `msg` (string) : 전송할 메서지
        * `title` (string) : 이메일 제목
    * 진행순서
        1. RFC - 522 포맷에 맞는 이메일 전송 헤더 생성. title 의 경우 base64로 인코딩
        2. `readTemplate` 함수를 아용하여 이메일 전송 포맷 문자열 가져와서 헤더와 병합
        3. 이메일 전송 진행