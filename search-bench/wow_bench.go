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

package main

import "fmt"
import "trie"
import "io/ioutil"
import "strings"
import "time"

func findTree() {
    tree := trie.NewTrie()
    tree.AddEntry("APPEARANCE OF A HUGE CYLINDER", "1")
    tree.AddEntry("APPEARANCES OF THE MARKINGS", "2")
    tree.AddEntry("ITS STRANGE APPEARANCE", "3")
    tree.AddEntry("WIMBLEDON PARTICULARLY HAD SUFFERED", "4")

    // get the file contents
    contents, _ := ioutil.ReadFile("war of the worlds.txt")
    strcontents := strings.ToUpper(string(contents))
    strcontents = strings.Replace(strcontents, "\n", " ", -1)
    strcontents = strings.Replace(strcontents, "\r", " ", -1)
    words := strings.Split(strcontents, " ")
    tmp := ""
    foundEntries := make([]string, 0)
    for x := 0; x < len(words); x++ {
        tmp = ""
        out := ""
        y := 0
        for ;y < 100 && (x+y) < len(words); y++ {
            if y > 0 {
                tmp += " "
            }
            wrd := words[x+y]
            // strip off some common grammar
            if len(wrd) > 1 {
                if wrd[len(wrd)-1] == '.' {
                    wrd = wrd[:len(wrd)-1]
                }
                if wrd[len(wrd)-1] == ',' {
                    wrd = wrd[:len(wrd)-1]
                }
            }
            tmp += wrd
            value, validPath := tree.GetEntry(tmp)
            if !validPath {
                break
            }
            if value != nil {
                out = value.(string)
            }
        }
        if out != "" {
            foundEntries = append(foundEntries, out)
            x += y -1
        }
    }
    fmt.Print("Found: ")
    fmt.Println(foundEntries)
}

func findHashMap() {
    tree := make(map[string]string,5)
    tree["APPEARANCE OF A HUGE CYLINDER"] = "1"
    tree["APPEARANCES OF THE MARKINGS"] = "2"
    tree["ITS STRANGE APPEARANCE"] = "3"
    tree["WIMBLEDON PARTICULARLY HAD SUFFERED"] ="4"

    // get the file contents
    contents, _ := ioutil.ReadFile("war of the worlds.txt")
    strcontents := strings.ToUpper(string(contents))
    strcontents = strings.Replace(strcontents, "\n", " ", -1)
    strcontents = strings.Replace(strcontents, "\r", " ", -1)
    words := strings.Split(strcontents, " ")
    tmp := ""
    foundEntries := make([]string, 0, len(words))
    for x := 0; x < len(words); x++ {
        tmp = ""
        out := ""
        y := 0
        for ; y < 5 && (x+y) < len(words); y++ {
            if y > 0 {
                tmp += " "
            }
            wrd := words[x+y]
            // strip off some common grammar
            if len(wrd) > 1 {
                if wrd[len(wrd)-1] == '.' {
                    wrd = wrd[:len(wrd)-1]
                }
                if wrd[len(wrd)-1] == ',' {
                    wrd = wrd[:len(wrd)-1]
                }
            }
            tmp += wrd
            value, found := tree[tmp]
            if found {
                out = value
            }
        }
        if out != "" {
            foundEntries = append(foundEntries, out)
            x += y -1
        }
    }
    fmt.Print("Found: ")
    fmt.Println(foundEntries)
}

func main() {
    fmt.Println(time.Now())
    fmt.Println("Test Hashmap:")
    findHashMap()
    fmt.Println(time.Now())
    fmt.Println("Test Trie:")
    findTree()
    fmt.Println(time.Now())

}
