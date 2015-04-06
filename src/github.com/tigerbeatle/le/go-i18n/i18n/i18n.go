// Package i18n supports string translations with variable substitution and CLDR pluralization.
// It is intended to be used in conjunction with github.com/finapps/go-i18n/goi18n,
// although that is not strictly required.
//
// Initialization
//
// Your Go program should load translations during its intialization.
//     i18n.MustLoadTranslationFile("path/to/fr-FR.all.json")
// If your translations are in a file format not supported by (Must)?LoadTranslationFile,
// then you can use the AddTranslation function to manually add translations.
//
// Fetching a translation
//
// Use Tfunc or MustTfunc to fetch a TranslateFunc that will return the translated string for a specific locale.
// The TranslateFunc will be bound to the first valid locale passed to Tfunc.
//     userLocale = "ar-AR"     // user preference, accept header, language cookie
//     defaultLocale = "en-US"  // known valid locale
//     T, err := i18n.Tfunc(userLocale, defaultLocale)
//     fmt.Println(T("Hello world"))
//
// Usually it is a good idea to identify strings by a generic id rather than the English translation,
// but the rest of this documentation will continue to use the English translation for readability.
//     T("program_greeting")
//
// Variables
//
// TranslateFunc supports strings that have variables using the text/template syntax.
//     T("Hello {{.Person}}", map[string]interface{}{
//         "Person": "Bob",
//     })
//
// Pluralization
//
// TranslateFunc supports the pluralization of strings using the CLDR pluralization rules defined here:
// http://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html
//     T("You have {{.Count}} unread emails.", 2)
//     T("I am {{.Count}} meters tall.", "1.7")
//
// Plural strings may also have variables.
//     T("{{.Person}} has {{.Count}} unread emails", 2, map[string]interface{}{
//         "Person": "Bob",
//     })
//
// Compound plural strings can be created with nesting.
//     T("{{.Person}} has {{.Count}} unread emails in the past {{.Timeframe}}.", 3, map[string]interface{}{
//         "Person":    "Bob",
//         "Timeframe": T("{{.Count}} days", 2),
//     })
//
// Templates
//
// You can use the .Funcs() method of a text/template or html/template to register a TranslateFunc
// for usage inside of that template.
package i18n

import (
	"github.com/tigerbeatle/le/go-i18n/i18n/bundle"
	"github.com/tigerbeatle/le/go-i18n/i18n/locale"
	"github.com/tigerbeatle/le/go-i18n/i18n/translation"
)

// TranslateFunc returns the translation of the string identified by translationID.
//
// If translationID is a non-plural form, then the first variadic argument may be a map[string]interface{}
// that contains template data.
//
// If translationID is a plural form, then the first variadic argument must be an integer type
// (int, int8, int16, int32, int64) or a float formatted as a string (e.g. "123.45").
// The second variadic argument may be a map[string]interface{} that contains template data.
type TranslateFunc func(translationID string, args ...interface{}) string

// IdentityTfunc returns a TranslateFunc that always returns the translationID passed to it.
//
// It is a useful placeholder when parsing a text/template or html/template
// before the actual Tfunc is available.
func IdentityTfunc() TranslateFunc {
	return func(translationID string, args ...interface{}) string {
		return translationID
	}
}

var defaultBundle = bundle.New()

// MustLoadTranslationFile is similar to LoadTranslationFile
// except it panics if an error happens.
func MustLoadTranslationFile(filename string) {
	defaultBundle.MustLoadTranslationFile(filename)
}

// LoadTranslationFile loads the translations from filename into memory.
//
// The locale that the translations are associated with is parsed from the filename.
//
// Generally you should load translation files once during your program's initialization.
func LoadTranslationFile(filename string) error {
	return defaultBundle.LoadTranslationFile(filename)
}

// AddTranslation adds translations for a locale.
//
// It is useful if your translations are in a format not supported by LoadTranslationFile.
func AddTranslation(locale *locale.Locale, translations ...translation.Translation) {
	defaultBundle.AddTranslation(locale, translations...)
}

// MustTfunc is similar to Tfunc except it panics if an error happens.
func MustTfunc(localeID string, localeIDs ...string) TranslateFunc {
	return TranslateFunc(defaultBundle.MustTfunc(localeID, localeIDs...))
}

// Tfunc returns a TranslateFunc that will be bound to the first valid locale from its parameters.
func Tfunc(localeID string, localeIDs ...string) (TranslateFunc, error) {
	tf, err := defaultBundle.Tfunc(localeID, localeIDs...)
	return TranslateFunc(tf), err
}
