package rpc

import (
    "fmt"
    "bytes"
    "encoding/json"
    "errors"
    "strconv"
)

type BaseMessage struct {
    Method string `json:"method"`
}

func EncodeMessage(msg any) string{
    content,err := json.Marshal(msg)
    if err!=nil{
        panic(err)
    }
    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(content),content)
}

func DecodeMessage(msg []byte) (string,[]byte,error) {
    header,content,found := bytes.Cut(msg, []byte{'\r','\n','\r','\n'})
    if !found {
        return "",nil,errors.New("Did not find seperator")
    }
    contentLengthBytes := header[len("Content-Header: "):]
    contentLength,err := strconv.Atoi(string(contentLengthBytes))
    if err!=nil {
        return "",nil,err
    }

    _ = content

    var baseMessage BaseMessage

    if err := json.Unmarshal(content[:contentLength], &baseMessage); err !=nil {
        return "",nil,err
    }

    return baseMessage.Method,content,nil
}
