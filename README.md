# bf

My adventures with the [Brainf**k](https://en.wikipedia.org/wiki/Brainfuck)  programming language.

## Overview

- Minimal, turing-complete programming language
- Has only 8 commands: `><+-.,[]`
- Represented by an array with 30,000 cells (initialized to 0)
- A data pointer points to the current cell

## Commands in detail

- `+`: Increments the value at the current cell by one
- `-`: Decrements the value at the current cell by one
- `>`: Moves the data pointer to the next cell (one to the right)
- `<`: Moves the data pointer to the previous cell (one to the left)
- `.`: Prints the ASCII value at the current cell (ie. 65 is 'A')
- `,`: Reads a single input char into the current cell
- `[`: If the value at the current cell is 0, skips ahead to the the corresponding `]`, otherwise move to the next cell
- `]`: If the value at the current cell is zero, move to the next cell, otherwise, move backwards to the corresponding `[`

## Observations

- `[` and `]` must be balanced in a valid BF program

## References

1. [Learn X (BF) in Y minutes](https://learnxinyminutes.com/docs/bf/)
2. [BF on Wikipedia]()

## License

```
BSD 3-Clause License

Copyright (c) 2017, Sandeep Raju Prabhakar
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
