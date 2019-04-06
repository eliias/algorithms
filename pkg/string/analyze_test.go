package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountOfVectorizedSingleCharString(t *testing.T) {
	result := Vectorize([]rune("a"), CreateASCIITable(ASCIIAlphabetLowercase))
	assert.Equal(t, 1.0, result[0], "vectorized string has 1 at index 0")
}

func TestCountOfVectorizedMultiCharString(t *testing.T) {
	result := Vectorize([]rune("aaaaa"), CreateASCIITable(ASCIIAlphabetLowercase))
	assert.Equal(t, 5.0, result[0], "vectorized string has 1 at index 0")
}

func TestStringsSimilar(t *testing.T) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	a := Vectorize([]rune("aaaaa"), alphabet)

	b := Vectorize([]rune("aaaaa"), alphabet)
	c := Vectorize([]rune("aaaa"), alphabet)
	d := Vectorize([]rune("aaa"), alphabet)

	assert.Equal(t, 1.0, a.Cosine(b), "vectorized strings should be similar")
	assert.Equal(t, 1.0, a.Cosine(c), "vectorized strings should be similar")
	assert.Equal(t, 1.0, a.Cosine(d), "vectorized strings should be similar")
}

func TestStringsNotSimilar(t *testing.T) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	a := Vectorize([]rune("aaaaa"), alphabet)

	b := Vectorize([]rune("aaaab"), alphabet)
	c := Vectorize([]rune("aaac"), alphabet)
	d := Vectorize([]rune(""), alphabet)

	assert.NotEqual(t, 1.0, a.Cosine(b), "vectorized strings should not be similar")
	assert.NotEqual(t, 1.0, a.Cosine(c), "vectorized strings should not be similar")
	assert.NotEqual(t, 1.0, a.Cosine(d), "vectorized strings should not be similar")
}

func BenchmarkVectorizeAsciiAlphabetString(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	r := []rune(ASCIIAlphabetLowercase)

	for n := 0; n < b.N; n++ {
		Vectorize(r, alphabet)
	}
}

func BenchmarkVectorizeRandomAsciiAlphabetString(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	r := []rune("olduksdiiifrprjfruufnohlyeihspml")

	for n := 0; n < b.N; n++ {
		Vectorize(r, alphabet)
	}
}

func BenchmarkVectorizeRandomLongAsciiAlphabetString(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	r := []rune("wnwcgzioomyoymeqzzuwsxtwfmlcywxwgtyacyizofglaquvxmezqhempxxapbfogjtajnqcariylxmucmwhbopittbhrpzwjxskejmjdxcaollbuuzkhndmmdzqhtmaldwctinqrfpxnxudnqirymdadiqwnaungyxejlueyqgftudvfzrabkxaqafpcwkuxdrxirkxaxzjskfblfuetrsnptrnibapyyxhsuoqjnuooiwekotplwfwvpqenlhdhkdqokrykqbiboftxdsafczzptivodroibmibtylnvcbpdhkzimipwterlnefoscyeinfolsykolsdjboxwpkpdojjdckzpwitaezuzyaohqryfofrdnyixjmcxhnbpkqeomeixzxlxvoejxwpbsfdhvarsgwofxqbkfhybpxvympecachnjwlvuipmnhiwxpeamqhglijoognoxsfuputnovtsrfucmfsydgcdlosxvghlhbyqcgukpzofnhshgncjulaoxpbaebqlhojilnnpgtuydfmgltqvvumcxepufxbqztwdmbyjqdbyxasjoaagufquwuuynfymniyrpwxsxllvyjovehjurxjctcyfzjnfenbenlmrakfmhnklktulzbjegdsvtbgwpbnhttalbgbydpwfygdelexlvykalnrgapeqfuhqqvofqrtveyiccyejsncvylousildbmnjmuliccnunsmmexzojmuwoyznwjmkiatmhhofswmzeldrrgitqaebswosswypyrhfpmqsngdkdvtlzcwaxhqmagmkabsfxgcdbniwnrjbsccpekmdzmzhfzsygmgugjifxahinnkcpmyuykpjpinuempiqrwaejfzsomzzzohyprhhrughsarrtnqwmpznsassqhvczlwquomsfwbrqiamsmjxbwxoittrupkrizkoaanvmsluzzbuobrlmcppmgykrpiatizwjqgwczfwpxjbfudnmxxuvjpkeoxopass")

	for n := 0; n < b.N; n++ {
		Vectorize(r, alphabet)
	}
}

func BenchmarkVectorizeAsciiNumbersString(b *testing.B) {
	alphabet := CreateASCIITable(ASCIINumbers)
	r := []rune(ASCIINumbers)

	for n := 0; n < b.N; n++ {
		Vectorize(r, alphabet)
	}
}

func BenchmarkVectorizeAsciiPunctuationString(b *testing.B) {
	alphabet := CreateASCIITable(ASCIINumbers)
	r := []rune(ASCIIPunctuation)

	for n := 0; n < b.N; n++ {
		Vectorize(r, alphabet)
	}
}

func BenchmarkTestStringSimilarityAsciiAlphabet(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	v1 := Vectorize([]rune(ASCIIAlphabetLowercase), alphabet)
	v2 := Vectorize([]rune(ASCIIAlphabetLowercase), alphabet)

	for n := 0; n < b.N; n++ {
		v1.Cosine(v2)
	}
}

func BenchmarkTestRandomStringSimilarityAsciiAlphabet(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	v1 := Vectorize([]rune("olduksdiiifrprjfruufnohlyeihspml"), alphabet)
	v2 := Vectorize([]rune("ezjwdffjhdlgflprxytfvmgamxdpvikq"), alphabet)

	for n := 0; n < b.N; n++ {
		v1.Cosine(v2)
	}
}

func BenchmarkTestRandomLongStringSimilarityAsciiAlphabet(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIAlphabetLowercase)
	v1 := Vectorize([]rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf"), alphabet)
	v2 := Vectorize([]rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs"), alphabet)

	for n := 0; n < b.N; n++ {
		v1.Cosine(v2)
	}
}

func BenchmarkTestStringSimilarityAsciiNumbers(b *testing.B) {
	alphabet := CreateASCIITable(ASCIINumbers)
	v1 := Vectorize([]rune(ASCIINumbers), alphabet)
	v2 := Vectorize([]rune(ASCIINumbers), alphabet)

	for n := 0; n < b.N; n++ {
		v1.Cosine(v2)
	}
}

func BenchmarkTestStringSimilarityAsciiPunctuation(b *testing.B) {
	alphabet := CreateASCIITable(ASCIIPunctuation)
	v1 := Vectorize([]rune(ASCIIPunctuation), alphabet)
	v2 := Vectorize([]rune(ASCIIPunctuation), alphabet)

	for n := 0; n < b.N; n++ {
		v1.Cosine(v2)
	}
}

func ExampleVectorize() {
	alphabet := CreateASCIITable(ASCIIPunctuation)
	Vectorize([]rune(ASCIIPunctuation), alphabet)
}
