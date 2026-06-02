package regexpx

// Unicode普通类
const (
	U_C  = `[\p{C}]`  // U_C matches Unicode category Other. Unicode 其他（Other）
	U_Cc = `[\p{Cc}]` // U_Cc matches control characters. 控制字符（control）
	U_Cf = `[\p{Cf}]` // U_Cf matches format characters. 格式字符（format）
	U_Co = `[\p{Co}]` // U_Co matches private-use characters. 私人使用区（private use）
	U_Cs = `[\p{Cs}]` // U_Cs matches surrogate characters. 代理区（surrogate）

	U_L  = `[\p{L}]`  // U_L matches letters. Unicode 字母（letter）
	U_Ll = `[\p{Ll}]` // U_Ll matches lowercase letters. 小写字母（lowercase letter）
	U_Lm = `[\p{Lm}]` // U_Lm matches modifier letters. 修饰字母（modifier letter）
	U_Lo = `[\p{Lo}]` // U_Lo matches other letters. 其它字母（other letter）
	U_Lt = `[\p{Lt}]` // U_Lt matches titlecase letters. 首字母大写（titlecase letter）
	U_Lu = `[\p{Lu}]` // U_Lu matches uppercase letters. 大写字母（uppercase letter）

	U_M  = `[\p{M}]`  // U_M matches marks. Unicode 标记（mark）
	U_Mc = `[\p{Mc}]` // U_Mc matches spacing combining marks. 间距标记（spacing mark）
	U_Me = `[\p{Me}]` // U_Me matches enclosing marks. 封闭标记（enclosing mark）
	U_Mn = `[\p{Mn}]` // U_Mn matches non-spacing marks. 非间距标记（non-spacing mark）

	U_N  = `[\p{N}]`  // U_N matches numbers. Unicode 数字（number）
	U_Nd = `[\p{Nd}]` // U_Nd matches decimal digits. 十进制数字（decimal number）
	U_Nl = `[\p{Nl}]` // U_Nl matches letter numbers. 字母数字（letter number）
	U_No = `[\p{No}]` // U_No matches other numbers. 其它数字（other number）

	U_P  = `[\p{P}]`  // U_P matches punctuation. Unicode 标点（punctuation）
	U_Pc = `[\p{Pc}]` // U_Pc matches connector punctuation. 连接符标点（connector punctuation）
	U_Pd = `[\p{Pd}]` // U_Pd matches dash punctuation. 破折号标点（dash punctuation）
	U_Pe = `[\p{Pe}]` // U_Pe matches close punctuation. 闭合标点（close punctuation）
	U_Pf = `[\p{Pf}]` // U_Pf matches final punctuation. 结束标点（final punctuation）
	U_Pi = `[\p{Pi}]` // U_Pi matches initial punctuation. 起始标点（initial punctuation）
	U_Po = `[\p{Po}]` // U_Po matches other punctuation. 其它标点（other punctuation）
	U_Ps = `[\p{Ps}]` // U_Ps matches open punctuation. 开放标点（open punctuation）

	U_S  = `[\p{S}]`  // U_S matches symbols. Unicode 符号（symbol）
	U_Sc = `[\p{Sc}]` // U_Sc matches currency symbols. 货币符号（currency symbol）
	U_Sk = `[\p{Sk}]` // U_Sk matches modifier symbols. 修饰符号（modifier symbol）
	U_Sm = `[\p{Sm}]` // U_Sm matches math symbols. 数学符号（math symbol）
	U_So = `[\p{So}]` // U_So matches other symbols. 其它符号（other symbol）

	U_Z  = `[\p{Z}]`  // U_Z matches separators. Unicode 分隔符（separator）
	U_Zl = `[\p{Zl}]` // U_Zl matches line separators. 行分隔符（line separator）
	U_Zp = `[\p{Zp}]` // U_Zp matches paragraph separators. 段落分隔符（paragraph separator）
	U_Zs = `[\p{Zs}]` // U_Zs matches space separators. 空白分隔符（space separator）

)

// Unicode类
const (
	UnicodeArabic             = `[\p{Arabic}]`              // UnicodeArabic matches Arabic script. 阿拉伯文
	UnicodeArmenian           = `[\p{Armenian}]`            // UnicodeArmenian matches Armenian script. 亚美尼亚文
	UnicodeBalinese           = `[\p{Balinese}]`            // UnicodeBalinese matches Balinese script. 巴厘岛文
	UnicodeBengali            = `[\p{Bengali}]`             // UnicodeBengali matches Bengali script. 孟加拉文
	UnicodeBopomofo           = `[\p{Bopomofo}]`            // UnicodeBopomofo matches Bopomofo script. 汉语拼音字母
	UnicodeBraille            = `[\p{Braille}]`             // UnicodeBraille matches Braille script. 盲文
	UnicodeBuginese           = `[\p{Buginese}]`            // UnicodeBuginese matches Buginese script. 布吉文
	UnicodeBuhid              = `[\p{Buhid}]`               // UnicodeBuhid matches Buhid script. 布希德文
	UnicodeCanadianAboriginal = `[\p{Canadian_Aboriginal}]` // UnicodeCanadianAboriginal matches Canadian Aboriginal script. 加拿大土著文
	UnicodeCarian             = `[\p{Carian}]`              // UnicodeCarian matches Carian script. 卡里亚文
	UnicodeCham               = `[\p{Cham}]`                // UnicodeCham matches Cham script. 占族文
	UnicodeCherokee           = `[\p{Cherokee}]`            // UnicodeCherokee matches Cherokee script. 切诺基文
	UnicodeCommon             = `[\p{Common}]`              // UnicodeCommon matches script-common characters. 通用字符（非特定书写系统）
	UnicodeCoptic             = `[\p{Coptic}]`              // UnicodeCoptic matches Coptic script. 科普特文
	UnicodeCuneiform          = `[\p{Cuneiform}]`           // UnicodeCuneiform matches Cuneiform script. 楔形文字
	UnicodeCypriot            = `[\p{Cypriot}]`             // UnicodeCypriot matches Cypriot script. 塞浦路斯文
	UnicodeCyrillic           = `[\p{Cyrillic}]`            // UnicodeCyrillic matches Cyrillic script. 西里尔文
	UnicodeDeseret            = `[\p{Deseret}]`             // UnicodeDeseret matches Deseret script. 犹他州文
	UnicodeDevanagari         = `[\p{Devanagari}]`          // UnicodeDevanagari matches Devanagari script. 天城文（梵文等）
	UnicodeEthiopic           = `[\p{Ethiopic}]`            // UnicodeEthiopic matches Ethiopic script. 埃塞俄比亚文
	UnicodeGeorgian           = `[\p{Georgian}]`            // UnicodeGeorgian matches Georgian script. 格鲁吉亚文
	UnicodeGlagolitic         = `[\p{Glagolitic}]`          // UnicodeGlagolitic matches Glagolitic script. 格拉哥里文
	UnicodeGothic             = `[\p{Gothic}]`              // UnicodeGothic matches Gothic script. 哥特文
	UnicodeGreek              = `[\p{Greek}]`               // UnicodeGreek matches Greek script. 希腊文
	UnicodeGujarati           = `[\p{Gujarati}]`            // UnicodeGujarati matches Gujarati script. 古吉拉特文
	UnicodeGurmukhi           = `[\p{Gurmukhi}]`            // UnicodeGurmukhi matches Gurmukhi script. 果鲁穆奇文
	UnicodeHan                = `[\p{Han}]`                 // UnicodeHan matches Han script. 汉文
	UnicodeHangul             = `[\p{Hangul}]`              // UnicodeHangul matches Hangul script. 韩文
	UnicodeHanunoo            = `[\p{Hanunoo}]`             // UnicodeHanunoo matches Hanunoo script. 哈努诺文
	UnicodeHebrew             = `[\p{Hebrew}]`              // UnicodeHebrew matches Hebrew script. 希伯来文
	UnicodeHiragana           = `[\p{Hiragana}]`            // UnicodeHiragana matches Hiragana script. 平假名
	UnicodeInherited          = `[\p{Inherited}]`           // UnicodeInherited matches inherited script characters. 继承前一字符的书写系统
	UnicodeKannada            = `[\p{Kannada}]`             // UnicodeKannada matches Kannada script. 卡纳达文
	UnicodeKatakana           = `[\p{Katakana}]`            // UnicodeKatakana matches Katakana script. 片假名
	UnicodeKayahLi            = `[\p{Kayah_Li}]`            // UnicodeKayahLi matches Kayah Li script. 克耶字母
	UnicodeKharoshthi         = `[\p{Kharoshthi}]`          // UnicodeKharoshthi matches Kharoshthi script. 卡罗须提文
	UnicodeKhmer              = `[\p{Khmer}]`               // UnicodeKhmer matches Khmer script. 高棉文
	UnicodeLao                = `[\p{Lao}]`                 // UnicodeLao matches Lao script. 老挝文
	UnicodeLatin              = `[\p{Latin}]`               // UnicodeLatin matches Latin script. 拉丁文
	UnicodeLepcha             = `[\p{Lepcha}]`              // UnicodeLepcha matches Lepcha script. 雷布查文
	UnicodeLimbu              = `[\p{Limbu}]`               // UnicodeLimbu matches Limbu script. 林布文
	UnicodeLinearB            = `[\p{Linear_B}]`            // UnicodeLinearB matches Linear B script. 线形文字 B
	UnicodeLycian             = `[\p{Lycian}]`              // UnicodeLycian matches Lycian script. 利西亚文
	UnicodeLydian             = `[\p{Lydian}]`              // UnicodeLydian matches Lydian script. 吕底亚文
	UnicodeMalayalam          = `[\p{Malayalam}]`           // UnicodeMalayalam matches Malayalam script. 马拉雅拉姆文
	UnicodeMongolian          = `[\p{Mongolian}]`           // UnicodeMongolian matches Mongolian script. 蒙古文
	UnicodeMyanmar            = `[\p{Myanmar}]`             // UnicodeMyanmar matches Myanmar script. 缅甸文
	UnicodeNewTaiLue          = `[\p{New_Tai_Lue}]`         // UnicodeNewTaiLue matches New Tai Lue script. 新傣仂文
	UnicodeNko                = `[\p{Nko}]`                 // UnicodeNko matches Nko script. Nko 文
	UnicodeOgham              = `[\p{Ogham}]`               // UnicodeOgham matches Ogham script. 欧甘文
	UnicodeOlChiki            = `[\p{Ol_Chiki}]`            // UnicodeOlChiki matches Ol Chiki script. 桑塔利文
	UnicodeOldItalic          = `[\p{Old_Italic}]`          // UnicodeOldItalic matches Old Italic script. 古意大利文
	UnicodeOldPersian         = `[\p{Old_Persian}]`         // UnicodeOldPersian matches Old Persian script. 古波斯文
	UnicodeOriya              = `[\p{Oriya}]`               // UnicodeOriya matches Oriya script. 奥里亚文
	UnicodeOsmanya            = `[\p{Osmanya}]`             // UnicodeOsmanya matches Osmanya script. 奥斯曼亚文
	UnicodePhagsPa            = `[\p{Phags_Pa}]`            // UnicodePhagsPa matches Phags-pa script. 八思巴文
	UnicodePhoenician         = `[\p{Phoenician}]`          // UnicodePhoenician matches Phoenician script. 腓尼基文
	UnicodeRejang             = `[\p{Rejang}]`              // UnicodeRejang matches Rejang script. 拉让文
	UnicodeRunic              = `[\p{Runic}]`               // UnicodeRunic matches Runic script. 卢恩文
	UnicodeSaurashtra         = `[\p{Saurashtra}]`          // UnicodeSaurashtra matches Saurashtra script. 索拉什特拉文
	UnicodeShavian            = `[\p{Shavian}]`             // UnicodeShavian matches Shavian script. 萧伯纳文
	UnicodeSinhala            = `[\p{Sinhala}]`             // UnicodeSinhala matches Sinhala script. 僧伽罗文
	UnicodeSundanese          = `[\p{Sundanese}]`           // UnicodeSundanese matches Sundanese script. 巽他文
	UnicodeSylotiNagri        = `[\p{Syloti_Nagri}]`        // UnicodeSylotiNagri matches Syloti Nagri script. 锡尔赫特文
	UnicodeSyriac             = `[\p{Syriac}]`              // UnicodeSyriac matches Syriac script. 叙利亚文
	UnicodeTagalog            = `[\p{Tagalog}]`             // UnicodeTagalog matches Tagalog script. 塔加拉文
	UnicodeTagbanwa           = `[\p{Tagbanwa}]`            // UnicodeTagbanwa matches Tagbanwa script. 塔格巴努亚文
	UnicodeTaiLe              = `[\p{Tai_Le}]`              // UnicodeTaiLe matches Tai Le script. 傣仂文
	UnicodeTamil              = `[\p{Tamil}]`               // UnicodeTamil matches Tamil script. 泰米尔文
	UnicodeTelugu             = `[\p{Telugu}]`              // UnicodeTelugu matches Telugu script. 泰卢固文
	UnicodeThaana             = `[\p{Thaana}]`              // UnicodeThaana matches Thaana script. 塔安那文
	UnicodeThai               = `[\p{Thai}]`                // UnicodeThai matches Thai script. 泰文
	UnicodeTibetan            = `[\p{Tibetan}]`             // UnicodeTibetan matches Tibetan script. 藏文
	UnicodeTifinagh           = `[\p{Tifinagh}]`            // UnicodeTifinagh matches Tifinagh script. 提非纳文
	UnicodeUgaritic           = `[\p{Ugaritic}]`            // UnicodeUgaritic matches Ugaritic script. 乌加里特文
	UnicodeVai                = `[\p{Vai}]`                 // UnicodeVai matches Vai script. 瓦伊文
	UnicodeYi                 = `[\p{Yi}]`                  // UnicodeYi matches Yi script. 彝文
)
