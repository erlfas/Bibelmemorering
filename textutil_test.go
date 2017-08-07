package main

import (
	"html/template"
	"testing"
)

func TestMakeFourthTransformation(t *testing.T) {
	sal23 := "1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile. 3 Han gir meg nytt liv. Han fører meg på rettferdighets stier for sitt navns skyld."
	trans := makeFourthTransformation(sal23)
	if trans != "1 E [...] . H [...] . 2 H [...] . 3 H [...] . H [...] ." {
		t.Error("Wrong output. Got:", trans)
	}
}

func TestIsBeginning3(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isb := [30]bool{true, true, false, false, false, true, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	for i := range words {
		is := isBeginning(words[i], i, words)
		if isb[i] != is {
			t.Error("Expected: ", isb[i], "but was: ", is, "for index", i)
		}
	}
}

func TestIsBeginning2(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile. 3 Han gir meg nytt liv. Han fører meg på rettferdighets stier for sitt navns skyld.")
	isb := isBeginning(words[1], 1, words) // En
	if isb != true {
		t.Error("Expected true")
	}
}

func TestIsBeginning(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile. 3 Han gir meg nytt liv. Han fører meg på rettferdighets stier for sitt navns skyld.")
	isb := isBeginning(words[0], 0, words) // 1
	if isb != true {
		t.Error("Expected true")
	}
}

func TestIsWord4(t *testing.T) {
	text := "David."
	isWord, ending := isWord(text)
	if isWord != true {
		t.Error("Expected isWord = true")
	}
	if ending != "." {
		t.Error("Expected ending = '.'")
	}
}

func TestMakeThirdTransformation(t *testing.T) {
	sal23 := "1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile. 3 Han gir meg nytt liv. Han fører meg på rettferdighets stier for sitt navns skyld."
	trans := makeThirdTransformation(sal23)

	if trans != "1 E [...] D. H [...] h, j [...] n. 2 H [...] e, h [...] h. 3 H [...] l. H [...] s." {
		t.Error("Got:", trans)
	}
}
func TestMakeSecondTransformation(t *testing.T) {
	sal23 := "1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile. 3 Han gir meg nytt liv. Han fører meg på rettferdighets stier for sitt navns skyld."
	trans := makeSecondTransformation(sal23)
	if trans != "1 E _ _ D. H _ _ h, j _ _ n. 2 H _ _ _ _ _ e, h _ _ _ _ _ _ _ h. 3 H _ _ _ l. H _ _ _ _ _ _ _ _ s." {
		t.Error("Got:", trans)
	}
}

func TestIsAtBoundary14(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[21], 21, words) // Han
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[21])
	}
}

func TestIsAtBoundary13(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[20], 20, words) // Han
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[20])
	}
}

func TestIsAtBoundary12(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[14], 14, words) // Han
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[14])
	}
}

func TestIsAtBoundary11(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[13], 13, words) // 2
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[13])
	}
}

func TestIsAtBoundary10(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[12], 12, words) // noe.
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[12])
	}
}

func TestIsAtBoundary9(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[10], 10, words) // mangler
	if isAtBoundary != false {
		t.Error("Expected isAtBoundary = false for", words[10])
	}
}

func TestIsAtBoundary8(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[9], 9, words) // jeg
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[9])
	}
}

func TestIsAtBoundary7(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[8], 8, words) // hyrde,
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[8])
	}
}

func TestIsAtBoundary6(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[5], 5, words) // Herren
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[5])
	}
}

func TestIsAtBoundary5(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[4], 4, words)
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true for", words[4])
	}
}

func TestIsAtBoundary4(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[3], 3, words)
	if isAtBoundary != false {
		t.Error("Expected isAtBoundary = false")
	}
}

func TestIsAtBoundary3(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[2], 2, words)
	if isAtBoundary != false {
		t.Error("Expected isAtBoundary = false")
	}
}

func TestIsAtBoundary2(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[1], 1, words)
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true")
	}
}

func TestIsAtBoundary(t *testing.T) {
	words := getTokens("1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile.")
	isAtBoundary := isAtBoundary(words[0], 0, words)
	if isAtBoundary != true {
		t.Error("Expected isAtBoundary = true")
	}
}

func TestIsWord2(t *testing.T) {
	isWord, ending := isWord("David.")
	if isWord != true {
		t.Error("Expected isWord = true")
	}
	if ending != "." {
		t.Error("Expected endsWithPunct = true")
	}
}

func TestIsWord(t *testing.T) {
	isWord, ending := isWord("salme")
	if isWord != true {
		t.Error("Expected isWord = true")
	}
	if ending != "" {
		t.Error("Expected endsWithPunct = false")
	}
}

func TestGetTokens(t *testing.T) {
	text := `Herren er min hyrde,
          jeg mangler ikke noe.`
	tokens := getTokens(text)
	expectedTokens := []string{"Herren", "er", "min", "hyrde,\n", "jeg", "mangler", "ikke", "noe."}
	for i := range tokens {
		if tokens[i] != expectedTokens[i] {
			t.Error(tokens[i], " != ", expectedTokens[i], "for index", i)
		}
	}
}

func TestEndsWithPunct3(t *testing.T) {
	text := "hyrde"
	ends := endsWithPunct(text)
	if ends != false {
		t.Error("Expected false")
	}
}

func TestEndsWithPunct2(t *testing.T) {
	text := "hyrde,"
	ends := endsWithPunct(text)
	if ends != true {
		t.Error("Expected true")
	}
}

func TestEndsWithPunct(t *testing.T) {
	text := "hyrde,\n"
	ends := endsWithPunct(text)
	if ends != true {
		t.Error("Expected true")
	}
}

func TestIsWord3(t *testing.T) {
	w := "hyrde,\n"
	isWord, ending := isWord(w)
	if isWord != true {
		t.Error("Expected isWord = true")
	}

	if ending != ",\n" {
		t.Error("Expected endsWithPunct = true")
	}
}

func TestMakeFirstTransformation3(t *testing.T) {
	text := `Herren er min hyrde,
          jeg mangler ikke noe.`
	trans := makeFirstTransformation(text)
	if trans != `H e m h,
 j m i n.` {
		t.Error("Wrong output")
	}
}

func TestMakeFirstTransformation2(t *testing.T) {
	sal23 := "David."
	trans := makeFirstTransformation(sal23)
	if trans != "D." {
		t.Error("Expected 'D.' but was:", trans)
	}
}

func TestMakeFirstTransformation(t *testing.T) {
	sal23 := "1 En salme av David. Herren er min hyrde, jeg mangler ikke noe. 2 Han lar meg ligge i grønne enger, han leder meg til vann der jeg finner hvile."
	trans := makeFirstTransformation(sal23)
	if trans != "1 E s a D. H e m h, j m i n. 2 H l m l i g e, h l m t v d j f h." {
		t.Error("Wrong output. Expected:", "1 E s a D. H e m h, j m i n. 2 H l m l i g e, h l m t v d j f h.")
	}
}

func TestIsWord5(t *testing.T) {
	text := "på"

	isWord, ending := isWord(text)
	if isWord != true {
		t.Error("Expected true")
	}
	if ending != "" {
		t.Error("Expected '' but got", ending)
	}
}

func TestMakeFirstTransformation4(t *testing.T) {
	sal23 := `1 En salme av David.
        
           Herren er min hyrde,
          jeg mangler ikke noe.
          
     2 Han lar meg ligge i grønne enger,
          han leder meg til vann der jeg finner hvile.
          
     3 Han gir meg nytt liv.
          Han fører meg på rettferdighets stier
          for sitt navns skyld.`

	makeFirstTransformation(sal23)
}

func TestHtmlFilter(t *testing.T) {
	txt := `1 En salme av David.
        
           Herren er min hyrde,
          jeg mangler ikke noe.
          
     2 Han lar meg ligge i grønne enger,
          han leder meg til vann der jeg finner hvile.
          
     3 Han gir meg nytt liv.
          Han fører meg på rettferdighets stier
          for sitt navns skyld.`

	output := template.HTML(htmlFilter(makeFirstTransformation(txt)))

	if output != "1 E s a D.<br/> <br/> H e m h,<br/> j m i n.<br/> <br/> 2 H l m l i g e,<br/> h l m t v d j f h.<br/> <br/> 3 H g m n l.<br/> H f m p r s<br/> f s n s." {
		t.Error("Wrong output. Got: ", output)
	}
}
