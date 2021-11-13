package inflect

import (
	"testing"
)

// test data

var SingularToPlural = map[string]string{
	"search":      "searches",
	"switch":      "switches",
	"fix":         "fixes",
	"box":         "boxes",
	"process":     "processes",
	"address":     "addresses",
	"case":        "cases",
	"stack":       "stacks",
	"wish":        "wishes",
	"fish":        "fish",
	"jeans":       "jeans",
	"funky jeans": "funky jeans",
	"category":    "categories",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"diagnosis_a": "diagnosis_as",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"bus":         "buses",
	"status":      "statuses",
	"status_code": "status_codes",
	"mouse":       "mice",
	"louse":       "lice",
	"house":       "houses",
	"octopus":     "octopi",
	"virus":       "viri",
	"alias":       "aliases",
	"portfolio":   "portfolios",
	"vertex":      "vertices",
	"matrix":      "matrices",
	"matrix_fu":   "matrix_fus",
	"axis":        "axes",
	"testis":      "testes",
	"crisis":      "crises",
	"rice":        "rice",
	"shoe":        "shoes",
	"horse":       "horses",
	"prize":       "prizes",
	"edge":        "edges",
	"database":    "databases",
}

var CapitalizeMixture = map[string]string{
	"product":               "Product",
	"special_guest":         "Special_guest",
	"applicationController": "ApplicationController",
	"Area51Controller":      "Area51Controller",
}

var CamelToUnderscore = map[string]string{
	"Product":               "product",
	"SpecialGuest":          "special_guest",
	"ApplicationController": "application_controller",
	"Area51Controller":      "area51_controller",
}

var UnderscoreToLowerCamel = map[string]string{
	"product":                "product",
	"special_guest":          "specialGuest",
	"application_controller": "applicationController",
	"area51_controller":      "area51Controller",
}

var CamelToUnderscoreWithoutReverse = map[string]string{
	"HTMLTidy":          "html_tidy",
	"HTMLTidyGenerator": "html_tidy_generator",
	"FreeBsd":           "free_bsd",
	"HTML":              "html",
}

var ClassNameToForeignKeyWithUnderscore = map[string]string{
	"Person":  "person_id",
	"Account": "account_id",
}

var PluralToForeignKeyWithUnderscore = map[string]string{
	"people":   "person_id",
	"accounts": "account_id",
}

var ClassNameToForeignKeyWithoutUnderscore = map[string]string{
	"Person":  "personid",
	"Account": "accountid",
}

var ClassNameToTableName = map[string]string{
	"PrimarySpokesman": "primary_spokesmen",
	"NodeChild":        "node_children",
}

var StringToParameterized = map[string]string{
	"Donald E. Knuth":                     "donald-e-knuth",
	"Random text with *(bad)* characters": "random-text-with-bad-characters",
	"Allow_Under_Scores":                  "allow_under_scores",
	"Trailing bad characters!@#":          "trailing-bad-characters",
	"!@#Leading bad characters":           "leading-bad-characters",
	"Squeeze   separators":                "squeeze-separators",
	"Test with + sign":                    "test-with-sign",
	"Test with malformed utf8 \251":       "test-with-malformed-utf8",
}

var StringToParameterizeWithNoSeparator = map[string]string{
	"Donald E. Knuth":                     "donaldeknuth",
	"With-some-dashes":                    "with-some-dashes",
	"Random text with *(bad)* characters": "randomtextwithbadcharacters",
	"Trailing bad characters!@#":          "trailingbadcharacters",
	"!@#Leading bad characters":           "leadingbadcharacters",
	"Squeeze   separators":                "squeezeseparators",
	"Test with + sign":                    "testwithsign",
	"Test with malformed utf8 \251":       "testwithmalformedutf8",
}

var StringToParameterizeWithUnderscore = map[string]string{
	"Donald E. Knuth":                     "donald_e_knuth",
	"Random text with *(bad)* characters": "random_text_with_bad_characters",
	"With-some-dashes":                    "with-some-dashes",
	"Retain_underscore":                   "retain_underscore",
	"Trailing bad characters!@#":          "trailing_bad_characters",
	"!@#Leading bad characters":           "leading_bad_characters",
	"Squeeze   separators":                "squeeze_separators",
	"Test with + sign":                    "test_with_sign",
	"Test with malformed utf8 \251":       "test_with_malformed_utf8",
}

var StringToParameterizedAndNormalized = map[string]string{
	"Malmö":         "malmo",
	"Garçons":       "garcons",
	"Opsů":          "opsu",
	"Ærøskøbing":    "aeroskobing",
	"Aßlar":         "asslar",
	"Japanese: 日本語": "japanese",
}

var UnderscoreToHuman = map[string]string{
	"employee_salary": "Employee salary",
	"employee_id":     "Employee",
	"underground":     "Underground",
}

var MixtureToTitleCase = map[string]string{
	"active_record":      "Active Record",
	"ActiveRecord":       "Active Record",
	"action web service": "Action Web Service",
	"Action Web Service": "Action Web Service",
	"Action web service": "Action Web Service",
	"actionwebservice":   "Actionwebservice",
	"Actionwebservice":   "Actionwebservice",
	"david's code":       "David's Code",
	"David's code":       "David's Code",
	"david's Code":       "David's Code",
}

var OrdinalNumbers = map[string]string{
	"-1":    "-1st",
	"-2":    "-2nd",
	"-3":    "-3rd",
	"-4":    "-4th",
	"-5":    "-5th",
	"-6":    "-6th",
	"-7":    "-7th",
	"-8":    "-8th",
	"-9":    "-9th",
	"-10":   "-10th",
	"-11":   "-11th",
	"-12":   "-12th",
	"-13":   "-13th",
	"-14":   "-14th",
	"-20":   "-20th",
	"-21":   "-21st",
	"-22":   "-22nd",
	"-23":   "-23rd",
	"-24":   "-24th",
	"-100":  "-100th",
	"-101":  "-101st",
	"-102":  "-102nd",
	"-103":  "-103rd",
	"-104":  "-104th",
	"-110":  "-110th",
	"-111":  "-111th",
	"-112":  "-112th",
	"-113":  "-113th",
	"-1000": "-1000th",
	"-1001": "-1001st",
	"0":     "0th",
	"1":     "1st",
	"2":     "2nd",
	"3":     "3rd",
	"4":     "4th",
	"5":     "5th",
	"6":     "6th",
	"7":     "7th",
	"8":     "8th",
	"9":     "9th",
	"10":    "10th",
	"11":    "11th",
	"12":    "12th",
	"13":    "13th",
	"14":    "14th",
	"20":    "20th",
	"21":    "21st",
	"22":    "22nd",
	"23":    "23rd",
	"24":    "24th",
	"100":   "100th",
	"101":   "101st",
	"102":   "102nd",
	"103":   "103rd",
	"104":   "104th",
	"110":   "110th",
	"111":   "111th",
	"112":   "112th",
	"113":   "113th",
	"1000":  "1000th",
	"1001":  "1001st",
}

var UnderscoresToDashes = map[string]string{
	"street":                "street",
	"street_address":        "street-address",
	"person_street_address": "person-street-address",
}

var Irregularities = map[string]string{
	"person": "people",
	"man":    "men",
	"child":  "children",
	"sex":    "sexes",
	"move":   "moves",
}

type AcronymCase struct {
	camel string
	under string
	human string
	title string
}

var AcronymCases = []*AcronymCase{
	//           camelize             underscore            humanize              titleize
	&AcronymCase{"API", "api", "API", "API"},
	&AcronymCase{"APIController", "api_controller", "API controller", "API Controller"},
	&AcronymCase{"Nokogiri::HTML", "nokogiri/html", "Nokogiri/HTML", "Nokogiri/HTML"},
	&AcronymCase{"HTTPAPI", "http_api", "HTTP API", "HTTP API"},
	&AcronymCase{"HTTP::Get", "http/get", "HTTP/get", "HTTP/Get"},
	&AcronymCase{"SSLError", "ssl_error", "SSL error", "SSL Error"},
	&AcronymCase{"RESTful", "restful", "RESTful", "RESTful"},
	&AcronymCase{"RESTfulController", "restful_controller", "RESTful controller", "RESTful Controller"},
	&AcronymCase{"IHeartW3C", "i_heart_w3c", "I heart W3C", "I Heart W3C"},
	&AcronymCase{"PhDRequired", "phd_required", "PhD required", "PhD Required"},
	&AcronymCase{"IRoRU", "i_ror_u", "I RoR u", "I RoR U"},
	&AcronymCase{"RESTfulHTTPAPI", "restful_http_api", "RESTful HTTP API", "RESTful HTTP API"},
	// misdirection
	&AcronymCase{"Capistrano", "capistrano", "Capistrano", "Capistrano"},
	&AcronymCase{"CapiController", "capi_controller", "Capi controller", "Capi Controller"},
	&AcronymCase{"HttpsApis", "https_apis", "Https apis", "Https Apis"},
	&AcronymCase{"Html5", "html5", "Html5", "Html5"},
	&AcronymCase{"Restfully", "restfully", "Restfully", "Restfully"},
	&AcronymCase{"RoRails", "ro_rails", "Ro rails", "Ro Rails"},
}

// tests

func TestPluralizePlurals(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "plurals", rs.Pluralize("plurals"); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := "Plurals", rs.Pluralize("Plurals"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestPluralizeEmptyString(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "", rs.Pluralize(""); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestUncountables(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for _, rule := range rs.uncountables {
		word := rule.match
		if got := rs.Singularize(word); got != word {
			t.Error("want", word, "got", got)
		}
		if got := rs.Pluralize(word); got != word {
			t.Error("want", word, "got", got)
		}
	}
}

func TestUncountableWordIsNotGreedy(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	uncountableWord := "ors"
	countableWord := "sponsor"

	rs.AddUncountable(uncountableWord)

	if want, got := uncountableWord, rs.Singularize(uncountableWord); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := uncountableWord, rs.Pluralize(uncountableWord); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := rs.Pluralize(uncountableWord), rs.Singularize(uncountableWord); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := "sponsor", rs.Singularize(countableWord); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := "sponsors", rs.Pluralize(countableWord); got != want {
		t.Error("want", want, "got", got)
	}
	if want, got := "sponsor", rs.Singularize(rs.Pluralize(countableWord)); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestPluralizeSingular(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for singular, plural := range SingularToPlural {
		if want, got := plural, rs.Pluralize(singular); got != want {
			t.Error("want", want, "got", got)
		}
		if want, got := rs.Capitalize(plural), rs.Capitalize(rs.Pluralize(singular)); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestSingularizePlural(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for singular, plural := range SingularToPlural {
		if want, got := singular, rs.Singularize(plural); got != want {
			t.Error("want", want, "got", got)
		}
		if want, got := rs.Capitalize(singular), rs.Capitalize(rs.Singularize(plural)); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestPluralizePlural(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for _, plural := range SingularToPlural {
		if want, got := plural, rs.Pluralize(plural); got != want {
			t.Error("want", want, "got", got)
		}
		if want, got := rs.Capitalize(plural), rs.Capitalize(rs.Pluralize(plural)); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestOverwritePreviousInflectors(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "series", rs.Singularize("series"); got != want {
		t.Error("want", want, "got", got)
	}
	// previously exact wasnt needed because add would *remove* from uncountables
	// but since we want to be able to easily "rollback" additions to rules
	// the algorithm has changed slightly
	rs.AddSingularExact("series", "replaced", true)
	if want, got := "replaced", rs.Singularize("series"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestTitleize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for before, titleized := range MixtureToTitleCase {
		if want, got := titleized, rs.Titleize(before); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestCapitalize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for lower, capitalized := range CapitalizeMixture {
		if want, got := capitalized, rs.Capitalize(lower); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestCamelize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for camel, underscore := range CamelToUnderscore {
		if want, got := camel, rs.Camelize(underscore); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestCamelizeWithLowerDowncasesTheFirstLetter(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "capital", rs.CamelizeDownFirst("Capital"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestCamelizeWithUnderscores(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "CamelCase", rs.Camelize("Camel_Case"); got != want {
		t.Error("want", want, "got", got)
	}
}

// func TestAcronyms(t *testing.T) {
// 	rs := AddDefaultRules(&Ruleset{})
// 	rs.AddAcronym("API")
// 	rs.AddAcronym("HTML")
// 	rs.AddAcronym("HTTP")
// 	rs.AddAcronym("RESTful")
// 	rs.AddAcronym("W3C")
// 	rs.AddAcronym("PhD")
// 	rs.AddAcronym("RoR")
// 	rs.AddAcronym("SSL")
// 	// each in table
// 	for _, x := range AcronymCases {
// 		if want, got := x.camel, rs.Camelize(x.under); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.camel, rs.Camelize(x.camel); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.under, rs.Underscore(x.under); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.under, rs.Underscore(x.camel); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.title, rs.Titleize(x.under); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.title, rs.Titleize(x.camel); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 		if want, got := x.human, Humanize(x.under); got != want {
// 			t.Error("want", want, "got", got)
// 		}
// 	}
// }

// func TestAcronymOverride(t *testing.T) {
// rs := AddDefaultRules(&Ruleset{})
//     rs.AddAcronym("API")
//     rs.AddAcronym("LegacyApi")
//     if want, got := "LegacyApi", rs.Camelize("legacyapi"))
//     if want, got := "LegacyAPI", rs.Camelize("legacy_api"))
//     if want, got := "SomeLegacyApi", rs.Camelize("some_legacyapi"))
//     if want, got := "Nonlegacyapi", rs.Camelize("nonlegacyapi"))
// }

// func TestAcronymsCamelizeLower(t *testing.T) {
// rs := AddDefaultRules(&Ruleset{})
//     rs.AddAcronym("API")
//     rs.AddAcronym("HTML")
//     if want, got := "htmlAPI", CamelizeDownFirst("html_api"))
//     if want, got := "htmlAPI", CamelizeDownFirst("htmlAPI"))
//     if want, got := "htmlAPI", CamelizeDownFirst("HTMLAPI"))
// }

func TestUnderscoreAcronymSequence(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	rs.AddAcronym("API")
	rs.AddAcronym("HTML5")
	rs.AddAcronym("HTML")
	if want, got := "html5_html_api", rs.Underscore("HTML5HTMLAPI"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestUnderscore(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	rs.AddAcronym("HTML")
	for camel, underscore := range CamelToUnderscore {
		if want, got := underscore, rs.Underscore(camel); got != want {
			t.Error("want", want, "got", got)
		}
	}
	for camel, underscore := range CamelToUnderscoreWithoutReverse {
		if want, got := underscore, rs.Underscore(camel); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestForeignKey(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for klass, foreignKey := range ClassNameToForeignKeyWithUnderscore {
		if want, got := foreignKey, rs.ForeignKey(klass); got != want {
			t.Error("want", want, "got", got)
		}
	}
	for word, foreignKey := range PluralToForeignKeyWithUnderscore {
		if want, got := foreignKey, rs.ForeignKey(word); got != want {
			t.Error("want", want, "got", got)
		}
	}
	for klass, foreignKey := range ClassNameToForeignKeyWithoutUnderscore {
		if want, got := foreignKey, rs.ForeignKeyCondensed(klass); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestTableize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for klass, table := range ClassNameToTableName {
		if want, got := table, rs.Tableize(klass); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestParameterize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for str, parameterized := range StringToParameterized {
		if want, got := parameterized, rs.Parameterize(str); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestParameterizeAndNormalize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for str, parameterized := range StringToParameterizedAndNormalized {
		if want, got := parameterized, rs.Parameterize(str); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestParameterizeWithCustomSeparator(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for str, parameterized := range StringToParameterizeWithUnderscore {
		if want, got := parameterized, rs.ParameterizeJoin(str, "_"); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestTypeify(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for klass, table := range ClassNameToTableName {
		if want, got := klass, rs.Typeify(table); got != want {
			t.Error("want", want, "got", got)
		}
		if want, got := klass, rs.Typeify("table_prefix."+table); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestTypeifyWithLeadingSchemaName(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	if want, got := "FooBar", rs.Typeify("schema.foo_bar"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestHumanize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for underscore, human := range UnderscoreToHuman {
		if want, got := human, rs.Humanize(underscore); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestHumanizeByString(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	rs.AddHuman("col_rpted_bugs", "reported bugs")
	if want, got := "90 reported bugs recently", rs.Humanize("90 col_rpted_bugs recently"); got != want {
		t.Error("want", want, "got", got)
	}
}

func TestOrdinal(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for number, ordinalized := range OrdinalNumbers {
		if want, got := ordinalized, rs.Ordinalize(number); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestDasherize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for underscored, dasherized := range UnderscoresToDashes {
		if want, got := dasherized, rs.Dasherize(underscored); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestUnderscoreAsReverseOfDasherize(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for underscored := range UnderscoresToDashes {
		if want, got := underscored, rs.Underscore(rs.Dasherize(underscored)); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestUnderscoreToLowerCamel(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for underscored, lower := range UnderscoreToLowerCamel {
		if want, got := lower, rs.CamelizeDownFirst(underscored); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestIrregularityBetweenSingularAndPlural(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for singular, plural := range Irregularities {
		rs.AddIrregular(singular, plural)
		if want, got := singular, rs.Singularize(plural); got != want {
			t.Error("want", want, "got", got)
		}
		if want, got := plural, rs.Pluralize(singular); got != want {
			t.Error("want", want, "got", got)
		}
	}
}

func TestPluralizeOfIrregularity(t *testing.T) {
	rs := AddDefaultRules(&Ruleset{})
	for singular, plural := range Irregularities {
		rs.AddIrregular(singular, plural)
		if want, got := plural, rs.Pluralize(plural); got != want {
			t.Error("want", want, "got", got)
		}
	}
}
