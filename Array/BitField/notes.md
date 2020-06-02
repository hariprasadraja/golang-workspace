# BitField

In golang, BitField implemented using iota.

## Shift Operation - How it works

It is taken from stack overflow question https://stackoverflow.com/questions/5801008/go-and-operators

```
153

The super (possibly over) simplified definition is just that << is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.

So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.

All the other answers give the more technical definition, but nobody laid it out really bluntly and I thought you might want that.

```
