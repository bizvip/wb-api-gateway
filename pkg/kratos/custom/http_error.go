package custom

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
)

type ValidErr struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := kratoshttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 如果是参数校验错误
	// if se.GetReason() == "VALIDATOR" {
	// 	e := v1.ErrorInvalidInput("参数校验失败")
	// 	r := &ValidErr{
	// 		Code: uint32(e.GetCode()),
	// 		Msg:  e.GetMessage(),
	// 		Data: e.GetMetadata(),
	// 	}
	// 	body, err = json.Marshal(r)
	// }

	w.Header().Set("Content-Type", ContentType(codec.Name()))
	// w.WriteHeader(int(se.Code))
	// 强制全部业务逻辑错误返回http status code 为200
	w.WriteHeader(200)
	_, _ = w.Write(body)
}
