# 크론 작업 알람 서비스

## 프로젝트 설명
기존 크론 작업 시 오류 발생 로그를 확인하기 위해 로그가 기록된 cloudwatch 혹은 datadog 등 직접 접속하여 확인할 수 밖에 없어 실시간으로 확인 및 error fix 하는데 제약이 발생. 
이러한 불편사항을 개선하기 위해 크론 작업 시 오류 및 알람사항 발생하면 실시간으로 알람을 제공하는 서비스 제공.
프로젝트 기획 시 크론 작업을 대상으로 만들기는 하였지만, 크론 잡 이외에도 API 서버, 클라이언트 등 실시간 로그 확인이 필요한 곳에서 사용할 수 있음.

## 프로젝트 아키텍처
<img width="817" alt="Image" src="https://github.com/user-attachments/assets/9310d2fb-cb8b-43c5-bb1c-f0b6f6448c6b" />

## 프로젝트 구성

### [cmd](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/cmd)
  main.go 실행파일이 위치하는 곳으로 애플리케이션의 시작 패키지.
### [config](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/config)
  환경변수 및 설정값을 정의해놓은 패키지.
### [internal](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/internal)
  프로젝트 내부에서 사용되는 핵심 애플리케이션 로직이 포함되어 있는 패키지. router, controller, usecase, repository 등 프로젝트의 핵심 패키지가 포함되어 있음.
### [pkg](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/pkg)
  프로젝트에서 공통적으로 사용될 유틸리티 패키지. database툴 및 기타 툴 등이 포함되어 있음.
  
