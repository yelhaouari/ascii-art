ASCII-Art (Standard)
Objectives

ascii-art is a Go program that takes a string as input and outputs a graphical representation using the standard ASCII banner. The program "writes big" by converting text into ASCII art.

Features

Supports letters, numbers, spaces, special characters, and line breaks (\n).

Each character has a height of 8 lines.

Uses the standard banner only.


Example Output

### Usage

```console
test$ go run . "" | cat -e
test$ go run . "\n" | cat -e
$
test$ go run . "Hello" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $

```


Learning Objectives:

Practice using the Go file system (fs) API

Work on data manipulation

Apply good Go coding practices

Implement unit testing

Allowed Packages:

Only standard Go packages.
