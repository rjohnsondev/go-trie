
package trie

import (
    //"strings"
    "fmt"
)

const startLetter = '.'
const endLetter = 'Z'
const noLetters= int(endLetter) - int(startLetter) +1 // optimised for english... with a couple left over

type branch struct {
    children []*branch
    value interface{}
    shortcut []byte
}

type Trie struct {
    tree *branch
    unicodeMap map[int] int
    nextIndex int
}

func (this *Trie) GetKey(ch byte) int {
    ir := int(ch)
    index := -1
    if ir >= startLetter && ir <= endLetter {
        // we use it's key..
        index = ir - startLetter
    } else {
        mapindex, exists := this.unicodeMap[ir]
        if !exists {
            index = this.nextIndex
            this.nextIndex++
            this.unicodeMap[ir] = index
        } else {
            index = mapindex
        }
    }
    return index
}

func (this *Trie) EnsureCapacity(children []*branch, index int) []*branch {
    if len(children) < index+1 {
        for x := len(children); x < index+1; x++ {
            children = append(children, nil)
        }
    }
    return children
}

func (this *Trie) AddEntry(entry string, value interface{}) {
    //entry = strings.ToUpper(entry)
    this.AddToBranch(this.tree, []byte(entry), value)
}

func (this *Trie) AddToBranch(t *branch, remEntry []byte, value interface{}) {

    // can we cheat?
    if t.shortcut == nil {
        t.shortcut = remEntry
        t.value = value
        t.children = nil // not needed, but it helps think things through
        return
    }

    shortcut := t.shortcut

    // are we on the right branch yet?
    if len(remEntry) == 0 {
        // we are here, set it and forget it
        t.value = value
        return
    } else {

        // find common prefix
        smallestLen := len(remEntry)
        if smallestLen > len(shortcut) {
            smallestLen = len(shortcut)
        }
        var x int
        for x = 0; x < smallestLen && shortcut[x] == remEntry[x]; x++ {

        }
        commonPrefix := shortcut[0:x]
        if x < len(shortcut) {
            // we can assign the t to a child
            ttail := shortcut[x+1:len(shortcut)]
            tkey := this.GetKey(shortcut[x])
            newTBranch := &branch {
                children: t.children,
                value: t.value,
                shortcut: ttail,
            }
            t.children = make([]*branch, noLetters, noLetters)
            t.children = this.EnsureCapacity(t.children, tkey)
            t.children[tkey] = newTBranch
            t.shortcut = commonPrefix
            t.value = nil
        } else {
            // the value of t remains
        }
        if x < len(remEntry) {
            // we can assign the v to a child
            vkey := this.GetKey(remEntry[x])
            vtail := remEntry[x+1:len(remEntry)]
            t.children = this.EnsureCapacity(t.children, vkey)
            if t.children[vkey] == nil {
                newVBranch := &branch {
                    children: nil,
                    value: nil,
                    shortcut: nil,
                }
                t.children[vkey] = newVBranch
            }
            this.AddToBranch(t.children[vkey], vtail, value)
        } else {
            // the value of v now takes up the position
            t.value = value
        }

    }
}

func (this *Trie) DumpTree() {
    fmt.Printf("\n\n")
    this.DumpBranch(this.tree, 1)
}

func (this *Trie) DumpBranch(t *branch, depth int) {
    // attempt to output a textual view of the tree
    for x := 0; x < depth; x ++ { fmt.Print("  ") }
    fmt.Printf("- cheat: %s\n", t.shortcut)
    for x := 0; x < depth; x ++ { fmt.Print("  ") }
    fmt.Printf("- value: %s\n",t.value)
    if t.children != nil {
        for x := 0; x < depth; x ++ { fmt.Print("  ") }
        fmt.Printf("- children:\n")
        for y := 0; y < len(t.children); y++ {
            if t.children[y] != nil {
                for x := 0; x < depth; x ++ { fmt.Print("  ") }
                charb := make([]byte, 1)
                charb[0] = byte(y)+startLetter
                fmt.Printf(" - %s\n", string(charb))
                this.DumpBranch(t.children[y], depth+1)
            }
        }
    }
}

func (this *Trie) GetEntry(entry string) (value interface{}, validPath bool) {
    t := this.tree
    //entry = strings.ToUpper(entry)
    eb := []byte(entry)
    // it's <= here to ensure we get to the cheat comparison nil on a valid path
    for x := 0; x <= len(eb); x++ {
        // if the current branch has a cheat, make sure we match it
        s := t.shortcut
        var y int
        for y = 0; y < len(s); y++ {
            if x+y >= len(eb) {
                return nil, true
            }
            if s[y] != eb[x+y] {
                return nil, false
            }
        }
        x += y
        if x < len(eb) {
            // we got through the cheat!
            index := -1
            // we don't use GetKey here as we don't want reads to pollute our hashmap
            ir := int(eb[x])
            if ir >= startLetter && ir <= endLetter {
                index = ir - startLetter
            } else {
                mapindex, exists := this.unicodeMap[ir]
                if !exists {
                    return nil, false // no mapping :/
                } else {
                    index = mapindex
                }
            }
            if index > len(t.children)-1 || t.children[index] == nil {
                return nil, false
            }
            t = t.children[index]
            eb = eb[x:]
            x = 0
        }
    }
    return t.value, true
}

func NewTrie() *Trie {
    t := &Trie {
        tree: &branch {
            children: nil,
            value: nil,
            shortcut: nil,
        },
        unicodeMap: make(map[int]int),
        nextIndex: 26,
    }
    return t
}
