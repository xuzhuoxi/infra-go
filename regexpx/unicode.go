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

//Unicode类
const (
	U_Arabic              = `[\p{Arabic}]`              //阿拉伯文
	U_Armenian            = `[\p{Armenian}]`            //亚美尼亚文
	U_Balinese            = `[\p{Balinese}]`            //巴厘岛文
	U_Bengali             = `[\p{Bengali}]`             //孟加拉文
	U_Bopomofo            = `[\p{Bopomofo}]`            //汉语拼音字母
	U_Braille             = `[\p{Braille}]`             //盲文
	U_Buginese            = `[\p{Buginese}]`            //布吉文
	U_Buhid               = `[\p{Buhid}]`               //布希德文
	U_Canadian_Aboriginal = `[\p{Canadian_Aboriginal}]` //加拿大土著文
	U_Carian              = `[\p{Carian}]`              //卡里亚文
	U_Cham                = `[\p{Cham}]`                //占族文
	U_Cherokee            = `[\p{Cherokee}]`            //切诺基文
	U_Common              = `[\p{Common}]`              //普通的，字符不是特定于一个脚本
	U_Coptic              = `[\p{Coptic}]`              //科普特文
	U_Cuneiform           = `[\p{Cuneiform}]`           //楔形文字
	U_Cypriot             = `[\p{Cypriot}]`             //塞浦路斯文
	U_Cyrillic            = `[\p{Cyrillic}]`            //斯拉夫文
	U_Deseret             = `[\p{Deseret}]`             //犹他州文
	U_Devanagari          = `[\p{Devanagari}]`          //梵文
	U_Ethiopic            = `[\p{Ethiopic}]`            //衣索比亚文
	U_Georgian            = `[\p{Georgian}]`            //格鲁吉亚文
	U_Glagolitic          = `[\p{Glagolitic}]`          //格拉哥里文
	U_Gothic              = `[\p{Gothic}]`              //哥特文
	U_Greek               = `[\p{Greek}]`               //希腊
	U_Gujarati            = `[\p{Gujarati}]`            //古吉拉特文
	U_Gurmukhi            = `[\p{Gurmukhi}]`            //果鲁穆奇文
	U_Han                 = `[\p{Han}]`                 //汉文
	U_Hangul              = `[\p{Hangul}]`              //韩文
	U_Hanunoo             = `[\p{Hanunoo}]`             //哈鲁喏文
	U_Hebrew              = `[\p{Hebrew}]`              //希伯来文
	U_Hiragana            = `[\p{Hiragana}]`            //平假名（日语）
	U_Inherited           = `[\p{Inherited}]`           //继承前一个字符的脚本
	U_Kannada             = `[\p{Kannada}]`             //坎那达文
	U_Katakana            = `[\p{Katakana}]`            //片假名（日语）
	U_Kayah_Li            = `[\p{Kayah_Li}]`            //克耶字母
	U_Kharoshthi          = `[\p{Kharoshthi}]`          //卡罗须提文
	U_Khmer               = `[\p{Khmer}]`               //高棉文
	U_Lao                 = `[\p{Lao}]`                 //老挝文
	U_Latin               = `[\p{Latin}]`               //拉丁文
	U_Lepcha              = `[\p{Lepcha}]`              //雷布查文
	U_Limbu               = `[\p{Limbu}]`               //林布文
	U_Linear_B            = `[\p{Linear_B}]`            //B类线形文字（古希腊）
	U_Lycian              = `[\p{Lycian}]`              //利西亚文
	U_Lydian              = `[\p{Lydian}]`              //吕底亚文
	U_Malayalam           = `[\p{Malayalam}]`           //马拉雅拉姆文
	U_Mongolian           = `[\p{Mongolian}]`           //蒙古文
	U_Myanmar             = `[\p{Myanmar}]`             //缅甸文
	U_New_Tai_Lue         = `[\p{New_Tai_Lue}]`         //新傣仂文
	U_Nko                 = `[\p{Nko}]`                 //Nko文
	U_Ogham               = `[\p{Ogham}]`               //欧甘文
	U_Ol_Chiki            = `[\p{Ol_Chiki}]`            //桑塔利文
	U_Old_Italic          = `[\p{Old_Italic}]`          //古意大利文
	U_Old_Persian         = `[\p{Old_Persian}]`         //古波斯文
	U_Oriya               = `[\p{Oriya}]`               //奥里亚文
	U_Osmanya             = `[\p{Osmanya}]`             //奥斯曼亚文
	U_Phags_Pa            = `[\p{Phags_Pa}]`            //八思巴文
	U_Phoenician          = `[\p{Phoenician}]`          //腓尼基文
	U_Rejang              = `[\p{Rejang}]`              //拉让文
	U_Runic               = `[\p{Runic}]`               //古代北欧文字
	U_Saurashtra          = `[\p{Saurashtra}]`          //索拉什特拉文（印度县城）
	U_Shavian             = `[\p{Shavian}]`             //萧伯纳文
	U_Sinhala             = `[\p{Sinhala}]`             //僧伽罗文
	U_Sundanese           = `[\p{Sundanese}]`           //巽他文
	U_Syloti_Nagri        = `[\p{Syloti_Nagri}]`        //锡尔赫特文
	U_Syriac              = `[\p{Syriac}]`              //叙利亚文
	U_Tagalog             = `[\p{Tagalog}]`             //塔加拉文
	U_Tagbanwa            = `[\p{Tagbanwa}]`            //塔格巴努亚文
	U_Tai_Le              = `[\p{Tai_Le}]`              //德宏傣文
	U_Tamil               = `[\p{Tamil}]`               //泰米尔文
	U_Telugu              = `[\p{Telugu}]`              //泰卢固文
	U_Thaana              = `[\p{Thaana}]`              //塔安那文
	U_Thai                = `[\p{Thai}]`                //泰文
	U_Tibetan             = `[\p{Tibetan}]`             //藏文
	U_Tifinagh            = `[\p{Tifinagh}]`            //提非纳文
	U_Ugaritic            = `[\p{Ugaritic}]`            //乌加里特文
	U_Vai                 = `[\p{Vai}]`                 //瓦伊文
	U_Yi                  = `[\p{Yi}]`                  // 彝文
)
