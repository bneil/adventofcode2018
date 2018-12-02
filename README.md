# Advent of code 2018

## Day 1
Quick start doing a simple loop over given inputs and tracking the culmination of said inputs

## Day 2

### Part 1:

Was mostly looping through a list and identifying occurences of a letter

#### rules
  - only 2 and 3 duplicate character occurences will be counted
  - if multiple 2 or 3 duplicate char occurences happen its only counted as one

### Part 2:

Lots of really neat strategies could be used to make this process more efficent - I was thinking about doing a presort of list of inputs and channels around those presorted buckets. However I ended up just doing a large O(N2)^2...

#### non rules
  - don't keep comparing already compared words
