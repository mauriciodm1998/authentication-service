package repositories

// import (
// 	"authentication-service/internal/canonical"
// 	"context"
// 	"errors"
// 	"regexp"
// 	"testing"
// 	"time"

// 	"github.com/pashagolub/pgxmock"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	repository LoginRepository

// 	Mock pgxmock.PgxPoolIface
// )

// func init() {
// 	mock, _ := pgxmock.NewPool()

// 	Mock = mock
// 	repository = NewUserRepo(mock)
// }

// func TestGetUserByLogin(t *testing.T) {
// 	ctx := context.Background()

// 	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

// 	user := canonical.User{
// 		Id:            "asdfasda",
// 		Login:         "fulano@",
// 		Password:      "fulanos",
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelId", "AccessedAt"}).
// 		AddRow("asdfasda", "fulano@", "fulanos", 1, nil))

// 	_, err := repository.GetUserByLogin(ctx, user.Login)

// 	assert.Nil(t, err)
// 	Mock.ExpectationsWereMet()
// }

// func TestGetUserByLoginQueryError(t *testing.T) {
// 	ctx := context.Background()

// 	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

// 	user := canonical.User{
// 		Id:            "asdfasda",
// 		Login:         "fulano@",
// 		Password:      "fulanos",
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

// 	_, err := repository.GetUserByLogin(ctx, user.Login)

// 	assert.Error(t, err)
// 	Mock.ExpectationsWereMet()
// }

// func TestGetUserByLoginNotFound(t *testing.T) {
// 	ctx := context.Background()

// 	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

// 	user := canonical.User{
// 		Id:            "asdfasda",
// 		Login:         "fulano@",
// 		Password:      "fulanos",
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelId", "AccessedAt"}))
// 	_, err := repository.GetUserByLogin(ctx, user.Login)

// 	assert.Equal(t, errors.New("entity not found"), err)
// 	Mock.ExpectationsWereMet()
// }

// func TestGetUserByLoginScanError(t *testing.T) {
// 	ctx := context.Background()

// 	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

// 	user := canonical.User{
// 		Id:            "asdfasda",
// 		Login:         "fulano@",
// 		Password:      "fulanos",
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelId", "AccessedAt"}).
// 		AddRow("asdfasda", "fulano@", "fulanos", 1, nil).RowError(0, errors.New("generic error")))

// 	_, err := repository.GetUserByLogin(ctx, user.Login)

// 	assert.Error(t, err)
// 	Mock.ExpectationsWereMet()
// }

// func TestSaveAccess(t *testing.T) {
// 	ctx := context.Background()

// 	time := time.Now()
// 	sqlStatement := "INSERT INTO \"Access\" (Id, AccessLevelId, UserID, accessedAt) VALUES ($1, $2, $3, $4)"

// 	access := canonical.Access{
// 		ID:            "asdasdas",
// 		USERID:        "asdasdsa",
// 		AccessedAt:    time,
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnResult(pgxmock.NewResult("INSERT", 1))

// 	err := repository.SaveAccess(ctx, access)

// 	assert.Nil(t, err)
// 	Mock.ExpectationsWereMet()
// }

// func TestSaveAccessError(t *testing.T) {
// 	ctx := context.Background()

// 	time := time.Now()
// 	sqlStatement := "INSERT INTO \"Access\" (Id, AccessLevelId, UserID, accessedAt) VALUES ($1, $2, $3, $4)"

// 	access := canonical.Access{
// 		ID:            "asdasdas",
// 		USERID:        "asdasdsa",
// 		AccessedAt:    time,
// 		AccessLevelID: 1,
// 	}

// 	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

// 	err := repository.SaveAccess(ctx, access)

// 	assert.Error(t, err)
// 	Mock.ExpectationsWereMet()
// }
