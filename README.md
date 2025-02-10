# SkillsForge-2025-Assessment
SkillsForge Technical assessment for 2025 Internship role
Track: Backend Development

### File Structure

├── section-one/
│   ├── non-repeating-char.go
│   └── string-compression.go
├── logservice/      
│   └── main.go
└── README.md


## Section one

**Question One**
Implement a string compression. For example, aaaabbbccddddddee would become a4b3c2d6e2. If the
length of the string is not reduced, return the original string.

Test Cases.
---
section_one.StringCompress("bbcceeee") == ‘b2c2e4’
section_one.StringCompress("aaabbbcccaaa") == ‘a3b3c3a3’
section_one.StringCompress("a") = a

**Solution** : In file string-compression.go - Function StringCompress 
![s1.png](../../Pictures/s1.png)

**Question Two**
Given a string, write a function to find the first non-repeating character. If all characters
repeat, return -1.

Test Cases.
---
section_one.FirstNonRepeating("swiss") == 'w'

**Solution** : In file non-repeating-char.go - Function FirstNonRepeating
![s2.png](../../Pictures/s2.png)

