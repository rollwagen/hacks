import argparse
from langchain_community.document_loaders import YoutubeLoader


def main():
    parser = argparse.ArgumentParser(description="YouTube Transcript Downloader")
    parser.add_argument("url", help="YouTube video URL")
    parser.add_argument("--lang", "-l", default="en", help="Language code for transcript (default: en)")
    args = parser.parse_args()

    loader = YoutubeLoader.from_youtube_url(
        args.url, add_video_info=False, language=args.lang
    )
    d = loader.load()
    print(d[0].page_content)


if __name__ == "__main__":
    main()
