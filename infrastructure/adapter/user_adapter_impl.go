package adapter

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/domain/adapter"
	"ddd-demo-go/domain/entity"
	"ddd-demo-go/infrastructure/common/httpclient"
	"ddd-demo-go/infrastructure/common/logit"
)

type UserAdapterImpl struct {
	logger logit.LoggerInterface
}

type UserServiceResult struct {
	ErrNo  int64  `json:"errno"`
	ErrMsg string `json:"err_msg"`
	UserID int64  `json:"id"`
}

func NewUserAdapter() adapter.UserAdapter {
	return &UserAdapterImpl{}
}

func (adp *UserAdapterImpl) GetUserByID(ctx *gin.Context, id int64) (*entity.UserEntity, error) {
	reqParam := map[string]any{
		"userId": id,
	}

	jsonData, err := json.Marshal(reqParam)
	if err != nil {
		return nil, err
	}

	myHTTPClient := httpclient.NewMyHttpClient()
	myHTTPClient.Retry = 1
	myHTTPClient.Timeout = 3000
	hd := http.Header{}
	hd.Add("Content-Type", "application/json;charset=utf-8")

	client := http.Client{
		Timeout: time.Duration(myHTTPClient.Timeout) * time.Millisecond,
	}
	res, err := myHTTPClient.SendRequest(ctx, client, "", true, jsonData, hd)
	if err != nil {
		return nil, err
	}

	userInfo := UserServiceResult{}
	if err = json.Unmarshal([]byte(res), &userInfo); err != nil {
		return nil, err
	}

	resNew := &entity.UserEntity{}
	resNew.UserID = userInfo.UserID

	return resNew, nil
}
