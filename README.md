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

## Day 3

### Part 1:

Create a fabric then allow areas to be claimed, each entry will enumerate with the number of claims on the given area. Then simply tally those which have 2 or more claims

### Part 2:

I thought of this one like dodgeball, you really only care about which position in the fabric is occupied. If a section is, you can eliminate both claims from the fabric since any conflict destroys each claim.
Really not an efficient task, but I really liked how I messed up some of the golang string interpolation.

- I was using a list to track the claims that were marked as erased, which was a terrible idea. Ended up with way to many entries and fell back to map
- I was doing s := string(posX)+":"+string(posY) and the string conversion was failing - so thats took me some time to go back and correct.
