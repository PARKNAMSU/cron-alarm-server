package platform_usecase

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/config"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/log_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/platform_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/common_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/encrypt_tool"
)

type PlatformUsecaseImpl interface {
	ApiKeyIssue(input ApiKeyIssueInput) (ApiKeyIssueOutput, *types.CustomError)
}

type platformUsecase struct {
	userRepository     user_repository.UserRepositoryImpl
	logRepository      log_repository.LogRepositoryImpl
	platformRepository platform_repository.PlatformRepositoryImpl
}

var (
	usecase *platformUsecase
)

func NewUsecase(
	userRepo user_repository.UserRepositoryImpl,
	logRepo log_repository.LogRepositoryImpl,
	platformRepo platform_repository.PlatformRepositoryImpl,
) PlatformUsecaseImpl {
	if usecase == nil {
		usecase = &platformUsecase{
			userRepository:     userRepo,
			logRepository:      logRepo,
			platformRepository: platformRepo,
		}
	}
	return usecase
}

func (u *platformUsecase) ApiKeyIssue(input ApiKeyIssueInput) (ApiKeyIssueOutput, *types.CustomError) {
	// 1. 사용자의 플랫폼 전체 목록을 조회한다.
	list := u.platformRepository.GetPlatform(platform_repository.GetPlatformInput{
		UserId:      &input.UserData.UserId,
		SearchType:  user_repository.GET_USER_KEY_ID,
		IsGetUsable: false,
	})
	// 2. 사용자의 최대 플랫폼 생성 개수가 넘었는지 확인한다.
	if len(list) >= input.UserData.MaxPlatformCnt {
		return ApiKeyIssueOutput{}, &types.CustomError{
			Code:   "MAX-PLATFORM-CNT",
			Msg:    "플랫폼 생성 개수를 초과하였습니다.",
			Status: 400,
		}
	}

	for _, platform := range list {
		if platform.Hostname == input.Hostname {
			return ApiKeyIssueOutput{}, &types.CustomError{
				Code:   "PLATFORM-EXIST",
				Msg:    "이미 존재하는 플랫폼입니다.",
				Status: 400,
			}
		}
	}

	var encryptKey string
	// 3. 넘지 않았을 경우 api key를 생성한다.
	if apiKey, err := common_tool.RandomString(32); err != nil {
		return ApiKeyIssueOutput{}, &types.CustomError{
			Code:   "INTERNAL-SERVER-ERROR",
			Msg:    "internal server error",
			Status: 500,
		}
	} else {
		encryptKey, _ = encrypt_tool.Encrypt([]byte(apiKey), config.USER_API_ENCRYPT_KEY)
	}

	if len(encryptKey) <= 0 {
		return ApiKeyIssueOutput{}, &types.CustomError{
			Code:   "INTERNAL-SERVER-ERROR",
			Msg:    "internal server error",
			Status: 500,
		}
	}

	expiredAt := time.Now().Add(config.API_KEY_AVAILABLE_PERIOD)

	// 4. 생성한 api key를 플랫폼 테이블에 저장한다.
	if err := u.platformRepository.InserPlatform(platform_repository.InserPlatformInput{
		UserId:    input.UserData.UserId,
		Hostname:  input.Hostname,
		ApiKey:    encryptKey,
		ExpiredAt: expiredAt,
	}); err != nil {
		return ApiKeyIssueOutput{}, &types.CustomError{
			Code:   "INTERNAL-SERVER-ERROR",
			Msg:    "internal server error",
			Status: 500,
		}
	}

	// 5. 생성 로그를 기록한다.
	if err := u.logRepository.InsertLogUserApiKey(log_repository.InsertLogUserApiKeyInput{
		UserId:   input.UserData.UserId,
		Hostname: input.Hostname,
		ApiKey:   encryptKey,
		IpAddr:   input.IpAddr,
		Action:   "issue",
	}); err != nil {
		return ApiKeyIssueOutput{}, &types.CustomError{
			Code:   "INTERNAL-SERVER-ERROR",
			Msg:    "internal server error",
			Status: 500,
		}
	}

	// 6. 생성한 api key를 반환한다.
	return ApiKeyIssueOutput{
		ApiKey:    encryptKey,
		ExpiredAt: expiredAt,
		Hostname:  input.Hostname,
	}, nil
}
