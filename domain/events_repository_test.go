package domain

import (
	"testing"
)

func Test_deve_retornar_array_vazio_de_eventos(t *testing.T) {
	events := NewEventsRepository().GetAll()

	for _ = range *events {
		t.Errorf("Array de eventos deveria estar vazio.")
		return
	}
}
