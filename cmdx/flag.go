//
//Created by xuzhuoxi
//on 2019-06-07.
//@author xuzhuoxi
//
package cmdx

import (
	"errors"
	"flag"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

func NewFlagSetExtend(name string, errorHandling flag.ErrorHandling) *FlagSetExtend {
	return &FlagSetExtend{
		FlagSet: *flag.NewFlagSet(name, errorHandling), errorHandling: errorHandling}
}

type FlagSetExtend struct {
	flag.FlagSet
	keyList       []string
	errorHandling flag.ErrorHandling
}

func (fs *FlagSetExtend) Parse(arguments []string) error {
	err := fs.FlagSet.Parse(arguments)
	if nil != err {
		return err
	}
	for _, kv := range arguments {
		s := strings.IndexByte(kv, '-')
		e := strings.LastIndexByte(kv, '=')
		if -1 == s || -1 == e {
			return errors.New("bad flag syntax: " + kv)
		}
		key := kv[s+1 : e]
		fs.keyList = append(fs.keyList, key)
	}
	return nil
}

func (fs *FlagSetExtend) CheckKey(key string) bool {
	_, ok := slicex.IndexString(fs.keyList, strings.ToLower(key))
	return ok
}
