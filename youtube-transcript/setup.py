from setuptools import setup, find_packages

DEPENDENCIES = ["langchain_community", "youtube-transcript-api"]

setup(
    name="ytt",
    version="0.1",
    packages=find_packages(),
    entry_points={"console_scripts": ["ytt=ytt.main:main"]},
    install_requires=DEPENDENCIES
)
