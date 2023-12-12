package str

import (
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

type Str struct {
	sb *strings.Builder
}

// New creates a new instance of Str.
func New(value string) *Str {
	var sb strings.Builder
	sb.WriteString(value)

	return &Str{sb: &sb}
}

func (s *Str) getValue() string {
	return s.sb.String()
}

func (s *Str) setValue(value ...string) {
	s.sb.Reset()

	for _, v := range value {
		s.sb.WriteString(v)
	}
}

// After returns a new instance of Str containing the substring after the first occurrence of the given substring.
func (s *Str) After(substring string) *Str {
	v := s.getValue()

	index := strings.Index(v, substring)
	if index != -1 {
		s.setValue(v[index+len(substring):])
	}

	return s
}

// AfterLast returns a new instance of Str containing the substring after the last occurrence of the given substring.
func (s *Str) AfterLast(substring string) *Str {
	v := s.getValue()

	index := strings.LastIndex(v, substring)
	if index != -1 {
		s.setValue(v[index+len(substring):])
	}

	return s
}

// Append appends the given string to the end of the current string and returns a pointer to the current instance of Str.
func (s *Str) Append(value string) *Str {
	s.sb.WriteString(value)
	return s
}

// At returns the character at the given index as a string.
func (s *Str) At(index int) string {
	if index >= 0 && index < s.sb.Len() {
		return string([]rune(s.getValue())[index])
	}

	return ""
}

// Before returns a new instance of Str containing the substring before the first occurrence of the given substring.
func (s *Str) Before(substring string) *Str {
	v := s.getValue()

	index := strings.Index(v, substring)
	if index != -1 {
		s.setValue(v[:index])
	}

	return s
}

// BeforeLast returns a new instance of Str containing the substring before the last occurrence of the given substring.
func (s *Str) BeforeLast(substring string) *Str {
	v := s.getValue()

	index := strings.LastIndex(v, substring)
	if index != -1 {
		s.setValue(v[:index])
	}

	return s
}

// Between returns a new instance of Str containing the substring between the first occurrence of the "from" substring and the first occurrence of the "to" substring.
func (s *Str) Between(from string, to string) *Str {
	return s.After(from).Before(to)
}

// Camel returns a new instance of Str in camelCase.
func (s *Str) Camel() *Str {
	v := s.getValue()

	if v == "" {
		return s
	}

	value := strings.Fields(v)
	first := strings.ToLower(value[0])
	new := first

	for _, word := range value[1:] {
		new += capitalize(word)
	}

	s.setValue(new)
	return s
}

// Capitalize returns a new instance of Str with the first character capitalized.
func (s *Str) Capitalize() *Str {
	s.setValue(capitalize(s.getValue()))
	return s
}

// Contains returns true if the current string contains the given substring.
func (s *Str) Contains(substring string) bool {
	return contains(true, s.getValue(), substring)
}

// ContainsAny returns true if the current string contains any of the given substrings.
func (s *Str) ContainsAny(substring []string) bool {
	return contains(true, s.getValue(), substring...)
}

// ContainsAll returns true if the current string contains all of the given substrings.
func (s *Str) ContainsAll(substrings []string) bool {
	return contains(false, s.getValue(), substrings...)
}

// Clone returns a new instance of Str with the same value as the current instance.
func (s *Str) Clone() *Str {
	return New(s.getValue())
}

// Count returns the number of occurrences of the given substring in the current string.
func (s *Str) Count(substring string) int {
	return strings.Count(s.getValue(), substring)
}

// EndsWith returns true if the current string ends with the given substring.
func (s *Str) EndsWith(substring string) bool {
	v := s.getValue()

	index := strings.LastIndex(v, substring)

	if index == -1 {
		return false
	}

	return index == -1 || v[index:] == substring && substring != ""
}

// Exactly returns true if the current string is exactly the same as the given string.
func (s *Str) Exactly(value string) bool {
	return s.getValue() == value
}

// First returns s new instance of Str containing the first character of the current string.
func (s *Str) First() *Str {
	s.setValue(s.At(0))

	return s
}

// FirstWord returns a new instance of Str containing the first word of the current string.
func (s *Str) FirstWord() *Str {
	words := strings.Fields(s.getValue())

	if len(words) > 0 {
		s.setValue(words[0])
	}

	return s
}

// Last returns s new instance of Str containing the last character of the current string.
func (s *Str) Last() *Str {
	s.setValue(s.At(utf8.RuneCountInString(s.getValue()) - 1))

	return s
}

// LastWord returns a new instance of Str containing the last word of the current string.
func (s *Str) LastWord() *Str {
	v := s.getValue()

	if v != "" {
		s.setValue(strings.Fields(v)[len(strings.Fields(v))-1])
	}

	return s
}

// Length returns the length of the current string.
func (s *Str) Length() int {
	return utf8.RuneCountInString(s.getValue())
}

// Lower returns a new instance of Str with all characters in lowercase.
func (s *Str) Lower() *Str {
	s.setValue(strings.ToLower(s.getValue()))
	return s
}

// IsEmpty returns true if the current string is empty.
func (s *Str) IsEmpty() bool {
	return s.getValue() == ""
}

// IsJson returns true if the current string is valid JSON.
func (s *Str) IsJson() bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s.getValue()), &js) == nil
}

// IsNotEmpty returns true if the current string is not empty.
func (s *Str) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// IsUrl returns true if the current string is a valid URL.
func (s *Str) IsUrl() bool {
	var schemes = []string{
		"aaa", "aaas", "about", "acap", "acct", "acd", "acr", "adiumxtra", "adt", "afp", "afs", "aim", "amss", "android", "appdata", "apt", "ark", "attachment", "aw", "barion", "beshare", "bitcoin", "bitcoincash", "blob", "bolo", "browserext", "calculator", "callto", "cap", "cast", "casts", "chrome", "chrome-extension", "cid", "coap", "coap\\+tcp", "coap\\+ws", "coaps", "coaps\\+tcp", "coaps\\+ws", "com-eventbrite-attendee", "content", "conti", "crid", "cvs", "dab", "data", "dav", "diaspora", "dict", "did", "dis", "dlna-playcontainer", "dlna-playsingle", "dns", "dntp", "dpp", "drm", "drop", "dtn", "dvb", "ed2k", "elsi", "example", "facetime", "fax", "feed", "feedready", "file", "filesystem", "finger", "first-run-pen-experience", "fish", "fm", "ftp", "fuchsia-pkg", "geo", "gg", "git", "gizmoproject", "go", "gopher", "graph", "gtalk", "h323", "ham", "hcap", "hcp", "http", "https", "hxxp", "hxxps", "hydrazone", "iax", "icap", "icon", "im", "imap", "info", "iotdisco", "ipn", "ipp", "ipps", "irc", "irc6", "ircs", "iris", "iris\\.beep", "iris\\.lwz", "iris\\.xpc", "iris\\.xpcs", "isostore", "itms", "jabber", "jar", "jms", "keyparc", "lastfm", "ldap", "ldaps", "leaptofrogans", "lorawan", "lvlt", "magnet", "mailserver", "mailto", "maps", "market", "message", "mid", "mms", "modem", "mongodb", "moz", "ms-access", "ms-browser-extension", "ms-calculator", "ms-drive-to", "ms-enrollment", "ms-excel", "ms-eyecontrolspeech", "ms-gamebarservices", "ms-gamingoverlay", "ms-getoffice", "ms-help", "ms-infopath", "ms-inputapp", "ms-lockscreencomponent-config", "ms-media-stream-id", "ms-mixedrealitycapture", "ms-mobileplans", "ms-officeapp", "ms-people", "ms-project", "ms-powerpoint", "ms-publisher", "ms-restoretabcompanion", "ms-screenclip", "ms-screensketch", "ms-search", "ms-search-repair", "ms-secondary-screen-controller", "ms-secondary-screen-setup", "ms-settings", "ms-settings-airplanemode", "ms-settings-bluetooth", "ms-settings-camera", "ms-settings-cellular", "ms-settings-cloudstorage", "ms-settings-connectabledevices", "ms-settings-displays-topology", "ms-settings-emailandaccounts", "ms-settings-language", "ms-settings-location", "ms-settings-lock", "ms-settings-nfctransactions", "ms-settings-notifications", "ms-settings-power", "ms-settings-privacy", "ms-settings-proximity", "ms-settings-screenrotation", "ms-settings-wifi", "ms-settings-workplace", "ms-spd", "ms-sttoverlay", "ms-transit-to", "ms-useractivityset", "ms-virtualtouchpad", "ms-visio", "ms-walk-to", "ms-whiteboard", "ms-whiteboard-cmd", "ms-word", "msnim", "msrp", "msrps", "mss", "mtqp", "mumble", "mupdate", "mvn", "news", "nfs", "ni", "nih", "nntp", "notes", "ocf", "oid", "onenote", "onenote-cmd", "opaquelocktoken", "openpgp4fpr", "pack", "palm", "paparazzi", "payto", "pkcs11", "platform", "pop", "pres", "prospero", "proxy", "pwid", "psyc", "pttp", "qb", "query", "redis", "rediss", "reload", "res", "resource", "rmi", "rsync", "rtmfp", "rtmp", "rtsp", "rtsps", "rtspu", "s3", "secondlife", "service", "session", "sftp", "sgn", "shttp", "sieve", "simpleledger", "sip", "sips", "skype", "smb", "sms", "smtp", "snews", "snmp", "soap\\.beep", "soap\\.beeps", "soldat", "spiffe", "spotify", "ssh", "steam", "stun", "stuns", "submit", "svn", "tag", "teamspeak", "tel", "teliaeid", "telnet", "tftp", "tg", "things", "thismessage", "tip", "tn3270", "tool", "ts3server", "turn", "turns", "tv", "udp", "unreal", "urn", "ut2004", "v-event", "vemmi", "ventrilo", "videotex", "vnc", "view-source", "wais", "webcal", "wpid", "ws", "wss", "wtai", "wyciwyg", "xcon", "xcon-userid", "xfire", "xmlrpc\\.beep", "xmlrpc\\.beeps", "xmpp", "xri", "ymsgr", "z39\\.50", "z39\\.50r", "z39\\.50s",
	}

	v := s.getValue()

	var scheme = strings.Split(v, "://")[0]

	var isValidScheme = false

	for _, p := range schemes {
		if scheme == p {
			isValidScheme = true
			break
		}
	}

	if !isValidScheme {
		return false
	}

	u, err := url.ParseRequestURI(v)

	if err != nil {
		return false
	}

	host, _, err := net.SplitHostPort(u.Host)

	if err != nil || host == "" {
		host = u.Host
	}

	punycode, err := idna.ToASCII(host)

	if err != nil {
		return false
	}

	if punycode == "localhost" {
		return true
	}

	var ip = net.ParseIP(punycode)

	if ip != nil {
		return true
	}

	var tld string
	index := strings.LastIndex(punycode, ".")
	if index != -1 {
		tld = punycode[index+len("."):]
	}

	if tld == "" {
		return false
	}

	regex := regexp.MustCompile(`^[a-zA-Z0-9.-]+$`)

	return regex.MatchString(punycode)
}

// Kebab returns a new lowercased instance of Str with all spaces and underscores replaced with hyphens.
func (s *Str) Kebab() *Str {
	s.setValue(snake(s.getValue(), "-"))

	return s
}

// Pascal returns a new instance of Str in PascalCase.
func (s *Str) Pascal() *Str {
	s.Title().Remove(" ")

	return s
}

// Prepend prepends the given string to the beginning of the current string and returns a pointer to the current instance of Str.
func (s *Str) Prepend(prependStr string) *Str {
	s.setValue(prependStr, s.getValue())

	return s
}

// Repeat returns a new instance of Str with the current string repeated the given number of times.
func (s *Str) Repeat(count int) *Str {
	s.setValue(strings.Repeat(s.getValue(), count))
	return s
}

// Remove returns a new instance of Str with all occurrences of the given substring removed.
func (s *Str) Remove(old string) *Str {
	s.setValue(strings.ReplaceAll(s.getValue(), old, ""))

	return s
}

// Replace returns a new instance of Str with all occurrences of the given substring replaced with the given replacement string.
func (s *Str) Replace(old string, new string) *Str {
	s.setValue(strings.ReplaceAll(s.getValue(), old, new))
	return s
}

func (s *Str) Screaming() *Str {
	v := s.getValue()

	// If already uppercased, we lowercase to avoid splitting on each uppercased character
	if strings.ToUpper(v) == v {
		v = strings.ToLower(v)
	}

	s.setValue(strings.ToUpper(snake(v, "_")))

	return s
}

// Slug returns a new Str with slugified version of the current string. The separator is used to replace all non-alphanumeric characters.
func (s *Str) Slug(separator string) *Str {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	processed := re.ReplaceAllString(s.getValue(), separator)

	processed = strings.Trim(processed, "_")
	processed = strings.Trim(processed, "-")
	processed = strings.Trim(processed, separator)
	processed = strings.ToLower(processed)

	s.setValue(processed)

	return s
}

// Snake returns a new lowercased instance of Str with all spaces and hyphens replaced with underscores.
func (s *Str) Snake() *Str {
	s.setValue(snake(s.getValue(), "_"))

	return s
}

// Split returns a slice of strings containing the substrings of the current string separated by the given separator.
func (s *Str) Split(separator string) []string {
	return strings.Split(s.getValue(), separator)
}

// StartsWith returns true if the current string starts with the given substring.
func (s *Str) StartsWith(substring string) bool {
	return strings.Index(s.getValue(), substring) == 0 && substring != ""
}

// String returns the current value as a string.
func (s *Str) String() string {
	return s.getValue()
}

// Swap returns a new instance of Str where all occurrences of keys are replaces with their corresponding values.
func (s *Str) Swap(pairs map[string]string) *Str {
	for key, value := range pairs {
		s.setValue(strings.ReplaceAll(s.getValue(), key, value))
	}

	return s
}

// Title returns a new instance of Str with the first letter of each word capitalized.
func (s *Str) Title() *Str {
	var parts = strings.Split(s.getValue(), " ")
	var mappedParts []string

	for _, part := range parts {
		mappedParts = append(mappedParts, capitalize(part))
	}

	s.setValue(strings.Join(mappedParts, " "))
	return s
}

// TrimAll returns a new instance of Str with all consecutive spaces replaced with a single space.
func (s *Str) TrimAll() *Str {
	s.setValue(strings.Join(strings.Fields(s.getValue()), " "))

	return s
}

// TrimBoth returns a new instance of Str with all leading and trailing spaces removed.
func (s *Str) TrimBoth() *Str {
	s.setValue(strings.Trim(s.getValue(), "\t\n\r\v\f "))

	return s
}

// TrimLeft returns a new instance of Str with all leading spaces removed.
func (s *Str) TrimLeft() *Str {
	s.setValue(strings.TrimLeft(s.getValue(), "\t\n\r\v\f "))

	return s
}

// TrimRight returns a new instance of Str with all trailing spaces removed.
func (s *Str) TrimRight() *Str {
	s.setValue(strings.TrimRight(s.getValue(), "\t\n\r\v\f "))

	return s
}

// Upper returns a new instance of Str with all characters in uppercase.
func (s *Str) Upper() *Str {
	s.setValue(strings.ToUpper(s.getValue()))

	return s
}

// When invokes the given callback if the given condition is true and returns a pointer to the current instance of Str.
func (s *Str) When(condition bool, callback func(s *Str) *Str) *Str {
	if condition {
		return callback(s)
	}

	return s
}

// WhenContains invokes the given callback if the current string contains the given value and returns a pointer to the current instance of Str.
func (s *Str) WhenContains(value string, callback func(s *Str) *Str) *Str {
	return s.When(s.Contains(value), callback)
}

// WhenContainsAny invokes the given callback if the current string contains any of the given values and returns a pointer to the current instance of Str.
func (s *Str) WhenContainsAny(values []string, callback func(s *Str) *Str) *Str {
	return s.When(s.ContainsAny(values), callback)
}

// WhenContainsAll invokes the given callback if the current string contains all of the given values and returns a pointer to the current instance of Str.
func (s *Str) WhenContainsAll(values []string, callback func(s *Str) *Str) *Str {
	return s.When(s.ContainsAll(values), callback)
}

// WhenEndsWith invokes the given callback if the current string ends with the given value and returns a pointer to the current instance of Str.
func (s *Str) WhenEndsWith(value string, callback func(s *Str) *Str) *Str {
	return s.When(s.EndsWith(value), callback)
}

// WhenEmpty invokes the given closure if the current string is empty and returns a pointer to the current instance of Str.
func (s *Str) WhenEmpty(callback func(s *Str) *Str) *Str {
	return s.When(s.IsEmpty(), callback)
}

// WhenExactly invokes the given callback if the current string is exactly the same as the given value and returns a pointer to the current instance of Str.
func (s *Str) WhenExactly(value string, callback func(s *Str) *Str) *Str {
	return s.When(s.Exactly(value), callback)
}

// WhenFunc invokes the given callback if the given condition is true and returns a pointer to the current instance of Str. The condition function is passed a copy of Str.
func (s *Str) WhenFunc(condition func(s *Str) bool, callback func(s *Str) *Str) *Str {
	return s.When(condition(New(s.getValue())), callback)
}

// WhenNotEmpty invokes the given closure if the current string is not empty and returns a pointer to the current instance of Str.
func (s *Str) WhenNotEmpty(callback func(s *Str) *Str) *Str {
	return s.When(s.IsNotEmpty(), callback)
}

// WhenNotExactly invokes the given callback if the current string is not exactly the same as the given value and returns a pointer to the current instance of Str.
func (s *Str) WhenNotExactly(value string, callback func(s *Str) *Str) *Str {
	return s.When(!s.Exactly(value), callback)
}

// WhenStartsWith invokes the given callback if the current string starts with the given value and returns a pointer to the current instance of Str.
func (s *Str) WhenStartsWith(value string, callback func(s *Str) *Str) *Str {
	return s.When(s.StartsWith(value), callback)
}

// Words returns a slice of strings containing each word of the current string with all consecutive spaces removed.
func (s *Str) Words() []string {
	return strings.Fields(s.getValue())
}

/* Helpers */

func capitalize(s string) string {
	if s == "" {
		return s
	}

	firstRune, size := utf8.DecodeRuneInString(s)

	var sb strings.Builder

	sb.WriteRune(unicode.ToUpper(firstRune))
	sb.WriteString(s[size:])

	return sb.String()
}

func contains(returnOnFirstMatch bool, value string, substrings ...string) bool {
	var found []string

	if len(substrings) == 0 {
		return false
	}

	for _, sub := range substrings {
		if strings.Contains(value, sub) {
			found = append(found, sub)

			if returnOnFirstMatch {
				return true
			}
		}
	}

	return len(found) == len(substrings)
}

func snake(value string, separator string) string {
	var snake strings.Builder

	for i, r := range value {
		if i != 0 && (unicode.IsUpper(r)) {
			snake.WriteString(separator)
		}

		snake.WriteRune(unicode.ToLower(r))
	}

	return strings.ReplaceAll(snake.String(), " ", separator)
}
