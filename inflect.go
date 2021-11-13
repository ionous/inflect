//
package inflect

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Rule for transforming a single word.
type Rule struct {
	match, sub string
	exact      bool
}

// Ruleset of multiple word transformations..
type Ruleset struct {
	plurals, singulars, humans, acronyms, uncountables []Rule
}

// Rules - a default set of transformations.
// initialized at startup with AddDefaultRules() automatically.
var Rules Ruleset

func init() {
	AddDefaultRules(&Rules)
}

// AddDefaultRules of common English pluralization to the passed rules.
// Returns the same ruleset for easier statement chaining
func AddDefaultRules(rs *Ruleset) *Ruleset {
	rs.plurals = append(rs.plurals, []Rule{
		{match: "s", sub: "s"},
		{match: "testis", sub: "testes"},
		{match: "axis", sub: "axes"},
		{match: "octopus", sub: "octopi"},
		{match: "virus", sub: "viri"},
		{match: "octopi", sub: "octopi"},
		{match: "viri", sub: "viri"},
		{match: "alias", sub: "aliases"},
		{match: "status", sub: "statuses"},
		{match: "bus", sub: "buses"},
		{match: "buffalo", sub: "buffaloes"},
		{match: "tomato", sub: "tomatoes"},
		{match: "tum", sub: "ta"},
		{match: "ium", sub: "ia"},
		{match: "ta", sub: "ta"},
		{match: "ia", sub: "ia"},
		{match: "sis", sub: "ses"},
		{match: "lf", sub: "lves"},
		{match: "rf", sub: "rves"},
		{match: "afe", sub: "aves"},
		{match: "bfe", sub: "bves"},
		{match: "cfe", sub: "cves"},
		{match: "dfe", sub: "dves"},
		{match: "efe", sub: "eves"},
		{match: "gfe", sub: "gves"},
		{match: "hfe", sub: "hves"},
		{match: "ife", sub: "ives"},
		{match: "jfe", sub: "jves"},
		{match: "kfe", sub: "kves"},
		{match: "lfe", sub: "lves"},
		{match: "mfe", sub: "mves"},
		{match: "nfe", sub: "nves"},
		{match: "ofe", sub: "oves"},
		{match: "pfe", sub: "pves"},
		{match: "qfe", sub: "qves"},
		{match: "rfe", sub: "rves"},
		{match: "sfe", sub: "sves"},
		{match: "tfe", sub: "tves"},
		{match: "ufe", sub: "uves"},
		{match: "vfe", sub: "vves"},
		{match: "wfe", sub: "wves"},
		{match: "xfe", sub: "xves"},
		{match: "yfe", sub: "yves"},
		{match: "zfe", sub: "zves"},
		{match: "hive", sub: "hives"},
		{match: "quy", sub: "quies"},
		{match: "by", sub: "bies"},
		{match: "cy", sub: "cies"},
		{match: "dy", sub: "dies"},
		{match: "fy", sub: "fies"},
		{match: "gy", sub: "gies"},
		{match: "hy", sub: "hies"},
		{match: "jy", sub: "jies"},
		{match: "ky", sub: "kies"},
		{match: "ly", sub: "lies"},
		{match: "my", sub: "mies"},
		{match: "ny", sub: "nies"},
		{match: "py", sub: "pies"},
		{match: "qy", sub: "qies"},
		{match: "ry", sub: "ries"},
		{match: "sy", sub: "sies"},
		{match: "ty", sub: "ties"},
		{match: "vy", sub: "vies"},
		{match: "wy", sub: "wies"},
		{match: "xy", sub: "xies"},
		{match: "zy", sub: "zies"},
		{match: "x", sub: "xes"},
		{match: "ch", sub: "ches"},
		{match: "ss", sub: "sses"},
		{match: "sh", sub: "shes"},
		{match: "matrix", sub: "matrices"},
		{match: "vertix", sub: "vertices"},
		{match: "indix", sub: "indices"},
		{match: "matrex", sub: "matrices"},
		{match: "vertex", sub: "vertices"},
		{match: "index", sub: "indices"},
		{match: "mouse", sub: "mice"},
		{match: "louse", sub: "lice"},
		{match: "mice", sub: "mice"},
		{match: "lice", sub: "lice"},
		{match: "ox", sub: "oxen", exact: true},
		{match: "oxen", sub: "oxen", exact: true},
		{match: "quiz", sub: "quizzes", exact: true},
	}...)
	rs.singulars = append(rs.singulars, []Rule{
		{match: "s", sub: ""},
		{match: "news", sub: "news"},
		{match: "ta", sub: "tum"},
		{match: "ia", sub: "ium"},
		{match: "analyses", sub: "analysis"},
		{match: "bases", sub: "basis"},
		{match: "diagnoses", sub: "diagnosis"},
		{match: "parentheses", sub: "parenthesis"},
		{match: "prognoses", sub: "prognosis"},
		{match: "synopses", sub: "synopsis"},
		{match: "theses", sub: "thesis"},
		{match: "analyses", sub: "analysis"},
		{match: "aves", sub: "afe"},
		{match: "bves", sub: "bfe"},
		{match: "cves", sub: "cfe"},
		{match: "dves", sub: "dfe"},
		{match: "eves", sub: "efe"},
		{match: "gves", sub: "gfe"},
		{match: "hves", sub: "hfe"},
		{match: "ives", sub: "ife"},
		{match: "jves", sub: "jfe"},
		{match: "kves", sub: "kfe"},
		{match: "lves", sub: "lfe"},
		{match: "mves", sub: "mfe"},
		{match: "nves", sub: "nfe"},
		{match: "oves", sub: "ofe"},
		{match: "pves", sub: "pfe"},
		{match: "qves", sub: "qfe"},
		{match: "rves", sub: "rfe"},
		{match: "sves", sub: "sfe"},
		{match: "tves", sub: "tfe"},
		{match: "uves", sub: "ufe"},
		{match: "vves", sub: "vfe"},
		{match: "wves", sub: "wfe"},
		{match: "xves", sub: "xfe"},
		{match: "yves", sub: "yfe"},
		{match: "zves", sub: "zfe"},
		{match: "hives", sub: "hive"},
		{match: "tives", sub: "tive"},
		{match: "lves", sub: "lf"},
		{match: "rves", sub: "rf"},
		{match: "quies", sub: "quy"},
		{match: "bies", sub: "by"},
		{match: "cies", sub: "cy"},
		{match: "dies", sub: "dy"},
		{match: "fies", sub: "fy"},
		{match: "gies", sub: "gy"},
		{match: "hies", sub: "hy"},
		{match: "jies", sub: "jy"},
		{match: "kies", sub: "ky"},
		{match: "lies", sub: "ly"},
		{match: "mies", sub: "my"},
		{match: "nies", sub: "ny"},
		{match: "pies", sub: "py"},
		{match: "qies", sub: "qy"},
		{match: "ries", sub: "ry"},
		{match: "sies", sub: "sy"},
		{match: "ties", sub: "ty"},
		{match: "vies", sub: "vy"},
		{match: "wies", sub: "wy"},
		{match: "xies", sub: "xy"},
		{match: "zies", sub: "zy"},
		{match: "series", sub: "series"},
		{match: "movies", sub: "movie"},
		{match: "xes", sub: "x"},
		{match: "ches", sub: "ch"},
		{match: "sses", sub: "ss"},
		{match: "shes", sub: "sh"},
		{match: "mice", sub: "mouse"},
		{match: "lice", sub: "louse"},
		{match: "buses", sub: "bus"},
		{match: "oes", sub: "o"},
		{match: "shoes", sub: "shoe"},
		{match: "crises", sub: "crisis"},
		{match: "axes", sub: "axis"},
		{match: "testes", sub: "testis"},
		{match: "octopi", sub: "octopus"},
		{match: "viri", sub: "virus"},
		{match: "statuses", sub: "status"},
		{match: "aliases", sub: "alias"},
		{match: "oxen", sub: "ox", exact: true},
		{match: "vertices", sub: "vertex"},
		{match: "indices", sub: "index"},
		{match: "matrices", sub: "matrix"},
		{match: "quizzes", sub: "quiz", exact: true},
		{match: "databases", sub: "database"},
	}...)
	rs.AddIrregular("person", "people")
	rs.AddIrregular("man", "men")
	rs.AddIrregular("child", "children")
	rs.AddIrregular("sex", "sexes")
	rs.AddIrregular("move", "moves")
	rs.AddIrregular("zombie", "zombies")
	rs.uncountables = append(rs.uncountables, []Rule{
		{exact: true, match: "equipment"},
		{exact: true, match: "fish"},
		{exact: true, match: "information"},
		{exact: true, match: "jeans"},
		{exact: true, match: "money"},
		{exact: true, match: "police"},
		{exact: true, match: "rice"},
		{exact: true, match: "series"},
		{exact: true, match: "sheep"},
		{exact: true, match: "species"},
	}...)

	return rs
}

// add a pluralization rule
func (rs *Ruleset) AddPlural(suffix, replacement string) {
	rs.AddPluralExact(suffix, replacement, false)
}

// add a pluralization rule with full string match
func (rs *Ruleset) AddPluralExact(suffix, replacement string, exact bool) {
	rs.plurals = append(rs.plurals, Rule{
		match: suffix,
		sub:   replacement,
		exact: exact,
	})
}

// add a singular rule
func (rs *Ruleset) AddSingular(suffix, replacement string) {
	rs.AddSingularExact(suffix, replacement, false)
}

// same as AddSingular but you can set `exact` to force  a full string match
func (rs *Ruleset) AddSingularExact(suffix, replacement string, exact bool) {
	rs.singulars = append(rs.singulars, Rule{
		match: suffix,
		sub:   replacement,
		exact: exact,
	})
}

// Human rules are applied by humanize to show more friendly versions of words
func (rs *Ruleset) AddHuman(suffix, replacement string) {
	rs.humans = append(rs.humans, Rule{
		match: suffix,
		sub:   replacement,
	})
}

// Add any inconsistent plural/singular rules to the set here.
func (rs *Ruleset) AddIrregular(singular, plural string) {
	rs.AddPlural(singular, plural)
	rs.AddPlural(plural, plural)
	rs.AddSingular(plural, singular)
}

// if you use acronym you may need to add them to the ruleset
// to prevent Underscored words of things like "HTML" coming out
// as "h_t_m_l"
func (rs *Ruleset) AddAcronym(word string) {
	rs.acronyms = append(rs.acronyms, Rule{
		match: word,
		sub:   rs.Titleize(strings.ToLower(word)),
	})
}

// add a word to this ruleset that has the same singular and plural form
// for example: "rice"
func (rs *Ruleset) AddUncountable(word string) {
	rs.uncountables = append(rs.uncountables, Rule{match: word, exact: true})
}

// handle multiple words by using the last one
func (rs *Ruleset) isUncountable(word string) bool {
	words := strings.Split(word, " ")
	last := strings.ToLower(words[len(words)-1])
	_, exact := find(rs.uncountables, last)
	return exact
}

// returns the plural form of a singular word
func (rs *Ruleset) Pluralize(word string) (ret string) {
	if len(word) > 0 {
		if p, exact := find(rs.plurals, word); exact {
			ret = p
		} else if rs.isUncountable(word) {
			ret = word
		} else if len(p) > 0 {
			ret = p // inexact match
		} else {
			ret = word + "s"
		}
	}
	return
}

func find(rules []Rule, word string) (ret string, exact bool) {
	for i := len(rules) - 1; i >= 0; i-- {
		if rule := rules[i]; rule.exact {
			if word == rule.match {
				ret = rule.sub
				exact = true
				break
			}
		} else {
			if trimmed := strings.TrimSuffix(word, rule.match); len(trimmed) < len(word) {
				ret = trimmed + rule.sub
				break
			}
		}
	}
	return
}

// returns the singular form of a plural word
func (rs *Ruleset) Singularize(word string) (ret string) {
	if len(word) > 0 {
		if p, exact := find(rs.singulars, word); exact {
			ret = p
		} else if rs.isUncountable(word) {
			ret = word
		} else if len(p) > 0 {
			ret = p // inexact match
		} else {
			ret = word
		}
	}
	return
}

// uppercase first character
func (rs *Ruleset) Capitalize(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}

// "dino_party" -> "DinoParty"
func (rs *Ruleset) Camelize(word string) string {
	return splitAtCaseChange(word, "", true)
}

// same as Camelcase but with first letter downcased
func (rs *Ruleset) CamelizeDownFirst(word string) string {
	word = rs.Camelize(word)
	return strings.ToLower(word[:1]) + word[1:]
}

// Capitalize every word in sentance "hello there" -> "Hello There"
func (rs *Ruleset) Titleize(word string) string {
	return splitAtCaseChange(word, " ", true)
}

func (rs *Ruleset) safeCaseAcronyms(word string) string {
	// convert an acroymn like HTML into Html
	// forward searches not sure why.
	for _, rule := range rs.acronyms {
		word = strings.Replace(word, rule.match, rule.sub, -1)
	}
	return word
}

func (rs *Ruleset) seperatedWords(word, sep string) string {
	word = rs.safeCaseAcronyms(word)
	return splitAtCaseChange(word, sep, false)
}

// Underscore lowercase version "BigBen" -> "big_ben"
func (rs *Ruleset) Underscore(word string) string {
	return rs.seperatedWords(word, "_")
}

// First letter of sentence capitalized
// Uses custom friendly replacements via AddHuman()
func (rs *Ruleset) Humanize(word string) string {
	if trimmed := strings.TrimSuffix(word, "_id"); len(trimmed) < len(word) {
		word = trimmed // strip foreign key kinds
	}
	// replace and strings in humans list
	for i := len(rs.humans) - 1; i >= 0; i-- {
		rule := rs.humans[i]
		word = strings.Replace(word, rule.match, rule.sub, -1)
	}
	sentance := rs.seperatedWords(word, " ")
	return strings.ToUpper(sentance[:1]) + sentance[1:]
}

// an underscored foreign key name "Person" -> "person_id"
func (rs *Ruleset) ForeignKey(word string) string {
	return rs.Underscore(rs.Singularize(word)) + "_id"
}

// a foreign key (with an underscore) "Person" -> "personid"
func (rs *Ruleset) ForeignKeyCondensed(word string) string {
	return rs.Underscore(word) + "id"
}

// Rails style pluralized table names: "SuperPerson" -> "super_people"
func (rs *Ruleset) Tableize(word string) string {
	return rs.Pluralize(rs.Underscore(rs.Typeify(word)))
}

var notUrlSafe *regexp.Regexp = regexp.MustCompile(`[^\w\d\-_ ]`)

// param safe dasherized names like "my-param"
func (rs *Ruleset) Parameterize(word string) string {
	return rs.ParameterizeJoin(word, "-")
}

// param safe dasherized names with custom seperator
func (rs *Ruleset) ParameterizeJoin(word, sep string) string {
	word = strings.ToLower(word)
	word = rs.Asciify(word)
	word = notUrlSafe.ReplaceAllString(word, "")
	word = strings.Replace(word, " ", sep, -1)
	if len(sep) > 0 {
		squash, err := regexp.Compile(sep + "+")
		if err == nil {
			word = squash.ReplaceAllString(word, sep)
		}
	}
	word = strings.Trim(word, sep+" ")
	return word
}

type lookalike struct {
	sub   string
	match *regexp.Regexp
}

var lookalikes []lookalike

// transforms latin characters like é -> e
func (rs *Ruleset) Asciify(word string) string {
	if len(lookalikes) == 0 {
		lookalikes = []lookalike{
			{"A", regexp.MustCompile(`À|Á|Â|Ã|Ä|Å`)},
			{"AE", regexp.MustCompile(`Æ`)},
			{"C", regexp.MustCompile(`Ç`)},
			{"E", regexp.MustCompile(`È|É|Ê|Ë`)},
			{"G", regexp.MustCompile(`Ğ`)},
			{"I", regexp.MustCompile(`Ì|Í|Î|Ï|İ`)},
			{"N", regexp.MustCompile(`Ñ`)},
			{"O", regexp.MustCompile(`Ò|Ó|Ô|Õ|Ö|Ø`)},
			{"S", regexp.MustCompile(`Ş`)},
			{"U", regexp.MustCompile(`Ù|Ú|Û|Ü`)},
			{"Y", regexp.MustCompile(`Ý`)},
			{"ss", regexp.MustCompile(`ß`)},
			{"a", regexp.MustCompile(`à|á|â|ã|ä|å`)},
			{"ae", regexp.MustCompile(`æ`)},
			{"c", regexp.MustCompile(`ç`)},
			{"e", regexp.MustCompile(`è|é|ê|ë`)},
			{"g", regexp.MustCompile(`ğ`)},
			{"i", regexp.MustCompile(`ì|í|î|ï|ı`)},
			{"n", regexp.MustCompile(`ñ`)},
			{"o", regexp.MustCompile(`ò|ó|ô|õ|ö|ø`)},
			{"s", regexp.MustCompile(`ş`)},
			{"u", regexp.MustCompile(`ù|ú|û|ü|ũ|ū|ŭ|ů|ű|ų`)},
			{"y", regexp.MustCompile(`ý|ÿ`)},
		}
	}
	for _, el := range lookalikes {
		word = el.match.ReplaceAllString(word, el.sub)
	}
	return word
}

var tablePrefix *regexp.Regexp = regexp.MustCompile(`^[^.]*\.`)

// "something_like_this" -> "SomethingLikeThis"
func (rs *Ruleset) Typeify(word string) string {
	word = tablePrefix.ReplaceAllString(word, "")
	return rs.Camelize(rs.Singularize(word))
}

// "SomeText" -> "some-text"
func (rs *Ruleset) Dasherize(word string) string {
	return rs.seperatedWords(word, "-")
}

// "1031" -> "1031st"
func (rs *Ruleset) Ordinalize(str string) (ret string) {
	if number, e := strconv.Atoi(str); e != nil {
		ret = str
	} else {
		var ext string
		switch abs(number) % 100 {
		case 11, 12, 13:
			ext = "th"
		default:
			switch abs(number) % 10 {
			case 1:
				ext = "st"
			case 2:
				ext = "nd"
			case 3:
				ext = "rd"
			default:
				ext = "th"
			}
		}
		ret = str + ext
	}
	return ret
}

func AddPlural(suffix, replacement string) {
	Rules.AddPlural(suffix, replacement)
}

func AddSingular(suffix, replacement string) {
	Rules.AddSingular(suffix, replacement)
}

func AddHuman(suffix, replacement string) {
	Rules.AddHuman(suffix, replacement)
}

func AddIrregular(singular, plural string) {
	Rules.AddIrregular(singular, plural)
}

func AddAcronym(word string) {
	Rules.AddAcronym(word)
}

func AddUncountable(word string) {
	Rules.AddUncountable(word)
}

func Pluralize(word string) string {
	return Rules.Pluralize(word)
}

func Singularize(word string) string {
	return Rules.Singularize(word)
}

func Capitalize(word string) string {
	return Rules.Capitalize(word)
}

func Camelize(word string) string {
	return Rules.Camelize(word)
}

func CamelizeDownFirst(word string) string {
	return Rules.CamelizeDownFirst(word)
}

func Titleize(word string) string {
	return Rules.Titleize(word)
}

func Underscore(word string) string {
	return Rules.Underscore(word)
}

func Humanize(word string) string {
	return Rules.Humanize(word)
}

func ForeignKey(word string) string {
	return Rules.ForeignKey(word)
}

func ForeignKeyCondensed(word string) string {
	return Rules.ForeignKeyCondensed(word)
}

func Tableize(word string) string {
	return Rules.Tableize(word)
}

func Parameterize(word string) string {
	return Rules.Parameterize(word)
}

func ParameterizeJoin(word, sep string) string {
	return Rules.ParameterizeJoin(word, sep)
}

func Typeify(word string) string {
	return Rules.Typeify(word)
}

func Dasherize(word string) string {
	return Rules.Dasherize(word)
}

func Ordinalize(word string) string {
	return Rules.Ordinalize(word)
}

func Asciify(word string) string {
	return Rules.Asciify(word)
}

// helper funcs

func isSpacerChar(c rune) bool {
	return c == '_' ||
		c == ' ' ||
		c == ':' ||
		c == '-'
}

func splitAtCaseChange(s, sep string, allowCaps bool) string {
	var word, words strings.Builder
	for _, c := range s {
		spacer := isSpacerChar(c)
		lower := unicode.ToLower(c)
		if word.Len() > 0 {
			if spacer || lower != c {
				if words.Len() > 0 {
					words.WriteString(sep)
				}
				words.WriteString(word.String())
				word.Reset()
			}
		}
		if !spacer {
			if !allowCaps || word.Len() > 0 {
				word.WriteRune(lower) // write lower case in the middle of the string
			} else if lower != c {
				word.WriteRune(c) // on edges, if c was already uppercase; great.
			} else {
				word.WriteRune(unicode.ToUpper(c)) // on edge, if c was lower, uppercase it.
			}
		}
	}
	// flow the last pending bit... always at least one rune long
	if words.Len() > 0 {
		words.WriteString(sep)
	}
	words.WriteString(word.String())
	return words.String()
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
