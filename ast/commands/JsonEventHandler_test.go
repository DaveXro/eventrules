package commands

import "testing"

func TestDecodeJsonFile(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}
}

func TestFindJsonStringField(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	eventId, _ := handler.StringField("EventId")

	if len(eventId) == 0 {
		t.Failed()
	}

}

func TestFindJsonBodyStringField(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	eventId, _ := handler.StringField("Reference")

	if len(eventId) == 0 {
		t.Failed()
	}
}

func TestFindJsonBodyIntField(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	eventId, _ := handler.IntegerField("StaffCount")

	if eventId != 40 {
		t.Failed()
	}
}
func TestFindJsonBodyIntFieldAsString(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	staffCount, _ := handler.StringField("StaffCount")

	if len(staffCount) == 0 {
		t.Failed()
	}
}

func TestFindJsonSubBodyStringField(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	rating, _ := handler.StringField("ExperianData.ExperianRating")

	if len(rating) == 0 {
		t.Failed()
	}
}

func TestFindJsonSubBodyStringFieldAndConvert(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	rating, _ := handler.IntegerField("ExperianData.ExperianRating")

	if rating < 40 {
		t.Failed()
	}
}

func TestFindJsonMissingStringField(t *testing.T) {
	eventPackage := &EventPackage{}

	eventPackage.DecodeEventFromFile("../../LeadEvent.json")

	if len(eventPackage.FieldMap) == 0 {
		t.Failed()
	}

	handler := IEventFieldMap(eventPackage)

	_, err := handler.StringField("ference")

	if err == nil {
		t.Failed()
	}
}
