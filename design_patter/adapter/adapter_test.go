package adapter

func ExampleAdapter() {
	voice := NewVoicePlayer()
	adapter := NewAdapter(voice)
	adapter.playMv()
	// Output:
	// play voice
}
