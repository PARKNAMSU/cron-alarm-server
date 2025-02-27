# 크론 작업 알람 서비스

## 프로젝트 설명
기존 크론 작업 시 오류 발생 로그를 확인하기 위해 로그가 기록된 cloudwatch 혹은 datadog 등 직접 접속하여 확인할 수 밖에 없어 실시간으로 확인 및 error fix 하는데 제약이 발생. 
이러한 불편사항을 개선하기 위해 크론 작업 시 오류 및 알람사항 발생하면 실시간으로 알람을 제공하는 서비스 제공

## 프로젝트 구성

* [cmd](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/cmd)
* [config](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/config)
* [internal](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/internal)
* [pkg](https://github.com/PARKNAMSU/cron-alarm-server/tree/main/app/pkg)
