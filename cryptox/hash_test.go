//
//Created by xuzhuoxi
//on 2019-02-05.
//@author xuzhuoxi
//
package cryptox

import (
	"fmt"
	"testing"
)

type hashTest struct {
	out       string
	in        string
	halfState string // marshaled hash state after first half of in written, used by TestGoldenMarshal
}

var md5Arr = []hashTest{
	{"d41d8cd98f00b204e9800998ecf8427e", "", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"0cc175b9c0f1b6a831c399e269772661", "a", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"187ef4436122d1cc2f40dc2b92f0eba0", "ab", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tva\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"900150983cd24fb0d6963f7d28e17f72", "abc", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tva\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"e2fc714c4727ee9395f324cd2e7f331f", "abcd", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvab\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"ab56b4d92b40713acc5af89985d4b786", "abcde", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvab\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"e80b5017098950fc58aad83c8c14978e", "abcdef", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvabc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"7ac66c0f148de9519b8bd264312c4d64", "abcdefg", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvabc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"e8dc4081b13434b45189a720b77b6818", "abcdefgh", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvabcd\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"8aa99b1f439ff71293e95357bac6fd94", "abcdefghi", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvabcd\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"a925576942e94b2ef57a066101b48876", "abcdefghij", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvabcde\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05"},
	{"d747fc1719c7eacb84058196cfe56d57", "Discard medicine more than two years old.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvDiscard medicine mor\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x14"},
	{"bff2dcb37ef3a44ba43ab144768ca837", "He who has a shady past knows that nice guys finish last.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvHe who has a shady past know\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"0441015ecb54a7342d017ed1bcfdbea5", "I wouldn't marry him with a ten foot pole.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvI wouldn't marry him \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x15"},
	{"9e3cac8e9e9757a60c3ea391130d3689", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvFree! Free!/A trip/to Mars/f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"a0f04459b031f916a59a35cc482dc039", "The days of the digital watch are numbered.  -Tom Stoppard", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvThe days of the digital watch\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1d"},
	{"e7a48e0fe884faf31475d2a04b1362cc", "Nepal premier won't resign.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvNepal premier\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\r"},
	{"637d2fe925c07c113800509964fb0e06", "For every action there is an equal and opposite government program.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvFor every action there is an equa\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00!"},
	{"834a8d18d5c6562119cf4c7f5086cb71", "His money is twice tainted: 'taint yours and 'taint mine.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvHis money is twice tainted: \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"de3a4d2fd6c73ec2db2abad23b444281", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvThere is no reason for any individual to hav\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00,"},
	{"acf203f997e2cf74ea3aff86985aefaf", "It's a tiny change to the code and not completely disgusting. - Bob Manchek", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvIt's a tiny change to the code and no\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00%"},
	{"e1c1384cb4d2221dfdd7c795a4222c9a", "size:  a.out:  bad magic", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tvsize:  a.out\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\f"},
	{"c90f3ddecc54f34228c063d7525bf644", "The major problem is with sendmail.  -Mark Horton", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvThe major problem is wit\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x18"},
	{"cdf7ab6c1fd49bd9933c43f3ea5af185", "Give me a rock, paper and scissors and I will move the world.  CCFestoon", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvGive me a rock, paper and scissors a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00$"},
	{"83bc85234942fc883c063cbd7f0ad5d0", "If the enemy is within range, then so are you.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvIf the enemy is within \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x17"},
	{"277cbe255686b48dd7e8f389394d9299", "It's well we cannot hear the screams/That we create in others' dreams.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvIt's well we cannot hear the scream\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00#"},
	{"fd3fb0a7ffb8af16603f3d3af98f8e1f", "You remind me of a TV show, but that's all right: I watch it anyway.", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvYou remind me of a TV show, but th\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\""},
	{"469b13a78ebf297ecda64d4723655154", "C is as portable as Stonehedge!!", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvC is as portable\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10"},
	{"63eb3a2f466410104731c4b037600110", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvEven if I could be Shakespeare, I think I sh\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00,"},
	{"72c2ed7592debca1c90fc0100f931a2f", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", "md5\x01\xa7\xc9\x18\x9b\xc3E\x18\xf2\x82\xfd\xf3$\x9d_\v\nem\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00B"},
	{"132f7619d33b523b1d9e5bd8e0928355", "How can you write a big system without C++?  -Paul Glick", "md5\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102TvHow can you write a big syst\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
}
var sha1Arr = []hashTest{
	{"76245dbf96f661bd221046197ab8b9f063f11bad", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n", "sha\x01\v\xa0)I\xdeq(8h\x9ev\xe5\x88[\xf8\x81\x17\xba4Daaaaaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x96"},
	{"da39a3ee5e6b4b0d3255bfef95601890afd80709", "", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", "a", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"da23614e02469a0d7c7bd1bdab5c9c474b1904dc", "ab", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"a9993e364706816aba3e25717850c26c9cd0d89d", "abc", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"81fe8bfe87576c3ecb22426f8e57847382917acf", "abcd", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0ab\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"03de6c570bfe24bfc328ccd7ca46b76eadaf4334", "abcde", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0ab\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"1f8ac10f23c5b5bc1167bda84b833e5c057a77d2", "abcdef", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"2fb5e13419fc89246865e7a324f476ec624e8740", "abcdefg", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"425af12a0743502b322e93a015bcf868e324d56a", "abcdefgh", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcd\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1", "abcdefghi", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcd\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"d68c19a0a345b7eab78d5e11e991c026ec60db63", "abcdefghij", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcde\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05"},
	{"ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7", "Discard medicine more than two years old.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Discard medicine mor\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x14"},
	{"e5dea09392dd886ca63531aaa00571dc07554bb6", "He who has a shady past knows that nice guys finish last.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0He who has a shady past know\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"45988f7234467b94e3e9494434c96ee3609d8f8f", "I wouldn't marry him with a ten foot pole.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0I wouldn't marry him \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x15"},
	{"55dee037eb7460d5a692d1ce11330b260e40c988", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Free! Free!/A trip/to Mars/f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"b7bc5fb91080c7de6b582ea281f8a396d7c0aee8", "The days of the digital watch are numbered.  -Tom Stoppard", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0The days of the digital watch\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1d"},
	{"c3aed9358f7c77f523afe86135f06b95b3999797", "Nepal premier won't resign.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Nepal premier\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\r"},
	{"6e29d302bf6e3a5e4305ff318d983197d6906bb9", "For every action there is an equal and opposite government program.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0For every action there is an equa\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00!"},
	{"597f6a540010f94c15d71806a99a2c8710e747bd", "His money is twice tainted: 'taint yours and 'taint mine.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0His money is twice tainted: \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"6859733b2590a8a091cecf50086febc5ceef1e80", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0There is no reason for any individual to hav\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00,"},
	{"514b2630ec089b8aee18795fc0cf1f4860cdacad", "It's a tiny change to the code and not completely disgusting. - Bob Manchek", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0It's a tiny change to the code and no\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00%"},
	{"c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69", "size:  a.out:  bad magic", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0size:  a.out\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\f"},
	{"74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a", "The major problem is with sendmail.  -Mark Horton", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0The major problem is wit\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x18"},
	{"0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66", "Give me a rock, paper and scissors and I will move the world.  CCFestoon", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Give me a rock, paper and scissors a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00$"},
	{"3ae7937dd790315beb0f48330e8642237c61550a", "If the enemy is within range, then so are you.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0If the enemy is within \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x17"},
	{"410a2b296df92b9a47412b13281df8f830a9f44b", "It's well we cannot hear the screams/That we create in others' dreams.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0It's well we cannot hear the scream\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00#"},
	{"841e7c85ca1adcddbdd0187f1289acb5c642f7f5", "You remind me of a TV show, but that's all right: I watch it anyway.", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0You remind me of a TV show, but th\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\""},
	{"163173b825d03b952601376b25212df66763e1db", "C is as portable as Stonehedge!!", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0C is as portable\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10"},
	{"32b0377f2687eb88e22106f133c586ab314d5279", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Even if I could be Shakespeare, I think I sh\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00,"},
	{"0885aaf99b569542fd165fa44e322718f4a984e0", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", "sha\x01x}\xf4\r\xeb\xf2\x10\x87\xe8[\xb2JA$D\xb7\u063ax8em\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00B"},
	{"6627d6904d71420b0bf3886ab629623538689f45", "How can you write a big system without C++?  -Paul Glick", "sha\x01gE#\x01\xef\u036b\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0How can you write a big syst\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c"},
}

func TestMd5(t *testing.T) {
	for _, val := range md5Arr {
		out := Md5String(val.in)
		fmt.Println(val.in, "\n\t", out, val.out)
		if out != val.out {
			t.Fatalf("md5(%s) = %s want %s", val.in, out, val.out)
		}
	}
}

func TestSha1(t *testing.T) {
	for _, val := range sha1Arr {
		out := Sha1String(val.in)
		fmt.Println(val.in, "\n\t", out, val.out)
		if out != val.out {
			t.Fatalf("sha1(%s) = %s want %s", val.in, out, val.out)
		}
	}
}
