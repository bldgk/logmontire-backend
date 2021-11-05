package http

type Validation struct {
	Key      string
	Type     string
	Rules    []string
	Required bool
}

type ValidationRules []Validation

func CreateValidationRules() ValidationRules {
	return ValidationRules{}
}

func (vr ValidationRules) AddValidation(key string, rules []string, typed string, required bool) ValidationRules {
	return append(vr, Validation{
		Key:      key,
		Rules:    rules,
		Type:     typed,
		Required: required,
	})
}
