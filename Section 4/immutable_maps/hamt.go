package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/fnv"
)

const (
	fanoutLog2       = 6
	fanout      uint = 1 << fanoutLog2
	fanMask     uint = fanout - 1
	maxDepth         = 60 / fanoutLog2
	keyNotFound      = "Key not found"
)

type Key []byte

type node interface {
	assoc(shift int, hash uint64, key Key, value interface{}) (last node, leaf *valueNode)
	without(shift int, hash uint64, key Key) node
	find(shift int, hash uint64, key Key) (value interface{}, err error)
	pos() uint64
}

type PersistentMap struct {
	root *bitmapNode
	//	collision map[uint]interface{}
}

type valueNode struct {
	key    Key
	hash   uint64
	value  interface{}
	bitpos uint64
}

type bitmapNode struct {
	childBitmap uint64
	children    []node
	bitpos      uint64
}

func (n *valueNode) assoc(shift int, hash uint64, key Key, val interface{}) (last node, leaf *valueNode) {
	if n.hash == hash {
		n.value = val
		last = n
		leaf = n
	} else {
		nn := &bitmapNode{0, make([]node, 2, 2), n.pos()}
		last = nn
		nn.assoc(shift, n.hash, key, n.value)
		_, leaf = nn.assoc(shift, hash, key, val)
	}

	return last, leaf
}

func (n *valueNode) without(shift int, hash uint64, key Key) node {
	return n
}

func (n *valueNode) find(shift int, hash uint64, key Key) (value interface{}, err error) {
	if hash == n.hash {
		value = n.value
	} else {
		err = errors.New(keyNotFound)
	}
	return value, err
}

func (n *valueNode) pos() uint64 {
	return n.bitpos
}

func (n *bitmapNode) assoc(shift int, hash uint64, key Key, val interface{}) (last node, leaf *valueNode) {
	bitsToShift := uint(shift * fanoutLog2)
	pos := bitpos(hash, bitsToShift)

	if (pos & n.childBitmap) == 0 { //nothing in slot, not found
		//mark our slot taken and xpand our children
		n.childBitmap |= pos
		newChildren := make([]node, (len(n.children) + 1))

		newChildIndex := n.index(pos)
		newChild := &valueNode{key, hash, val, pos}
		newChildren[newChildIndex] = newChild

		for _, c := range n.children {
			if c != nil {
				oldChildNewIndex := n.index(c.pos())
				newChildren[oldChildNewIndex] = c
			}
		}

		n.children = newChildren
		last = n
		leaf = newChild
	} else {
		index := n.index(pos)
		nodeAtIndex := n.children[index]
		last, leaf = nodeAtIndex.assoc(shift+1, hash, key, val)

		if _, isValNode := nodeAtIndex.(*valueNode); isValNode {
			n.children[index] = last
		}
	}
	return last, leaf
}

func (n *bitmapNode) without(shift int, hash uint64, key Key) node {
	return n
}

func (n *bitmapNode) find(shift int, hash uint64, key Key) (value interface{}, err error) {
	bitsToShift := uint(shift * fanoutLog2)
	pos := bitpos(hash, bitsToShift)
	if cMap := n.childBitmap; (pos & cMap) == 0 { //nothing in slot, not found
		err = errors.New(keyNotFound)
	} else {
		index := n.index(pos)
		if int(index) >= len(n.children) {
			err = errors.New("Keys computed index is larger than children")
		} else {
			value, err = n.children[index].find(shift+1, hash, key)
		}
	}

	return value, err
}

func (n *bitmapNode) pos() uint64 {
	return n.bitpos
}

//Shift key hash until leaf with matching key is found or key is not found
func (t *PersistentMap) Get(key Key) (value interface{}, err error) {
	//Hash our key and look for it in the root
	hash := hash(key)
	value, err = t.root.find(0, hash, key)

	return value, err
}

func (t *PersistentMap) Insert(key Key, value interface{}) (n node) {
	hash := hash(key)
	_, n = t.root.assoc(0, hash, key, value)

	return n
}

func New() *PersistentMap {
	return &PersistentMap{
		root: &bitmapNode{},
	}
}

func StringKey(k string) Key {
	return Key([]byte(k))
}

func IntKey(i int) Key {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	return Key(buf.Bytes())
}

func hash(a []byte) uint64 {
	h := fnv.New64()
	h.Write(a)

	return h.Sum64()
}

func shift(hash uint64, shift uint) uint64 {
	if shift == 0 {
		return hash
	}
	return hash >> shift
}

func mask(hash uint64, bshift uint) uint {
	return uint(shift(hash, bshift) & uint64(fanMask))
}

func bitpos(hash uint64, bshift uint) uint64 {
	return 1 << mask(hash, bshift)
}

func (n *bitmapNode) index(onebitset uint64) uint {
	return popcount_2(n.childBitmap & (onebitset - 1))
}
