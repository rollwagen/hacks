# ytt

ytt is a Python package that provides a command-line tool to fetch and print the
transcript of a given YouTube video URL. It utilizes the
`langchain_community.document_loaders.YoutubeLoader` to load the transcript from
the YouTube API.

## Installation

To install the package, you can use `pipx` (a tool to install Python packages in an isolated environment, see <https://github.com/pypa/pipx>):

```bash
pipx install --force .
```

## Usage

To use the `ytt` command-line tool, simply provide the YouTube video URL as an argument:

```bash
ytt <video_url>
```

Example:

```bash
ytt "https://www.youtube.com/watch?v=8clH7cbnIQw"
```

This will print the transcript of the specified YouTube video to the console.

_Note:_ depending on the URL/your shell, you might have to wrap the video URL in quotes (' or ")
