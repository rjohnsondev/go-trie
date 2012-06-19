/*
Copyright (c) 2012, Richard Johnson
All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

 - Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
 - Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package trie

import "testing"
import "fmt"
import "os"
import "container/list"
import "strings"
//import "math/rand"
//import "bufio"
import "encoding/binary"

func TestMulipleAdditions(t *testing.T) {
    trie := NewTrie()
    trie.AddEntry("", "0")
    val, subpath := trie.GetEntry("")
    if val.(string) != "0" {
        t.Errorf("Failed to retrieve entry with empty key")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("booboo", "1")
    val, subpath = trie.GetEntry("booboo")
    if val.(string) != "1" {
        t.Errorf("Failed to retrieve first entry")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("boogoo", "2")
    val, subpath = trie.GetEntry("boogoo")
    if val.(string) != "2" {
        t.Errorf("Failed to retrieve entry after mid split")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("boodoo", "3")
    val, subpath = trie.GetEntry("boodoo")
    if val.(string) != "3" {
        t.Errorf("Failed to retrieve entry after additional mid split")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("boodod", "4")
    val, subpath = trie.GetEntry("boodod")
    if val.(string) != "4" {
        t.Errorf("Failed to retrieve entry after tail variation")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("aoodod", "5")
    val, subpath = trie.GetEntry("aoodod")
    if val.(string) != "5" {
        t.Errorf("Failed to retrieve entry after lead variation")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    trie.AddEntry("你好世界", "6")
    val, subpath = trie.GetEntry("你好世界")
    if val.(string) != "6" {
        t.Errorf("Failed to retrieve unicode entry")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

    val, subpath = trie.GetEntry("")
    if val.(string) != "0" {
        t.Errorf("Second sweep: Failed to retrieve entry with empty key")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("booboo")
    if val.(string) != "1" {
        t.Errorf("Second sweep: Failed to retrieve first entry")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("boogoo")
    if val.(string) != "2" {
        t.Errorf("Second sweep: Failed to retrieve entry after mid split")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("boodoo")
    if val.(string) != "3" {
        t.Errorf("Second sweep: Failed to retrieve entry after additional mid split")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("boodod")
    if val.(string) != "4" {
        t.Errorf("Second sweep: Failed to retrieve entry after tail variation")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("aoodod")
    if val.(string) != "5" {
        t.Errorf("Second sweep: Failed to retrieve entry after lead variation")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }
    val, subpath = trie.GetEntry("你好世界")
    if val.(string) != "6" {
        t.Errorf("Second sweep: Failed to retrieve unicode entry")
    }
    if subpath != true {
        t.Errorf("Valid response with invalid subpath")
    }

}

func TestValidPaths(t *testing.T) {
    trie := NewTrie()

    trie.AddEntry("aaaaa", "1")
    val, validPath := trie.GetEntry("aaa")
    if val != nil {
        t.Errorf("Value returned for subpath")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }
    trie.AddEntry("aab", "2")
    val, validPath = trie.GetEntry("aaa")
    if val != nil {
        t.Errorf("Value returned for subpath "+val.(string))
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }
    trie.AddEntry("aaba", "3")
    val, validPath = trie.GetEntry("aaa")
    if val != nil {
        t.Errorf("Value returned for subpath "+val.(string))
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }
    trie.AddEntry("abaa", "4")
    val, validPath = trie.GetEntry("abb")
    if val != nil {
        t.Errorf("Value returned for subpath "+val.(string))
    }
    if validPath == true {
        t.Errorf("Valid false positive")
    }

}

func TestMins(t *testing.T) {
    trie := NewTrie()

    trie.AddEntry("a", "1")
    val, validPath := trie.GetEntry("a")
    if val == nil {
        t.Errorf("Couldn't get a")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    trie.AddEntry("b", "2")
    val, validPath = trie.GetEntry("b")
    if val == nil {
        t.Errorf("Couldn't get b")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    trie.AddEntry("aa", "3")
    val, validPath = trie.GetEntry("aa")
    if val == nil {
        t.Errorf("Couldn't get aa")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }
}

func TestStuff(t *testing.T) {
    trie := NewTrie()

    trie.AddEntry("ebay", "1")
    val, validPath := trie.GetEntry("ebay")
    if val.(string) != "1" {
        t.Errorf("Unable to retrieve ebay 1")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    trie.AddEntry("ebays", "2")
    val, validPath = trie.GetEntry("ebays")
    if val.(string) != "2" {
        t.Errorf("Unable to retrieve ebays 2")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    trie.AddEntry("eba", "4")
    trie.AddEntry("ebay asdf", "5")

    val, validPath = trie.GetEntry("ebay")
    if val.(string) != "1" {
        t.Errorf("Unable to retrieve ebay 3")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

}

func TestValidBranching(t *testing.T) {

    trie := NewTrie()
    trie.AddEntry("shure asdf", "7")
    trie.AddEntry("shure qwer", "8")
    trie.AddEntry("shurtrax max-pax", "9")
    trie.AddEntry("shura no toki", "10")
    trie.AddEntry("shure", "6")

    val, validPath := trie.GetEntry("shure")
    if val.(string) != "6" {
        t.Errorf("Unable to retrieve shure 6")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    val, validPath = trie.GetEntry("shure asdf")
    if val.(string) != "7" {
        t.Errorf("Unable to retrieve shure 7")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    val, validPath = trie.GetEntry("shure qwer")
    if val.(string) != "8" {
        t.Errorf("Unable to retrieve shure 8")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    val, validPath = trie.GetEntry("shurtrax max-pax")
    if val.(string) != "9" {
        t.Errorf("Unable to retrieve shure 9")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

    val, validPath = trie.GetEntry("shura no toki")
    if val.(string) != "10" {
        t.Errorf("Unable to retrieve shure 10")
    }
    if validPath != true {
        t.Errorf("Valid subpath not identified")
    }

}

func TestExteme(t *testing.T) {
    trie := NewTrie()

    trie.AddEntry("+", "1")
    trie.AddEntry(",", "1")
    trie.AddEntry(".", "1")
    trie.AddEntry(".a", "1")
    trie.AddEntry("..", "1")
    trie.AddEntry("[", "1")
    trie.GetEntry("[")
    trie.GetEntry(".")
    trie.GetEntry("..")
    trie.GetEntry(".+")
    trie.GetEntry("[")
}


/**
 * A couple of simple benchmarks
 */

var phraseFile string = "/home/richard/Documents/tsw/120"

func BenchmarkInsertTrie(b *testing.B) {
    b.StopTimer()
    trie := NewTrie()
    f, err := os.Open(phraseFile)
    keys := list.New()
    for x := 0; x < b.N; x++ {
        if err != nil {
            f.Seek(0,0)
        }
        var length uint32
        err = binary.Read(f, binary.LittleEndian, &length)
        str := make([]byte, length)
        _, err = f.Read(str)
        var gc uint32
        err = binary.Read(f, binary.LittleEndian, &gc)
        sstr := strings.ToUpper(string(str))
        sstr += fmt.Sprintf("%d", x)
        keys.PushBack(sstr)
    }
    b.StartTimer()
    x := 0
    for e := keys.Front(); e != nil; e = e.Next() {
        x++
        trie.AddEntry(e.Value.(string), x)
    }
}

func BenchmarkInsertHash(b *testing.B) {
    b.StopTimer()
    f, err := os.Open(phraseFile)
    keys := list.New()
    for x := 0; x < b.N; x++ {
        if err != nil {
            f.Seek(0,0)
        }
        var length uint32
        err = binary.Read(f, binary.LittleEndian, &length)
        str := make([]byte, length)
        _, err = f.Read(str)
        var gc uint32
        err = binary.Read(f, binary.LittleEndian, &gc)
        sstr := strings.ToUpper(string(str))
        sstr += fmt.Sprintf("%d", x)
        keys.PushBack(sstr)
    }
    b.StartTimer()
    var x uint32 = 0
    values := make(map[string]uint32, b.N)
    for e := keys.Front(); e != nil; e = e.Next() {
        x++
        values[e.Value.(string)] = x
    }
}

func BenchmarkFetchTrie(b *testing.B) {
    b.StopTimer()
    trie := NewTrie()
    f, err := os.Open(phraseFile)
    keys := list.New()
    for x := 0; x < b.N; x++ {
        if err != nil {
            f.Seek(0,0)
        }
        var length uint32
        err = binary.Read(f, binary.LittleEndian, &length)
        str := make([]byte, length)
        _, err = f.Read(str)
        var gc uint32
        err = binary.Read(f, binary.LittleEndian, &gc)
        sstr := strings.ToUpper(string(str))
        sstr += fmt.Sprintf("%d", x)
        trie.AddEntry(sstr, gc)
        keys.PushBack(sstr)
    }

    b.StartTimer()
    for e := keys.Front(); e != nil; e = e.Next() {
        _, _ = trie.GetEntry(e.Value.(string))
    }
}

func BenchmarkFetchHashmap(b *testing.B) {
    b.StopTimer()
    f, err := os.Open(phraseFile)
    keys := list.New()
    values := make(map[string]uint32, b.N)
    for x := 0; x < b.N; x++ {
        if err != nil {
            f.Seek(0,0)
        }
        var length uint32
        err = binary.Read(f, binary.LittleEndian, &length)
        str := make([]byte, length)
        _, err = f.Read(str)
        var gc uint32
        err = binary.Read(f, binary.LittleEndian, &gc)
        sstr := strings.ToUpper(string(str))
        sstr += fmt.Sprintf("%d", x)
        values[sstr] = gc
        keys.PushBack(sstr)
    }

    b.StartTimer()
    for e := keys.Front(); e != nil; e = e.Next() {
        _ = values[e.Value.(string)]
    }
}
