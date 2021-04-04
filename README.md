# Pac Go
A terminal-based implementation of the classic Pac-Man game, written in Go. Inspired by [danicat](https://github.com/danicat)'s [_Pac Go_](https://github.com/danicat/pacgo).

## Installation
1. Ensure that you have [Go](https://golang.org/doc/install) installed on your machine
2. Clone this repository
3. In the project directory, run:
```sh
$ go build
```

## Usage
Once built, you can execute the binary by running:
```sh
$ path/to/binary/pac-go
```
> Note: Executing the binary with no options specified will start a game using the default options.

## Gameplay
- Executing the binary (with or without command-line options) begins a new game.
- Use the arrow keys to move. Pac-Man does not move automatically.
- Pac-Man has 3 lives by default. You can customize the number of lives using the `-l`, `--player-lives` flag.
- Ghosts move randomly, a single space at a time, at 200ms intervals by default. You can customize this interval to increase/decrease gameplay speed (and difficuly!) using the `-f`, `--framerate` flag. Note that regardless of the specified framerate, each Ghost moves a single space per space moved by Pac-Man.
- Collecting a Dot scores 1 point.
- Collecting a Pill scores 10 points, and renders each Ghost vulnerable for 20 seconds by default. You can customize each of these values using the `-p`, `--pill-score` and `-d`, `--pill-duration` flags.
- Collecting a Pill while one is already active stacks the Pill duration. Use them wisely!
- Ghosts can only be defeated by colliding with one while a Pill is active. Defeating a Ghost scores 15 points by default. You can customize this value using the `-g`, `--ghost-defeat-score` flag.
- Colliding with a Ghost while no Pill is active results in a lost life.
- Games have no time limit.
- The player is victorious when all Dot and Pill characters are collected by Pac-Man.
- The player is defeated when Pac-Man loses all lives, or when the player quits the game using either `ESC` or `⌃` + `C`.

## Customization
### Command Line Flags
You can modify several aspects of the game using command line flags while executing the binary.
| Flag                         | Type       | Default               | Description                                                                                                                                                                                                  |
|------------------------------|------------|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `-c`, `--config-file`        | `string`   | `"./lib/config.json"` | A relative path to a custom JSON configuration file. See below for [configuration file formatting specifications](#config-files).                                        |
| `-m`, `--maze-file`          | `string`   | `"./lib/maze01.txt"`  | A relative path to a custom maze definition file. See below for [more information on maze definition files](#maze-definition-files).                                     |
| `-l`, `--player-lives`       | `int`      | `3`                   | The amount of lives with which Pac-Man begins the game.                                                                                                                                                      |
| `-p`, `--pill-score`         | `int`      | `10`                  | The amount of points scored when reaching a Pill (normal Dots are always worth 1 point).                                                                                                                     |
| `-d`, `--pill-duration`      | `duration` | `20s`                 | The amount of time for which a Pill is effective, rendering Ghosts vulnerable to Pac-Man. This option accepts any value acceptable to [`time.ParseDuration`](https://pkg.go.dev/time?tab=doc#ParseDuration). |
| `-g`, `--ghost-defeat-score` | `int`      | `15`                  | The amount of points scored when defeating a vulnerable Ghost.                                                                                                                                               |
| `-f`, `--framerate`          | `duration` | `200ms`               | The speed at which to render the game. This option accepts any value acceptable to [`time.ParseDuration`](https://pkg.go.dev/time?tab=doc#ParseDuration).                                                    |
| `-h`, `--help`               |            |                       | Prints usage instructions.                                                                                                                                                                                   |

### Config Files
The default configuration file (located at `lib/config.json`) is formatted with the following structure:
```json
{
  "player": "😃",
  "ghost": "👻",
  "ghost_blue": "💩",
  "wall": "  ",
  "dot": "• ",
  "pill": "💊",
  "death": "💀",
  "space": "  ",
  "use_emoji": true
}
```
Each field is required, but you're welcome to change the values to customize your experience. Pay special attention to the `"use_emoji"` field, as emoji must be rendered as double-width characters. If the value of this field is not consistent with the values of the other fields (single-width ASCII characters vs. emoji), the game will render with incorrect spacing. The configutation file located at [`lib/config_noemoji.json`](./lib/config_noemoji.json) provides an example using regular ASCII characters and `"use_emoji": false`.

### Maze Definition Files
The default maze definition file (located at [`lib/maze01.txt`](./lib/maze01.txt)) looks like:
```txt
############################
#............##............#
#.####.#####.##.#####.####.#
#X####.#####.##.#####.####X#
#..........................#
#.####.##.########.##.####.#
#......##....##....##......#
######.##### ## #####.######
     #.##          ##.#
     #.## ###--### ##.#
######.## # GGGG # ##.######
      .   # GGGG #   .      |
######.## # GGGG # ##.######
     #.## ######## ##.#
     #.##    P     ##.#
######.## ######## ##.######
#............##............#
#.####.#####.##.#####.####.#
#X..##................##..X#
###.##.##.########.##.##.###
#......##....##....##......#
#.##########.##.##########.#
#..........................#
############################
```
Note the `|` character on line 12. Whitespace must exist up to the `|` in order for the wrap-around effect to work correctly. The `|` character is not present in the actual [`maze01.txt` file](./lib/maze01.txt). When designing a custom maze definition file, if the wrap-around effect is desired, the row(s) on which it should occur must be the maximum width of the maze. Similarly, you can create a wrap-around effect by including whitespace at the top and bottom of the same column, and each column in which the wrap-around should occur must be the maximum height of the maze.

Each of the different characters used in the above file renders as a specific object. These same characters must be used when writing a custom maze definition file. For reference, the available characters and their resulting values are:
| Character |  Result |
|:---------:|:-------:|
|    `P`    | Pac-Man |
|    `G`    |  Ghost  |
|    `X`    |   Pill  |
|    `.`    |   Dot   |
|    `#`    |   Wall  |
|   other   |  Space  |
