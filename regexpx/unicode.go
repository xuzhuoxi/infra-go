package regexpx

// Unicode普通类
const (
	U_C  = `[\p{C}]`  //-其他-(other)
	U_Cc = `[\p{Cc}]` //控制字符(control)
	U_Cf = `[\p{Cf}]` //格式(format)
	U_Co = `[\p{Co}]` //私人使用区(privateuse)
	U_Cs = `[\p{Cs}]` //代理区(surrogate)

	U_L  = `[\p{L}]`  //-字母-(letter)
	U_Ll = `[\p{Ll}]` //小写字母(lowercaseletter)
	U_Lm = `[\p{Lm}]` //修饰字母(modifierletter)
	U_Lo = `[\p{Lo}]` //其它字母(otherletter)
	U_Lt = `[\p{Lt}]` //首字母大写字母(titlecaseletter)
	U_Lu = `[\p{Lu}]` //大写字母(uppercaseletter)

	U_M  = `[\p{M}]`  //-标记-(mark)
	U_Mc = `[\p{Mc}]` //间距标记(spacingmark)
	U_Me = `[\p{Me}]` //关闭标记(enclosingmark)
	U_Mn = `[\p{Mn}]` //非间距标记(non-spacingmark)

	U_N  = `[\p{N}]`  //-数字-(number)
	U_Nd = `[\p{Nd}]` //十進制数字(decimalnumber)
	U_Nl = `[\p{Nl}]` //字母数字(letternumber)
	U_No = `[\p{No}]` //其它数字(othernumber)

	U_P  = `[\p{P}]`  //-标点-(punctuation)
	U_Pc = `[\p{Pc}]` //连接符标点(connectorpunctuation)
	U_Pd = `[\p{Pd}]` //破折号标点符号(dashpunctuation)
	U_Pe = `[\p{Pe}]` //关闭的标点符号(closepunctuation)
	U_Pf = `[\p{Pf}]` //最后的标点符号(finalpunctuation)
	U_Pi = `[\p{Pi}]` //最初的标点符号(initialpunctuation)
	U_Po = `[\p{Po}]` //其他标点符号(otherpunctuation)
	U_Ps = `[\p{Ps}]` //开放的标点符号(openpunctuation)

	U_S  = `[\p{S}]`  //-符号-(symbol)
	U_Sc = `[\p{Sc}]` //货币符号(currencysymbol)
	U_Sk = `[\p{Sk}]` //修饰符号(modifiersymbol)
	U_Sm = `[\p{Sm}]` //数学符号(mathsymbol)
	U_So = `[\p{So}]` //其他符号(othersymbol)

	U_Z  = `[\p{Z}]`  //-分隔符-(separator)
	U_Zl = `[\p{Zl}]` //行分隔符(lineseparator)
	U_Zp = `[\p{Zp}]` //段落分隔符(paragraphseparator)
	U_Zs = `[\p{Zs}]` //空白分隔符(spaceseparator)

)

// Unicode类
const (
	UnicodeArabic             = `[\p{Arabic}]`              //阿拉伯文
	UnicodeArmenian           = `[\p{Armenian}]`            //亚美尼亚文
	UnicodeBalinese           = `[\p{Balinese}]`            //巴厘岛文
	UnicodeBengali            = `[\p{Bengali}]`             //孟加拉文
	UnicodeBopomofo           = `[\p{Bopomofo}]`            //汉语拼音字母
	UnicodeBraille            = `[\p{Braille}]`             //盲文
	UnicodeBuginese           = `[\p{Buginese}]`            //布吉文
	UnicodeBuhid              = `[\p{Buhid}]`               //布希德文
	UnicodeCanadianAboriginal = `[\p{Canadian_Aboriginal}]` //加拿大土著文
	UnicodeCarian             = `[\p{Carian}]`              //卡里亚文
	UnicodeCham               = `[\p{Cham}]`                //占族文
	UnicodeCherokee           = `[\p{Cherokee}]`            //切诺基文
	UnicodeCommon             = `[\p{Common}]`              //普通的，字符不是特定于一个脚本
	UnicodeCoptic             = `[\p{Coptic}]`              //科普特文
	UnicodeCuneiform          = `[\p{Cuneiform}]`           //楔形文字
	UnicodeCypriot            = `[\p{Cypriot}]`             //塞浦路斯文
	UnicodeCyrillic           = `[\p{Cyrillic}]`            //斯拉夫文
	UnicodeDeseret            = `[\p{Deseret}]`             //犹他州文
	UnicodeDevanagari         = `[\p{Devanagari}]`          //梵文
	UnicodeEthiopic           = `[\p{Ethiopic}]`            //衣索比亚文
	UnicodeGeorgian           = `[\p{Georgian}]`            //格鲁吉亚文
	UnicodeGlagolitic         = `[\p{Glagolitic}]`          //格拉哥里文
	UnicodeGothic             = `[\p{Gothic}]`              //哥特文
	UnicodeGreek              = `[\p{Greek}]`               //希腊
	UnicodeGujarati           = `[\p{Gujarati}]`            //古吉拉特文
	UnicodeGurmukhi           = `[\p{Gurmukhi}]`            //果鲁穆奇文
	UnicodeHan                = `[\p{Han}]`                 //汉文
	UnicodeHangul             = `[\p{Hangul}]`              //韩文
	UnicodeHanunoo            = `[\p{Hanunoo}]`             //哈鲁喏文
	UnicodeHebrew             = `[\p{Hebrew}]`              //希伯来文
	UnicodeHiragana           = `[\p{Hiragana}]`            //平假名（日语）
	UnicodeInherited          = `[\p{Inherited}]`           //继承前一个字符的脚本
	UnicodeKannada            = `[\p{Kannada}]`             //坎那达文
	UnicodeKatakana           = `[\p{Katakana}]`            //片假名（日语）
	UnicodeKayahLi            = `[\p{Kayah_Li}]`            //克耶字母
	UnicodeKharoshthi         = `[\p{Kharoshthi}]`          //卡罗须提文
	UnicodeKhmer              = `[\p{Khmer}]`               //高棉文
	UnicodeLao                = `[\p{Lao}]`                 //老挝文
	UnicodeLatin              = `[\p{Latin}]`               //拉丁文
	UnicodeLepcha             = `[\p{Lepcha}]`              //雷布查文
	UnicodeLimbu              = `[\p{Limbu}]`               //林布文
	UnicodeLinearB            = `[\p{Linear_B}]`            //B类线形文字（古希腊）
	UnicodeLycian             = `[\p{Lycian}]`              //利西亚文
	UnicodeLydian             = `[\p{Lydian}]`              //吕底亚文
	UnicodeMalayalam          = `[\p{Malayalam}]`           //马拉雅拉姆文
	UnicodeMongolian          = `[\p{Mongolian}]`           //蒙古文
	UnicodeMyanmar            = `[\p{Myanmar}]`             //缅甸文
	UnicodeNewTaiLue          = `[\p{New_Tai_Lue}]`         //新傣仂文
	UnicodeNko                = `[\p{Nko}]`                 //Nko文
	UnicodeOgham              = `[\p{Ogham}]`               //欧甘文
	UnicodeOlChiki            = `[\p{Ol_Chiki}]`            //桑塔利文
	UnicodeOldItalic          = `[\p{Old_Italic}]`          //古意大利文
	UnicodeOldPersian         = `[\p{Old_Persian}]`         //古波斯文
	UnicodeOriya              = `[\p{Oriya}]`               //奥里亚文
	UnicodeOsmanya            = `[\p{Osmanya}]`             //奥斯曼亚文
	UnicodePhagsPa            = `[\p{Phags_Pa}]`            //八思巴文
	UnicodePhoenician         = `[\p{Phoenician}]`          //腓尼基文
	UnicodeRejang             = `[\p{Rejang}]`              //拉让文
	UnicodeRunic              = `[\p{Runic}]`               //古代北欧文字
	UnicodeSaurashtra         = `[\p{Saurashtra}]`          //索拉什特拉文（印度县城）
	UnicodeShavian            = `[\p{Shavian}]`             //萧伯纳文
	UnicodeSinhala            = `[\p{Sinhala}]`             //僧伽罗文
	UnicodeSundanese          = `[\p{Sundanese}]`           //巽他文
	UnicodeSylotiNagri        = `[\p{Syloti_Nagri}]`        //锡尔赫特文
	UnicodeSyriac             = `[\p{Syriac}]`              //叙利亚文
	UnicodeTagalog            = `[\p{Tagalog}]`             //塔加拉文
	UnicodeTagbanwa           = `[\p{Tagbanwa}]`            //塔格巴努亚文
	UnicodeTaiLe              = `[\p{Tai_Le}]`              //德宏傣文
	UnicodeTamil              = `[\p{Tamil}]`               //泰米尔文
	UnicodeTelugu             = `[\p{Telugu}]`              //泰卢固文
	UnicodeThaana             = `[\p{Thaana}]`              //塔安那文
	UnicodeThai               = `[\p{Thai}]`                //泰文
	UnicodeTibetan            = `[\p{Tibetan}]`             //藏文
	UnicodeTifinagh           = `[\p{Tifinagh}]`            //提非纳文
	UnicodeUgaritic           = `[\p{Ugaritic}]`            //乌加里特文
	UnicodeVai                = `[\p{Vai}]`                 //瓦伊文
	UnicodeYi                 = `[\p{Yi}]`                  // 彝文
)
