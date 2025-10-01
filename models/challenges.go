package models

import (
	"math/rand"
	"strings"
	"time"
)

type Challenge struct {
	Question string
	Answer   string
	Type     string
	Hint     string
}

var Challenges = []Challenge{
	// riddles
	{"من هو الحيوان الذي لا يلد ولا يبيض؟", "ذكر", "Riddle", "فكر في الحيوانات الأسطورية."},
	{"ما هو الشيء الذي يجب كسره قبل استخدامه؟", "البيضة", "Riddle", "يُستخدم في الطهي."},
	{"ما هو الشيء الذي يمكن كسره دون أن نلمسه؟", "الوعد", "Riddle", "يتعلق بالثقة."},
	{"ما هو الشيء الذي يصدق دائماً ويكذب فقط حين يجوع؟", "الساعة", "Riddle", "يُستخدم لمعرفة الوقت."},
	{"ما هو الشيء الذي يتحرك بدون أرجل ويبكي بدون عيون؟", "السحاب", "Riddle", "يتعلق بالطقس."},
	{"ما هو الشيء الذي ليس بإنسان ولا حيوان ولا جماد؟", "نبات", "Riddle", "يُستخدم في الزراعة."},
	{"ما هو الشيء الذي لا يخلو منه أي بيت؟", "حرف الياء", "Riddle", "يُستخدم في الكتابة."},
	{"ما هو الشيء الذي يؤخذ منك قبل أن يعطي لك؟", "الصورة الفوتوغرافية", "Riddle", "يُستخدم في التصوير."},
	{"ما هو الشيء الذي درجة حرارته ثابتة سواء وضعته في الثلاجة أو على النار؟", "الفلفل", "Riddle", "يُستخدم في الطهي."},
	{"ما هو الشيء الذي يمشي ويقف وليس له أرجل؟", "الساعة", "Riddle", "يُستخدم لمعرفة الوقت."},
	{"تأتي بالنعومة وتخبر بدون صوت، تُشعر ولكن لا تُرى، ما هي؟", "لمسة", "Riddle", "يدان أو شفاه تقوم بها."},
	// 5
	{"رائحة صغيرة من زجاجة تُعيد الذكريات وتُغري الحواس، ما هي؟", "عطر", "Riddle", "زهري، فانيليا أو خشبي."},
	// 6
	{"يضيء الجو بلطف، يجعل اللحظة حميمية لكنه يذوب إن قربته كثيرًا، ما هو؟", "شمعة", "Riddle", "ضوء خافت ورومانسية."},
	{"🔮 I am taken from the mine, locked in a wooden case, never released but used by every scribe. What am I?", "pencil lead", "Riddle", "Think of writing tools."},
	{"🧪 The more you take away from me, the bigger I get. What am I?", "hole", "Riddle", "Look into the void."},
	{"🕯️ I have a heart that never beats, I grow without life, I am dead yet consume life. What am I?", "candle", "Riddle", "Melts as it burns."},
	//Mixed
	{"Unscramble: 'ELPPA'", "apple", "Decode", "It’s a fruit."},
	{"marra hedhy ma fama chay , dima nhebbek w bch nab9a nhebbek ale toul", "done", "Memory", "I love you till death ( ekteb done )"},
	{"Write your favorite emoji 3 times in chat.", "done", "Dare", "Fun and easy!"},
	{"شيء له قلب ولا ينبض، ما هو؟", "الخس", "Riddle", "خضر ويؤكل."},
	//Memory
	// ---------------- Tounsi ----------------
	{"Chnowa kenet awel 7aja 3tithelek?", "parfume", "Memory", "Tawa tadhker awel cadeau."},
	{"Chnowa ahsen film tfarejt fih maaye?", "Joker", "Memory", "film aawedneh 4 marrat fi ciné."},
	{"Wa9teh kenet l'moment li badlet vision lel relation?", "manouba", "Memory", "blasa fergha w 9riba"},
	{"Akther moment tdahhek w nafs l'wa9t ter3eb fi hyetna", "khzena", "Memory", "blasa nhoto feha dbach"},
	{"blasa awel mara we had fun w hasina beb3adhna bel behy", "bhar", "Memory", "blasa mchinelha w ahna te3bin"},
	{"في bardo شنوة أول حاجة درناها كيف دخلنا؟", "hug", "Memory", "Warm moment."},
	//Spy
	// 🧙 Witch Spy
	{"🔮 A witch cast a spell: 'EVOL'. What is the real word hidden in reverse?", "love", "Spy", "Spell it backwards."},
	{"🕯️ The grimoire asks: Which element calms your soul most – fire, water, earth, or air?", "water", "Spy", "Think of your vibe."},
	{"🌙 If we met in a secret ritual under the moon, which symbol would you draw to bind us?", "heart", "Spy", "It represents love."},

	// 🎤 K-pop Spy
	{"🎤 You’re undercover in a K-pop group. Which role would you take – leader, main vocal, dancer, or visual?", "dancer", "Spy", "Pick your vibe."},
	{"💎 If I were your idol 'bias', what stage name would you give me?", "prince", "Spy", "Romantic nickname works here."},
	{"🎶 Complete the lyric code: 'Cause I-I-I’m in the stars tonight…'", "so watch me bring the fire and set the night alight", "Spy", "It’s BTS – Dynamite."},

	// 🕵️ Crime / Noir Spy
	{"🕵️ A priceless jewel was stolen. The police found only one clue: lipstick on the glass. Who is the thief?", "me", "Spy", "It’s always the closest one."},
	{"🔎 In the case file, it says: 'Suspect is guilty of…' — complete the sentence.", "loving you", "Spy", "Romantic confession."},
	{"💣 Agent, the bomb password is a 3-word phrase that melts hearts. What is it?", "i love you", "Spy", "Classic key phrase."},

	// 🧪 Dexter-Style Spy
	{"🩸 Blood tells the truth. In Dexter’s code, what color always betrays the killer?", "red", "Spy", "Obvious but deadly."},
	{"🔪 Dexter keeps trophies of his work. What tiny object does he collect?", "blood slide", "Spy", "Glass with a drop on it."},
	{"🧠 Double life test: by day he’s a forensic scientist, by night…?", "serial killer", "Spy", "His secret identity."},
	{"📂 If I were a case file, what single word would be written in red ink?", "love", "Spy", "Dark but romantic twist."},
	{"🧬 Dexter follows 'The Code' given by his father. What’s the first rule?", "never get caught", "Spy", "Basic survival rule."},
	{"Which color makes me blush the most?", "red", "Spy", "Hint: think romantic."},
	{"What is my favorite sweet treat?", "chocolate", "Spy", "Yummy and tempting!"},
	{"Who do I secretly admire the most?", "you", "Spy", "Hint: maybe it’s you 😉"},
	// Caesar cipher (2 only)
	{"Decode: 'Khoor'", "hello", "Decode", "Caesar shift +3."},
	{"Decode: 'Wklv lv d whvw'", "this is a test", "Decode", "Caesar shift +3."},
	{"🎤 Reverse lyric: 'thgil a ni thgin eht'", "the night in a light", "Decode", "Lyrics backwards."},
	{"🔮 Secret: 'EVIL+LIVE'", "evil", "Decode", "Two sides of the same word."},
	{"🩸 Forensic clue: 'EDROB'", "blood", "Decode", "Think like Dexter."},
	{"🕵️ Crime word: 'GNIHCTAW'", "watching", "Decode", "Someone’s spying."},
	{"🕵️ Secret name scrambled: 'REDTXE'", "dexter", "Decode", "Dark passenger."},

	// Reversed words
	{"Decode: 'dlrow olleh'", "hello world", "Decode", "Read it backwards."},
	{"Decode: 'atad'", "data", "Decode", "Backwards spelling."},

	// Scrambled letters (English)
	{"Unscramble: 'ELPPA'", "apple", "Decode", "It’s a fruit."},
	{"Unscramble: 'ANANAB'", "banana", "Decode", "Yellow and sweet."},
	{"Unscramble: 'ORENG'", "green", "Decode", "A color."},
	{"Unscramble: 'OGL'", "log", "Decode", "Something from a tree."},

	// Simple Arabic challenges
	{"Decode: 'ىلع فرغ'", "غرف علي", "Decode", "Read the letters from right to left."},
	{"Unscramble: 'مكتبة'", "مكتبة", "Decode", "A place with books."},

	// More playful
	{"Unscramble: 'LVOE'", "love", "Decode", "Hint: romantic word."},
	{"Unscramble: 'NOSMIE'", "moines", "Decode", "A tricky mix."},
	{"Decode: 'sdrawkcab gninrael'", "learning backwards", "Decode", "Read backwards!"},
	{"Unscramble: 'RAEHB'", "behar", "Decode", "Arabic word, sea."},
	//Dares
	{"Send a selfie making a silly flirty face.", "done", "Dare", "No hint"},
	{"Send a message to mom t9olha i love you mom", "done", "Dare", "No hint"},
	{"Whisper a naughty word in a voice note.", "done", "Dare", "No hint"},
	{"Ab3eth msg tefah l' ay had men maarfek and screenshot it.", "done", "Dare", "No hint"},
	{"Send an emoji that shows your mischievous side.", "done", "Dare", "No hint"},
	{"Pretend to kiss your hand and send a selfie.", "done", "Dare", "No hint"},
	{"Send a message with your funniest pickup line.", "done", "Dare", "No hint"},
	{"Record a voice note saying 'I’m guilty of flirting!'", "done", "Dare", "No hint"},
	{"Send a selfie wearing the most daring outfit you have.", "done", "Dare", "No hint"},
	{"Text a flirty question and see the response.", "done", "Dare", "No hint"},
	{"Write a one-line naughty joke.", "done", "Dare", "No hint"},
	{"Send a selfie with your tongue out (funny/flirty).", "done", "Dare", "No hint"},
	{"Pretend to be seductive in a 5-second video.", "done", "Dare", "No hint"},
	{"Send an emoji combination that represents flirting.", "done", "Dare", "No hint"},
	{"Send the beautiful jucy Ass :<)", "done", "Dare", "No hint"},
	{"Send a screenshot of a fun, flirty meme.", "done", "Dare", "Keep it playful."},
	{"Record yourself saying a cheesy pick-up line.", "done", "Dare", "No hint"},
	{"Send a photo of your hand in a 'kiss' pose.", "done", "Dare", "No hint"},
	{"Send a text confessing the silliest secret you’ve kept.", "done", "Dare", "No hint"},
	// Quiz (medium-hard, English/Arabic mix)
	{"Which element has the chemical symbol 'Fe'?", "iron", "Quiz", "It’s used to make steel."},
	{"Who painted the Mona Lisa?", "leonardo da vinci", "Quiz", "Famous Italian Renaissance artist."},
	{"في أي عام بدأ الحرب العالمية الثانية؟", "1939", "Quiz", "It started قبل الهولوكوست."},
	{"What is the square root of 225?", "15", "Quiz", "It’s a perfect square."},
	{"Which planet has the most moons?", "saturn", "Quiz", "It has the famous rings."},
	{"ما هو أطول نهر في العالم؟", "nile", "Quiz", "يمر عبر إفريقيا الشمالية."},
	{"What is the hardest natural substance on Earth?", "diamond", "Quiz", "Used in jewelry and drills."},
	{"Who developed the theory of general relativity?", "einstein", "Quiz", "Famous physicist with E=mc^2."},
	{"كم عدد الدول العربية في قارة إفريقيا؟", "10", "Quiz", "Think of North African countries."},
	{"Which language has the most native speakers worldwide?", "mandarin", "Quiz", "Not English or Spanish."},
	{"Who wrote 'One Hundred Years of Solitude'?", "gabriel garcia marquez", "Quiz", "Famous Colombian author."},
	{"ما هي عاصمة اليابان؟", "tokyo", "Quiz", "أكبر مدينة في اليابان."},
	{"What is the smallest prime number greater than 50?", "53", "Quiz", "Prime numbers only."},
	{"Which chemical element is a noble gas and used in neon signs?", "neon", "Quiz", "It glows brightly."},
	{"في أي سنة نزل أول إنسان على القمر؟", "1969", "Quiz", "Neil Armstrong was the first."},
	{"Which organ in the human body produces insulin?", "pancreas", "Quiz", "Regulates blood sugar."},
	{"ما هو الحيوان الذي يمكنه البقاء دون ماء لأطول فترة؟", "camel", "Quiz", "يعيش في الصحراء."},
	{"What is the value of π (Pi) up to 3 decimal places?", "3.142", "Quiz", "Used in circle calculations."},
	{"Which country won the first FIFA World Cup?", "uruguay", "Quiz", "Year: 1930."},
	{"من هو مؤلف كتاب 'الأمير'؟", "machiavelli", "Quiz", "فيلسوف إيطالي."},
}

func FilterByType(challenges []Challenge, answered map[string]bool, mode string) []Challenge {
	mode = strings.ToLower(mode)
	var result []Challenge
	for _, c := range challenges {
		if strings.ToLower(c.Type) == mode {
			if answered == nil || !answered[c.Question] {
				result = append(result, c)
			}
		}
	}
	return result
}

func RandomChallenge(challenges []Challenge, answered map[string]bool) Challenge {
	var pool []Challenge
	for _, c := range challenges {
		if answered == nil || !answered[c.Question] {
			pool = append(pool, c)
		}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return pool[r.Intn(len(pool))]
}

func MatchAnswer(given, expected string) bool {
	if expected == "" {
		return true
	}
	return strings.TrimSpace(strings.ToLower(given)) == strings.TrimSpace(strings.ToLower(expected))
}
