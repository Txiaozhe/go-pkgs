package met

import (
	"crypto/md5"
	"crypto/sha1"

	"github.com/spaolacci/murmur3"
)

type Source struct {
	Str string
}

func (src *Source) Md5Hash() [16]byte {
	return md5.Sum([]byte(src.Str))
}

func (src *Source) Sha1Hash() [20]byte {
	return sha1.Sum([]byte(src.Str))
}

func (src *Source) Murmur32() uint32 {
	return murmur3.Sum32([]byte(src.Str))
}

func (src *Source) Murmur64() uint64 {
	return murmur3.Sum64([]byte(src.Str))
}
