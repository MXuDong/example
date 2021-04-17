package pkg

// i18n support
type LanguageCollect struct {
	Name   string                   `json:"name"`
	Values map[string]LanguageValue `json:"values"`
}

type LanguageValue struct {
	Values map[string]I18nType `json:"values"`
}

type I18nType struct {
	Values map[string]KeyValue `json:"values"`
}

type KeyValue map[string]string

// InitI18n will init the i18n value
func InitI18n(map[string]LanguageCollect) {

}
