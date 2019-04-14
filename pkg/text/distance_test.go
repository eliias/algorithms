package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistanceDifferentShort(t *testing.T) {
	a := []rune("Hannes")
	b := []rune("Mooser")

	distance := Hamming(a, b)
	assert.Equal(t, 5, distance, "hamming distance is 5")
}

func TestLevenhsteinEditDistanceDifferentShort(t *testing.T) {
	a := []rune("Hannes")
	b := []rune("Moser")

	distance := LevenshteinTwoMatrixRows(a, b)
	assert.Equal(t, 5, distance, "edit distance is 5")
}

func TestLevenhsteinEditDistanceSameButOneShort(t *testing.T) {
	a := []rune("aaaa")
	b := []rune("aaab")

	distance := LevenshteinTwoMatrixRows(a, b)
	assert.Equal(t, 1, distance, "edit distance is 1")
}

func TestLevenhsteinMatrixEditDistanceSameShort(t *testing.T) {
	a := []rune("aaaa")
	b := []rune("aaaa")

	distance := LevenshteinTwoMatrixRows(a, b)
	assert.Equal(t, 0, distance, "edit distance is 0")
}

func TestLevenhsteinMatrixEditDistanceDifferentLong(t *testing.T) {
	a := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	b := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	distance := LevenshteinTwoMatrixRows(a, b)
	assert.Equal(t, 903, distance, "levenhstein edit distance is 903")
}

func TestLevenhsteinEditDistanceSameShort(t *testing.T) {
	a := []rune("aaaa")
	b := []rune("aaaa")

	distance := Levenshtein(a, b)
	assert.Equal(t, 0, distance, "edit distance is 0")
}

func TestLevenhsteinEditDistanceDifferentLong(t *testing.T) {
	a := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	b := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	distance := Levenshtein(a, b)
	assert.Equal(t, 903, distance, "levenhstein edit distance is 903")
}

func TestEnhancedUkkonenEditDistanceDifferentLong(t *testing.T) {
	a := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	b := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	distance := EnhancedUkkonen(a, b)
	assert.Equal(t, 903, distance, "levenhstein edit distance is 903")
}

func TestLevenhsteinIosifovichEditDistanceSameShort(t *testing.T) {
	t.Skip("Iosifovich needs some fighting")
	//a := []rune("aaaa")
	//b := []rune("aaaa")
	//
	//distance := LevenshteinIosifovich(a, b)
	//assert.Equal(t, 0, distance, "levenhstein edit distance is 1")
}

func BenchmarkTestLevenhsteinMatrixEditDistanceDifferentShort(b *testing.B) {
	s1 := []rune("Hannes")
	s2 := []rune("Moser")

	for n := 0; n < b.N; n++ {
		LevenshteinTwoMatrixRows(s1, s2)
	}
}

func BenchmarkTestLevenhsteinMatrixEditDistanceDifferentLong(b *testing.B) {
	s1 := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	s2 := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	for n := 0; n < b.N; n++ {
		LevenshteinTwoMatrixRows(s1, s2)
	}
}

func BenchmarkTestLevenhsteinEditDistanceDifferentShort(b *testing.B) {
	s1 := []rune("Hannes")
	s2 := []rune("Moser")

	for n := 0; n < b.N; n++ {
		Levenshtein(s1, s2)
	}
}

func BenchmarkTestLevenhsteinEditDistanceDifferentLong(b *testing.B) {
	s1 := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	s2 := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	for n := 0; n < b.N; n++ {
		Levenshtein(s1, s2)
	}
}

func BenchmarkTestEnhancedUkkonenEditDistanceDifferentShort(b *testing.B) {
	s1 := []rune("Hannes")
	s2 := []rune("Moser")

	for n := 0; n < b.N; n++ {
		EnhancedUkkonen(s1, s2)
	}
}

func BenchmarkTestEnhancedUkkonenEditDistanceDifferentLong(b *testing.B) {
	s1 := []rune("kvctspmqeapctcyklebxrenjdfhzqzejzopgwwoahitfxpbfhmegxbpobhtcqeubzfwpyuoyewxvnvlkieegvozqgfvbfebatagyoamoepgqlylhmeabpkbyahoydnsaktdyiuskdtinkqsfmwwpqzyabnkvvsizakhyorhzytilzfhbygxhetlxwkdtrwlahjuquwbupwtbbwnpfnqsyscsrlqaxedgtxxwqfbbapcienpzpuyqaesbxbbyhbxctwsbxpaufnuomqlykwonaasdbcrlfvemxqsniysxqzauyxgznhkpqocklftmdmmpescxfyppcchohatoatltjvqzrcivojkfuwghhrdsdaamjpwuaaowyyojpykrmokgzxuavwtiqhkgcblkiozlkazualjzckyohnwjiohrbsctzxhnjvxswcccogeaughhyevarhiptxqeqcmlmfizjalgxuiclvhbmjbtlniamglrvcosrxcedkbkzchqoygcoqlcazuxkwhnmovwuhzkbwhxzjmzgmfbaqegxgfciregezjqgrnafrfgaxfpecmqedgzzamktxpbwpyoxqathlkpkrzpaihaxphshxkaygrbpnggwqfdmyeksfkxlerxbsumhdfwdjqnbskwkazrszsbdfnkuaowfkeruexbjvkpljbbsgjhwdjykyohiaxkdrnwyfihvbdgfzsglduddtqjqcmmimkwytzmdjbusztefjpqvesqnztsjrsqkjijzaunblmliwpsmliapttofnwidraresmzxvpjrzxgmileecxlsbqnaqwaxyekelfkyaoawifegjbhhzoguwanlowiqdwmjosaxjmfimhgphrgrhpjszjaczrumqpxqfzgwvthqvxegtgygtawpolmdxvkxqlkeemnptiyfepfkjiajlexoyismdpmasaonkmrnemvsbcvbnzjdwynzkledawhmhnwofvrejsafnyneggqkewbtmxexjdrezxctxwf")
	s2 := []rune("xofwplnpdugcibsscuyyrimbjhlqeuqzouubfysdzbsznxihvvabylyzspvfevqlqtsiyzlsnnovsjdzcvwrugqaerzogtfgelgfsiefcpmtfspkoqpcnruexiclltntwmbxqzpyjqwqmhwrljxgydmwygwgamzjsmrqivpgjwwdqzfdpctqslumbelknnyssoeotmtycnnnplaxxwzabanzqebpsqcdawvimdgtikmtjjgmycfrsubcalihqwgayrhvvmlqjockcpngaxdmagjdirtqcbkgrawzctdpypdgonitgexbueldqyctvhdtktulwfqnesrazvyepqgzcarcvzmzfvslenomerymqrthusjpcptolanwuajnonsjvjacajbxkprgfxpjqybpeopfqdwkpwfdbqohauwjuxqmfczowmkdljxqxmgyvtegmeuvpxfxkmrzblffigciutfjspjsteoupyzzxrjeoyiofsrgfmuinconfogiedjlkepayutmcjndmxlikuwjpvjnnxaaqwmkfqjdrfletkfvrnjqrncbzhobcnkzmvhimzbynxbjbrkesmftugmsxwostvjuwfvvcnrdwedqcxtiftnukzyixjgixglgpihgbemylatzlwlvszxbrnztqgizofxtyrpxfzzdpbwojurrpgivduslkhtbqqcluzgzoecysfnizgazparzioqlpgrsnmrhamtgusvipjewfklcjckwxypkyjqkjczqozwfrxfoyxcbndcvvqphcbsrqaeinkfyywjmkglpusxibavlaydhchrtkdacjwspbkbokilpgvhjgnvknstmfhtkehegduxogtjkcxmkzxrtgrwvqzlwohjjtkepknhoidrnzuqctvuuwpkoybgrdjtgbggeeqivgbquyphwvhuvehnfdixdsjoikvbvgaootmndhaqnekkvwrmymuiwslkdjtzknbdgmhawaqvjrenwzckvrbfhnrbopfdruduphegs")

	for n := 0; n < b.N; n++ {
		EnhancedUkkonen(s1, s2)
	}
}