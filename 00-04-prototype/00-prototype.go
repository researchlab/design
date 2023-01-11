package prototype

import (
	"encoding/json"
	"time"
)

// prototype 用于深拷贝和分层拷贝比较复杂的克隆对象的场景， 被克隆对象提供一个clone方法， 返回该对象的拷贝

// Keyword 搜索关键字
type Keyword struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

// Clone 使用序列化/反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords 关键字map
type Keywords map[string]*Keyword

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}
	for k, v := range words {
		// 浅拷贝， 直接拷贝了地址
		newKeywords[k] = v
	}

	// 替换需要更新的字段, 这里用的深拷贝
	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
