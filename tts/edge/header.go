package edge

import (
	"bufio"
	"bytes"
)

type AudioMp3MessageHead struct {
	SourceMsg   string
	XRequestId  string
	ContentType string
	XStreamId   string
	Path        string
}

func NewAudioMp3MessageHead(data []byte) *AudioMp3MessageHead {
	head := &AudioMp3MessageHead{
		SourceMsg: string(data),
	}
	head.analyzeByte(data)
	return head
}

func (head *AudioMp3MessageHead) analyzeByte(data []byte) {
	reader := bufio.NewReader(bytes.NewBuffer(data))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = line[:len(line)-1] // 去掉换行符
		if line != "" {
			split := bytes.SplitN([]byte(line), []byte(":"), 2)
			if len(split) == 2 {
				key := string(bytes.TrimSpace(split[0]))
				value := string(bytes.TrimSpace(split[1]))

				switch key {
				case "X-RequestId":
					head.XRequestId = value
				case "Content-Type":
					head.ContentType = value
				case "X-StreamId":
					head.XStreamId = value
				case "Path":
					head.Path = value
				}
			}
		}
	}
}
