package usecase

import (
	"vk-film-library/internal/entity/global"
	"vk-film-library/internal/entity/user"
	"vk-film-library/internal/transaction"
	"vk-film-library/rimport"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
}

func NewAuth(log *logrus.Logger, ri rimport.RepositoryImports) *AuthUsecase {
	return &AuthUsecase{
		log:               log,
		RepositoryImports: ri,
	}
}

func (u *AuthUsecase) RegisterUser(ts transaction.Session, us user.User) (err error) {
	if us.Login == "" || us.Password == "" || us.Role == "" {
		err = global.ErrParamsIncorect
		return
	}

	if !us.IsValidRole() {
		err = global.ErrParamsIncorect
		return
	}

	regUser := user.RegisteredUser{
		Login: us.Login,
		Role:  us.Role,
	}

	regUser.Password, err = bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorln("не удалось захэшировать пароль, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	if err = u.Repository.Auth.RegisterUser(ts, regUser); err != nil {
		u.log.Errorln("не удалось создать нового пользователя, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.Infoln("пользователь успешно зарегистрирован")
	return
}

func (u *AuthUsecase) Login(ts transaction.Session, us user.User) (jwtTokenString string, err error) {
	if us.Login == "" || us.Password == "" {
		err = global.ErrParamsIncorect
		return
	}

	var registeredUser user.RegisteredUser

	registeredUser, err = u.Repository.Auth.GetUseInfo(ts, us.Login)
	switch err {
	case nil:
	case global.ErrNoData:
		u.log.Debugln("пользователь не найден")
		return
	default:
		u.log.Errorln("не удалось найти пользователя, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	if err = bcrypt.CompareHashAndPassword(registeredUser.Password, []byte(us.Password)); err != nil {
		u.log.Errorln("неверный логин или пароль")
		err = user.ErrUnAuthorized
		return
	}

	jwtTokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": registeredUser.Login,
		"role":  registeredUser.Role,
	}).SignedString([]byte(user.JwtKey))
	if err != nil {
		u.log.Errorln("не удалось сгенерировать jwt токен, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.Infof("пользователь %s успешно авторизован", registeredUser.Login)
	return
}
