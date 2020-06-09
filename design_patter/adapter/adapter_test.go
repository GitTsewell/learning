package adapter

func ExampleAdapter() {
	voice := NewVoicePlayer()
	adapter := NewAdapter(voice)
	adapter.playVoice()
	// Output:
	// play voice
}
