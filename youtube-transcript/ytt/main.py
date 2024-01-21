import sys
from langchain_community.document_loaders import YoutubeLoader


def main():
    if len(sys.argv) <= 1:
        print("YouTube Transript: Please provide a video url as first argument", file=sys.stderr)
        sys.exit(1)

    url = sys.argv[1]

    loader = YoutubeLoader.from_youtube_url(
        url, add_video_info=False
        # 'https://www.youtube.com/watch?v=8clH7cbnIQw', add_video_info=False
    )
    d = loader.load()
    print(d[0].page_content)


if __name__ == "__main__":
    main()
