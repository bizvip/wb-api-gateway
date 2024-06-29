/******************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.                         *
 * Author ORCID: https://orcid.org/0009-0003-8150-367X                        *
 ******************************************************************************/

package custom

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goccy/go-json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ResponseEncoder 自定义ResponseEncoder，去掉proto定义http返回后会加上的 @type 字段，并且用了兼容官方json的更高性能三方json库
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 提取编码器，如果找不到则忽略错误并使用默认编码器
	codec, _ := http.CodecForRequest(r, "Accept")

	var data []byte
	var err error

	switch m := v.(type) {
	case proto.Message:
		options := protojson.MarshalOptions{
			EmitUnpopulated: true,  // 默认值不忽略
			UseProtoNames:   false, // 使用proto name返回http字段
			UseEnumNumbers:  true,  // 将枚举值作为数字发出，默认为枚举值的字符串
		}
		data, err = options.Marshal(m)
		if err != nil {
			return err
		}

		// 解析 JSON
		var jsonObj map[string]interface{}
		if err = json.Unmarshal(data, &jsonObj); err != nil {
			return err
		}

		// 删除 data 中的 @type 字段
		if data, exists := jsonObj["data"]; exists {
			if dataMap, ok := data.(map[string]interface{}); ok {
				delete(dataMap, "@type")
			}
		}

		data, err = json.Marshal(jsonObj)
		if err != nil {
			return err
		}
	default:
		data, err = codec.Marshal(v)
		if err != nil {
			return err
		}
	}

	w.Header().Set("Content-Type", ContentType(codec.Name()))
	_, err = w.Write(data)
	return err
}
