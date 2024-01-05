/*
The "Find All Anagrams in a String" problem involves searching for all occurrences of
 anagrams of a given non-empty string p within a larger string s. An anagram is a word or 
 phrase formed by rearranging the letters of another, such as "cinema" and "iceman." 
 In this problem, you are required to find all starting indices of anagrams of p in s.

Problem Statement:
Given a string s and a non-empty string p, find all the starting indices of anagrams of p in s.

Example:
plaintext
Copy code
Input: s = "cbaebabacd", p = "abc"
Output: [0, 6]
Explanation:
The substring with starting index 0 is "cba," which is an anagram of "abc."
The substring with starting index 6 is "bac," which is an anagram of "abc

// understanding problem :
means we have to take two string one is of large size and second is of small size.
then we have to check that how many anagram of smaller string is present in larger string.
and then put that in another empty string.
Algorithm :
    1. take one larger string as a input and one smaller string as input 
    2. find all the anagram of smaller string in larger string.

*/

package main 

import (
    "fmt"
    
)

func main() {
    // declaration of variable
    var first_string string
    var anagram_string string
    //taking input for this
    fmt.Println("Enter the string")
    fmt.Scan(&first_string)
    fmt.Println("The inputed string :",first_string)
    fmt.Println("Input the string whose anagram is to find")
    fmt.Scan(&anagram_string)
    fmt.Println("The anagram string :",anagram_string)
    

}