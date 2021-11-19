package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

type NoLinkHashTable struct {
	Bucket   []Bucket
	Capacity int32
	Used     int32
}

type Bucket struct {
	Hkey string
	Data int
}

func CreateNoLinkHashTable(length int32) {
	ht := NoLinkHashTable{}
	ht.Capacity = length
	ht.Bucket = make([]Bucket, length)

	for i := 0; i < 5; i++ {
		ht.hSet("abc", i)
	}

	fmt.Println(ht)

	one := ht.hGet("abc")
	fmt.Println(one)
}

func (ht *NoLinkHashTable) hSet(key string, value int) {
	bucket := Bucket{key, value}

	hashCode := GetHashCode(key, 1)
	index := hashCode % (ht.Capacity - 1)
	ht.addBucket(index, key, bucket)
}

func (ht *NoLinkHashTable) hGet(key string) Bucket {
	hashCode := GetHashCode(key, 1)
	index := hashCode % (ht.Capacity - 1)
	return ht.findBucket(index, key)
}

func (ht *NoLinkHashTable) findBucket(index int32, key string) Bucket {
	var empty = Bucket{}
	if ok := ht.Bucket[index]; ok != empty {
		if ht.Bucket[index].Hkey == key {
			return ht.Bucket[index]
		} else {
			index++
		}

		for {
			if ok := ht.Bucket[index]; ok != empty {
				if ht.Bucket[index].Hkey == key {
					return ht.Bucket[index]
				} else {
					index++
				}
			} else {
				return empty
			}
		}
	}

	return empty
}

func (ht *NoLinkHashTable) addBucket(index int32, key string, bucket Bucket) {
	if ht.Capacity == ht.Used {
		panic("Table is full")
	}
	var empty = Bucket{}
	if ok := ht.Bucket[index]; ok != empty {
		if ok.Hkey == key {
			ht.Bucket[index] = bucket
			return
		}
		index++
		for {
			if index >= ht.Capacity {
				index = 0
			}
			if ok := ht.Bucket[index]; ok != empty {
				if ok.Hkey == key {
					ht.Bucket[index] = bucket
					return
				}
				index++
			} else {
				ht.Bucket[index] = bucket
				ht.Used++
				break
			}
		}
	} else {
		ht.Bucket[index] = bucket
		ht.Used++
	}
}

// 感觉这个函数有问题
func GetHashCode(key string, hashType int) int32 {
	var Result []byte
	switch hashType {
	case 1:
		Md5Inst := md5.New()
		Md5Inst.Write([]byte(key))
		Result = Md5Inst.Sum([]byte(""))
	case 2:
		Sha1Inst := sha1.New()
		Sha1Inst.Write([]byte(key))
		Result = Sha1Inst.Sum([]byte(""))
	}
	bin_buf := bytes.NewBuffer(Result)
	var x int32
	binary.Read(bin_buf, binary.BigEndian, &x)
	return x
}
