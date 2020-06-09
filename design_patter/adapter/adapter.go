package adapter

import "fmt"

// 思路流程
// 我有一个MP4 能播放歌曲mv 所以他有一个playMv的方法 运行的很完美
// 后面有有一段音频 只有音乐没有画面 但是mp4 只有playMv 并没有playVoice这个方法 但是mp4实际上是能播放voice的 这个时候怎么办呢 难道我重新设计我的MP4吗
// 这种情况下 就可以用适配器模式 写一个adapter的playMv 去调用playVoice就可以了

type mv interface {
	playMv()
}

type voice interface {
	playVoice()
}

type voicePlayer struct {
}

func NewVoicePlayer() voice {
	return &voicePlayer{}
}

func (p *voicePlayer) playVoice() {
	fmt.Print("play voice\n")
}

type adapter struct {
	voice
}

func NewAdapter(adaptee voice) voice {
	return &adapter{voice: adaptee}
}

func (a adapter) playMv() {
	a.voice.playVoice()
}
