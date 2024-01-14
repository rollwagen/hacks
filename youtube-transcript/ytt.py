import sys
from langchain_community.document_loaders import YoutubeLoader

url = sys.argv[1]
print(f'{url=}')

loader = YoutubeLoader.from_youtube_url(
    url, add_video_info=False
    # 'https://www.youtube.com/watch?v=8clH7cbnIQw', add_video_info=False
)
d = loader.load()
print(d[0].page_content)
