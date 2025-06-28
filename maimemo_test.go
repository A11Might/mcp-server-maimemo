package main

import (
	"testing"
)

func TestMaiMemoClient_GetNotepad(t *testing.T) {
	type args struct {
		notepadId string
	}
	tests := []struct {
		name    string
		args    args
		want    *Notepad
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				notepadId: "np-XwvN41I0JhepK3a3rq-HI6V6AsIw3PmpswmVwQrYlnu4-Cr3RuGFYYstbL9zrBO9",
			},
			want: &Notepad{
				ID: "np-XwvN41I0JhepK3a3rq-HI6V6AsIw3PmpswmVwQrYlnu4-Cr3RuGFYYstbL9zrBO9",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefaultMaimemoClient.GetNotepad(tt.args.notepadId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaiMemoClient.GetNotepad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("MaiMemoClient.GetNotepad() = %v, want %v", got, tt.want)
			}
		})
	}
}
