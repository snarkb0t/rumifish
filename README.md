# rumifish
a deadfish interpreter / language shell written in go !
supports the standard `i`, `d`, `s` and `o` commands as well as the optionally implemented `h` (halt) command.

## usage
to enter the REPL:
```
rumifish
```
you can enter `q`, `quit` or `exit` to end the session.

to run a program written in deadfish:
```
rumifish path/to/program
```

## technical stuff
this aims to stay as faithful as possible to the original interpreter following the specifications listed in the [esolang wiki](https://esolangs.org/wiki/Deadfish). as such rumifish only checks if the accumulator == 256 or -1, rather than >= or <=, respectively.

## notes
this does **NOT** support compiling into machine code (yet?)
<br>this project came about because i was interested in writing a brainfuck interpreter and i stumbled upon [this video](https://www.youtube.com/watch?v=4uNM73pfJn0), which briefly touched on writing one for deadfish (albeit in python.)
<br>credits to my dog, rumi, for gracefully lending her name to be used in this project !
