package edge

import (
	"testing"

	"github.com/pp-group/edge-tts-go/biz/service/tts/edge"
)

func TestGenTTS(t *testing.T) {
	fileName, err := genTTS("如果喜欢这个项目的话请点个 Star 吧。")
	if err != nil {
		t.Errorf("genTTS fail, err: %v", err)
		return
	}
	t.Logf("genTTS success, fileName: %s", fileName)
}

// genTTS template use github.com/pp-group/edge-tts-go
func genTTS(text string) (string, error) {
	c, err := edge.NewCommunicate(text)
	if err != nil {
		return "", err
	}

	speech, err := NewLocalSpeech(c, "tmp")
	if err != nil {
		return "", err
	}

	_, callback := speech.GenTTS()

	err = callback()
	if err != nil {
		return "", err
	}

	return speech.URL(speech.FileName)
}